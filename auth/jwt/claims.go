package jwt

import "github.com/dgrijalva/jwt-go"

type Claim struct {
	UserID uint64 `json:"user_id"`
	jwt.StandardClaims
}

// Makes a new JWT Claim
func NewClaim(userId uint64, expiry int64) *Claim {
	return &Claim{
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiry,
		},
	}
}
