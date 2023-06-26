package apis

import "github.com/gin-gonic/gin"

type Config interface {
	Get(*gin.Context)
}
