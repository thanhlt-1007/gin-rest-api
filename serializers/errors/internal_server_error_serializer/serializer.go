package internal_server_error_serializer

import (
	"fmt"
    "gin-rest-api/responses/errors/internal_server_error_response"
)

func Serializer(err error) internal_server_error_response.Response {
    return internal_server_error_response.Response {
        Message: fmt.Sprintf("UNKNOW_ERROR: [%v] - [%#v], Message: %s", err, err, err.Error()),
		Code: "INTERNAL_SERVER_ERROR",
    }
}
