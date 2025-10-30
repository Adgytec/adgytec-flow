package usermanagement

import (
	"context"
	"time"

	"github.com/Adgytec/adgytec-flow/database/db"
)

func (s *userManagementService) getManagementUsersQuery(
	ctx context.Context,
	searchQuery string,
	limit int32,
) ([]db.GlobalUserDetails, error) {
	return s.db.Queries().GetManagementUsersByQuery(
		ctx,
		db.GetManagementUsersByQueryParams{
			Limit: limit,
			Query: searchQuery,
		},
	)
}

func (s *userManagementService) getManagementUsersInitialLatestFirst(
	ctx context.Context,
	limit int32,
) ([]db.GlobalUserDetails, error) {
	return s.db.Queries().GetManagementUsersLatestFirst(
		ctx,
		limit,
	)
}

func (s *userManagementService) getManagementUsersInitialOldestFirst(
	ctx context.Context,
	limit int32,
) ([]db.GlobalUserDetails, error) {
	return s.db.Queries().GetManagementUsersOldestFirst(
		ctx,
		limit,
	)
}

func (s *userManagementService) getManagementUsersGreaterThanCursorLatestFirst(
	ctx context.Context,
	cursor time.Time,
	limit int32,
) ([]db.GlobalUserDetails, error) {
	return s.db.Queries().GetManagementUsersLatestFirstGreaterThanCursor(
		ctx,
		db.GetManagementUsersLatestFirstGreaterThanCursorParams{
			Cursor: cursor,
			Limit:  limit,
		},
	)
}

func (s *userManagementService) getManagementUsersGreaterThanCursorOldestFirst(
	ctx context.Context,
	cursor time.Time,
	limit int32,
) ([]db.GlobalUserDetails, error) {
	return s.db.Queries().GetManagementUsersOldestFirstGreaterThanCursor(
		ctx,
		db.GetManagementUsersOldestFirstGreaterThanCursorParams{
			Cursor: cursor,
			Limit:  limit,
		},
	)
}

func (s *userManagementService) getManagementUsersLesserThanCursorLatestFirst(
	ctx context.Context,
	cursor time.Time,
	limit int32,
) ([]db.GlobalUserDetails, error) {
	return s.db.Queries().GetManagementUsersLatestFirstLesserThanCursor(
		ctx,
		db.GetManagementUsersLatestFirstLesserThanCursorParams{
			Cursor: cursor,
			Limit:  limit,
		},
	)
}

func (s *userManagementService) getManagementUsersLesserThanCursorOldestFirst(
	ctx context.Context,
	cursor time.Time,
	limit int32,
) ([]db.GlobalUserDetails, error) {
	return s.db.Queries().GetManagementUsersOldestFirstLesserThanCursor(
		ctx,
		db.GetManagementUsersOldestFirstLesserThanCursorParams{
			Cursor: cursor,
			Limit:  limit,
		},
	)
}
