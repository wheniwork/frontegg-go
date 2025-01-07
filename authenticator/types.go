package authenticator

type AuthenticateRequest struct {
	ClientID string `json:"clientId"`
	Secret   string `json:"secret"`
}

type AuthenticateResponse struct {
	Token     string `json:"token"`
	ExpiresIn int    `json:"expiresIn"`
}
