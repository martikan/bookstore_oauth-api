package access_token

import (
	"strings"
	"time"

	"github.com/martikan/bookstore_oauth-api/errors"
)

const (
	expirationTime = 24

	// GrantTypes
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for `Password` GrantType
	Username string `json:"username"`
	Password string `json:"password"`

	// Used for `Client credentials` GrantType
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"` // Identification for client (fore example Web or Phone, etc.)
	Expires     int64  `json:"expires"`
}

func (at *AccessTokenRequest) Validate() *errors.RestError {

	switch at.GrantType {
	case grantTypeClientCredentials:
		break
	case grantTypePassword:
		break
	default:
		return errors.NewBadRequestError("Invalid grant_type parameter")
	}

	return nil
}

func (at *AccessToken) Validate() *errors.RestError {

	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("Invalid access token")
	}

	if at.UserId <= 0 {
		return errors.NewBadRequestError("Invalid user id")
	}

	if at.ClientId <= 0 {
		return errors.NewBadRequestError("Invalid client id")
	}

	if at.Expires <= 0 {
		return errors.NewBadRequestError("Invalid expiration time")
	}

	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Generate(id int64) {
	at.UserId = id

	// TODO: Generate Access Token
}
