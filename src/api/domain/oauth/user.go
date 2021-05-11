package oauth

// модель пользователя в нашем приложения

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}
