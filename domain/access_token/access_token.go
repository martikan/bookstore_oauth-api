package access_token

import (
	"strings"
	"time"

	"github.com/martikan/bookstore_oauth-api/errors"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"` // For identificate client (fore example Web or Phone, etc.)
	Expires     int64  `json:"expires"`
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
