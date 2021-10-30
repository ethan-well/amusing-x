package promotionalbook

import (
	"context"
	"github.com/ItsWewin/superfactory/xerror"
)

type PromotionalBook struct {
	//plutoservice.UnimplementedPromotionalObj

	BookID int64
}

func NewPromotionalBook(bookID int64) *PromotionalBook {
	return &PromotionalBook{BookID: bookID}
}

func (pb *PromotionalBook) InventoryCacheInit(ctx context.Context) *xerror.Error {
	err := NewBookInventoryCache().CacheInit(ctx)
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
	if err != nil {
		return -1, err
	}

	return iv, nil
}

func (pb *PromotionalBook) InventoryLock(ctx context.Context) (int64, *xerror.Error) {
	return -1, nil
}

func (pb *PromotionalBook) InventoryUnlock(ctx context.Context) (int64, *xerror.Error) {
	return -1, nil
}
