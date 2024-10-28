package usecase

import (
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type UserUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return UserUseCase{
		repository: repo,
	}
}

func (uu *UserUseCase) CreateUser(user model.User) (model.User, error) {

	_, err := uu.repository.CreateUser(user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
