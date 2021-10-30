package plutoservice

import (
	"amusingx.fit/amusingx/protos/amusingplutoservice/plutoservice"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *PlutoService) InventoryUnlock(context.Context, *plutoservice.InventoryUnlockRequest) (*plutoservice.InventoryUnlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InventoryUnlock not implemented")
}
