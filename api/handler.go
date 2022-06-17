package api

import (
	"github.com/gorilla/mux"
	"mini-project/api/gql"
	"mini-project/api/rest"
	"net/http"
)

type Handler struct {
	restHandler *rest.RestHandler
	gqlHandler  *gql.GqlHandler
}

func NewHandler(restHandler *rest.RestHandler, gqlHandler *gql.GqlHandler) *Handler {
	handler := Handler{
		restHandler: restHandler,
		gqlHandler:  gqlHandler,
	}
	return &handler
}

func (h Handler) InitHandlers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", h.restHandler.GetAllUserHandler).Methods(http.MethodGet)
	router.HandleFunc("/user/{nama}", h.restHandler.GetUserHandler).Methods(http.MethodGet)
	router.HandleFunc("/user/", h.restHandler.GetUserHandler).Methods(http.MethodGet).Queries("nama", "{nama}")
	router.HandleFunc("/user/create_user", h.restHandler.CreateUserHandler).Methods(http.MethodPost)
	return router
}
