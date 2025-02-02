package sign_up_request

type Request struct {
    Email string `form:"email" binding:"required,email"`
    Password string `form:"password" binding:"required"`
}
