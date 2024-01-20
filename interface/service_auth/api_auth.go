package serviceauth

import "github.com/gin-gonic/gin"

type ApiAuth interface {
	Login(*gin.Context)
}
