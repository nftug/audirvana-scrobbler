package customerr

import (
	"audirvana-scrobbler/app/shared/enums"
	"encoding/json"
	"fmt"

	"github.com/go-task/task/v3/errors"
)

type ErrorResponse struct {
	content *errorContent
	inner   error
}

type errorContent struct {
	Code enums.ErrorCode `json:"code"`
	Data *errorData      `json:"data"`
}

type errorData struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *ErrorResponse) Error() string {
	errJSON, _ := json.Marshal(e.content)
	return string(errJSON)
}

func (e *ErrorResponse) Code() enums.ErrorCode { return e.content.Code }

func (e *ErrorResponse) Content() *errorContent { return e.content }

func NewValidationError(field string, format string, a ...any) error {
	return &ErrorResponse{
		content: &errorContent{
			Code: enums.ValidationError,
			Data: &errorData{field, fmt.Sprintf(format, a...)},
		},
		inner: fmt.Errorf(format, a...),
	}
}

func NewInternalError(format string, a ...any) error {
	return &ErrorResponse{
		content: &errorContent{
			Code: enums.InternalError,
			Data: &errorData{"error", fmt.Sprintf(format, a...)},
		},
		inner: fmt.Errorf(format, a...),
	}
}

func NewNotFoundError() error {
	return &ErrorResponse{
		content: &errorContent{Code: enums.NotFoundError},
		inner:   errors.New("Not found"),
	}
}
