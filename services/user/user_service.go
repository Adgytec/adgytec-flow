package user

import (
	"net/http"

	"github.com/Adgytec/adgytec-flow/config/auth"
	"github.com/Adgytec/adgytec-flow/config/cache"
	"github.com/Adgytec/adgytec-flow/config/cdn"
	"github.com/Adgytec/adgytec-flow/config/database"
	"github.com/Adgytec/adgytec-flow/config/serializer"
	"github.com/Adgytec/adgytec-flow/database/models"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/services/media"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/pagination"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type userServiceParams interface {
	Database() database.Database
	Auth() auth.Auth
	IAMService() iam.IAMServicePC
	CDN() cdn.CDN
	CacheClient() cache.CacheClient
	MediaWithTransaction() media.MediaServicePC
}

type userServiceMuxParams interface {
	userServiceParams
	Middleware() core.MiddlewarePC
}

type userService struct {
	db               database.Database
	auth             auth.Auth
	iam              iam.IAMServicePC
	cdn              cdn.CDN
	media            media.MediaServicePC
	getUserCache     cache.Cache[models.GlobalUser]
	getUserListCache cache.Cache[pagination.ResponsePagination[models.GlobalUser]]
}

func (s *userService) getSocialLinkIDFromRequest(r *http.Request) (uuid.UUID, error) {
	socialLinkID := chi.URLParam(r, "socialLinkID")
	socialLinkUUID, socialLinkIDErr := uuid.Parse(socialLinkID)
	if socialLinkIDErr != nil {
		return uuid.Nil, &InvalidSocialLinkIDError{
			InvalidSocialLinkID: socialLinkID,
		}
	}

	return socialLinkUUID, nil
}

func newUserService(params userServiceParams) *userService {
	return &userService{
		db:               params.Database(),
		auth:             params.Auth(),
		iam:              params.IAMService(),
		cdn:              params.CDN(),
		media:            params.MediaWithTransaction(),
		getUserCache:     cache.NewCache(params.CacheClient(), serializer.NewGobSerializer[models.GlobalUser](), "user"),
		getUserListCache: cache.NewCache(params.CacheClient(), serializer.NewGobSerializer[pagination.ResponsePagination[models.GlobalUser]](), "user-list"),
	}
}
