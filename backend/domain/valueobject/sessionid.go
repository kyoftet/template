package valueobject

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type SessionID struct {
	value string
}

func NewSessionID(userId string) *SessionID {
	sessionId := fmt.Sprintf("%x", sha256.Sum256([]byte(userId+time.Now().String())))
	return &SessionID{value: sessionId}
}

func (s *SessionID) Value() string {
	return s.value
}
