package helpers

type QueryKey string

const (
	NextCursor  QueryKey = "next_cursor"
	PrevCursor  QueryKey = "prev_cursor"
	Sort        QueryKey = "sort"
	SearchQuery QueryKey = "search"
)
