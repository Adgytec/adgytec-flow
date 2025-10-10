package media

import (
	"context"
	"log"

	"github.com/Adgytec/adgytec-flow/config/database"
)

type MediaServiceActions interface {
	NewMediaItem(ctx context.Context, input NewMediaItemInfoWithStorageDetails) (*MediaUploadDetails, error)
	NewMediaItems(ctx context.Context, input []NewMediaItemInfoWithStorageDetails) ([]MediaUploadDetails, error)
}

type mediaServiceActions struct {
	service *mediaService
}

type MediaServicePC interface {
	WithTransaction(db database.Database) MediaServiceActions
}

type mediaServicePC struct {
	params mediaServiceParams
}

func (pc *mediaServicePC) WithTransaction(tx database.Database) MediaServiceActions {
	return &mediaServiceActions{
		service: newMediaServiceWithTx(pc.params, tx),
	}
}

func NewMediaServicePC(params mediaServiceParams) MediaServicePC {
	log.Printf("creating %s-service PC", serviceName)
	return &mediaServicePC{
		params: params,
	}
}
