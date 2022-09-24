package auth

import (
	"os"
)

// JwtKey is the key used to sign JWT tokens.
var JwtKey = []byte(os.Getenv("JWT_KEY"))

// Users is a list of users.
