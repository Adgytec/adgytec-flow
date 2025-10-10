package storage

import (
	"fmt"
	"net/url"

	"github.com/google/uuid"
)

const objectInitialStatus = "temp"

func newObjectTag(id uuid.UUID) string {
	return fmt.Sprintf("status=%s,id=%s", url.QueryEscape(objectInitialStatus), url.QueryEscape(id.String()))
}
