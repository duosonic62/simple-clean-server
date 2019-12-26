package factoryimple

import (
	"github.com/google/uuid"
)

type IDFactoryImpl struct {
}

func (factory IDFactoryImpl) Gen() (string, error)  {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}