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

type CategoryHandler struct {
	categoryUseCase usecases.CategoryUsecase
}

func CategoryController(r *mux.Router, service usecases.CategoryUsecase) {
	categoryHandler := CategoryHandler{service}

	r.Use(middleware.ActivityLogMiddleware)

	category := r.PathPrefix("/category").Subrouter()
	// book.Use(middleware.TokenValidationMiddleware)
	category.HandleFunc("", categoryHandler.AllCategory).Methods(http.MethodGet)
}

func (s CategoryHandler) AllCategory(w http.ResponseWriter, r *http.Request) {
	category, err := s.categoryUseCase.GetAllCategory()
	if err != nil {
		w.Write([]byte("Data Not Found"))
		log.Print(err)
		return
	}
	var response response.MyResponse = response.MyResponse{"Successfully display category data", http.StatusOK, category}
	byteOfBuku, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfBuku)
}
