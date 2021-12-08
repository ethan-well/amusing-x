package github

const (
	ProviderGitHub = "github"
	RequestErrKey  = "error"
)

type AccessTokenResponse struct {
	AccessToken      string `json:"access_token"`
	Scope            string `json:"scope"`
	TokenType        string `json:"token_type"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorUri         string `json:"error_uri"`
}

type AccessTokenErrResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorUri         string `json:"error_uri"`
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
