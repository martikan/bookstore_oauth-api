package rest

import (
	"encoding/json"
	"github.com/martikan/bookstore_oauth-api/domain/user"
	"github.com/martikan/bookstore_oauth-api/errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"time"
)

var (
	userRestClient = rest.RequestBuilder{
		BaseURL: "https:api.bookstore.com",
		Timeout: 100 * time.Millisecond,
	}
)

func NewRepository() UserRepository {
	return &userRepository{}
}

type userRepository struct {
}

type UserRepository interface {
	SignIn(string, string) (*user.User, *errors.RestError)
}

func (u *userRepository) SignIn(email string, passwd string) (*user.User, *errors.RestError) {

	request := user.SignInRequest{
		Email:    email,
		Password: passwd,
	}

	response := userRestClient.Post("/users/sign-in", request)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("Invalid response when trying to sign in")
	}

	if response.StatusCode > 299 { // When error occur
		var restErr errors.RestError
		if err := json.Unmarshal(response.Bytes(), &restErr); err != nil {
			return nil, errors.NewInternalServerError("Invalid error interface when trying to sign in")
		}

		return nil, &restErr
	}

	var usr user.User
	if err := json.Unmarshal(response.Bytes(), &usr); err != nil {
		return nil, errors.NewInternalServerError("Error when trying to unmarshal UserDTO response")
	}
	
	return &usr, nil
}
