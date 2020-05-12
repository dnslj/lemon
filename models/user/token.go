package user

type Token struct {
	Token    string `json:"token"`
	ExpireIn int    `json:"expires_in"`
}
