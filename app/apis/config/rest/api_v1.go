//go:build v1

package rest

import (
	"cupcake/app/databases"
	"cupcake/entities"
	"cupcake/interface/apis"
	"cupcake/interface/gateways"

	"cupcake/internal/helpers"
	"cupcake/pkg"

	"github.com/gin-gonic/gin"
)

type Api struct {
	conf *pkg.Config
}

func NewApi(conf *pkg.Config) apis.Config {
	return &Api{
		conf: conf,
	}
}

func (a *Api) Get(g *gin.Context) {

	db := gateways.NewDatabase(databases.NewMongoWrapper(), &gateways.DatabaseOption{
		Username: a.conf.Databases["mongo"].Username,
		Password: a.conf.Databases["mongo"].Password,
		Database: a.conf.Databases["mongo"].Database,
		Host:     a.conf.Databases["mongo"].Host,
		Port:     a.conf.Databases["mongo"].Port,
	})
	defer db.Close()

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
