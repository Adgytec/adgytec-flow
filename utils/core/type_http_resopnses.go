package core

type ResponseHTTPError struct {
	HTTPStatusCode int                `json:"-"`
	Message        *string            `json:"message,omitempty"`
	FieldErrors    *map[string]string `json:"fieldErrors,omitempty"`
}
