package valueobject

import (
	"crypto/sha256"
	"encoding/hex"
)

type Password struct {
	value string
	hash  string
}

func NewPassword(password string) *Password {
	hash := sha256.Sum256([]byte(password))
	return &Password{
		value: password,
		hash:  hex.EncodeToString(hash[:]),
	}
}

func (p *Password) Value() string {
	return p.value
}

func (p *Password) Hash() string {
	return p.hash
}
