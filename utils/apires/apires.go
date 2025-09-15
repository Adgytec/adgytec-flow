package apires

type ErrorResponse interface {
	HTTPResponse() ErrorDetails
}

type ErrorDetails struct {
	HTTPStatusCode int     `json:"-"`
	Message        *string `json:"message,omitempty"`
	FieldErrors    error   `json:"fieldErrors,omitempty"`
}
