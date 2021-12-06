package github

type AccessTokenRequest struct {
	ClientID string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code string `json:"code"`
	RedirectUlr string `json:"redirect_ulr"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope string `json:"scope"`
	TokenType string `json:"token_type"`
}