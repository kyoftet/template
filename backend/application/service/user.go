package service

import (
	"backend/application/dto"
	"backend/application/dto/mapper"
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
	user := mapper.ToUserEntity(dtoUser)
	// TODO: 同じid、メールアドレスの場合はエラー
	return u.userRepository.Insert(user)
}

func (u *user) Delete(userId string) error {
	return u.userRepository.Delete(userId)
}

func (u *user) Update(dtoUser *dto.User) error {
	// TODO: 同じid, メールアドレスの場合はエラー
	user := mapper.ToUserEntity(dtoUser)
	return u.userRepository.Update(user)
}
