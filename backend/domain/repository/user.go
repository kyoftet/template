package repository

import "backend/domain/entity"

type User interface {
	Insert(*entity.User) error
	Delete(userId string) error
	Update(*entity.User) error
}
