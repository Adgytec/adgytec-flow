package core

type IErrorResponse interface {
	HTTPResponse() ResponseHTTPError
}
