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

func (uu *UserUseCase) CreateSession(session model.Session) (*model.Session, error) {
	_, err := uu.repository.CreateSession(session.Email, session.Password)
	if err != nil {
		return nil, err
	}

	return &session, nil
}
