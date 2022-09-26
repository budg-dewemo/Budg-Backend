package models

// JwtToken is a JWT token.
type JwtToken struct {
	Token string `json:"token"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
