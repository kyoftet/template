package repository

import "backend/domain/entity"

type Auth interface {
	Set(*entity.Auth) error
	Delete(string) error
	Get(string) (string, error)
}
