package types

type Response struct {
	Status       int         `json:"status"`
	ResponseData interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Gender   string `json:"gender" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

type UserUpdateRequest struct {
	Name   string `json:"name" validate:"required"`
	Gender string `json:"gender" validate:"required"`
	Email  string `json:"email" validate:"required"`
}

type CreatePostRequest struct {
  Content string `json:"content" validate:"required"`
}

type UpdatePostRequest struct {
  Content string `json:"content" validate:"required"`
}

type CreateCommentRequest struct {
  Content string `json:"content" validate:"required"`
}
