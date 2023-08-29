package ginadaptor

import (
	"github.com/gin-gonic/gin"
)

type RouterFunc func(router *gin.RouterGroup)

func V1() RouterFunc {
	return func(router *gin.RouterGroup) {
		router.
			POST("/ping", func(c *gin.Context) {
				c.JSON(200, "pong")
			})
	}
}
