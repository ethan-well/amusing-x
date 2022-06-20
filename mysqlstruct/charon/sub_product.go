package charon

import "time"

const (
	ProductLevelProduct = iota
	ProductLevelSubProduct
)

type SubProduct struct {
	ID         int64     `db:"id"`
	Name       string    `db:"name"`
	Desc       string    `db:"description"`
	ProductId  int64     `db:"product_id"`
	Currency   string    `db:"currency"`
	Price      int64     `db:"price"`
	Stock      int64     `db:"stock"`
	MinNum     int64     `db:"min_num"`
	MaxNum     int64     `db:"max_num"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}

type SubProductWithProductStock struct {
	ID                 int64     `db:"id"`
	Name               string    `db:"name"`
	Desc               string    `db:"description"`
	ProductId          int64     `db:"product_id"`
	Currency           string    `db:"currency"`
	Price              int64     `db:"price"`
	Stock              int64     `db:"stock"`
	MinNum             int64     `db:"min_num"`
	MaxNum             int64     `db:"max_num"`
	RealInventory      int64     `db:"real_inventory"`
	AvailableInventory int64     `db:"available_inventory"`
	CreateTime         time.Time `db:"create_time"`
	UpdateTime         time.Time `db:"update_time"`
}
