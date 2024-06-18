package schema

import "myapp/domain/apperror"

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewErrorResponse(err error) ErrorResponse {
	appErr, ok := err.(*apperror.AppError)
	if !ok {
		return ErrorResponse{
			Code:    int(apperror.CodeInternalServer),
			Message: err.Error(),
		}
	}

	return ErrorResponse{
		Code:    int(appErr.Code),
		Message: appErr.Message,
	}
}
