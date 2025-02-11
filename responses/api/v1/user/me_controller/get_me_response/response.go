package get_me_response

type Response struct {
	Data Data `json:"data"`
}

type Data struct {
	ID    uint    `json:"id" example:"1"`
	Email string  `json:"email" example:"user@example.com"`
	Name  *string `json:"name" example:"USER"`
}
