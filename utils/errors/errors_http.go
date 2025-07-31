package app_errors

// var (
// 	ErrServer           = fmt.Errorf("server-error")
// 	ErrNetwork          = fmt.Errorf("networ-error")
// 	ErrTooManyRequests  = fmt.Errorf("too-many-requests-error")
// 	ErrAuthentication   = fmt.Errorf("authentication-error")
// 	ErrAuthorization    = fmt.Errorf("authorization-error")
// 	ErrNotFound         = fmt.Errorf("not-found-error")
// 	ErrMethodNotAllowed = fmt.Errorf("method-not-allowed-error")
// 	ErrFormField        = fmt.Errorf("form-field-error")
// 	ErrFormAction       = fmt.Errorf("form-action-error")
// 	ErrUnknown          = fmt.Errorf("unknown-error")
// )

type RequestDecodeError struct {
	Status  int
	Message string
}

func (e *RequestDecodeError) Error() string {
	return e.Message
}
