package entity

import "time"

type User struct {
	Id        int64
	Username  string
	Name      string
	Surname   string
	Password  string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
