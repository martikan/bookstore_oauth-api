package rest

import (
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

// TestMain Entry point
func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestSignIn_Successful(t *testing.T) {

	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https:api.bookstore.com",
		ReqBody:      `{"email": "email@gmail.com", "password": "aaa"}`,
		RespBody:     `{"id": 1, "first_name": "John", "last_name": "Doe", "email": "jd@gmail.com"}`,
		RespHTTPCode: http.StatusOK,
	})

	repository := userRepository{}
	usr, err := repository.SignIn("email@gmail.com", "asd")

	assert.Nil(t, err)
	assert.NotNil(t, usr)
	assert.EqualValues(t, 1, usr.Id)
	assert.EqualValues(t, "John", usr.FirstName)
	assert.EqualValues(t, "Doe", usr.LastName)
	assert.EqualValues(t, "jd@gmail.com", usr.Email)
}

func TestSignIn_TimeoutFromApi(t *testing.T) {

	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https:api.bookstore.com",
		ReqBody:      `{"email": "email@gmail.com", "password": "aaa"}`,
		RespBody:     `{}`,
		RespHTTPCode: -1,
	})

	repository := userRepository{}
	usr, err := repository.SignIn("email@gmail.com", "asd")

	assert.Nil(t, usr)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid response when trying to sign in", err.Message)
}

func TestSignIn_InvalidErrorInterface(t *testing.T) {

	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https:api.bookstore.com",
		ReqBody:      `{"email": "email@gmail.com", "password": "aaa"}`,
		RespBody:     `{"status": "404", "error": "not_found", "message": "Invalid user credentials"}`,
		RespHTTPCode: http.StatusNotFound,
	})

	repository := userRepository{}
	usr, err := repository.SignIn("email@gmail.com", "asd")

	assert.Nil(t, usr)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid error interface when trying to sign in", err.Message)
}

func TestSignIn_InvalidSignInCredentials(t *testing.T) {

	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https:api.bookstore.com",
		ReqBody:      `{"email": "email@gmail.com", "password": "aaa"}`,
		RespBody:     `{"status": 404, "error": "not_found", "message": "Invalid user credentials"}`,
		RespHTTPCode: http.StatusNotFound,
	})

	repository := userRepository{}
	usr, err := repository.SignIn("email@gmail.com", "asd")

	assert.Nil(t, usr)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Invalid user credentials", err.Message)
}

func TestSignIn_InvalidJsonResponse(t *testing.T) {

	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https:api.bookstore.com",
		ReqBody:      `{"email": "email@gmail.com", "password": "aaa"}`,
		RespBody:     `{"id": "1", "first_name": "John", "last_name": "Doe", "email": "jd@gmail.com"}`,
		RespHTTPCode: http.StatusOK,
	})

	repository := userRepository{}
	usr, err := repository.SignIn("email@gmail.com", "asd")

	assert.Nil(t, usr)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Error when trying to unmarshal UserDTO response", err.Message)
}
