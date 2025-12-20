package bindings

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
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

var defaultValidator = validator.New()

func Validate(s any) error {
	if err := defaultValidator.Struct(s); err != nil {
		errData := make([]ErrorData, 0)
		for _, err := range err.(validator.ValidationErrors) {
			errData = append(errData, ErrorData{
				Field:   strings.ToLower(err.Field()),
				Message: fmt.Sprintf("Validation error on %s: %s %s", err.Field(), err.Tag(), err.Param()),
			})
		}
		return &ErrorResponse{
			Code: ValidationError,
			Data: errData,
		}
	}
	return nil
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
