package charon

import "time"

type ProductImage struct {
	Id           int64     `db:"id"`
	ProductId    int64     `db:"product_id"`
	ProductLevel int       `db:"product_level"`
	Url          string    `db:"url"`
	UploaderType string    `db:"uploader_type"`
	CreateTime   time.Time `db:"create_time"`
	UpdateTime   time.Time `db:"update_time"`
}
