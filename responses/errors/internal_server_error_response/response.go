package internal_server_error_response

type Response struct {
    Message string `json:"message"`
    Code string `json:"code"`
}
