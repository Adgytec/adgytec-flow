package core

import "time"

type IPaginationItem interface {
	GetCreatedAt() time.Time
}
