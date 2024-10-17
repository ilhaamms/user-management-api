package response

type UserRegisterResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserLoginResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}
