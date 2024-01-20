package rest

import (
	serviceauth "cupcake/interface/service_auth"
	"cupcake/internal/helpers"
	"cupcake/usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type ImplApiAuth struct {
	db *mongo.Database
}

func NewApiAuth(db *mongo.Database) serviceauth.ApiAuth {
	return &ImplApiAuth{
		db: db,
	}
}

type requestAuthLogin struct {
	Username string
	Password string
}

func (a *ImplApiAuth) Login(c *gin.Context) {

	r := requestAuthLogin{}

	err := helpers.GetGinContextBody(c, &r)
	if err != nil {
		c.AbortWithStatusJSON(400, helpers.HttpResponse{
			Data: err.Error(),
		})
		return
	}

	err = usecase.NewAuthentication(a.db).AttemptLogin(r.Username, r.Password)
	if err != nil {
		c.AbortWithStatusJSON(400, helpers.HttpResponse{
			Data: err.Error(),
		})
		return
	}

	c.JSON(200, helpers.HttpResponse{
		Data: "successfully login",
	})
	return

}
