package dto

type User struct {
	ID           string `json:"user_id"`
	Name         string `json:"username"`
	Password     string `json:"password"`
	MailAddress  string `json:"mail_address"`
	Introduction string `json:"introduction"`
}
