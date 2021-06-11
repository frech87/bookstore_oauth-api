package app

import (
	"github.com/frech87/bookstore_oauth-api/src/http"
	"github.com/frech87/bookstore_oauth-api/src/repository/db"
	"github.com/frech87/bookstore_oauth-api/src/service"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	//dbRepository := db.New()
	//atService := access_token.NewService(dbRepository)

	atService := service.NewService(db.New())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.POST("/oauth/access_token/:access_token_id", atHandler.Update)
	router.Run("127.0.0.1:8080")
}
