package get_ping_serializer

import (
    "gin-rest-api/responses/ping_controller/get_ping_response"
)

func Serializer() get_ping_response.Response {
    data := get_ping_response.DataResponse{
        Message: "pong",
    }

    return get_ping_response.Response {
        Data: data,
    }
}
