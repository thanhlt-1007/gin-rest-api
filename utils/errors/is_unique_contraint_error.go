package errors

import (
    "regexp"
    "strings"
    "gin-rest-api/errors"
)

func IsUniqueContraintError(err error) (*errors.UniqueContraintError, bool) {
    reg, _ := regexp.Compile("UNIQUE constraint failed: ([a-z.]+)")
    matches := reg.FindStringSubmatch(err.Error())
    if len(matches) == 0 {
        return nil, false
    }

    substrings := strings.Split(matches[1], ".")
    table_name := substrings[0]
    fields := strings.Split(substrings[1], "_")

    unique_contraint_error := errors.UniqueContraintError {
        RootError: err,
        TableName: table_name,
        Fields: fields,
    }
    return &unique_contraint_error, true
}
