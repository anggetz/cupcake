package serviceauth

import "github.com/gin-gonic/gin"

type ApiUser interface {
	Get(*gin.Context)
}
