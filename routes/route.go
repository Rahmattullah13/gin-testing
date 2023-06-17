package routes

import (
	"gin-test/services/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello Wordl",
		})
	})

	r.POST("/user", users.ControllerRegister)
	r.POST("/user/login", users.ControllerLogin)
	return r
}
