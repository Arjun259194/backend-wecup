package controllers

type errorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type reponse struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	ReponseData interface{} `json:"data"`
}

type registerRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	Email    string `json:"email"`
}
