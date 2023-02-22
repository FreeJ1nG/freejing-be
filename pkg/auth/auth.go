package auth

type User struct {
	Id            string `json:"id"`
	Name          string `json:"username"`
	Email         string `json:"email"`
	PhotoUrl      string `json:"photoUrl"`
	EmailVerified bool   `json:"emailVerified"`
}
