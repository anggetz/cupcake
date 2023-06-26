//go:build v1

package rest

import (
	"cupcake/adapters/apis"

	"github.com/gin-gonic/gin"
)

type Api struct{}

func NewApi() apis.Config {
	return &Api{}
}

func (a *Api) Get(g *gin.Context) {
	g.JSON(200, "test v1")
}
