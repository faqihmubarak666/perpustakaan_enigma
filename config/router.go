package config

import (
	"fmt"
	"gomux/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}

func RunServer(router *mux.Router) {
	host := utils.ReadEnv("serverHost", "")
	port := utils.ReadEnv("serverPort", "")
	fmt.Printf("Setting Web Server at host : %v & port : %v ", host, port)
	hostPort := fmt.Sprintf("%s:%s", host, port)
	err := http.ListenAndServe(hostPort, router)
	if err != nil {
		log.Fatal(err)
	}
}
