package types

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	Username string `json:"username,omitempty"`
	Type     string `json:"type,omitempty"`
	jwt.StandardClaims
}
