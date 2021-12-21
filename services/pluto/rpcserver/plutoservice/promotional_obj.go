package plutoservice

import (
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

type PromotionalObj interface {
	InventoryCacheInit(context.Context) aerror.Error
	InventoryQuery(context.Context) (int64, aerror.Error)
	InventoryLock(context.Context, int) (int64, aerror.Error)
	InventoryUnlock(context.Context, int) (int64, aerror.Error)
}
