// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.2

package types

type LoginRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserName string `json:"userName"`
	UserId   int64  `json:"userId"`
	Token    string `json:"token"`
}

type RegisterRequest struct {
	UserName  string `json:"userName"`
	Password1 string `json:"password"`
	Password2 string `json:"password2"`
}

type RegisterResponse struct {
	UserName string `json:"userName"`
	UserId   int64  `json:"userId"`
}
