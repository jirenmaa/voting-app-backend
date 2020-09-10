package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

// Token struct
type Token struct {
	JwtAlgo   *jwt.SigningMethodHMAC
	JwtClaim  *Claim
	JwtSecret string
}

// Gets the signed token
func (t *Token) GetToken() (string, error) {
	token := jwt.NewWithClaims(t.JwtAlgo, t.JwtClaim)

	return token.SignedString([]byte(t.JwtSecret))
}
