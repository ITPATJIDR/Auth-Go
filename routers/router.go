package routers

import (
	controllers "go-jwt-api/controllers/auth"

	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine {
  r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Message": "I'm Alive...",
		})
	})

  user := r.Group("/user")
  {
    user.POST("/register",controllers.Register)
    user.POST("/login",controllers.Login)
  }

  return r
}
