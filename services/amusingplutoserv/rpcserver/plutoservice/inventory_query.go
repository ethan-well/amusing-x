package plutoservice

import (
	"amusingx.fit/amusingx/protos/amusingplutoservice/plutoservice"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *PlutoService) InventoryQuery(context.Context, *plutoservice.InventoryQueryRequest) (*plutoservice.InventoryQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InventoryQuery not implemented")
}
