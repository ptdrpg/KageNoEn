package model

type DeleteModel struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type RegisterResponse struct {
	Data  User   `json:"data"`
	Token string `json:"token"`
}
