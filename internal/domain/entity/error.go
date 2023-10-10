package entity

import (
	"fmt"
)

type AppError struct {
	State   int
	Code    ErrorCode
	Message string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Type: %d, Code: %d, Message: %s", e.State, e.Code, ErrorDetails[e.Code].Message)
}

func NewAppError(code ErrorCode, values ...interface{}) *AppError {
	detail := getFormattedErrorDetailByCode(code, values...)
	if detail == nil {
		return nil
	}

	return &AppError{
		Code:    code,
		State:   detail.State,
		Message: detail.Message,
	}
}

func getFormattedErrorDetailByCode(code ErrorCode, args ...interface{}) *ErrorDetail {
	detail, ok := ErrorDetails[code]
	if !ok {
		return nil
	}

	formattedMessage := fmt.Sprintf(detail.Message, args...)
	detail.Message = formattedMessage

	return &detail
}
