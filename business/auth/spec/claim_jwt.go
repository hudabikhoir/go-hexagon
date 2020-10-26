package spec

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// JWTClaimSpec to custom jwt claim
type JWTClaimSpec struct {
	UserID   string
	Username string
	RoleID   int
	jwt.StandardClaims
}
