package main

import (
	"context"
	"flag"
	"log"
	"mini-project/api"
	"mini-project/api/rest"
	"mini-project/graph"
	"mini-project/infra/db"
	"mini-project/infra/es"
	"mini-project/infra/redis"
	"mini-project/module/user"
	"mini-project/module/user/repo"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	redisClient := redis.NewRedisClient()
	yugaByteClient := db.NewYugabyteClient()
	esClient := es.NewElasticsearchClient()

	userRepo := repo.NewUserDBRepo(redisClient, yugaByteClient, esClient)
	userManager := user.NewUserManager(userRepo)

	restHandler := rest.NewRestHandler(userManager)
	gqlHander := graph.NewGQLHandler(userManager)
	handler := api.NewHandler(restHandler, gqlHander)

	router := handler.InitHandlers()

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	srv := &http.Server{
		Addr: "0.0.0.0:8181",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println("listen and serve")
			log.Println(err)
		}
	}()

	log.Println("API is running!")
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
