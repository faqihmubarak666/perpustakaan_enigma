package controllers

import (
	"encoding/json"
	"gomux/main/master/models"
	"gomux/main/master/response"
	"gomux/main/master/usecases"
	"gomux/utils"
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	userUseCase usecases.UserUsecase
}

func UserController(r *mux.Router, service usecases.UserUsecase) {
	userHandler := UserHandler{userUseCase: service}

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("", userHandler.getUser).Methods(http.MethodPost)
	auth.HandleFunc("", userHandler.allUser).Methods(http.MethodGet)
}

func (u *UserHandler) getUser(w http.ResponseWriter, r *http.Request) {
	var data models.User
	_ = json.NewDecoder(r.Body).Decode(&data)
	isValid, _ := u.userUseCase.GetUser(&data)

	if isValid {
		w.WriteHeader(http.StatusOK)
		token, err := utils.JwtEncoder(data.Username, "rahasiadong")
		if err != nil {
			http.Error(w, "Failed token generation", http.StatusUnauthorized)
		} else {
			var response response.MyResponse = response.MyResponse{"Successfully Login", http.StatusOK, token}
			byteOfUser, err := json.Marshal(response)
			if err != nil {
				w.Write([]byte("Oops something when wrong"))
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(byteOfUser)
		}
	} else {
		http.Error(w, "Invalid login", http.StatusUnauthorized)
	}
}

func (s UserHandler) allUser(w http.ResponseWriter, r *http.Request) {
	user, err := s.userUseCase.GetAllUser()
	if err != nil {
		w.Write([]byte("Data Not Found"))
		log.Print(err)
		return
	}
	var response response.MyResponse = response.MyResponse{"Successfully display user data", http.StatusOK, user}
	byteOfUser, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfUser)
}
