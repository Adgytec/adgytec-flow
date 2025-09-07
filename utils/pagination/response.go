package pagination

func newPaginationResponse[T PaginationItem](items []T, next, prev *T) ResponsePagination[T] {
	var pageInfo PageInfo

	if next != nil {
		// has next page
		pageInfo.HasNextPage = true
		pageInfo.NextCursor = encodeTimeToBase64((*next).GetCreatedAt())
	}

	if prev != nil {
		// has prev page
		pageInfo.HasPrevPage = true
		pageInfo.PrevCursor = encodeTimeToBase64((*prev).GetCreatedAt())
	}

	return ResponsePagination[T]{
		PageInfo:  pageInfo,
		PageItems: items,
	}
}
