package oauth

import (
	"auth-microservice/src/api/utils/errors"
	"fmt"
)

const (
	queryGetUserByUsernameAndPassword = "SELECT id, username FROM users WHERE username=? AND password=?;"
)

var (
	// типа БД с пользователем
	users = map[string]*User{
		"lam": {Id: 123, Username: "lam"},
	}
)

func GetUserByUsernameAndPassword(username string, password string) (*User, errors.ApiError) {
	user := users[username]
	if user == nil {
		return nil, errors.NewNotFoundApiError(fmt.Sprintf("no user with nickname '%s'", username))
	}

	return user, nil
}
