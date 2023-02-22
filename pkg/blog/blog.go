package blog

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/FreeJ1nG.com/freejing-be/httpm"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type RequestBody struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Response struct {
	Data       interface{} `json:"data,omitempty"`
	StatusCode int         `json:"statusCode"`
	Success    bool        `json:"success"`
	Error      string      `json:"errors,omitempty"`
}

type Post struct {
	Id         string `json:"id"`
	CreateDate string `json:"create_date"`
	Title      string `json:"title"`
	Content    string `json:"content"`
}

func getUuid() string {
	id := uuid.New()
	return id.String()
}

func getPostById(db *sql.DB, id string) (Post, error) {
	var post Post

	row := db.QueryRow("SELECT * FROM blog WHERE id = $1", id)
	if err := row.Scan(&post.Id, &post.Title, &post.Content, &post.CreateDate); err != nil {
		if err == sql.ErrNoRows {
			return post, fmt.Errorf("no such post")
		}
		return post, err
	}

	return post, nil
}

func getAllPosts(db *sql.DB) ([]Post, error) {
	var posts []Post

	rows, err := db.Query("SELECT * FROM blog")
	if err != nil {
		return posts, err
	}

	defer rows.Close()
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.CreateDate); err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return posts, err
	}

	return posts, nil
}

func deletePostById(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM blog WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func createPost(db *sql.DB, title string, content string) (Post, error) {
	var post Post

	id := getUuid()
	createDate := time.Now().String()

	_, err := db.Exec("INSERT INTO blog (id, title, content, create_date) VALUES ($1, $2, $3, $4)", id, title, content, createDate)
	if err != nil {
		return post, err
	}

	row := db.QueryRow("SELECT * FROM blog WHERE id = $1", id)
	if err := row.Scan(&post.Id, &post.Title, &post.Content, &post.CreateDate); err != nil {
		return post, err
	}

	return post, nil
}

func updatePost(db *sql.DB, id string, title string, content string) (Post, error) {
	var post Post

	stmt, err := db.Prepare("UPDATE blog SET title = $1, content = $2 WHERE id = $3")
	if err != nil {
		return post, err
	}

	defer stmt.Close()
	_, err = stmt.Exec(title, content, id)
	if err != nil {
		return post, err
	}

	row := db.QueryRow("SELECT * FROM blog WHERE id = $1", id)
	if err := row.Scan(&post.Id, &post.Title, &post.Content, &post.CreateDate); err != nil {
		return post, err
	}

	return post, nil
}

func CreatePostHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody RequestBody
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			w.Write(httpm.MakeErrorResponse(w, http.StatusInternalServerError, err))
			return
		}

		if requestBody.Title == "" {
			w.Write(httpm.MakeErrorResponse(w, http.StatusBadRequest, fmt.Errorf("title is missing from request body")))
			return
		}
		if requestBody.Content == "" {
			w.Write(httpm.MakeErrorResponse(w, http.StatusBadRequest, fmt.Errorf("content is missing from request body")))
			return
		}

		post, err := createPost(db, requestBody.Title, requestBody.Content)
		if err != nil {
			w.Write(httpm.MakeErrorResponse(w, http.StatusInternalServerError, err))
			return
		}

		w.Write(httpm.MakeSuccessResponse[Post](w, http.StatusCreated, post))
	}
}

func DeletePostHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if id == "" {
			w.Write(httpm.MakeErrorResponse(w, http.StatusBadRequest, fmt.Errorf("no post id present in url parameter")))
			return
		}

		err := deletePostById(db, id)
		if err != nil {
			w.Write(httpm.MakeErrorResponse(w, http.StatusInternalServerError, err))
			return
		}

		w.Write(httpm.MakeSuccessResponse[interface{}](w, http.StatusNoContent, nil))
	}
}

func UpdatePostHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if id == "" {
			w.Write(httpm.MakeErrorResponse(w, http.StatusBadRequest, fmt.Errorf("no post id present in url parameter")))
			return
		}

		var requestBody RequestBody
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			w.Write(httpm.MakeErrorResponse(w, http.StatusBadRequest, err))
			return
		}

		if requestBody.Title == "" {
			w.Write(httpm.MakeErrorResponse(w, http.StatusBadRequest, fmt.Errorf("title is missing from request body")))
			return
		}
		if requestBody.Content == "" {
			w.Write(httpm.MakeErrorResponse(w, http.StatusBadRequest, fmt.Errorf("content is missing from request body")))
			return
		}

		post, err := updatePost(db, id, requestBody.Title, requestBody.Content)
		if err != nil {
			w.Write(httpm.MakeErrorResponse(w, http.StatusBadRequest, err))
			return
		}

		w.Write(httpm.MakeSuccessResponse[Post](w, http.StatusOK, post))
	}
}

func GetPostByIdHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if id == "" {
			w.Write(httpm.MakeErrorResponse(w, http.StatusBadRequest, fmt.Errorf("no post id present in request")))
			return
		}

		post, err := getPostById(db, id)
		if err != nil {
			w.Write(httpm.MakeErrorResponse(w, http.StatusInternalServerError, err))
			return
		}

		w.Write(httpm.MakeSuccessResponse[Post](w, http.StatusOK, post))
	}
}

func GetPostsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		posts, err := getAllPosts(db)
		if err != nil {
			w.Write(httpm.MakeErrorResponse(w, http.StatusInternalServerError, err))
			return
		}

		w.Write(httpm.MakeSuccessResponse[Post](w, http.StatusOK, posts))
	}
}
