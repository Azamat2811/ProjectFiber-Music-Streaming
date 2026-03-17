package models

type Users struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	City     string `json:"city"`
}
