package ganymede

type User struct {
	ID             int64  `db:"id"`
	Name           string `db:"name"`
	Login          string `db:"login"`
	PasswordDigest string `db:"password_digest"`
	CreateTime     string `db:"create_time"`
	UpdateTime     string `db:"update_time"`
}
