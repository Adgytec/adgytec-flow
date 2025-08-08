package user

import "github.com/Adgytec/adgytec-flow/utils/core"

type userServicePC struct {
	service *userService
}

func (b *userServicePC) CreateUser(username string) error {
	return nil
}

func (b *userServicePC) GetUser(username string) (any, error) {
	return nil, nil
}

func (b *userServicePC) UpdateLastAccessed(username string) error {
	return nil
}

func CreateUserServicePC(params iUserServiceParams) core.IUserServicePC {
	return &userServicePC{
		service: &userService{
			db:   params.Database(),
			auth: params.Auth(),
		},
	}
}
