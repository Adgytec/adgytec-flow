package core

type ResponseHTTPError struct {
	HTTPStatusCode int                `json:"-"`
	Message        *string            `json:"message,omitempty"`
	FieldErrors    *map[string]string `json:"fieldErrors,omitempty"`
}

type PageInfo struct {
	HasNextPage bool   `json:"hasNextPage"`
	NextCursor  string `json:"nextCursor"`
	HasPrevPage bool   `json:"hasPrevPage"`
	PrevCursor  string `json:"prevCursor"`
}

type ResponsePagination[T any] struct {
	PageInfo  PageInfo `json:"pageInfo"`
	PageItems []T      `json:"pageItems"`
}
