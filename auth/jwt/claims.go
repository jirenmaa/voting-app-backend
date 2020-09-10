package jwt

import "github.com/dgrijalva/jwt-go"

type Claim struct {
	UserID uint16 `json:"user_id"`
	jwt.StandardClaims
}

// Makes a new JWT Claim
func NewClaim(userId uint16, expiry int64) *Claim {
	return &Claim{
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiry,
		},
	}
}
