package controllers

import (
	"encoding/json"
	"gomux/main/master/middleware"
	"gomux/main/master/response"
	"gomux/main/master/usecases"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type PublisherHandler struct {
	publisherUseCase usecases.PublisherUsecase
}

func PublisherController(r *mux.Router, service usecases.PublisherUsecase) {
	publisherHandler := PublisherHandler{service}

	r.Use(middleware.ActivityLogMiddleware)

	publisher := r.PathPrefix("/publisher").Subrouter()
	// book.Use(middleware.TokenValidationMiddleware)
	publisher.HandleFunc("", publisherHandler.allPublisher).Methods(http.MethodGet)
}

func (s PublisherHandler) allPublisher(w http.ResponseWriter, r *http.Request) {
	publisher, err := s.publisherUseCase.GetAllPublisher()
	if err != nil {
		w.Write([]byte("Data Not Found"))
		log.Print(err)
		return
	}
	var response response.MyResponse = response.MyResponse{"Successfully display publisher data", http.StatusOK, publisher}
	byteOfBuku, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfBuku)
}
