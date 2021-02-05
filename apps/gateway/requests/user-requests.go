package requests

import "net/http"

type LoginRequest struct {
	Login string
	Password string
}
func ParseLoginRequest(r *http.Request) *LoginRequest {
	return &LoginRequest{
		Login: r.FormValue("login"),
		Password: r.FormValue("password"),
	}
}

type SignupRequest struct {
	Login string
	Password string
}
func ParseSignupRequest(r *http.Request) *SignupRequest {
	return &SignupRequest{
		Login: r.FormValue("login"),
		Password: r.FormValue("password"),
	}
}
