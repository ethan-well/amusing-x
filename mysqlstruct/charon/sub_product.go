package charon

import "time"

type SubProduct struct {
	ID         int64     `db:"id"`
	Name       string    `db:"name"`
	Desc       string    `db:"description"`
	ProductId  int64     `db:"product_id"`
	Currency   string    `db:"currency"`
	Price      int64     `db:"price"`
	Stock      int64     `db:"stock"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}
