package unauthorized_error_serializer

import (
	"fmt"

	"gin-rest-api/consts"
	"gin-rest-api/errors"
	"gin-rest-api/responses/errors/unauthorized_error_response"
)

func Serializer(err errors.UnauthorizedError) unauthorized_error_response.Response {
	message := fmt.Sprintf("UNAUTHORIZED_ERROR: %s", err.Error())
	code := consts.UNAUTHORIZED_ERROR

	return unauthorized_error_response.Response{
		Message: message,
		Code:    code,
	}
}
