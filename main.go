package main

import (
	"transaction/routes"
  "transaction/db"
  
	"github.com/gin-gonic/gin"
)

func init() {
	db.Load()
}

func main() {

	router := gin.Default()

	routes.Avaiable(router)

	router.Run("localhost:8080")

}
