package core

type IUserServicePC interface {
	CreateUser(string) error
	GetUser(string) (any, error)
	UpdateLastAccessed(string) error
}
