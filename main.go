package main

import (
	"transaction/db"
	predefinedData "transaction/db/Data"
	"transaction/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	db.Load()
	predefinedData.Load(db.DB)
}

func main() {
	router := gin.Default()
	routes.Avaiable(router)
	router.Run("localhost:8080")
}
