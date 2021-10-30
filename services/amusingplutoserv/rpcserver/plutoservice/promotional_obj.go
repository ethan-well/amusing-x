package plutoservice

import (
	"context"
	"github.com/ItsWewin/superfactory/xerror"
)

type PromotionalObj interface {
	InventoryCacheInit(ctx context.Context) *xerror.Error
	InventoryQuery(ctx context.Context) (int64, *xerror.Error)
	InventoryLock(ctx context.Context) (int64, *xerror.Error)
	InventoryUnlock(ctx context.Context) (int64, *xerror.Error)
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
