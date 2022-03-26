package charon

import "time"

type Product struct {
	ID         int64     `db:"id"`
	Name       string    `db:"name"`
	Desc       string    `db:"description"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}
