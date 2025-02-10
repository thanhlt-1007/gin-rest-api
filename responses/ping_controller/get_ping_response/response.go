package get_ping_response

type Response struct {
    Data DataResponse `json:"data"`
}

type DataResponse struct {
    Message string `json:"message" example:"pong"`
}
