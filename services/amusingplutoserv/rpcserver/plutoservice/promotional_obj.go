package plutoservice

import (
	"context"
	"github.com/ItsWewin/superfactory/xerror"
)

type PromotionalObj interface {
	InventoryCacheInit(context.Context) *xerror.Error
	InventoryQuery(context.Context) (int64, *xerror.Error)
	InventoryLock(context.Context, int) (int64, *xerror.Error)
	InventoryUnlock(context.Context) (int64, *xerror.Error)
}

type UnimplementedPromotionalObj struct {
}

func (*UnimplementedPromotionalObj) InventoryQuery(ctx context.Context) (int64, *xerror.Error) {
	return 0, xerror.NewErrorf(nil, xerror.Code.SUnimplementedInterface, "unimplemented method: InventoryQuery")
}
func (*UnimplementedPromotionalObj) InventoryLock(ctx context.Context) (int64, *xerror.Error) {
	return 0, xerror.NewErrorf(nil, xerror.Code.SUnimplementedInterface, "unimplemented method: InventoryLock")
}
func (*UnimplementedPromotionalObj) InventoryUnlock(ctx context.Context) (int64, *xerror.Error) {
	return 0, xerror.NewErrorf(nil, xerror.Code.SUnimplementedInterface, "unimplemented method: InventoryUnlock")
}
