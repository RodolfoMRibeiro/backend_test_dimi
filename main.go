package main

import (
	"fmt"
	"os"
	"transaction/db"
	seed "transaction/db/Seed"
	"transaction/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	db.Load()
	seed.Handler(db.DB)
}

func main() {
	router := gin.Default()
	routes.Avaiable(router)
	router.Run(fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")))
}
