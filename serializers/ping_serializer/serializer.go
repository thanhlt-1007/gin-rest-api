package ping_serializer

import (
    "gin-rest-api/responses/ping_response"
)

func Serializer() ping_response.Response {
    return ping_response.Response {
        Message: "pong",
    }
}
