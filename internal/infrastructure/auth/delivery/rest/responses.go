package rest

type errorResponse struct {
	Error string `json:"error"`
	Msg   string `json:"msg"`
}

type tokenPairResponse struct {
	Access  string `json:"access"`
	Session string `json:"session"`
}

type accessTokenResponse struct {
	Access string `json:"access"`
}
