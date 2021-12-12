package github

type AccessTokenResponse struct {
	AccessToken      string `json:"access_token"`
	Scope            string `json:"scope"`
	TokenType        string `json:"token_type"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorUri         string `json:"error_uri"`
}

type AccessTokenRequest struct {
	AccessTokenUrl string `json:"access_token_url"`
	ClientID       string `json:"client_id"`
	ClientSecret   string `json:"client_secret"`
	Code           string `json:"code"`
	RedirectUrl    string `json:"redirect_url"`
}

type UserProfile struct {
	Login     string `json:"login"`
	ID        int64  `json:"id"`
	AvatarUrl string `json:"avatar_url"`
	Name      string `json:"name"`
	Company   string `json:"company"`
	Blog      string `json:"blog"`
	Email     string `json:"email"`
	Location  string `json:"location"`
}

func (resp *AccessTokenResponse) IsError() bool {
	return len(resp.Error) != 0
}
