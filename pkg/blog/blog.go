package blog

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/FreeJ1nG.com/freejing-be/httpm"
	"github.com/FreeJ1nG/freejing-be/dbquery"
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

func CreatePostHandler(queries *dbquery.Queries, ctx context.Context) http.HandlerFunc {
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

		post, err := queries.CreateBlog(ctx, dbquery.CreateBlogParams{
			Title:      requestBody.Title,
			Content:    requestBody.Content,
			CreateDate: time.Now(),
		})
		if err != nil {
			w.Write(httpm.MakeErrorResponse(w, http.StatusInternalServerError, err))
			return
		}

		w.Write(httpm.MakeSuccessResponse[Post](w, http.StatusCreated, post))
	}
}

func DeletePostHandler(queries *dbquery.Queries, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if id == "" {
			w.Write(httpm.MakeErrorResponse(w, http.StatusBadRequest, fmt.Errorf("no post id present in url parameter")))
			return
		}

		id_int, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			w.Write(httpm.MakeErrorResponse(w, http.StatusInternalServerError, err))
			return
		}

		err = queries.DeleteBlog(ctx, id_int)
		if err != nil {
			w.Write(httpm.MakeErrorResponse(w, http.StatusInternalServerError, err))
			return
		}

		w.Write(httpm.MakeSuccessResponse[interface{}](w, http.StatusNoContent, nil))
	}
}

func UpdatePostHandler(queries *dbquery.Queries, ctx context.Context) http.HandlerFunc {
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

		id_int, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			w.Write(httpm.MakeErrorResponse(w, http.StatusInternalServerError, err))
			return
		}

		post, err := queries.UpdateBlog(ctx, dbquery.UpdateBlogParams{
			ID: id_int, Title: requestBody.Title, Content: requestBody.Content,
		})
		if err != nil {
			w.Write(httpm.MakeErrorResponse(w, http.StatusBadRequest, err))
			return
		}

		w.Write(httpm.MakeSuccessResponse[Post](w, http.StatusOK, post))
	}
}

func GetPostByIdHandler(queries *dbquery.Queries, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if id == "" {
			w.Write(httpm.MakeErrorResponse(w, http.StatusBadRequest, fmt.Errorf("no post id present in request")))
			return
		}

		id_int, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			w.Write(httpm.MakeErrorResponse(w, http.StatusInternalServerError, err))
			return
		}

		post, err := queries.GetBlogById(ctx, id_int)
		if err != nil {
			w.Write(httpm.MakeErrorResponse(w, http.StatusInternalServerError, err))
			return
		}

		w.Write(httpm.MakeSuccessResponse[Post](w, http.StatusOK, post))
	}
}

func GetPostsHandler(queries *dbquery.Queries, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		posts, err := queries.GetBlogs(ctx)
		if err != nil {
			w.Write(httpm.MakeErrorResponse(w, http.StatusInternalServerError, err))
			return
		}

		w.Write(httpm.MakeSuccessResponse[Post](w, http.StatusOK, posts))
	}
}
