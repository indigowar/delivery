package rest

type startSessionRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type sessionTokenRequest struct {
	Token string `json:"token"`
}
