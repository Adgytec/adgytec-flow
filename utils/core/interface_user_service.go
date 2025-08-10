package core

type IUserServicePC interface {
	CreateUser(string) (string, error)
	UpdateLastAccessed(string) error
}
