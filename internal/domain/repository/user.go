package repository

import "github.com/duosonic/simple-clean-server/internal/domain/entity"

type UserRepository interface {
	Save(entity.User)
}