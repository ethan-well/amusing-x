package model

import (
	"amusingx.fit/amusingx/services/pluto/mysql"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/logger"
	"time"
)

type BookInventory struct {
	ID                 int64     `db:"id"`
	BookID             int64     `db:"book_id"`
	RealInventory      int64     `db:"real_inventory"`
	AvailableInventory int64     `db:"available_inventory"`
	LockedInventory    int64     `db:"locked_inventory"`
	CreateTime         time.Time `db:"create_time"`
	UpdateTime         time.Time `db:"update_time"`
}

func QueryBookInventoryByLimit(ctx context.Context, offSet, limit int64) ([]*BookInventory, aerror.Error) {
	sqlStr := `SELECT id, book_id, real_inventory, available_inventory, locked_inventory
		FROM book_inventory
		ORDER BY id
		LIMIT ?, ?
	`

	if mysql.PlutoDB == nil {
		logger.Errorf("pluto db is nil")
	}

	var ivs []*BookInventory
	err := mysql.PlutoDB.SelectContext(ctx, &ivs, sqlStr, offSet, limit)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "query book_inventory failed")
	}

	return ivs, nil
}

func QueryBookInventoryByID(ctx context.Context, bookID int64) (*BookInventory, aerror.Error) {
	sqlStr := `SELECT id, book_id, real_inventory, available_inventory, locked_inventory
		FROM book_inventory
		where book_id = ?
	`

	if mysql.PlutoDB == nil {
		logger.Errorf("pluto db is nil")
	}

	var ivs BookInventory
	err := mysql.PlutoDB.GetContext(ctx, &ivs, sqlStr, bookID)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "query book_inventory failed")
	}

	return &ivs, nil
}
