package main

import (
	"fmt"
	"os"
	config "transaction/configs"
	"transaction/db"
	seed "transaction/db/Seed"
	"transaction/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	config.Load()
}

func main() {
	seed.Handler(db.GetGormDB())

	router := gin.Default()
	routes.Avaiable(router)
	router.Run(fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")))
}
