package entry

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sdttttt/go-sso/util"
)

type AuthenticationModule struct{}

func (*AuthenticationModule) VerifyToken(token string) bool {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return util.SecetKey, nil
	})

	if err == nil {
		return true
	}

	return false
}

func (*AuthenticationModule) GenerateToken() string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: int64(time.Now().Unix() + 36000),
		Issuer:    "sdttttt",
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(util.SecetKey)

	if err != nil {
		log.Println(err)
		return ""
	}

	return tokenString
}
