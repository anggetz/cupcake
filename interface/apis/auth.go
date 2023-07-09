package apis

import "github.com/gin-gonic/gin"

type Auth interface {
	Login(*gin.Context)
}
