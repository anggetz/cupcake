//go:build v2

package rest

import (
	"cupcake/interface/apis"

	"cupcake/app/databases"
	"cupcake/entities"
	"cupcake/interface/gateways"
	"github.com/gin-gonic/gin"

	"cupcake/internal/helpers"
)

type Api struct{}

func NewApi() apis.Config {
	return &Api{}
}

func (a *Api) Get(g *gin.Context) {
	db := gateways.NewDatabase(databases.NewPgWrapper())

	user := []entities.User{}
	err := db.Get("users", &user, []gateways.DatabaseWhereQueryBuilder{
		{
			Op:    "eq",
			Field: "name",
			Value: "'test'",
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
