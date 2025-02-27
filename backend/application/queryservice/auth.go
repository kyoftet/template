package queryservice

type Auth interface {
	GetPasswordByUserID(userId string) (string, error)
}
