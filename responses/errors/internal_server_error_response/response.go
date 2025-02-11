package internal_server_error_response

type Response struct {
	Message string `json:"message" example:"Internal Server Error"`
	Code    string `json:"code" example:"INTERNAL_SERVER_ERROR"`
}
