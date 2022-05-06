package db

import (
	"github.com/gocql/gocql"
	"github.com/martikan/bookstore_oauth-api/client/cassandra"
	"github.com/martikan/bookstore_oauth-api/domain/access_token"
	"github.com/martikan/bookstore_oauth-api/errors"
)

const (
	getAccessToken = "SELECT access token, user_id, client_id, expires FROM access_tokens WHERE access_token = ?;"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
}

type dbRepository struct{}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestError) {

	session, err := cassandra.GetSession()
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer session.Close()

	var result access_token.AccessToken
	if err := session.Query(getAccessToken, id).Scan(&result.AccessToken, &result.UserId, &result.ClientId, &result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("No access token found with the given id")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &result, nil
}
