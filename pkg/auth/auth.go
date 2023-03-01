package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FreeJ1nG.com/freejing-be/httpm"
	"github.com/FreeJ1nG/freejing-be/dbquery"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type newUserRequestBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Id           string `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

func GenerateUuid() string {
	id := uuid.New()
	return id.String()
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return string(hash), err
	}

	return string(hash), nil
}

func (u *User) ValidatePasswordHash(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

// Createuser godoc
// @Summary Create new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body newUserRequestBody true "Create new user"
// @Success 201 {object} httpm.Response{data=dbquery.User} "User has been created"
// @Router /v1/auth [post]
func CreateUserHandler(queries *dbquery.Queries, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody newUserRequestBody
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			w.Write(httpm.MakeErrorResponse(w, http.StatusInternalServerError, err))
			return
		}

		if requestBody.Username == "" {
			w.Write(httpm.MakeErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("username is missing from request body")))
			return
		}
		if requestBody.Email == "" {
			w.Write(httpm.MakeErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("email is missing from request body")))
			return
		}
		if requestBody.Password == "" {
			w.Write(httpm.MakeErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("password is missing from request body")))
			return
		}

		passwordHash, err := HashPassword(requestBody.Password)
		if err != nil {
			w.Write(httpm.MakeErrorResponse(w, http.StatusInternalServerError, err))
			return
		}

		user, err := queries.CreateUser(ctx, dbquery.CreateUserParams{
			Username: requestBody.Username, Email: requestBody.Email, PasswordHash: passwordHash,
		})

		if err != nil {
			w.Write(httpm.MakeErrorResponse(w, http.StatusInternalServerError, err))
			return
		}

		w.Write(httpm.MakeSuccessResponse[dbquery.User](w, http.StatusCreated, user))
	}
}

// GetUserByUsername godoc
// @Summary Get user with a certain username
// @Tags user
// @Produce json
// @Success 200 {object} httpm.Response{data=dbquery.User} "User has been found"
// @Router /v1/user/{username} [get]
func GetUserHandler(queries *dbquery.Queries, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		username := vars["username"]

		if username == "" {
			w.Write(httpm.MakeErrorResponse(w, http.StatusBadRequest, fmt.Errorf("no username present in request")))
			return
		}

		user, err := queries.GetUser(ctx, username)
		if err != nil {
			w.Write(httpm.MakeErrorResponse(w, http.StatusNotFound, err))
			return
		}

		w.Write(httpm.MakeSuccessResponse[dbquery.User](w, http.StatusOK, user))
	}
}

// DeletePost godoc
// @Summary Delete user with a certain username
// @Tags user
// @Success 204 "User has been deleted"
// @Router /v1/auth/{username} [delete]
func DeleteUserHandler(queries *dbquery.Queries, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		username := vars["username"]

		if username == "" {
			w.Write(httpm.MakeErrorResponse(w, http.StatusBadRequest, fmt.Errorf("no username present in request")))
			return
		}

		err := queries.DeleteUser(ctx, username)
		if err != nil {
			w.Write(httpm.MakeErrorResponse(w, http.StatusInternalServerError, err))
			return
		}

		w.Write(httpm.MakeSuccessResponse[interface{}](w, http.StatusNoContent, nil))
	}
}

// UpdateUser godoc
// @Summary Update user with a certain username
// @Tags user
// @Accept json
// @Produce json
// @Param user body newUserRequestBody true "Update user with a certain username"
// @Success 200 {object} httpm.Response{data=dbquery.User} "User has been updated"
// @Router /v1/auth/{username} [put]
func UpdateUserHandler(queries *dbquery.Queries, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody newUserRequestBody
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			w.Write(httpm.MakeErrorResponse(w, http.StatusInternalServerError, err))
			return
		}

		if requestBody.Username == "" {
			w.Write(httpm.MakeErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("username is missing from request body")))
			return
		}
		if requestBody.Email == "" {
			w.Write(httpm.MakeErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("email is missing from request body")))
			return
		}
		if requestBody.Password == "" {
			w.Write(httpm.MakeErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("password is missing from request body")))
			return
		}

		vars := mux.Vars(r)
		username := vars["username"]

		if username == "" {
			w.Write(httpm.MakeErrorResponse(w, http.StatusBadRequest, fmt.Errorf("no username present in request")))
			return
		}

		passwordHash, err := HashPassword(requestBody.Password)
		if err != nil {
			w.Write(httpm.MakeErrorResponse(w, http.StatusInternalServerError, err))
			return
		}

		user, err := queries.UpdateUser(ctx, dbquery.UpdateUserParams{
			Username: username, Email: requestBody.Email, PasswordHash: passwordHash, Username_2: requestBody.Username,
		})
		if err != nil {
			w.Write(httpm.MakeErrorResponse(w, http.StatusInternalServerError, err))
			return
		}

		w.Write(httpm.MakeSuccessResponse[dbquery.User](w, http.StatusOK, user))
	}
}
