package dto

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email" example:"almarufhidayat99@gmail.com"`
	Password string `json:"password" binding:"min=8" example:"Admin@1234"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"almarufhidayat99@gmail.com"`
	Password string `json:"password" binding:"min=8" example:"Admin@1234"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
