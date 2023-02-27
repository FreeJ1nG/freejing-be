package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/FreeJ1nG.com/freejing-be/auth"
	"github.com/FreeJ1nG.com/freejing-be/websocket"
	"github.com/FreeJ1nG/freejing-be/blog"
	"github.com/FreeJ1nG/freejing-be/dbquery"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func logRequestFunc(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.URL.Path)
		f(w, r)
	}
}

func setupRoutes(db *sql.DB) {
	ctx := context.Background()
	pool := websocket.NewPool()
	router := mux.NewRouter()
	queries := dbquery.New(db)

	go pool.Start(queries, ctx)

	mainRouter := router.PathPrefix("/v1").Subrouter()

	mainRouter.HandleFunc("/blogs", logRequestFunc(blog.CreatePostHandler(queries, ctx))).Methods("POST")
	mainRouter.HandleFunc("/blogs", logRequestFunc(blog.GetPostsHandler(queries, ctx))).Methods("GET")
	mainRouter.HandleFunc("/blogs/{id}", logRequestFunc(blog.GetPostByIdHandler(queries, ctx))).Methods("GET")
	mainRouter.HandleFunc("/blogs/{id}", logRequestFunc(blog.DeletePostHandler(queries, ctx))).Methods("DELETE")
	mainRouter.HandleFunc("/blogs/{id}", logRequestFunc(blog.UpdatePostHandler(queries, ctx))).Methods("PATCH")

	mainRouter.HandleFunc("/auth/{username}", logRequestFunc(auth.GetUserHandler(queries, ctx))).Methods("GET")
	mainRouter.HandleFunc("/auth", logRequestFunc(auth.CreateUserHandler(queries, ctx))).Methods("POST")
	mainRouter.HandleFunc("/auth/{username}", logRequestFunc(auth.UpdateUserHandler(queries, ctx))).Methods("PATCH")
	mainRouter.HandleFunc("/auth/{username}", logRequestFunc(auth.DeleteUserHandler(queries, ctx))).Methods("DELETE")

	mainRouter.HandleFunc("/ws", logRequestFunc(websocket.WebsocketHandler(pool)))

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
