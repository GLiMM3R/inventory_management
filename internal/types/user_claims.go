package types

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	jwt.StandardClaims
}
