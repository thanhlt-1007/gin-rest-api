package get_ping

import (
    "gin-rest-api/responses/ping_response"
)

type SuccessResponse struct {
    Data ping_response.Response `json:"data"`
}
