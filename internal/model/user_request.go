package model

type UserRequest struct {
	UserId   uint   `json:"userId"`
	Username string `json:"username"`
}
