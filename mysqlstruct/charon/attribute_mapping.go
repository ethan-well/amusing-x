package charon

import "time"

type AttributeMapping struct {
	ID           int64     `db:"id"`
	AttrId       int64     `db:"attr_id"`
	AttrValue    string    `db:"attr_value"`
	SubProductId int64     `db:"sub_product_id"`
	CreateTime   time.Time `db:"create_time"`
	UpdateTime   time.Time `db:"update_time"`
}
