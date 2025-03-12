package bindings

type ErrorCode string

const (
	NotLoggedIn     = ErrorCode("NotLoggedIn")
	ValidationError = ErrorCode("ValidationError")
	NotFoundError   = ErrorCode("NotFound")
	InternalError   = ErrorCode("InternalError")
)

var ErrorCodes = []ErrorCode{ValidationError, NotFoundError, InternalError, NotLoggedIn}

func (e ErrorCode) TSName() string { return string(e) }
