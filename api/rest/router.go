package rest

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (h RestHandler) InitHandlers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", h.GetAllUserHandler).Methods(http.MethodGet)
	router.HandleFunc("/user/{nama}", h.GetUserHandler).Methods(http.MethodGet)
	router.HandleFunc("/user/", h.GetUserCurlHandler).Methods(http.MethodGet).Queries("nama", "{nama}")
	router.HandleFunc("/user/create_user", h.CreateUserHandler).Methods(http.MethodPost)
	return router
}
