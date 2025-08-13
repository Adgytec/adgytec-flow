package helpers

import "net/http"

func GetRequestQueryValue(r *http.Request, key QueryKey) string {
	queryVal := r.URL.Query().Get(string(key))
	return queryVal
}
