package main

import (
	"transaction/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.Avaiable(router)
	router.Run("localhost:8080")
}
