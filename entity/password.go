package entity

import "time"

type Password struct {
    Username string
    Password string
    For      string
    createAt time.Time
    UpdateAt time.Time
}
