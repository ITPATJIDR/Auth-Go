package user

import (
	"go-jwt-api/model"
	"go-jwt-api/utils/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetALlUser (c *gin.Context) {
  var users []model.User
  if err := db.DB.Find(&users).Error; err != nil{
    c.JSON(http.StatusOK, gin.H{"status":"500", "message": "Internal Server Errro"})
  }
  c.JSON(http.StatusOK, gin.H{"status":"200", "data": users})
}
