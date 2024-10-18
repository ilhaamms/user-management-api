package response

type UserRegisterResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserLoginResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type Pagination struct {
	CurrenPage int `json:"current_page"`
	TotalPage  int `json:"total_page"`
	Limit      int `json:"limit"`
}
