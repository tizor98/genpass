package service

import (
	"github.com/tizor98/genpass/entity"
)

func IsAuth() (*entity.User, bool) {
	user := GetActive()
	if user.Id != 0 {
		return &user, true
	}

	return nil, false
}
