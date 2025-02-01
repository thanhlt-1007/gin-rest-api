package responses

type PingResponse struct {
    Message string
}

func (PingResponse) GetPingResponse() PingResponse {
    return PingResponse {
        Message: "pong",
    }
}
