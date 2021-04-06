package amusinguser

type User struct {
	ID             int64  `db:"id"`
	Nickname       string `db:"nickname"`
	Phone          string `db:"phone"`
	PasswordDigest string `db:"password_digest"`
	CreateTime     string `db:"create_time"`
	UpdateTime     string `db:"update_time"`
	Salt           string `db:"salt"`
}
