package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type AccessDetails struct {
	UserId string
	Groups []interface{}
	Roles []interface{}
}

// ExtractToken extracts the JWT token from the header.
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	// normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// VerifyToken verifies the JWT token to ensure the public key was provided and it was signed via RSA.
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)

	key, err := jwt.ParseRSAPublicKeyFromPEM(getPublicKey())
	if err != nil {
		return nil, fmt.Errorf("error parsing RSA public key: %v\n", err)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token method conform to "SigningMethodRSA"
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// ExtractTokenMetadata extracts the data contained within the JWT token and returns an AccessDetails object.
func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId, ok := claims["sub"].(string)
		if !ok {
			return nil, err
		}

		groups, ok := claims["groups"].([]interface{})
		if !ok {
			return nil, err
		}

		roles, ok := claims["roles"].([]interface{})
		if !ok {
			return nil, err
		}

		return &AccessDetails{
			UserId: userId,
			Groups: groups,
			Roles: roles,
		}, nil
	}
	return nil, err
}

// getPublicKey returns the public key from the specified file
func getPublicKey() []byte {
	publicKey, err := ioutil.ReadFile("services/admin/backend/config/keycloak_realm_key.rsa.pub")
	if err != nil {
		log.Fatal(err.Error())
	}

	return publicKey
}
