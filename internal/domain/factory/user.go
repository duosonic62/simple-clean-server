package factory

import "github.com/duosonic/simple-clean-server/internal/domain/entity"

type UserFactory struct {
	idFactory IDFactory
}

// コンストラクタ
func NewUserFactory(factory IDFactory) UserFactory {
	return UserFactory{
		idFactory: factory,
	}
}

func (factory UserFactory) NewUser(name string, role int) (entity.User, error) {
	id, err := factory.idFactory.Gen()
	if err != nil {
		return entity.User{}, err
	}
	user := entity.User{
		ID:   id,
		Name: name,
		Role: role,
	}

	return user, nil
}
