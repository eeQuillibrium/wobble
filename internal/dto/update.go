package dto

type Update struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
