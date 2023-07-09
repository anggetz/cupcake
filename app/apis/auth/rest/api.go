package rest

import (
	"cupcake/app/databases"
	"cupcake/interface/apis"
	"cupcake/interface/gateways"
	"cupcake/internal/helpers"
	"cupcake/pkg"
	"cupcake/use_cases/authentication"

	"github.com/gin-gonic/gin"
)

type ImplApi struct {
	conf *pkg.Config
}

func NewApi(conf *pkg.Config) apis.Auth {
	return &ImplApi{
		conf: conf,
	}
}

func (a *ImplApi) Login(gin *gin.Context) {

	db := gateways.NewDatabase(databases.NewMongoWrapper(), &gateways.DatabaseOption{
		Username: a.conf.Databases["mongo"].Username,
		Password: a.conf.Databases["mongo"].Password,
		Database: a.conf.Databases["mongo"].Database,
		Host:     a.conf.Databases["mongo"].Host,
		Port:     a.conf.Databases["mongo"].Port,
	})
	defer db.Close()

	ret, err := authentication.UseCaseLogin(db, "test", "test")
	if err != nil {
		gin.JSON(200, helpers.HttpResponse{
			Data: err.Error(),
		})
		return
	}

	if !ret {
		gin.JSON(200, helpers.HttpResponse{
			Data: "username or password not correct",
		})
		return
	}

	gin.JSON(200, helpers.HttpResponse{
		Data: "login succesfully",
	})
	return
}
