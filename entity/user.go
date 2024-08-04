package entity

import "time"

type User struct {
    Id        uint64
    Name      string
    Surname   string
    Username  string
    Password  string
    CreatedAt time.Time
    UpdatedAt time.Time
}
