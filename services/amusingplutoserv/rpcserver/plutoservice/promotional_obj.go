package plutoservice

import (
	"context"
	"github.com/ItsWewin/superfactory/xerror"
)

type PromotionalObj interface {
	InventoryCacheInit(context.Context) *xerror.Error
	InventoryQuery(context.Context) (int64, *xerror.Error)
	InventoryLock(context.Context, int) (int64, *xerror.Error)
	InventoryUnlock(context.Context, int) (int64, *xerror.Error)
}
