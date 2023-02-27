package auth

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

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

func getUserObject(db *sql.DB, username string) (User, error) {
	var user User
	row := db.QueryRow("SELECT * FROM users WHERE username = $1", username)
	if err := row.Scan(&user.Id, &user.Username, &user.Email, &user.PasswordHash); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("no such user")
		}
		return user, err
	}
	return user, nil
}

// func addUserObject(db *sql.DB, email string, username string, password string) (User, error) {
// 	var user User
// 	user, err := getUserObject(db, username)
// 	if err != nil {
// 		return user, err
// 	}

// 	passwordHash, err := HashPassword(password)
// 	if err != nil {
// 		return user, err
// 	}

// 	newUser := User{
// 		Id:           GenerateUuid(),
// 		Username:     username,
// 		Email:        email,
// 		PasswordHash: passwordHash,
// 	}
// }
