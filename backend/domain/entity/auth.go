package entity

import (
	"backend/domain/valueobject"
	"time"
)

type Auth struct {
	UserID    string
	Password  *valueobject.Password
	SessionID *valueobject.SessionID
	DeadLine  time.Duration
}

func NewAuth(userId, password string) *Auth {
	return &Auth{
		UserID:    userId,
		Password:  valueobject.NewPassword(password),
		SessionID: valueobject.NewSessionID(userId),
		DeadLine:  24 * time.Hour,
	}
}
