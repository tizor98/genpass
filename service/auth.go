package service

import (
	"context"
	"github.com/tizor98/genpass/repository"
)

func IsAuth() (bool, bool) {
	useRepo := repository.UserRepository(context.Background())
	useRepo.GetUserByUsername("hola")
	return false, false
}
