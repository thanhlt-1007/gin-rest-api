package unique_contraint_error_response

import (
    "gin-rest-api/responses/errors/field_error_response"
)

type Response struct {
    Message string `json:"message"`
    Code string `json:"code"`
    Errors []field_error_response.Response `json:"errors"`
}
