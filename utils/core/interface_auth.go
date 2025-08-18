package core

type IAuth interface {
	CreateUser(string) error
	DisableUser(string) error
	EnableUser(string) error
}
