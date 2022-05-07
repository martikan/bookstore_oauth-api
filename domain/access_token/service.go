package access_token

import (
	"strings"

	"github.com/martikan/bookstore_oauth-api/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestError)
	Create(AccessToken) *errors.RestError
	UpdateExpirationTime(AccessToken) *errors.RestError
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestError)
	Create(AccessToken) *errors.RestError
	UpdateExpirationTime(AccessToken) *errors.RestError
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(at string) (*AccessToken, *errors.RestError) {

	accessTokenId := strings.TrimSpace(at)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("Invalid access token")
	}

	accessToken, err := s.repository.GetById(at)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (s *service) Create(at AccessToken) *errors.RestError {

	if err := at.Validate(); err != nil {
		return err
	}
	
	return s.repository.Create(at)
}

func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestError {

	if err := at.Validate(); err != nil {
		return err
	}

	return s.repository.UpdateExpirationTime(at)
}
