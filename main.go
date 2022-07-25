package main

import (
	"transaction/db"
	"transaction/routes"

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
