package app

import (
	"github.com/gin-gonic/gin"
	"github.com/martikan/bookstore_oauth-api/client/cassandra"
	"github.com/martikan/bookstore_oauth-api/domain/access_token"
	"github.com/martikan/bookstore_oauth-api/http"
	"github.com/martikan/bookstore_oauth-api/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {

	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	defer session.Close()

	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)

	router.Run(":8080")
}
