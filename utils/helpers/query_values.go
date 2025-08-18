package helpers

import (
	"net/http"
	"strings"
)

func GetRequestQueryValue(r *http.Request, key QueryKey) string {
	queryVal := r.URL.Query().Get(string(key))
	return strings.TrimSpace(queryVal)
}
