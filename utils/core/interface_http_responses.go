package core

type ErrorResponse interface {
	HTTPResponse() ResponseHTTPError
}
