package routers

import (
	controllers_auth "go-jwt-api/controllers/auth"
  controllers_user "go-jwt-api/controllers/user"

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
    user.POST("/register",controllers_auth.Register)
    user.POST("/login",controllers_auth.Login)
    user.GET("/getAllUser", controllers_user.GetALlUser)
  }

  return r
}
