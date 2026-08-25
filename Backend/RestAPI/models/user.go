package models

type User struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	CreatedAt string `json:"-"`
	Privilege int    `json:"privilege"`
}
