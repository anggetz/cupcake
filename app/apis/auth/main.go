package main

import (
	"cupcake/pkg"
	"encoding/json"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)

func main() {

	// register nats
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)
	nc.Subscribe("config.share", func(msg *nats.Msg) {
		err := json.Unmarshal(msg.Data, &pkg.Conf)
		if err != nil {
			panic(err)
		}

		log.Println("new config receive", pkg.Conf)
	})

	msg, err := nc.Request("config.get", []byte(""), time.Second*10)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(msg.Data, &pkg.Conf)
	if err != nil {
		panic(err)
	}

	log.Println("config receive", pkg.Conf)

	r := gin.Default()

	// let' try singleton bros

	// register router
	apiGroup := r.Group("/v1/auth")
	apiGroup.GET("/get")

	r.Run("localhost:8002")

	nc.Drain()
}
