package sign_up_request

type Request struct {
    Email string `form:"email" binding:"required,max=255,email,unique_in_model=User"`
    Password string `form:"password" binding:"required,max=255"`
    ConfirmPassword string `form:"confirm_password" binding:"required,eqfield=Password"`
}
