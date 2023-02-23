package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/FreeJ1nG.com/freejing-be/websocket"
	"github.com/FreeJ1nG/freejing-be/blog"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func logRequestFunc(f http.HandlerFunc, log string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.URL.Path)
		f(w, r)
	}
}

func setupRoutes(db *sql.DB) {
	pool := websocket.NewPool()
	go pool.Start(db)

	router := mux.NewRouter()

	mainRouter := router.PathPrefix("/v1").Subrouter()

	mainRouter.HandleFunc("/blogs", logRequestFunc(blog.CreatePostHandler(db), "POST /v1/blogs")).Methods("POST")
	mainRouter.HandleFunc("/blogs", logRequestFunc(blog.GetPostsHandler(db), "GET /v1/blogs")).Methods("GET")
	mainRouter.HandleFunc("/blogs/{id}", logRequestFunc(blog.GetPostByIdHandler(db), "GET /v1/blogs/{id}")).Methods("GET")
	mainRouter.HandleFunc("/blogs/{id}", logRequestFunc(blog.DeletePostHandler(db), "DELETE /v1/blogs/{id}")).Methods("DELETE")
	mainRouter.HandleFunc("/blogs/{id}", logRequestFunc(blog.UpdatePostHandler(db), "PATCH /v1/blogs/{id}")).Methods("PATCH")

	mainRouter.HandleFunc("/ws", logRequestFunc(websocket.WebsocketHandler(db, pool), "WEBSOCKET /v1/ws"))

	server := &http.Server{
		Addr:    ":7070",
		Handler: cors.AllowAll().Handler(router),
	}

	defer server.Close()
	server.ListenAndServe()
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
		fmt.Println(err)
	}
	defer db.Close()

	fmt.Println("Portofolio App v0.02")
	setupRoutes(db)
}
