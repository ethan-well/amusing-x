package ganymede

type Role struct {
	ID      int64  `db:"id"`
	Name    string `db:"name"`
	Service string `db:"service"`
	Desc    string `db:"desc"`
}
