package enums

type ErrorCode string

const (
	ValidationError = ErrorCode("ValidationError")
	NotFoundError   = ErrorCode("NotFound")
	InternalError   = ErrorCode("InternalError")
)

var ErrorCodes = []ErrorCode{ValidationError, NotFoundError, InternalError}

func (e ErrorCode) TSName() string { return string(e) }
