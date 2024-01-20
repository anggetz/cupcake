package main

import (
	"context"
	"cupcake/app/services/auth/rest"
	"cupcake/internal/helpers"
	"cupcake/pkg"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
)

var config pkg.Config
var mongoDatabase *mongo.Database

func setUpDB() error {
	var err error
	// mongoWrapper := databases.MongoWrapper{}
	if mongoDatabase != nil {
		mongoDatabase.Client().Disconnect(context.Background())
	}

	mongoDatabase, err = helpers.ConnectMongo(helpers.MongoDbOptions{
		Username: config.Databases["mongo"].Username,
		Password: config.Databases["mongo"].Password,
		Database: config.Databases["mongo"].Database,
		Host:     config.Databases["mongo"].Host,
		Port:     config.Databases["mongo"].Port,
	})

	if err != nil {
		return err
	}

	return nil
}

func main() {

	// register nats
	// Connect to a server
	nc, _ := nats.Connect(fmt.Sprintf("%s:%s", os.Getenv("CUPCAKE_NATS_HOST"), os.Getenv("CUPCAKE_NATS_PORT")))

	nc.Subscribe("config.share", func(msg *nats.Msg) {
		fmt.Println("new config received")

		err := json.Unmarshal(msg.Data, &config)
		if err != nil {
			panic(err)
		}

		err = setUpDB()
		if err != nil {
			fmt.Println("error setup mongo database")
		}
	})

	msg, err := nc.Request("config.get", []byte(""), time.Second*10)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(msg.Data, &config)
	if err != nil {
		panic(err)
	}

	log.Println("config receive", config)

	err = setUpDB()
	if err != nil {
		fmt.Println("error setup mongo database")
	}

	defer mongoDatabase.Client().Disconnect(context.Background())

	r := gin.Default()

	// register router

	userApiHandler := rest.NewApiUser(mongoDatabase)
	authApiHandler := rest.NewApiAuth(mongoDatabase)

	apiGroup := r.Group("/v1/auth")
	apiGroup.POST("/login", authApiHandler.Login)

	apiGroupUser := apiGroup.Group("/user")
	apiGroupUser.GET("/get", userApiHandler.Get)

	fmt.Println("service auth running on port", config.Services["auth"].Host+":"+config.Services["auth"].Port)

	r.Run(config.Services["auth"].Host)

	nc.Drain()
}
