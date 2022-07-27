package main

import (
	"transaction/db"
	seed "transaction/db/Seed"
	"transaction/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	db.LoadEnv()
}

func main() {
	db.ConnectDatabase()
	seed.Handler(db.DB)

	router := gin.Default()
	routes.Avaiable(router)
	router.Run("localhost:8080")
}
