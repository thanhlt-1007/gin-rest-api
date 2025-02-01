package auth

type SignUpRequest struct {
    Email string `form:"email" binding:"required"`
    Password string `form:"password" binding:"required"`
}
