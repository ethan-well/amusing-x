package ganymede

type OauthInfo struct {
	ID          int64  `db:"id"`
	UseID       int64  `db:"user_id"`
	Provider    string `db:"provider"`
	OuterID     int64  `db:"outer_id"`
	Login       string `db:"login"`
	AvatarUrl   string `db:"avatar_url"`
	Email       string `db:"email"`
	AccessToken string `db:"access_token"`
	Code        string `db:"code"`
	CreateTime  string `db:"create_time"`
	UpdateTime  string `db:"update_time"`
}
