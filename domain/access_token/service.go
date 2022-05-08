package access_token

import (
	"github.com/martikan/bookstore_oauth-api/domain/user"
	"strings"

	"github.com/martikan/bookstore_oauth-api/errors"
)

type DbRepository interface {
	GetById(string) (*AccessToken, *errors.RestError)
	Create(AccessToken) *errors.RestError
	UpdateExpirationTime(AccessToken) *errors.RestError
}

type RestUserRepository interface {
	SignIn(string, string) (*user.User, *errors.RestError)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestError)
	Create(AccessToken) *errors.RestError
	UpdateExpirationTime(AccessToken) *errors.RestError
}

type service struct {
	dbRepository       DbRepository
	restUserRepository RestUserRepository
}

func NewService(dbRepo DbRepository, restUserRepo RestUserRepository) *service {
	return &service{
		dbRepository:       dbRepo,
		restUserRepository: restUserRepo,
	}
}

func (s *service) GetById(at string) (*AccessToken, *errors.RestError) {

	accessTokenId := strings.TrimSpace(at)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("Invalid access token")
	}

	accessToken, err := s.dbRepository.GetById(at)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (s *service) Create(r AccessTokenRequest) (*AccessToken, *errors.RestError) {

	if err := r.Validate(); err != nil {
		return nil, err
	}

	// Authenticate the User against the user-api
	usr, err := s.restUserRepository.SignIn(r.Username, r.Password)
	if err != nil {
		return nil, err
	}

	// Generate new Access Token
	at := GetNewAccessToken()
	at.Generate(usr.Id)

	// Save new Access Token
	if err := s.dbRepository.Create(at); err != nil {
		return nil, err
	}

	return &at, nil
}

func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestError {

	if err := at.Validate(); err != nil {
		return err
	}

	return s.dbRepository.UpdateExpirationTime(at)
}

func (s *service) SignIn(email string, passwd string) (*user.User, *errors.RestError) {

	usr, err := s.restUserRepository.SignIn(email, passwd)
	if err != nil {
		return nil, err
	}

	return usr, nil
}
