package entity

import "backend/domain/valueobject"

type User struct {
	ID           string
	Name         string
	MailAddress  string
	Password     *valueobject.Password
	Introduction string
}

func NewUser(id, name, mailAddress, password, introduction string) *User {
	return &User{
		ID:           id,
		Name:         name,
		MailAddress:  mailAddress,
		Password:     valueobject.NewPassword(password),
		Introduction: introduction,
	}
}
