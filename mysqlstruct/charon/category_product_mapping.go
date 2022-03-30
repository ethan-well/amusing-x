package charon

type CategoryProductMapping struct {
	ID         int64  `db:"id"`
	CategoryId int64  `db:"category_id"`
	ProductId  int64  `db:"product_id"`
	CreateTime string `db:"create_time"`
	UpdateTime string `db:"update_time"`
}
