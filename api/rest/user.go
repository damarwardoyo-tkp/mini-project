package rest

import (
	"github.com/gorilla/mux"
	"mini-project/entity"
	"net/http"
)

func (h RestHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nama := vars["nama"]

	resp, err := h.manager.GetUser(nama)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func (h RestHandler) GetAllUserHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := h.manager.GetUserList()
	if err != nil || resp == "[]" {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))

}

func (h RestHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req entity.User
	err := ParseBody(r.Body, &req)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.manager.CreateUser(req); err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
