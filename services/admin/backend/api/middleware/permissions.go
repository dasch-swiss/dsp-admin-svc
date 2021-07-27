package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type AccessDetails struct {
	UserId string
	Permissions []string
}

// Requesting Party Token
type RPT struct {
	// Scopes Client authorization scopes (essentially the permission strings)
	Scopes []string `json:"scopes"`

	// ResourceId Client authorization resource id
	ResourceId string `json:"resourceId"`

	// ResourceName Client authorization resource name
	ResourceName string `json:"resourceName"`
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

// TokenValid checks if a JWT token is valid.
func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
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

		permissions, permissionsErr := getPermissions(r)
		if permissionsErr != nil {
			return nil, permissionsErr
		}

		return &AccessDetails{
			UserId: userId,
			Permissions: permissions,
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

// getPermissions returns a list of the users permissions
func getPermissions(r *http.Request) ([]string, error) {

	// check token to make sure it's valid
	_, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}

	endpoint := "https://auth.dasch.swiss/auth/realms/permissions-test/protocol/openid-connect/token"
	data := url.Values{}
	data.Set("grant_type", "urn:ietf:params:oauth:grant-type:uma-ticket")
	data.Set("audience", "projects-api")
	data.Set("response_mode", "permissions")

	client := &http.Client{}

	req, reqErr := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
	if reqErr != nil {
		log.Fatal(reqErr)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", r.Header.Get("Authorization"))
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	// make request to the endpoint to obtain the users RPT which contains the users permissions
	res, resErr := client.Do(req)
	if resErr != nil {
		return nil, resErr
	}
	defer res.Body.Close()

	// body is a byte array which is needed for unmarshalling (Golang term for deserialization)
	body, bodyErr := ioutil.ReadAll(res.Body)
	if bodyErr != nil {
		return nil, bodyErr
	}

	// users list of RPTs
	var rpt []RPT

	// unmarshal/deserialize the body from the response
	unmarshalError := json.Unmarshal(body, &rpt)
	if unmarshalError != nil {
		return nil, errors.New("user has no permissions")
	}

	// list of permission to return
	var permissions []string

	// loop through the list of users RPTs to create a list of all the permissions the user has (scope = permission)
	for _, tok := range rpt {
		for _, scope := range tok.Scopes {
			permissions = append(permissions, scope)
		}
	}

	return permissions, err
}
