package unknown_panic_error_serializer

import (
	"fmt"
    "gin-rest-api/responses/errors/internal_server_error_response"
)

func Serializer(err any) internal_server_error_response.Response {
    return internal_server_error_response.Response {
        Message: fmt.Sprintf("UNKNOW_PANIC: [%v] - [%#v]", err, err),
		Code: "INTERNAL_SERVER_ERROR",
    }
}
