package core

import "context"

type IUserServicePC interface {
	CreateUser(context.Context, string) (string, error)
	UpdateLastAccessed(context.Context, string) error
}
