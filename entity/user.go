package entity

import "time"

type User struct {
	Id        int64
	Username  string
	Password  string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
