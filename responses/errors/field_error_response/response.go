package field_error_response

type Response struct {
    Field string `json:"field" example:"email"`
    Tag string `json:"tag" example:"required"`
}
