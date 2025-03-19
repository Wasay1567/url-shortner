package main

import (
	"log"
	"net/http"

	"github.com/Wasay1567/url-shortner-golang/controllers"
)

type Api struct {
	addr string
}

func main() {
	a := Api{
		addr: ":8080",
	}

	router := http.NewServeMux()

	router.HandleFunc("POST /shorten", controllers.Shorten)
	router.HandleFunc("GET /redirect/{code}", controllers.Redirect)

	srv := http.Server{
		Addr:    a.addr,
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Error starting the server")
	}
}
