package usecase

import (
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/repository"
)

type SessionUseCase struct {
	repository repository.UserRepository
}

func NewSessionUseCase(repo repository.UserRepository) SessionUseCase {
	return SessionUseCase{
		repository: repo,
	}
}

func (uu *SessionUseCase) CreateSession(email string, password string) (*model.User, error) {
	user, err := uu.repository.CreateSession(email, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
