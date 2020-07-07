package main

import (
	"gomux/config"
	"gomux/main/master"
)

func main() {
	db, _ := config.Connection()
	router := config.CreateRouter()
	master.Init(router, db)
	config.RunServer(router)
}
