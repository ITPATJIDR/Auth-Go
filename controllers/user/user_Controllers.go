package user

import (
	"go-jwt-api/model"
	"go-jwt-api/utils/db"
	"go-jwt-api/utils/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetALlUser (c *gin.Context) {
  var users []model.User

  header := c.Request.Header.Get("Authorization")
  tokenString := strings.Split(header," ")
  id , err := token.VerifyJwtToken(tokenString[1])

  if err != nil {
    c.JSON(http.StatusOK, gin.H{"status":"500", "message":err.Error()})
  }else{
    if err := db.DB.Find(&users).Error; err != nil{
      c.JSON(http.StatusOK, gin.H{"status":"500", "message": "Internal Server Errro"})
    }

    c.JSON(http.StatusOK, gin.H{"status":"200", "data": users, "userId": id})
  }
}
