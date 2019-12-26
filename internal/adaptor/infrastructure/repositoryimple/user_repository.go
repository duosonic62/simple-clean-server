package repositoryimple

import (
	"github.com/duosonic/simple-clean-server/internal/domain/entity"
	"github.com/duosonic/simple-clean-server/internal/domain/repository"
)

type UserRepositoryImpl struct {

}

// コンストラクタ
func NewUserRpository() repository.UserRepository  {
	return UserRepositoryImpl{}
}

// とりあえずモック
func (repository UserRepositoryImpl) Save(user entity.User)  {
	return
}