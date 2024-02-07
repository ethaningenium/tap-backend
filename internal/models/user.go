package models

type UserRegister struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserRegisterRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	RefreshToken string `json:"access_token"`
}

type UserLogin struct {
	Email string `json:"email"`
	Password string `json:"password"`
}




