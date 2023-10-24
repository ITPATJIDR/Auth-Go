package controllers

import (
	"fmt"
	"go-jwt-api/model"
	"go-jwt-api/utils/db"
	"go-jwt-api/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {

	var registerData model.RegisterSturct

	if err := c.ShouldBindJSON(&registerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userExist model.User
	checkUserExist := db.DB.Where("username = ?", registerData.Username).First(&userExist)

	if checkUserExist.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{"status": "400", "message": "Username is already exist !!!"})
		return
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(registerData.Password), 10)

	user := model.User{Username: registerData.Username, Password: string(hashPassword), Fullname: registerData.Fullname, Avatar: registerData.Avatar}

	result := db.DB.Create(&user)

	if result.Error == nil {
		c.JSON(http.StatusOK, gin.H{"status": "200", "message": "Register Success..."})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "200", "message": "Register Failed !!!"})
	}
}

func Login(c *gin.Context) {

	var loginData model.LoginStruct
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userExist model.User
	if err := db.DB.Where("username = ?", loginData.Username).First(&userExist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

  fmt.Printf(userExist.Password)
  err := bcrypt.CompareHashAndPassword([]byte(userExist.Password),[]byte(loginData.Password))
  if err == nil {
    jwtToken, err := token.GenJwtToken(int(userExist.ID))
    if err == nil {
      c.JSON(http.StatusOK, gin.H{"status": "200", "message": "Login Success !!!", "token": jwtToken})
    }else{
      c.JSON(http.StatusOK, gin.H{"status": "400", "message": "Something Wrong !!!"})
    }
  }else{
    c.JSON(http.StatusOK, gin.H{"status": "400", "message": "username or password incrroct !!!"})
  }


}
