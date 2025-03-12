package bindings

import (
	"fmt"
)

type ErrorResponse struct {
	Code ErrorCode   `json:"code"`
	Data []ErrorData `json:"data"`
}

type ErrorData struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *ErrorResponse) Error() string {
	if len(e.Data) == 0 {
		return string(e.Code)
	} else {
		return e.Data[0].Message
	}
}

func NewValidationError(field string, format string, a ...any) *ErrorResponse {
	return &ErrorResponse{
		Code: ValidationError,
		Data: []ErrorData{{field, fmt.Sprintf(format, a...)}},
	}
}

func NewInternalError(format string, a ...any) *ErrorResponse {
	return &ErrorResponse{
		Code: InternalError,
		Data: []ErrorData{{"error", fmt.Sprintf(format, a...)}},
	}
}

func NewNotFoundError() *ErrorResponse {
	return &ErrorResponse{Code: NotFoundError}
}

func NewNotLoggedInError() *ErrorResponse {
	return &ErrorResponse{Code: NotLoggedIn}
}
