package model

type UserResponse struct {
	Id       uint   `json:"id"`
	UserId   uint   `json:"userId"`
	Username string `json:"username"`
}
