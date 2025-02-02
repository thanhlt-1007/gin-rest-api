package internal_server_error_serializer

import (
    "fmt"
    "gin-rest-api/consts"
    "gin-rest-api/responses/errors/internal_server_error_response"
)

func Serializer(err error) internal_server_error_response.Response {
    message := fmt.Sprintf("UNKNOW_ERROR Type: [%T], Error: [%v], Error: [%#v], Message: %s", err, err, err, err.Error())
    code := consts.INTERNAL_SERVER_ERROR_CODE

    return internal_server_error_response.Response {
        Message: message,
        Code: code,
    }
}
