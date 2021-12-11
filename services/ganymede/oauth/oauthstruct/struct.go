package oauthstruct

const (
	AccessTokenExpireNever = -1
)

type AccessToken struct {
	AccessToken  string
	Scope        string
	TokenType    string
	ExpiresIn    int64 // -1 表示永不过期
	RefreshToken string
	OpenID       string
	UnionID      string
}

type UserProfile struct {
	Login       string // 地方方网站唯一标识
	OuterUserID int64  // 用户外部 id
	AvatarUrl   string
	Name        string
	Company     string
	Blog        string
	Email       string
	Location    string
}
