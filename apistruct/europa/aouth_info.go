package europa

type OAuthInfo struct {
	Provider     string `json:"provider"`
	OauthUrl     string `json:"oauth_url"`
	ClientID     string `json:"client_id"`
	Scope        string `json:"scope"`
	State        string `json:"state"`
	GrantType    string `json:"grant_type"`
	RedirectUrl  string `json:"redirect_url"`
	CompletePath string `json:"complete_path"`
}
