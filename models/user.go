package models

type User struct {
	Id       int    `json:"id"`
	PhotoUrl string `json:"photoUrl"`
	UserName string `json:"userName"`
	FullName string `json:"fullName"`
	Password string `json:"password"`
}
