package oauth

type AccessToken struct {
	UserId      int64  `json:"user_id"`
	AccessToken string `json:"access_token"`
	Expires     int64  `json:"expires"` // timestamp
}
