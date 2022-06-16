package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"mini-project/entity"
	"net/http"
)

func (h RestHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nama := vars["nama"]
	log.Print(nama)
	fmt.Println(nama)
	h.manager.GetUser()
}

func (h RestHandler) GetUserCurlHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nama := vars["nama"]
	fmt.Println(nama)
	h.manager.GetUser()
}

func (h RestHandler) GetAllUserHandler(w http.ResponseWriter, r *http.Request) {
	h.manager.GetUserList()

}

func (h RestHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req entity.User
	err := ParseBody(r.Body, &req)
	if err != nil {
		log.Fatalln(err)
	}

	if err := h.manager.CreateUser(r.Context(), req); err != nil {
		log.Println("[CreateUserHandler]Gagal membuat user baru")
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error")
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
