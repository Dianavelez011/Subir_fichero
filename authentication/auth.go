package auth

type SignUpRequest struct {
    Name            string `json:"name" binding:"required"`
    Email           string `json:"email" binding:"required,email"`
    Password        string `json:"password" binding:"required"`
    ConfirmPassword string `json:"confirmPassword" binding:"required,eqfield=Password"`
}

type SignInRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

type JwtResponse struct {
    Token string `json:"token"`
}