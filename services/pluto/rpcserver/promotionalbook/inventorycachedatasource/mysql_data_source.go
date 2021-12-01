package inventorycachedatasource

import (
	"amusingx.fit/amusingx/services/pluto/mysql/plutodb/model"
	"context"
	"github.com/ItsWewin/superfactory/xerror"
)

type MysqlDataSource struct {
}

func NewRedisDataSource() *MysqlDataSource {
	return &MysqlDataSource{}
}

func (s *MysqlDataSource) GetInventories(ctx context.Context) (map[int64]int64, *xerror.Error) {
	offset := 0
	limit := 10000
	idInventoryMap := make(map[int64]int64)
	for offset != -1 {
		ivs, err := model.QueryBookInventoryByLimit(ctx, 0, 1000)
		if err != nil {
			offset = -1
			return nil, err
		}

		for _, iv := range ivs {
			idInventoryMap[iv.BookID] = iv.RealInventory
		}

		if len(ivs) < limit || len(ivs) == 0 {
			offset = -1
		}
	}

	return idInventoryMap, nil
}

func (s *MysqlDataSource) GetInventoriesByID(ctx context.Context, bookID int64) (map[int64]int64, *xerror.Error) {
	ivs, err := model.QueryBookInventoryByID(ctx, bookID)
	if err != nil {
		return nil, err
	}

	if ivs == nil {
		return map[int64]int64{bookID: 0}, nil
	}

	return map[int64]int64{ivs.ID: ivs.AvailableInventory}, nil
}
