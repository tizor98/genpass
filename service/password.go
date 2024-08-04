package service

import (
    "context"
    "fmt"
)

func NewPassword(ctx context.Context) string {
    if user := ctx.Value("user"); user != nil {

    }

    pass := "new password"
    fmt.Println(pass)
    return pass
}
