package main

import (
	"cupcake/app/apis/config/rest"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// register router
	apiGroup := r.Group("/v1/config")
	apiGroup.GET("/get", rest.NewApi().Get)

	// register nats
	// Connect to a server
	// nc, _ := nats.Connect(nats.DefaultURL)
	// nc.Subscribe("config.get", func(msg *nats.Msg) {

	// })

	r.Run("localhost:8000")
}
