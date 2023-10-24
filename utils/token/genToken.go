package token

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"

	"time"
)

var secret []byte
  
func GenJwtToken (userId int) (string, error) {
  secret = []byte(os.Getenv("JWT_SECRET_KEY"))
  token := jwt.NewWithClaims(jwt.SigningMethodHS256 ,jwt.MapClaims{
    "id": userId,
    "exp" : time.Now().Add(time.Minute * 10).Unix(),
  })

  tokenString, err := token.SignedString(secret)
  if err != nil {
    return "Something Wrong", err
  }

  return tokenString, nil
}

func VerifyJwtToken(tokenString string) (int, error) {
  secret = []byte(os.Getenv("JWT_SECRET_KEY"))
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("Unexpected siging method: %v", token.Header["alg"])
    }
    return secret, nil
  })

  if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    idClaims, ok := claims["id"].(float64)
    if !ok {
      return 0, err 
    }

    id := int(idClaims)
    return id, nil

  } else {
    return 0, err 
  }
}
