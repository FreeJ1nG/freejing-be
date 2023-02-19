package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/FreeJ1nG/ristek-oprec/blog"
	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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

	log.Fatal(http.ListenAndServe(":8080", cors.AllowAll().Handler(router)))
}

func main() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "ristek_oprec",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println("Backend was not able to connect to database")
	}

	fmt.Println("Ristek Oprec App v0.01")
	setupRoutes(db)
}
