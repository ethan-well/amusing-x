package charon

type ProductStock struct {
	ID                 int64 `db:"id"`
	SubProductId       int64 `db:"sub_product_id"`
	RealInventory      int64 `db:"real_inventory"`
	AvailableInventory int64 `db:"available_inventory"`
}
