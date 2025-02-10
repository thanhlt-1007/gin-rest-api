package get_ping_response

type Response struct {
    Data Data `json:"data"`
}

type Data struct {
    Message string `json:"message" example:"pong"`
}
