package oauth

import (
	"auth-microservice/src/api/utils/errors"
	"fmt"
)

var (
	tokens = make(map[string]*AccessToken, 0)
)

func (at *AccessToken) Save() errors.ApiError {
	// имитация сохранения токена в БД
	// создаем токен по данным пользователя
	at.AccessToken = fmt.Sprintf("%d", at.UserId)
	// сохраняем в БД
	tokens[at.AccessToken] = at

	return nil
}
