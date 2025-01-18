package bindings

import (
	"encoding/json"
	"fmt"
)

type ErrorResponse struct {
	Code ErrorCode  `json:"code"`
	Data *ErrorData `json:"data"`
}

type ErrorData struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *ErrorResponse) Error() string {
	errJSON, _ := json.Marshal(e.Data.Message)
	return string(errJSON)
}

func NewValidationError(field string, format string, a ...any) *ErrorResponse {
	return &ErrorResponse{
		Code: ValidationError,
		Data: &ErrorData{field, fmt.Sprintf(format, a...)},
	}
}

func NewInternalError(format string, a ...any) *ErrorResponse {
	return &ErrorResponse{
		Code: InternalError,
		Data: &ErrorData{"error", fmt.Sprintf(format, a...)},
	}
}

func NewNotFoundError() *ErrorResponse {
	return &ErrorResponse{Code: NotFoundError}
}
