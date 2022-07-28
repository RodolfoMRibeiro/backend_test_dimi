package main

import (
	"fmt"
	config "transaction/configs"
	"transaction/db"
	seed "transaction/db/Seed"
	"transaction/routes"
	"transaction/server"
)

func main() {
	config.Load()
	db.StartDatabase()
	seed.Handler(db.GetGormDB())
	createdServer := server.CreateServer()
	routes.Avaiable(createdServer.GetServerEngine())
	createdServer.GetServerEngine().Run(fmt.Sprintf("%s:%s", config.Server.HOST, config.Server.PORT))
}
