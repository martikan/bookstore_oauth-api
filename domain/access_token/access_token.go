package access_token

import (
	"fmt"
	"time"
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

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {

	now := time.Now().UTC()

	expirationTime := time.Unix(at.Expires, 0)
	fmt.Println("Expiration time:", expirationTime)

	return expirationTime.Before(now)
}
