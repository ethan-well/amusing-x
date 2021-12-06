package europa

type OAuthLogin struct {
	Code     string `json:"code"`
	Provider string `json:"provider"`
}
