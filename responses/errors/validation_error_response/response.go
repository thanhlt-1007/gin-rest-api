package validation_error_response

import (
    "gin-rest-api/responses/errors/field_error_response"
)

type Response struct {
    Message string `json:"message" example:"Validation Error"`
    Code string `json:"code" example:"VALIDATION_ERROR"`
    Errors []field_error_response.Response `json:"errors"`
}
