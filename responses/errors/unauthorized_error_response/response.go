package unauthorized_error_response

type Response struct {
	Message string `json:"message" example:"Unauthorized Error"`
	Code    string `json:"code" example:"UNAUTHORIZED_ERROR"`
}
