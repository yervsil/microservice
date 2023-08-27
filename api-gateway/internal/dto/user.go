package dto

type SignUp struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type SignIn struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}