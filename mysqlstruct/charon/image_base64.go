package charon

import "time"

type ProductImageBase64 struct {
	Id         int64     `db:"id"`
	Base64     string    `db:"base64"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}
