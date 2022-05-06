package access_token

import (
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	if expirationTime != 24 {
		t.Error("Expiration time should be 24 hours")
	}
}

func TestGetNewAccessToken(t *testing.T) {

	at := GetNewAccessToken()

	if at.IsExpired() {
		t.Error("Brand new access token should not be expired")
	}

	if at.AccessToken != "" {
		t.Error("Brand new access token should not have defined access token id")
	}

	if at.UserId != 0 {
		t.Error("Brand new access token should not have an associated user_id")
	}

}

func TestAccessTokenIsExpired(t *testing.T) {

	at := AccessToken{}

	if !at.IsExpired() {
		t.Error("Empty access token should be expired")
	}

	at.Expires = time.Now().Add(3 * time.Hour).Unix()
	if at.IsExpired() {
		t.Error("Access token expiring theree hours from now should NOT be expired")
	}

}
