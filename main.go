package main

import (
  "fmt"
	"go-jwt-api/routers"
	"go-jwt-api/utils/db"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)


func main() {

  err := godotenv.Load(".env")

  if err != nil {
    fmt.Print("Can't Load .env")
  }

  db.ConnectDatabase()
  r := routers.SetupRouter()
  r.Use(cors.Default())

  r.Run()
}
