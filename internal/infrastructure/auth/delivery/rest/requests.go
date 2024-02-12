package rest

type startSessionRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type sessionTokenRequest struct {
	token string `json:"token"`
}
