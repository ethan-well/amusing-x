package promotionalbook

import (
	"amusingx.fit/amusingx/etcd"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/xerror"
	"strconv"
	"time"
)

type PromotionalBook struct {
	BookID  int64
	lockKey string
}

func NewPromotionalBook(bookID int64) *PromotionalBook {
	return &PromotionalBook{BookID: bookID, lockKey: "/pluto:service:book:inventory:cache:upload:lock"}
}

func (pb *PromotionalBook) InventoryCacheInit(ctx context.Context) *xerror.Error {
	lock, err := etcd.Lock(pb.lockKey, 2*time.Second, 30)
	if err != nil {
		logger.Errorf("get lock failed")
		return err
	}
	defer lock.Unlock(ctx)

	err = NewBookInventoryCache().AllBookInventoryCacheInit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (pb *PromotionalBook) InventoryQuery(ctx context.Context) (int64, *xerror.Error) {
	if pb.BookID <= 0 {
		return -1, xerror.NewErrorf(nil, xerror.Code.CParamsError, "book id is invalid")
	}

	iv, err := NewBookInventoryCache().GetInventory(ctx, pb.BookID)
	if err != nil { // 是可能是没有缓存导致的失败
		return pb.trySetInventory(ctx, pb.BookID) // 尝试再吃缓存
	}

	return iv, nil
}

func (pb *PromotionalBook) InventoryLock(ctx context.Context, lockCount int) (int64, *xerror.Error) {
	cache := NewBookInventoryCache()
	iv, err := cache.InventoryDecrBy(ctx, pb.BookID, lockCount) // 获取到锁后，先尝试再次获取
	if err != nil {
		return -1, err
	}

	return iv, nil
}

func (pb *PromotionalBook) InventoryUnlock(ctx context.Context) (int64, *xerror.Error) {
	return -1, nil
}

// 如果是缓存过期，先获取分布式锁，尝试再次获取
// 没有获取到再 hset
func (pb *PromotionalBook) trySetInventory(ctx context.Context, bookID int64) (int64, *xerror.Error) {
	key := pb.lockKey + ":" + strconv.FormatInt(bookID, 64)

	lock, err := etcd.Lock(key, 2*time.Second, 5)
	if err != nil {
		return -1, xerror.NewErrorf(err, xerror.Code.SUnexpectedErr, "etcd lock failed")
	}
	defer lock.Unlock(ctx)

	cache := NewBookInventoryCache()
	iv, err := cache.GetInventory(ctx, pb.BookID) // 获取到锁后，先尝试再次获取
	if err == nil {
		return iv, nil
	}

	return cache.SetInventory(ctx, bookID) // 一个 key 最多出现一次缓存击穿
}