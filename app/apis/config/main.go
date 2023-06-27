package main

import (
	"cupcake/app/apis/config/rest"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
)

type config struct {
	AppName string
}

func main() {

	var conf config
	var configPath = flag.String("config", "app.yml", "please input config path")

	flag.Parse()

	v := viper.New()
	v.SetConfigFile(*configPath)

	if err := v.ReadInConfig(); err != nil {
		fmt.Errorf("unable to read config file. %s", err.Error())
	}

	if err := v.Unmarshal(&conf); err != nil {
		fmt.Errorf("unable to parse config file. %s", err.Error())
	}

	r := gin.Default()

	// register router
	apiGroup := r.Group("/v1/config")
	apiGroup.GET("/get", rest.NewApi().Get)

	// register nats
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)

	confMarshalled, err := json.Marshal(conf)
	if err != nil {
		panic(err)
	}
	nc.Subscribe("config.get", func(msg *nats.Msg) {
		msg.Respond(confMarshalled)
	})

	err = nc.Publish("config.share", confMarshalled)
	if err != nil {
		panic(err)
	}

	r.Run("localhost:8001")
}
