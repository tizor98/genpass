package service

import (
	"context"
	"github.com/tizor98/genpass/entity"
	"github.com/tizor98/genpass/repository"
)

func IsAuth() (*entity.User, bool) {
	ur := repository.UserRepository(context.Background())
	user := ur.GetActive()
	if user.Id != 0 {
		return &user, true
	}

	return nil, false
}
