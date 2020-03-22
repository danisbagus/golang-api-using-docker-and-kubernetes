package router

import (
	handler "go-project/handlers"
	middleware "go-project/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api")
	{	
		v1.GET("/user", middleware.Auth,handler.GetUser)
		v1.POST("/login", handler.Login)
		// v1.POST("/user", handler.CreateUser)
		// v1.PATCH("/user/:user_id", handler.UpdateUser)
		// v1.Delete("/user/:user_id", handler.DeleteUser)

	}
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "first page go-api",
		})
	})


	return r
}