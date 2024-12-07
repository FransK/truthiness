package auth

import (
	"crypto/rsa"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
)

var jwtKey = []byte("replace-me-with-secret-key")

var (
	verifyKey  *rsa.PublicKey
	signKey    *rsa.PrivateKey
	serverPort int
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}

func CreateToken(username string, role string) (string, error) {
	// create a signer with specified method
	t := jwt.New(jwt.SigningMethodHS256)

	// set our claims
	t.Claims = &Claims{
		jwt.RegisteredClaims{
			// set the expire time
			// see https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.4
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
		username,
		role,
	}

	// Create token string
	return t.SignedString(jwtKey)
}

func ValidateTokenAndGetRole(r *http.Request) (string, error) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	}, request.WithClaims(&Claims{}))

	// let caller handle error
	if err != nil {
		return "", err
	}

	// token is valid
	return token.Claims.(*Claims).Role, nil
}
