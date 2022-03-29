package charon

import "time"

type Product struct {
	ID         int64     `db:"id"`
	Name       string    `db:"name"`
	Desc       string    `db:"description"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}

type ProductWide struct {
	ID         int64  `db:"id"`
	CategoryId int64  `db:"category_id"`
	Name       string `db:"name"`
	Desc       string `db:"description"`
	CreateTime string `db:"create_time"`
	UpdateTime string `db:"update_time"`
}
