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

type AuthorHandler struct {
	authorUseCase usecases.AuthorUsecase
}

func AuthorController(r *mux.Router, service usecases.AuthorUsecase) {
	authorHandler := AuthorHandler{service}

	r.Use(middleware.ActivityLogMiddleware)

	Author := r.PathPrefix("/author").Subrouter()
	// book.Use(middleware.TokenValidationMiddleware)
	Author.HandleFunc("", authorHandler.allAuthor).Methods(http.MethodGet)
}

func (s AuthorHandler) allAuthor(w http.ResponseWriter, r *http.Request) {
	author, err := s.authorUseCase.GetAllAuthor()
	if err != nil {
		w.Write([]byte("Data Not Found"))
		log.Print(err)
		return
	}
	var response response.MyResponse = response.MyResponse{"Successfully display Author data", http.StatusOK, author}
	byteOfBuku, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfBuku)
}
