package app

import (
	"auth-microservice/src/api/controllers/oauth"
	"auth-microservice/src/api/controllers/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/oauth/access_token", oauth.CreateAccessToken)
}
