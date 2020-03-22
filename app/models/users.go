package models

type Users struct {
	Model
	Username	string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Group       string `json:"group"`
}
