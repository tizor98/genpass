package entity

import "time"

type Password struct {
	Id        int64
	ForEntity string
	Password  string
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
