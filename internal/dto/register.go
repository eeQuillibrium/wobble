package dto

type Register struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}
