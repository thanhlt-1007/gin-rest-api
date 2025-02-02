package unknown_panic_error_serializer

import (
    "fmt"
    "gin-rest-api/consts"
    "gin-rest-api/responses/errors/internal_server_error_response"
)

func Serializer(err any) internal_server_error_response.Response {
    message := fmt.Sprintf("UNKNOW_PANIC Type: [%T], Error: [%v], Error: [%#v]", err, err, err)
    code := consts.INTERNAL_SERVER_ERROR_CODE

    return internal_server_error_response.Response {
        Message: message,
		Code: code,
    }
}
