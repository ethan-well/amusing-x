package charon

import "time"

type Category struct {
	ID         int64     `db:"id"`
	ParentID   int64     `db:"parent_id"`
	Name       string    `db:"name"`
	Desc       string    `db:"description"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}

type CategoryWide struct {
	ID         int64     `db:"id"`
	ParentID   int64     `db:"parent_id"`
	Name       string    `db:"name"`
	Desc       string    `db:"description"`
	ProductId  int64     `db:"product_id"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}
