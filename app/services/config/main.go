package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"cupcake/pkg"

	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
)

func main() {

	// var configPath = flag.String("config", "app.yml", "please input config path")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databases := map[string]pkg.DatabaseConfig{}

	// database setup
	dbMongo := pkg.DatabaseConfig{
		Username: os.Getenv("CUPCAKE_DB_MONGO_USERNAME"),
		Password: os.Getenv("CUPCAKE_DB_MONGO_PASSWORD"),
		Database: os.Getenv("CUPCAKE_DB_MONGO_DATABASE"),
		Host:     os.Getenv("CUPCAKE_DB_MONGO_HOST"),
		Port:     os.Getenv("CUPCAKE_DB_MONGO_PORT"),
	}

	databases["mongo"] = dbMongo

	// service setup
	services := map[string]pkg.ServiceConfig{}

	serviceAuth := pkg.ServiceConfig{
		Host: os.Getenv("CUPCAKE_SERVICE_AUTH_HOST"),
	}

	services["auth"] = serviceAuth

	config := &pkg.Config{
		Databases: databases,
		Services:  services,
	}

	// register nats
	// Connect to a server
	nc, _ := nats.Connect(fmt.Sprintf("%s:%s", os.Getenv("CUPCAKE_NATS_HOST"), os.Getenv("CUPCAKE_NATS_PORT")))

	confMarshalled, err := json.Marshal(config)
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

	fmt.Println("service config running")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)
	defer signal.Stop(signals)

	<-signals // wait for signal
	go func() {
		<-signals // hard exit on second signal (in case shutdown gets stuck)
		os.Exit(1)
	}()
}
