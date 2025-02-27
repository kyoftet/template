package service

import (
	"backend/application/dto"
	"backend/domain/entity"
	"backend/domain/repository"
)

type User interface {
	SignUp(*dto.User) error
	Delete(userId string) error
	Update(*dto.User) error
}

type user struct {
	userRepository repository.User
}

func NewUser(userRepository repository.User) User {
	return &user{userRepository: userRepository}
}

func (u *user) SignUp(dtoUser *dto.User) error {
	user := dtoTransferEntity(dtoUser)
	// TODO: 同じid、メールアドレスの場合はエラー
	return u.userRepository.Insert(user)
}

func (u *user) Delete(userId string) error {
	return u.userRepository.Delete(userId)
}

func (u *user) Update(dtoUser *dto.User) error {
	// TODO: 同じid, メールアドレスの場合はエラー
	user := dtoTransferEntity(dtoUser)
	return u.userRepository.Update(user)
}

func dtoTransferEntity(dtoUser *dto.User) *entity.User {
	return entity.NewUser(
		dtoUser.ID,
		dtoUser.Name,
		dtoUser.MailAddress,
		dtoUser.Password,
		dtoUser.Introduction,
	)
}
