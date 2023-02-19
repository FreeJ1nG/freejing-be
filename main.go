package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/FreeJ1nG/ristek-oprec/blog"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func setupRoutes(db *sql.DB) {
	router := mux.NewRouter()

	mainRouter := router.PathPrefix("/v1").Subrouter()

	mainRouter.HandleFunc("/blogs", blog.CreatePostHandler(db)).Methods("POST")
	mainRouter.HandleFunc("/blogs", blog.GetPostsHandler(db)).Methods("GET")
	mainRouter.HandleFunc("/blogs/{id}", blog.GetPostByIdHandler(db)).Methods("GET")
	mainRouter.HandleFunc("/blogs/{id}", blog.DeletePostHandler(db)).Methods("DELETE")
	mainRouter.HandleFunc("/blogs/{id}", blog.UpdatePostHandler(db)).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8081", cors.AllowAll().Handler(router)))
}

func main() {
	connStr := "postgres://freejing:@localhost:5432/portofolio?sslmode=require"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Portofolio App v0.01")
	setupRoutes(db)
}
