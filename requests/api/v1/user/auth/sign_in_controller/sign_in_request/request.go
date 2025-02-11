package sign_in_request

type Request struct {
	Email    string `form:"email" json:"email" binding:"required" example:"user@example.com"`
	Password string `form:"password" json:"password" binding:"required" example:"Aa@123456"`
}
