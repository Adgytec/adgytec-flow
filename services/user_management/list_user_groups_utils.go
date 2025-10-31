package usermanagement

import (
	"context"
	"time"

	"github.com/Adgytec/adgytec-flow/database/db"
)

func (s *userManagementService) getUserGroupsQuery(
	ctx context.Context,
	searchQuery string,
	limit int32,
) ([]db.ManagementUserGroupDetails, error) {
	return s.db.Queries().GetUserGroupsByQuery(
		ctx,
		db.GetUserGroupsByQueryParams{
			Limit: limit,
			Query: searchQuery,
		},
	)
}

func (s *userManagementService) getUserGroupsInitialLatestFirst(
	ctx context.Context,
	limit int32,
) ([]db.ManagementUserGroupDetails, error) {
	return s.db.Queries().GetUserGroupsLatestFirst(
		ctx,
		limit,
	)
}

func (s *userManagementService) getUserGroupsInitialOldestFirst(
	ctx context.Context,
	limit int32,
) ([]db.ManagementUserGroupDetails, error) {
	return s.db.Queries().GetUserGroupsOldestFirst(
		ctx,
		limit,
	)
}

func (s *userManagementService) getUserGroupsGreaterThanCursorLatestFirst(
	ctx context.Context,
	cursor time.Time,
	limit int32,
) ([]db.ManagementUserGroupDetails, error) {
	return s.db.Queries().GetUserGroupsLatestFirstGreaterThanCursor(
		ctx,
		db.GetUserGroupsLatestFirstGreaterThanCursorParams{
			Cursor: cursor,
			Limit:  limit,
		},
	)
}

func (s *userManagementService) getUserGroupsGreaterThanCursorOldestFirst(
	ctx context.Context,
	cursor time.Time,
	limit int32,
) ([]db.ManagementUserGroupDetails, error) {
	return s.db.Queries().GetUserGroupsOldestFirstGreaterThanCursor(
		ctx,
		db.GetUserGroupsOldestFirstGreaterThanCursorParams{
			Cursor: cursor,
			Limit:  limit,
		},
	)
}

func (s *userManagementService) getUserGroupsLesserThanCursorLatestFirst(
	ctx context.Context,
	cursor time.Time,
	limit int32,
) ([]db.ManagementUserGroupDetails, error) {
	return s.db.Queries().GetUserGroupsLatestFirstLesserThanCursor(
		ctx,
		db.GetUserGroupsLatestFirstLesserThanCursorParams{
			Cursor: cursor,
			Limit:  limit,
		},
	)
}

func (s *userManagementService) getUserGroupsLesserThanCursorOldestFirst(
	ctx context.Context,
	cursor time.Time,
	limit int32,
) ([]db.ManagementUserGroupDetails, error) {
	return s.db.Queries().GetUserGroupsOldestFirstLesserThanCursor(
		ctx,
		db.GetUserGroupsOldestFirstLesserThanCursorParams{
			Cursor: cursor,
			Limit:  limit,
		},
	)
}
