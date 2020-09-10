package jwt

import (
	"fmt"
	"os"
	"time"

	gojwt "github.com/dgrijalva/jwt-go"
)

type JWT interface {
	GenerateToken(user uint16) (string, error)
	Validate(token string) (*gojwt.Token, error)
}

type Service struct {
	SecretKey string
}

type Claim struct {
	UserId uint16
	gojwt.StandardClaims
}

func NewJWTService() JWT {
	return &Service{
		SecretKey: os.Getenv("TOKEN_SECRET"),
	}
}

func (s *Service) GenerateToken(user uint16) (string, error) {
	claims := &Claim{
		UserId: user,
		StandardClaims: gojwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := gojwt.NewWithClaims(gojwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(s.SecretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *Service) Validate(token string) (*gojwt.Token, error) {
	return gojwt.Parse(token, func(token *gojwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*gojwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}

		return []byte(s.SecretKey), nil
	})
}
