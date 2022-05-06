package access_token

import (
	"github.com/martikan/bookstore_oauth-api/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestError)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestError)
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
	accessToken, err := s.repository.GetById(at)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}
