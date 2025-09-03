package core

import "time"

type PaginationItem interface {
	GetCreatedAt() time.Time
}
