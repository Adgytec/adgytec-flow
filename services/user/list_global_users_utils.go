package user

import (
	"context"
	"time"

	"github.com/Adgytec/adgytec-flow/database/db"
)

func (s *userService) getGlobalUsersQuery(
	ctx context.Context,
	searchQuery string,
	limit int32,
) ([]db.GlobalUserDetail, error) {
	return s.db.Queries().GetGlobalUsersByQuery(
		ctx,
		db.GetGlobalUsersByQueryParams{
			Limit: limit,
			Query: searchQuery,
		},
	)
}

func (s *userService) getGlobalUsersInitialLatestFirst(
	ctx context.Context,
	limit int32,
) ([]db.GlobalUserDetail, error) {
	return s.db.Queries().GetGlobalUsersLatestFirst(
		ctx,
		limit,
	)
}

func (s *userService) getGlobalUsersInitialOldestFirst(
	ctx context.Context,
	limit int32,
) ([]db.GlobalUserDetail, error) {
	return s.db.Queries().GetGlobalUsersOldestFirst(
		ctx,
		limit,
	)
}

func (s *userService) getGlobalUsersGreaterThanCursorLatestFirst(
	ctx context.Context,
	cursor time.Time,
	limit int32,
) ([]db.GlobalUserDetail, error) {
	return s.db.Queries().GetGlobalUsersLatestFirstGreaterThanCursor(
		ctx,
		db.GetGlobalUsersLatestFirstGreaterThanCursorParams{
			Cursor: cursor,
			Limit:  limit,
		},
	)
}

func (s *userService) getGlobalUsersGreaterThanCursorOldestFirst(
	ctx context.Context,
	cursor time.Time,
	limit int32,
) ([]db.GlobalUserDetail, error) {
	return s.db.Queries().GetGlobalUsersOldestFirstGreaterThanCursor(
		ctx,
		db.GetGlobalUsersOldestFirstGreaterThanCursorParams{
			Cursor: cursor,
			Limit:  limit,
		},
	)
}

func (s *userService) getGlobalUsersLesserThanCursorLatestFirst(
	ctx context.Context,
	cursor time.Time,
	limit int32,
) ([]db.GlobalUserDetail, error) {
	return s.db.Queries().GetGlobalUsersLatestFirstLesserThanCursor(
		ctx,
		db.GetGlobalUsersLatestFirstLesserThanCursorParams{
			Cursor: cursor,
			Limit:  limit,
		},
	)
}

func (s *userService) getGlobalUsersLesserThanCursorOldestFirst(
	ctx context.Context,
	cursor time.Time,
	limit int32,
) ([]db.GlobalUserDetail, error) {
	return s.db.Queries().GetGlobalUsersOldestFirstLesserThanCursor(
		ctx,
		db.GetGlobalUsersOldestFirstLesserThanCursorParams{
			Cursor: cursor,
			Limit:  limit,
		},
	)
}
