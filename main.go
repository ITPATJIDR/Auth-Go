package main

import (
	"go-jwt-api/routers"
	"go-jwt-api/utils/db"

	"github.com/gin-contrib/cors"
)


func main() {

  db.ConnectDatabase()
  r := routers.SetupRouter()
  r.Use(cors.Default())

  r.Run()
}
