package service

import (
	"backend/application/dto"
	"backend/application/queryservice"
	"backend/domain/entity"
	"backend/domain/repository"
	"errors"
)

type Auth interface {
	Login(*dto.Auth) (string, error)
	Logout(sessionId string) error
	GetSessionId(sessionId string) (string, error)
}

type auth struct {
	authRepository   repository.Auth
	authQueryService queryservice.Auth
}

func NewAuth(
	authRepository repository.Auth,
	authQueryService queryservice.Auth,
) Auth {
	return &auth{
		authRepository:   authRepository,
		authQueryService: authQueryService,
	}
}

func (a *auth) Login(dtoAuth *dto.Auth) (string, error) {
	auth := entity.NewAuth(dtoAuth.UserID, dtoAuth.Password)

	storedPassword, err := a.authQueryService.GetPasswordByUserID(
		auth.UserID,
	)
	if err != nil {
		return "", err
	}

	if storedPassword != auth.Password.Hash() {
		return "", errors.New("password is uncorrect")
	}

	err = a.authRepository.Set(auth)
	return auth.SessionID.Value(), err
}

func (a *auth) Logout(sessionId string) error {
	return a.authRepository.Delete(sessionId)
}

func (a *auth) GetSessionId(sessionId string) (string, error) {
	return a.authRepository.Get(sessionId)
}
