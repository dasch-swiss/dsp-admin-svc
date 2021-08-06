package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type UserInfo struct {
	UserId         string
	Groups         []interface{}
	Roles          []interface{}
	Projects       []string
	IsSystemAdmin  bool
	IsProjectAdmin bool
}

// ExtractToken extracts the JWT token from the header.
func ExtractToken(r *http.Request) string {
	log.Print("EXTRACT TOKEN")
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
	log.Print("VERIFY TOKEN")
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

// ExtractTokenMetadata extracts the data contained within the JWT token and returns an UserInfo object.
func ExtractTokenMetadata(r *http.Request) (*UserInfo, error) {
	log.Print("EXTRACT TOKEN METADATA")
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

		var isSysAdmin = false
		var isProjAdmin = false
		var projects []string

		// loop through the users groups
		for _, g := range groups {

			// split each group by ":"
			s := strings.Split(fmt.Sprintf("%v", g), ":")

			// check if last item is equal to "SystemAdmin"
			if s[len(s)-1] == "SystemAdmin" {
				isSysAdmin = true
			}

			// check if last item is equal to "ProjectAdmin"
			if s[len(s)-1] == "ProjectAdmin" {
				isProjAdmin = true

				// add the project id to list of projects user has access to
				projects = append(projects, s[len(s)-2])
			}
		}

		return &UserInfo{
			UserId:         userId,
			Groups:         groups,
			Roles:          roles,
			Projects:       projects,
			IsSystemAdmin:  isSysAdmin,
			IsProjectAdmin: isProjAdmin,
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
