//go:build v1

package rest

import (
	"cupcake/app/databases"
	"cupcake/entities"
	"cupcake/interface/apis"
	"cupcake/interface/gateways"

	"cupcake/internal/helpers"

	"github.com/gin-gonic/gin"
)

type Api struct{}

func NewApi() apis.Config {
	return &Api{}
}

func (a *Api) Get(g *gin.Context) {

	db := gateways.NewDatabase(databases.NewMongoWrapper())

	user := []entities.User{}
	err := db.Get("users", &user, []gateways.DatabaseWhereQueryBuilder{
		{
			Op:    "eq",
			Field: "Username",
			Value: "test",
		},
	})

	if err != nil {
		g.JSON(500, helpers.HttpResponse{
			Data: err.Error(),
		})
		return
	}

	g.JSON(200, helpers.HttpResponse{
		Data: user,
	})
}
