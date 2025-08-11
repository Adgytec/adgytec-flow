package user

import (
	"log"

	"github.com/Adgytec/adgytec-flow/config/cache"
	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type userServicePC struct {
	service *userService
}

func (pc *userServicePC) CreateUser(username string) (string, error) {
	return "", nil
}

func (pc *userServicePC) UpdateLastAccessed(username string) error {
	return nil
}

func CreateUserServicePC(params iUserServiceParams) core.IUserServicePC {
	log.Println("creating user-service PC")
	return &userServicePC{
		service: &userService{
			db:                    params.Database(),
			auth:                  params.Auth(),
			accessManagement:      params.AccessManagement(),
			getUserCache:          cache.CreateNewCache[db_actions.GlobalUser](params.CacheClient(), "user"),
			getUserListCache:      cache.CreateNewCache[[]db_actions.GlobalUser](params.CacheClient(), "user-list"),
			userLastAccessedCache: cache.CreateNewCache[bool](params.CacheClient(), "user-last-accessed"),
		},
	}
}
