package middleware

import (
	"go-jwt-api/utils/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)



func Vertify_token_middleware() gin.HandlerFunc {
  return func(c *gin.Context){
    header := c.Request.Header.Get("Authorization")
    tokenString := strings.Split(header," ")
    id, err := token.VerifyJwtToken(tokenString[1])
    if err != nil {
      c.JSON(http.StatusOK, gin.H{"status":"500", "message":err.Error()})
      c.Abort()
    }else{
      c.Set("id", id)
      c.Next()
    }
  }
}
