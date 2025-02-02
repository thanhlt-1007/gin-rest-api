package unique_contraint_error_serializer

import (
    "fmt"
    "gin-rest-api/responses/errors/field_error_response"
    "gin-rest-api/responses/errors/unique_contraint_error_response"
    "gin-rest-api/errors"
)

func Serializer(err errors.UniqueContraintError) unique_contraint_error_response.Response {
    message := fmt.Sprintf("UNIQUE_CONTRAINT_ERROR: %s", err.Error())
    code := "UNIQUE_CONTRAINT"
    var errors []field_error_response.Response
        for _, field := range err.Fields {
            err := field_error_response.Response{
                Field: field,
                Tag: "unique",
            }
            errors = append(errors, err)
        }

    return unique_contraint_error_response.Response {
        Message: message,
        Code: code,
        Errors: errors,
    }
}
