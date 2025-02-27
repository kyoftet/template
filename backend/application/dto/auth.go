package dto

type Auth struct {
	UserID      string `json:"user_id"`
	MailAddress string `json:"mail_address"`
	Password    string `json:"password"`
}
