package service

import (
	"context"
	"github.com/tizor98/genpass/entity"
	"github.com/tizor98/genpass/repository"
	"log"
)

func SaveNewPassword(password, forEntity string, user *entity.User) {
	pr := repository.PasswordRepository(context.Background())

	passId, err := pr.Create(password, forEntity, user.Id)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Password created %d for entity %s", passId, forEntity)
}
