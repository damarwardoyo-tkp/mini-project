package api

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"mini-project/api/rest"
	"mini-project/graph"
	"mini-project/graph/generated"
	"net/http"
)

type Handler struct {
	restHandler *rest.RestHandler
	gqlHandler  *graph.Resolver
}

func NewHandler(restHandler *rest.RestHandler, gqlHandler *graph.Resolver) *Handler {
	handler := Handler{
		restHandler: restHandler,
		gqlHandler:  gqlHandler,
	}
	return &handler
}

func (h Handler) InitHandlers() *mux.Router {
	router := mux.NewRouter()
	gql := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: h.gqlHandler}))
	router.Handle("/gql", playground.Handler("GraphQL playground", "/gql/query"))
	router.Handle("/gql/query", gql)
	router.HandleFunc("/", h.restHandler.GetAllUserHandler).Methods(http.MethodGet)
	router.HandleFunc("/user/{nama}", h.restHandler.GetUserHandler).Methods(http.MethodGet)
	router.HandleFunc("/user/", h.restHandler.GetUserHandler).Methods(http.MethodGet).Queries("nama", "{nama}")
	router.HandleFunc("/user/create_user", h.restHandler.CreateUserHandler).Methods(http.MethodPost)
	return router
}
