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
	host := utils.ReadEnv("host", "")
	port := utils.ReadEnv("port", "")
	hostPort := fmt.Sprintf("%s:%s", host, port)
	fmt.Println("Setting Web Server at port : " + hostPort)
	err := http.ListenAndServe(hostPort, router)
	if err != nil {
		log.Fatal(err)
	}
}
