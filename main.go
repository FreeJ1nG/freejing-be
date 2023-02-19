package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/FreeJ1nG/freejing-be/blog"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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

	log.Fatal(http.ListenAndServe(":7070", cors.AllowAll().Handler(router)))
}

func main() {
	err := godotenv.Load(".env.local")
	if err != nil {
		panic(err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("SELECT * FROM blog")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	fmt.Println("Portofolio App v0.01")
	setupRoutes(db)
}
