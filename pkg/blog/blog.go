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

// CreatePost godoc
// @Summary Create new blog post
// @Tags blog
// @Accept json
// @Produce json
// @Param post body dbquery.Blog true "Create Blog Post"
// @Success 200 {object} dbquery.Blog
// @Router /v1/blogs [post]
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

		w.Write(httpm.MakeSuccessResponse[dbquery.Blog](w, http.StatusCreated, post))
	}
}

// DeletePost godoc
// @Summary Delete blog post with a certain id
// @Tags blog
// @Success 204
// @Router /v1/blogs/{id} [delete]
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

// UpdatePost godoc
// @Summary Update blog post with a certain id
// @Tags blog
// @Accept json
// @Produce json
// @Param post body dbquery.Blog true "Update Blog Post"
// @Success 200 {object} dbquery.Blog
// @Router /v1/blogs/{id} [put]
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

		w.Write(httpm.MakeSuccessResponse[dbquery.Blog](w, http.StatusOK, post))
	}
}

// GetPostById godoc
// @Summary Get blog post with a certain id
// @Tags blog
// @Produce json
// @Success 200 {object} httpm.Response{data=dbquery.Blog}
// @Router /v1/blogs/{id} [get]
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

		w.Write(httpm.MakeSuccessResponse[dbquery.Blog](w, http.StatusOK, post))
	}
}

// GetPosts godoc
// @Summary Get blog posts
// @Tags blog
// @Produce json
// @Success 200 {array} dbquery.Blog
// @Router /v1/blogs [get]
func GetPostsHandler(queries *dbquery.Queries, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		posts, err := queries.GetBlogs(ctx)
		if err != nil {
			w.Write(httpm.MakeErrorResponse(w, http.StatusInternalServerError, err))
			return
		}

		w.Write(httpm.MakeSuccessResponse[dbquery.Blog](w, http.StatusOK, posts))
	}
}
