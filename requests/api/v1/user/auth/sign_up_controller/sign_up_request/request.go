package sign_up_request

type Request struct {
	Email           string `form:"email" json:"email" binding:"required,max=255,email,unique_in_model=User" example:"user@example.com"`
	Password        string `form:"password" json:"password" binding:"required,max=255" example:"Aa@123456"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required,eqfield=Password" example:"Aa@123456"`
}
