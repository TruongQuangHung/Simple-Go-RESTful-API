package model

type AddUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
