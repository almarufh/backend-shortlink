package dto

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email" example:"almarufhidayat99@gmail.com"`
	Password string `json:"password" binding:"min=7" example:"Admin@1234"`
}
