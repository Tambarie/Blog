package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tambarie/Blog/models"
	"net/http"
)



func main()  {
	//Initializing the go-chi router
	r :=  chi.NewRouter()
	r.Use(middleware.Logger)


	// Setting up my routes
	r.Get("/",models.Blog)
	r.Post("/create", models.CreateArticle)

	//r.Route("/create", func(r chi.Router) {
	//
	//})



	//Setting up the server
	server := http.Server{
		Addr:              "localhost:3000",
		Handler:           r,

	}
	server.ListenAndServe()
}