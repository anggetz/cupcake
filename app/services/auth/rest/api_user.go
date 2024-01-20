package rest

import (
	"context"
	"cupcake/entities"
	serviceauth "cupcake/interface/service_auth"
	"cupcake/internal/helpers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ImplApiUser struct {
	db *mongo.Database
}

func NewApiUser(db *mongo.Database) serviceauth.ApiUser {
	return &ImplApiUser{
		db: db,
	}
}

func (a *ImplApiUser) Get(gin *gin.Context) {

	users := []entities.User{}

	cur, err := a.db.Collection(new(entities.User).CollectionName()).Aggregate(context.Background(), []bson.M{})
	if err != nil {
		gin.AbortWithStatusJSON(400, helpers.HttpResponse{
			Data: err.Error(),
		})
		return
	}

	err = cur.All(context.Background(), &users)
	if err != nil {
		gin.AbortWithStatusJSON(400, helpers.HttpResponse{
			Data: err.Error(),
		})
		return
	}

	gin.JSON(200, helpers.HttpResponse{
		Data: users,
	})
	return

}
