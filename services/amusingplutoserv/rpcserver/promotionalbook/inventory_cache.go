package promotionalbook

import (
	"amusingx.fit/amusingx/services/amusingplutoserv/rpcserver/promotionalbook/inventorycachedatasource"
	"amusingx.fit/amusingx/services/amusingplutoserv/xredis"
	"context"
	"github.com/ItsWewin/superfactory/xerror"
	"strconv"
)

const (
	booksInventoryCacheKey = "pluto:service:promotion:books:inventory:cache:key"
)

type BookInventoryCache struct {
	key string
}

func NewBookInventoryCache() *BookInventoryCache {
	return &BookInventoryCache{key: booksInventoryCacheKey}
}

func (bc *BookInventoryCache) GetInventory(ctx context.Context, bookID int64) (int64, *xerror.Error) {
	result := xredis.Client.HGet(ctx, bc.key, strconv.FormatInt(bookID, 10))

	iv, err := result.Int64()
	if err != nil {
		return -1, xerror.NewErrorf(err, xerror.Code.SRedisExecuteErr, "")
	}

	return iv, nil
}

// redis load: book inventory
func (bc *BookInventoryCache) CacheInit(ctx context.Context) *xerror.Error {
	if xredis.Client == nil {
		return xerror.NewErrorf(nil, xerror.Code.SRedisExecuteErr, "redis client is nil")
	}

	xredis.Client.Del(ctx, bc.key)

	dataSource, xErr := inventorycachedatasource.NewDataSource(inventorycachedatasource.MysqlDBSource)
	if xErr != nil {
		return xErr
	}

	idInventoryMap, xErr := dataSource.GetInventories(ctx)
	if xErr != nil {
		return xErr
	}

	for id, iv := range idInventoryMap {
		err := xredis.Client.HSet(ctx, bc.key, id, iv).Err()
		if err != nil {
			return xerror.NewErrorf(err, xerror.Code.SRedisExecuteErr,
				"redis hset failed, key: %s, filed: %d, value: %d", bc.key, id, iv)
		}
	}

	return nil
}
