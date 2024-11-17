// model/auth.go
package model

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Bu strukturani qo'shish kerak
type LoginResponse struct {
	Token string `json:"token"`
}
