package apperror

import "errors"

type AppError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

// New - 新しいAppErrorを生成する
func New(code ErrorCode, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

// Is - エラーコードが一致するか判定する
func Is(err error, code ErrorCode) bool {
	var e *AppError
	if errors.As(err, &e) {
		return e.Code == code
	}
	return false
}

type ErrorCode int

const (
	CodeInvalidArgument ErrorCode = 400
	CodeUnauthorized    ErrorCode = 401
	CodeForbidden       ErrorCode = 403
	CodeNotFound        ErrorCode = 404
	CodeConflict        ErrorCode = 409
	CodeInternalServer  ErrorCode = 500
)
