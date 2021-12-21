package promotionalbook

import (
	"amusingx.fit/amusingx/etcd"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/logger"
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

func (pb *PromotionalBook) InventoryCacheInit(ctx context.Context) aerror.Error {
	lock, err := etcd.Lock(pb.lockKey, 1*time.Second, 30)
	if err != nil {
		logger.Errorf("get lock failed")
		return err
	}
	logger.Infof("get lock succed, key: %s", pb.lockKey)
	defer lock.Unlock(ctx)

	err = NewBookInventoryCache().AllBookInventoryCacheInit(ctx)
	if err != nil {
		logger.Errorf("book inventory cache init failed: %s", err)
		return err
	}

	return nil
}

func (pb *PromotionalBook) InventoryQuery(ctx context.Context) (int64, aerror.Error) {
	if pb.BookID <= 0 {
		return -1, aerror.NewErrorf(nil, aerror.Code.CParamsError, "book id is invalid")
	}

	iv, err := NewBookInventoryCache().GetInventory(ctx, pb.BookID)
	if err != nil { // 是可能是没有缓存导致的失败
		return pb.trySetInventory(ctx, pb.BookID) // 尝试再吃缓存
	}

	return iv, nil
}

func (pb *PromotionalBook) InventoryLock(ctx context.Context, lockCount int) (int64, aerror.Error) {
	cache := NewBookInventoryCache()
	iv, err := cache.InventoryDecrBy(ctx, pb.BookID, lockCount) // 获取到锁后，先尝试再次获取
	if err != nil {
		return -1, err
	}

	return iv, nil
}

func (pb *PromotionalBook) InventoryUnlock(ctx context.Context, unLockCount int) (int64, aerror.Error) {
	cache := NewBookInventoryCache()
	iv, err := cache.InventoryIncrBy(ctx, pb.BookID, unLockCount) // 获取到锁后，先尝试再次获取
	if err != nil {
		return -1, err
	}

	return iv, nil
}

// 如果是缓存过期，先获取分布式锁，尝试再次获取
// 没有获取到再 hset
func (pb *PromotionalBook) trySetInventory(ctx context.Context, bookID int64) (int64, aerror.Error) {
	key := pb.lockKey + ":" + strconv.FormatInt(bookID, 64)

	lock, err := etcd.Lock(key, 2*time.Second, 5)
	if err != nil {
		return -1, aerror.NewErrorf(err, aerror.Code.SUnexpectedErr, "etcd lock failed")
	}
	defer lock.Unlock(ctx)

	cache := NewBookInventoryCache()
	iv, err := cache.GetInventory(ctx, pb.BookID) // 获取到锁后，先尝试再次获取
	if err == nil {
		return iv, nil
	}

	return cache.SetInventory(ctx, bookID) // 一个 key 最多出现一次缓存击穿
}
