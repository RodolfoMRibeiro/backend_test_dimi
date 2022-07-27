package main

import (
	"transaction/db"
	seed "transaction/db/Seed"
	"transaction/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	db.Load()
	seed.Handler(db.GetGormDB())
}

func main() {
	router := gin.Default()
	routes.Avaiable(router)
	router.Run("localhost:8080")
}
