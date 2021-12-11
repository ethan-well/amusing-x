package wechat

const ProviderWeChat = "weChat"

type AccessTokenRequest struct {
	AppID     string `json:"appid"`
	Secret    string `json:"secret"`
	Code      string `json:"code"`
	GrantType string `json:"grant_type"`
}

type AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
	Unionid      string `json:"unionid"`
	Errcode      int64  `json:"errcode"`
	Errmsg       string `json:"errmsg"`
}

func (resp *AccessTokenResponse) IsError() bool {
	return resp.Errcode != 0
}

type UserProfile struct {
	Openid    string   `json:"openid"`
	Nickname  string   `json:"nickname"`
	Sex       string   `json:"sex"`
	Province  string   `json:"province"`
	City      string   `json:"city"`
	Country   string   `json:"country"`
	AvatarUrl string   `json:"headimgurl"`
	Privilege []string `json:"privilege"`
	Unionid   string   `json:"unionid"`
	Errcode   int64    `json:"errcode"`
	Errmsg    string   `json:"errmsg"`
}

func (resp *UserProfile) IsError() bool {
	return resp.Errcode != 0
}
