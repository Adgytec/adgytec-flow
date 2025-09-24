package media

import (
	"context"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/google/uuid"
)

func (s *mediaService) completeMediaItemsUpload(ctx context.Context, mediaIDs []uuid.UUID) error {
	qtx, tx, err := s.database.WithTransaction(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	mediaStatusUpdateErr := qtx.Queries().UpdateMediaItemsStatus(
		ctx,
		db.UpdateMediaItemsStatusParams{
			Status:   db.GlobalMediaStatusProcessing,
			MediaIds: mediaIDs,
		},
	)
	if mediaStatusUpdateErr != nil {
		return mediaStatusUpdateErr
	}

	_, mediaOutboxInsertErr := qtx.Queries().AddMediaItemsToOutbox(ctx, mediaIDs)
	if mediaOutboxInsertErr != nil {
		return mediaOutboxInsertErr
	}

	return tx.Commit(ctx)
}

func (pc *mediaServicePC) CompleteMediaItemsUpload(ctx context.Context, mediaIDs []uuid.UUID) error {
	return pc.service.completeMediaItemsUpload(ctx, mediaIDs)
}

func (pc *mediaServicePC) CompleteMediaItemUpload(ctx context.Context, mediaID uuid.UUID) error {
	return pc.service.completeMediaItemsUpload(ctx, []uuid.UUID{mediaID})
}
