package usecase

import (
	"context"
	"cupcake/entities"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type authenticationImpl struct {
	db *mongo.Database
}

func NewAuthentication(db *mongo.Database) *authenticationImpl {
	return &authenticationImpl{
		db: db,
	}
}

func (l *authenticationImpl) AttemptLogin(username, password string) error {
	user := []entities.User{}

	cur, err := l.db.Collection(new(entities.User).CollectionName()).Aggregate(context.Background(), []bson.M{
		{
			"$match": bson.M{
				"username": username,
			},
		},
	})
	if err != nil {
		fmt.Println("error", err.Error())
		return fmt.Errorf("username or password does not match")
	}

	err = cur.All(context.Background(), &user)
	if err != nil {
		fmt.Println("error", err.Error())
		return fmt.Errorf("username or password does not match")
	}

	if len(user) > 0 {
		return nil
	}

	return nil
}
