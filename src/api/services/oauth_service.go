package services

import (
	"auth-microservice/src/api/domain/oauth"
	"auth-microservice/src/api/utils/errors"
	"time"
)

type oauthService struct{}

type oauthServiceInterface interface {
	CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, errors.ApiError)
	GetAccessToken(accessToken string) (*oauth.AccessToken, errors.ApiError)
}

var (
	OauthService oauthServiceInterface
)

func init() {
	OauthService = &oauthService{}
}

func (s *oauthService) CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, errors.ApiError) {
	// валидируем входящие данные
	if err := request.Validate(); err != nil {
		return nil, err
	}

	// обращаем к БД за данными через ДАО
	user, err := oauth.GetUserByUsernameAndPassword(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	// создаем новый токен для пользователя
	token := oauth.AccessToken{
		UserId:  user.Id,
		Expires: time.Now().UTC().Add(24 * time.Hour).Unix(),
	}

	// сохраняем токен в БД
	if err := token.Save(); err != nil {
		return nil, err
	}

	return &token, nil
}

func (s *oauthService) GetAccessToken(accessToken string) (*oauth.AccessToken, errors.ApiError) {
	token, err := oauth.GetAccessToken(accessToken)

	if err != nil {
		return nil, err
	}

	if token.IsExpired() {
		return nil, errors.NewNotFoundApiError("no access token found")
	}

	return token, nil
}
