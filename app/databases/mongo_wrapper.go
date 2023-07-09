package databases

import (
	"context"
	"cupcake/interface/gateways"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoWrapper struct {
	dbClient *mongo.Client
	db       *mongo.Database
}

func NewMongoWrapper() gateways.Database {
	return &MongoWrapper{}
}

// whereClause is a string !!!
func (i *MongoWrapper) Get(tableName string, dest interface{}, matchAggregate interface{}) error {
	fmt.Println(matchAggregate)
	cur, err := i.db.Collection(tableName).Aggregate(context.Background(), []bson.M{
		{
			"$match": matchAggregate,
		},
	})
	if err != nil {
		return err
	}

	err = cur.All(context.Background(), dest)

	return err
}

func (i *MongoWrapper) DBClientName() string {
	return "mongo"
}

func (i *MongoWrapper) Close() error {
	return i.dbClient.Disconnect(context.Background())
}

func (i *MongoWrapper) Connect(dbOption gateways.DatabaseOption) (gateways.Database, error) {
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

	i.dbClient = client
	i.db = client.Database(dbOption.Database)

	return i, nil

}
