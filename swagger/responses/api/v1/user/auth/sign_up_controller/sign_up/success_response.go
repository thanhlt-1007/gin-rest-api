package sign_up

import (
    "gin-rest-api/responses/api/v1/user/auth/sign_up_response"
)

type SuccessResponse struct {
    Data sign_up_response.Response `json:"data"`
}
