package helpers

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbOptions struct {
	Username string
	Password string
	Database string
	Host     string
	Port     string
}

func ConnectMongo(dbOption MongoDbOptions) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// if we set the username or password , we will used cred
	mongoCred := ""
	if dbOption.Username != "" && dbOption.Password != "" {
		mongoCred = dbOption.Username + ":" + dbOption.Password + "@"
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+mongoCred+dbOption.Host+":"+dbOption.Port))

	if err != nil {
		return nil, err
	}

	return client.Database(dbOption.Database), nil
}
