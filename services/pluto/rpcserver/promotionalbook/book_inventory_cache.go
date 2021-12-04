package promotionalbook

import (
	"amusingx.fit/amusingx/services/pluto/rpcserver/promotionalbook/inventorycachedatasource"
	"amusingx.fit/amusingx/services/pluto/xredis"
	"context"
	"fmt"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/xerror"
	"github.com/go-redis/redis/v8"
	"strconv"
)

const (
	booksInventoryCacheKey = "pluto:service:promotion:books:inventory:cache:key"
	// -1 参数错误
	InventoryDecrScriptCodeParamsErr = -1
	// -2 库存不足
	InventoryDecrScriptCodeInventoryShortage = -2
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

func (bc *BookInventoryCache) SetInventory(ctx context.Context, bookID int64, inventory ...int64) (int64, *xerror.Error) {
	if xredis.Client == nil {
		return -1, xerror.NewErrorf(nil, xerror.Code.SRedisExecuteErr, "redis client is nil")
	}

	xredis.Client.Del(ctx, bc.key)

	var iv int64
	if len(inventory) > 0 {
		iv = inventory[0]
	} else {
		dataSource, xErr := inventorycachedatasource.NewDataSource(inventorycachedatasource.LocalVariableSource)
		if xErr != nil {
			return -1, xErr
		}

		idInventoryMap, xErr := dataSource.GetInventoriesByID(ctx, bookID)
		if xErr != nil {
			return -1, xErr
		}

		var ok bool
		iv, ok = idInventoryMap[bookID]
		if !ok {
			iv = 0
		}
	}

	err := xredis.Client.HSet(ctx, bc.key, bookID, iv).Err()
	if err != nil {
		return -1, xerror.NewErrorf(err, xerror.Code.SRedisExecuteErr, "redis hset failed, key: %s, filed: %d, value: %d", bc.key, bookID, iv)
	}

	return iv, nil
}

// redis load: book inventory
func (bc *BookInventoryCache) AllBookInventoryCacheInit(ctx context.Context) *xerror.Error {
	if xredis.Client == nil {
		return xerror.NewErrorf(nil, xerror.Code.SRedisExecuteErr, "redis client is nil")
	}
	
	xredis.Client.Del(ctx, bc.key)
	
	dataSource, xErr := inventorycachedatasource.NewDataSource(inventorycachedatasource.LocalVariableSource)
	if xErr != nil {
		return xErr
	}

	idInventoryMap, xErr := dataSource.GetInventories(ctx)
	if xErr != nil {
		return xErr
	}

	m := map[int64]*redis.IntCmd{}
	pipe := xredis.Client.Pipeline()
	for id, iv := range idInventoryMap {
		m[id] = pipe.HSet(ctx, bc.key, id, iv)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		return xerror.NewErrorf(err, xerror.Code.SRedisExecuteErr, "Redis command exec failed")
	}

	errMap := map[int64]error{}
	for k, v := range m {
		err := v.Err()
		if err != nil {
			errMap[k] = err
		}
	}

	if len(errMap) != 0 {
		return xerror.NewErrorf(err, xerror.Code.SRedisExecuteErr, fmt.Sprintf("%v", errMap))
	}

	return nil
}

func (bc *BookInventoryCache) InventoryDecrBy(ctx context.Context, bookID int64, decr int) (int64, *xerror.Error) {
	// -1 参数错误
	// -2 库存不足
	// >= 0 操作成功，返回剩余库存
	logger.Infof("InventoryDecrBy, key: %s, bookID: %d, Decr: %d", bc.key, bookID, decr)
	var script = `
local incr = tonumber(ARGV[2])
local inventory = tonumber(redis.call('HGET', KEYS[1], ARGV[1]))

if incr < 0 then
	return -1 
end

if inventory >= incr then
	return redis.call('HINCRBY', KEYS[1], ARGV[1], -1 * incr)
else
	return -2
end
`
	cmd := xredis.Client.Eval(ctx, script, []string{bc.key}, bookID, decr)
	r, err := cmd.Result()
	if err != nil {
		return 0, xerror.NewError(err, xerror.Code.SRedisExecuteErr, "减库存失败")
	}
	code, ok := r.(int64)
	if !ok {
		return 0, xerror.NewErrorf(err, xerror.Code.SRedisExecuteErr, "减库存失败")
	}
	
	switch code {
	case InventoryDecrScriptCodeParamsErr:
		return 0, xerror.NewErrorf(err, xerror.Code.SRedisExecuteErr, "参数错误")
	case InventoryDecrScriptCodeInventoryShortage:
		return 0, xerror.NewErrorf(err, xerror.Code.SRedisExecuteErr, "库存不足")
	default:
		logger.Infof("减库存成功，剩余库存 %d", code)
		return code, nil
	}
}

func (bc *BookInventoryCache) InventoryIncrBy(ctx context.Context, bookID int64, incr int) (int64, *xerror.Error) {
	// -1 参数错误
	// -2 库存不足
	// >= 0 操作成功，返回剩余库存
	var script = `
local incr = tonumber(ARGV[2])
local inventory = tonumber(redis.call('HGET', KEYS[1], ARGV[1]))

if incr < 0 then
	return -1 
end

if inventory < 0 then
	return -2
end

return redis.call('HINCRBY', KEYS[1], ARGV[1], incr)
`
	cmd := xredis.Client.Eval(ctx, script, []string{bc.key}, bookID, incr)
	r, err := cmd.Result()
	if err != nil {
		return -1, xerror.NewError(err, xerror.Code.SRedisExecuteErr, "加库存失败")
	}
	code, ok := r.(int64)
	if !ok {
		return -1, xerror.NewErrorf(err, xerror.Code.SRedisExecuteErr, "加库存失败")
	}

	switch code {
	case InventoryDecrScriptCodeParamsErr:
		return -1, xerror.NewErrorf(err, xerror.Code.SRedisExecuteErr, "参数错误")
	case InventoryDecrScriptCodeInventoryShortage:
		return -1, xerror.NewErrorf(err, xerror.Code.SRedisExecuteErr, "原本库存小于 0")
	default:
		return code, nil
	}
}
