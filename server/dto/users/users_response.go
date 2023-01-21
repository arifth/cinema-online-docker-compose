package usersdto

type UserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email"`
	Password string `json:"password" form:"password" validate:"required"`
	// Phone          string `json:"phone" form:"phone" validate:"required"`
	// Address        string `json:"address" form:"address" validate:"required"`
	// Image          string `json:"image" form:"image" validate:"required"`
	// Transaction_id int    `json:"transaction_id" form:"transaction_id" validate:"required"`
}
