package core

type IAuth interface {
	CreateUser(string) (string, error)
	DisableUser(string) error
	EnableUser(string) error
}
