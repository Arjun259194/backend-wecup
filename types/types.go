package types

type Response struct {
	Status       int         `json:"status"`
	ResponseData interface{} `json:"data"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Gender   string `json:"gender" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
