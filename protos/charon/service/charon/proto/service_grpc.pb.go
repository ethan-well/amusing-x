// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package charonservice

import (
	proto "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CharonServClient is the client API for CharonServ service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CharonServClient interface {
	Pong(ctx context.Context, in *BlankParams, opts ...grpc.CallOption) (*PongResponse, error)
	Create(ctx context.Context, in *CategoryCreateRequest, opts ...grpc.CallOption) (*CategoryCreateResponse, error)
	Categories(ctx context.Context, in *proto.CategoryListRequest, opts ...grpc.CallOption) (*proto.CategoryListResponse, error)
	Delete(ctx context.Context, in *CategoryDeleteRequest, opts ...grpc.CallOption) (*CategoryDeleteResponse, error)
	Category(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error)
	Update(ctx context.Context, in *CategoryUpdateRequest, opts ...grpc.CallOption) (*CategoryUpdateResponse, error)
	CategoriesDelete(ctx context.Context, in *proto.CategoriesDeleteRequest, opts ...grpc.CallOption) (*proto.CategoriesDeleteResponse, error)
	ProductCreate(ctx context.Context, in *proto.ProductCreateRequest, opts ...grpc.CallOption) (*proto.Product, error)
	ProductDelete(ctx context.Context, in *proto.ProductDeleteRequest, opts ...grpc.CallOption) (*proto.ProductDeleteResponse, error)
	ProductsDelete(ctx context.Context, in *proto.ProductsDeleteRequest, opts ...grpc.CallOption) (*proto.ProductsDeleteResponse, error)
	Products(ctx context.Context, in *proto.ProductListRequest, opts ...grpc.CallOption) (*proto.ProductListResponse, error)
	Product(ctx context.Context, in *proto.ProductRequest, opts ...grpc.CallOption) (*proto.Product, error)
	ProductUpdate(ctx context.Context, in *proto.ProductUpdateRequest, opts ...grpc.CallOption) (*proto.Product, error)
	SubProductCreate(ctx context.Context, in *proto.SubProductCreateRequest, opts ...grpc.CallOption) (*proto.SubProduct, error)
	SubProductDelete(ctx context.Context, in *proto.SubProductDeleteRequest, opts ...grpc.CallOption) (*proto.SubProductDeleteResponse, error)
	SubProductsDelete(ctx context.Context, in *proto.SubProductsDeleteRequest, opts ...grpc.CallOption) (*proto.SubProductsDeleteResponse, error)
	SubProducts(ctx context.Context, in *proto.SubProductListRequest, opts ...grpc.CallOption) (*proto.SubProductListResponse, error)
	SubProduct(ctx context.Context, in *proto.SubProductRequest, opts ...grpc.CallOption) (*proto.SubProduct, error)
	SubProductUpdate(ctx context.Context, in *proto.SubProductUpdateRequest, opts ...grpc.CallOption) (*proto.SubProduct, error)
	AttributeCreate(ctx context.Context, in *proto.AttributeCreateRequest, opts ...grpc.CallOption) (*proto.Attribute, error)
	AttributeDelete(ctx context.Context, in *proto.AttributeDeleteRequest, opts ...grpc.CallOption) (*proto.AttributeDeleteResponse, error)
	Attributes(ctx context.Context, in *proto.AttributeListRequest, opts ...grpc.CallOption) (*proto.AttributeListResponse, error)
	Attribute(ctx context.Context, in *proto.AttributeRequest, opts ...grpc.CallOption) (*proto.Attribute, error)
	AttributeUpdate(ctx context.Context, in *proto.AttributeUpdateRequest, opts ...grpc.CallOption) (*proto.Attribute, error)
}

type charonServClient struct {
	cc grpc.ClientConnInterface
}

func NewCharonServClient(cc grpc.ClientConnInterface) CharonServClient {
	return &charonServClient{cc}
}

func (c *charonServClient) Pong(ctx context.Context, in *BlankParams, opts ...grpc.CallOption) (*PongResponse, error) {
	out := new(PongResponse)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/Pong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charonServClient) Create(ctx context.Context, in *CategoryCreateRequest, opts ...grpc.CallOption) (*CategoryCreateResponse, error) {
	out := new(CategoryCreateResponse)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charonServClient) Categories(ctx context.Context, in *proto.CategoryListRequest, opts ...grpc.CallOption) (*proto.CategoryListResponse, error) {
	out := new(proto.CategoryListResponse)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/Categories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charonServClient) Delete(ctx context.Context, in *CategoryDeleteRequest, opts ...grpc.CallOption) (*CategoryDeleteResponse, error) {
	out := new(CategoryDeleteResponse)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charonServClient) Category(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error) {
	out := new(CategoryResponse)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/Category", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charonServClient) Update(ctx context.Context, in *CategoryUpdateRequest, opts ...grpc.CallOption) (*CategoryUpdateResponse, error) {
	out := new(CategoryUpdateResponse)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charonServClient) CategoriesDelete(ctx context.Context, in *proto.CategoriesDeleteRequest, opts ...grpc.CallOption) (*proto.CategoriesDeleteResponse, error) {
	out := new(proto.CategoriesDeleteResponse)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/CategoriesDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charonServClient) ProductCreate(ctx context.Context, in *proto.ProductCreateRequest, opts ...grpc.CallOption) (*proto.Product, error) {
	out := new(proto.Product)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/ProductCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charonServClient) ProductDelete(ctx context.Context, in *proto.ProductDeleteRequest, opts ...grpc.CallOption) (*proto.ProductDeleteResponse, error) {
	out := new(proto.ProductDeleteResponse)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/ProductDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charonServClient) ProductsDelete(ctx context.Context, in *proto.ProductsDeleteRequest, opts ...grpc.CallOption) (*proto.ProductsDeleteResponse, error) {
	out := new(proto.ProductsDeleteResponse)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/ProductsDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charonServClient) Products(ctx context.Context, in *proto.ProductListRequest, opts ...grpc.CallOption) (*proto.ProductListResponse, error) {
	out := new(proto.ProductListResponse)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/Products", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charonServClient) Product(ctx context.Context, in *proto.ProductRequest, opts ...grpc.CallOption) (*proto.Product, error) {
	out := new(proto.Product)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/Product", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charonServClient) ProductUpdate(ctx context.Context, in *proto.ProductUpdateRequest, opts ...grpc.CallOption) (*proto.Product, error) {
	out := new(proto.Product)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/ProductUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charonServClient) SubProductCreate(ctx context.Context, in *proto.SubProductCreateRequest, opts ...grpc.CallOption) (*proto.SubProduct, error) {
	out := new(proto.SubProduct)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/SubProductCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charonServClient) SubProductDelete(ctx context.Context, in *proto.SubProductDeleteRequest, opts ...grpc.CallOption) (*proto.SubProductDeleteResponse, error) {
	out := new(proto.SubProductDeleteResponse)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/SubProductDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charonServClient) SubProductsDelete(ctx context.Context, in *proto.SubProductsDeleteRequest, opts ...grpc.CallOption) (*proto.SubProductsDeleteResponse, error) {
	out := new(proto.SubProductsDeleteResponse)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/SubProductsDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charonServClient) SubProducts(ctx context.Context, in *proto.SubProductListRequest, opts ...grpc.CallOption) (*proto.SubProductListResponse, error) {
	out := new(proto.SubProductListResponse)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/SubProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charonServClient) SubProduct(ctx context.Context, in *proto.SubProductRequest, opts ...grpc.CallOption) (*proto.SubProduct, error) {
	out := new(proto.SubProduct)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/SubProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charonServClient) SubProductUpdate(ctx context.Context, in *proto.SubProductUpdateRequest, opts ...grpc.CallOption) (*proto.SubProduct, error) {
	out := new(proto.SubProduct)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/SubProductUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charonServClient) AttributeCreate(ctx context.Context, in *proto.AttributeCreateRequest, opts ...grpc.CallOption) (*proto.Attribute, error) {
	out := new(proto.Attribute)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/AttributeCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charonServClient) AttributeDelete(ctx context.Context, in *proto.AttributeDeleteRequest, opts ...grpc.CallOption) (*proto.AttributeDeleteResponse, error) {
	out := new(proto.AttributeDeleteResponse)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/AttributeDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charonServClient) Attributes(ctx context.Context, in *proto.AttributeListRequest, opts ...grpc.CallOption) (*proto.AttributeListResponse, error) {
	out := new(proto.AttributeListResponse)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/Attributes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charonServClient) Attribute(ctx context.Context, in *proto.AttributeRequest, opts ...grpc.CallOption) (*proto.Attribute, error) {
	out := new(proto.Attribute)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/Attribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charonServClient) AttributeUpdate(ctx context.Context, in *proto.AttributeUpdateRequest, opts ...grpc.CallOption) (*proto.Attribute, error) {
	out := new(proto.Attribute)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/AttributeUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CharonServServer is the server API for CharonServ service.
// All implementations must embed UnimplementedCharonServServer
// for forward compatibility
type CharonServServer interface {
	Pong(context.Context, *BlankParams) (*PongResponse, error)
	Create(context.Context, *CategoryCreateRequest) (*CategoryCreateResponse, error)
	Categories(context.Context, *proto.CategoryListRequest) (*proto.CategoryListResponse, error)
	Delete(context.Context, *CategoryDeleteRequest) (*CategoryDeleteResponse, error)
	Category(context.Context, *CategoryRequest) (*CategoryResponse, error)
	Update(context.Context, *CategoryUpdateRequest) (*CategoryUpdateResponse, error)
	CategoriesDelete(context.Context, *proto.CategoriesDeleteRequest) (*proto.CategoriesDeleteResponse, error)
	ProductCreate(context.Context, *proto.ProductCreateRequest) (*proto.Product, error)
	ProductDelete(context.Context, *proto.ProductDeleteRequest) (*proto.ProductDeleteResponse, error)
	ProductsDelete(context.Context, *proto.ProductsDeleteRequest) (*proto.ProductsDeleteResponse, error)
	Products(context.Context, *proto.ProductListRequest) (*proto.ProductListResponse, error)
	Product(context.Context, *proto.ProductRequest) (*proto.Product, error)
	ProductUpdate(context.Context, *proto.ProductUpdateRequest) (*proto.Product, error)
	SubProductCreate(context.Context, *proto.SubProductCreateRequest) (*proto.SubProduct, error)
	SubProductDelete(context.Context, *proto.SubProductDeleteRequest) (*proto.SubProductDeleteResponse, error)
	SubProductsDelete(context.Context, *proto.SubProductsDeleteRequest) (*proto.SubProductsDeleteResponse, error)
	SubProducts(context.Context, *proto.SubProductListRequest) (*proto.SubProductListResponse, error)
	SubProduct(context.Context, *proto.SubProductRequest) (*proto.SubProduct, error)
	SubProductUpdate(context.Context, *proto.SubProductUpdateRequest) (*proto.SubProduct, error)
	AttributeCreate(context.Context, *proto.AttributeCreateRequest) (*proto.Attribute, error)
	AttributeDelete(context.Context, *proto.AttributeDeleteRequest) (*proto.AttributeDeleteResponse, error)
	Attributes(context.Context, *proto.AttributeListRequest) (*proto.AttributeListResponse, error)
	Attribute(context.Context, *proto.AttributeRequest) (*proto.Attribute, error)
	AttributeUpdate(context.Context, *proto.AttributeUpdateRequest) (*proto.Attribute, error)
	mustEmbedUnimplementedCharonServServer()
}

// UnimplementedCharonServServer must be embedded to have forward compatible implementations.
type UnimplementedCharonServServer struct {
}

func (UnimplementedCharonServServer) Pong(context.Context, *BlankParams) (*PongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pong not implemented")
}
func (UnimplementedCharonServServer) Create(context.Context, *CategoryCreateRequest) (*CategoryCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCharonServServer) Categories(context.Context, *proto.CategoryListRequest) (*proto.CategoryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Categories not implemented")
}
func (UnimplementedCharonServServer) Delete(context.Context, *CategoryDeleteRequest) (*CategoryDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCharonServServer) Category(context.Context, *CategoryRequest) (*CategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Category not implemented")
}
func (UnimplementedCharonServServer) Update(context.Context, *CategoryUpdateRequest) (*CategoryUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCharonServServer) CategoriesDelete(context.Context, *proto.CategoriesDeleteRequest) (*proto.CategoriesDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoriesDelete not implemented")
}
func (UnimplementedCharonServServer) ProductCreate(context.Context, *proto.ProductCreateRequest) (*proto.Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductCreate not implemented")
}
func (UnimplementedCharonServServer) ProductDelete(context.Context, *proto.ProductDeleteRequest) (*proto.ProductDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductDelete not implemented")
}
func (UnimplementedCharonServServer) ProductsDelete(context.Context, *proto.ProductsDeleteRequest) (*proto.ProductsDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductsDelete not implemented")
}
func (UnimplementedCharonServServer) Products(context.Context, *proto.ProductListRequest) (*proto.ProductListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Products not implemented")
}
func (UnimplementedCharonServServer) Product(context.Context, *proto.ProductRequest) (*proto.Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Product not implemented")
}
func (UnimplementedCharonServServer) ProductUpdate(context.Context, *proto.ProductUpdateRequest) (*proto.Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductUpdate not implemented")
}
func (UnimplementedCharonServServer) SubProductCreate(context.Context, *proto.SubProductCreateRequest) (*proto.SubProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubProductCreate not implemented")
}
func (UnimplementedCharonServServer) SubProductDelete(context.Context, *proto.SubProductDeleteRequest) (*proto.SubProductDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubProductDelete not implemented")
}
func (UnimplementedCharonServServer) SubProductsDelete(context.Context, *proto.SubProductsDeleteRequest) (*proto.SubProductsDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubProductsDelete not implemented")
}
func (UnimplementedCharonServServer) SubProducts(context.Context, *proto.SubProductListRequest) (*proto.SubProductListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubProducts not implemented")
}
func (UnimplementedCharonServServer) SubProduct(context.Context, *proto.SubProductRequest) (*proto.SubProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubProduct not implemented")
}
func (UnimplementedCharonServServer) SubProductUpdate(context.Context, *proto.SubProductUpdateRequest) (*proto.SubProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubProductUpdate not implemented")
}
func (UnimplementedCharonServServer) AttributeCreate(context.Context, *proto.AttributeCreateRequest) (*proto.Attribute, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttributeCreate not implemented")
}
func (UnimplementedCharonServServer) AttributeDelete(context.Context, *proto.AttributeDeleteRequest) (*proto.AttributeDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttributeDelete not implemented")
}
func (UnimplementedCharonServServer) Attributes(context.Context, *proto.AttributeListRequest) (*proto.AttributeListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Attributes not implemented")
}
func (UnimplementedCharonServServer) Attribute(context.Context, *proto.AttributeRequest) (*proto.Attribute, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Attribute not implemented")
}
func (UnimplementedCharonServServer) AttributeUpdate(context.Context, *proto.AttributeUpdateRequest) (*proto.Attribute, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttributeUpdate not implemented")
}
func (UnimplementedCharonServServer) mustEmbedUnimplementedCharonServServer() {}

// UnsafeCharonServServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CharonServServer will
// result in compilation errors.
type UnsafeCharonServServer interface {
	mustEmbedUnimplementedCharonServServer()
}

func RegisterCharonServServer(s grpc.ServiceRegistrar, srv CharonServServer) {
	s.RegisterService(&CharonServ_ServiceDesc, srv)
}

func _CharonServ_Pong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlankParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).Pong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/Pong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).Pong(ctx, req.(*BlankParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharonServ_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).Create(ctx, req.(*CategoryCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharonServ_Categories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.CategoryListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).Categories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/Categories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).Categories(ctx, req.(*proto.CategoryListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharonServ_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).Delete(ctx, req.(*CategoryDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharonServ_Category_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).Category(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/Category",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).Category(ctx, req.(*CategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharonServ_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).Update(ctx, req.(*CategoryUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharonServ_CategoriesDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.CategoriesDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).CategoriesDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/CategoriesDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).CategoriesDelete(ctx, req.(*proto.CategoriesDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharonServ_ProductCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.ProductCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).ProductCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/ProductCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).ProductCreate(ctx, req.(*proto.ProductCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharonServ_ProductDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.ProductDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).ProductDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/ProductDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).ProductDelete(ctx, req.(*proto.ProductDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharonServ_ProductsDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.ProductsDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).ProductsDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/ProductsDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).ProductsDelete(ctx, req.(*proto.ProductsDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharonServ_Products_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.ProductListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).Products(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/Products",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).Products(ctx, req.(*proto.ProductListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharonServ_Product_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).Product(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/Product",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).Product(ctx, req.(*proto.ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharonServ_ProductUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.ProductUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).ProductUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/ProductUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).ProductUpdate(ctx, req.(*proto.ProductUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharonServ_SubProductCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.SubProductCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).SubProductCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/SubProductCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).SubProductCreate(ctx, req.(*proto.SubProductCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharonServ_SubProductDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.SubProductDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).SubProductDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/SubProductDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).SubProductDelete(ctx, req.(*proto.SubProductDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharonServ_SubProductsDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.SubProductsDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).SubProductsDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/SubProductsDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).SubProductsDelete(ctx, req.(*proto.SubProductsDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharonServ_SubProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.SubProductListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).SubProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/SubProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).SubProducts(ctx, req.(*proto.SubProductListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharonServ_SubProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.SubProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).SubProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/SubProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).SubProduct(ctx, req.(*proto.SubProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharonServ_SubProductUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.SubProductUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).SubProductUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/SubProductUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).SubProductUpdate(ctx, req.(*proto.SubProductUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharonServ_AttributeCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.AttributeCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).AttributeCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/AttributeCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).AttributeCreate(ctx, req.(*proto.AttributeCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharonServ_AttributeDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.AttributeDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).AttributeDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/AttributeDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).AttributeDelete(ctx, req.(*proto.AttributeDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharonServ_Attributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.AttributeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).Attributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/Attributes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).Attributes(ctx, req.(*proto.AttributeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharonServ_Attribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.AttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).Attribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/Attribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).Attribute(ctx, req.(*proto.AttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharonServ_AttributeUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.AttributeUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).AttributeUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/AttributeUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).AttributeUpdate(ctx, req.(*proto.AttributeUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CharonServ_ServiceDesc is the grpc.ServiceDesc for CharonServ service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CharonServ_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "charonservice.CharonServ",
	HandlerType: (*CharonServServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pong",
			Handler:    _CharonServ_Pong_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CharonServ_Create_Handler,
		},
		{
			MethodName: "Categories",
			Handler:    _CharonServ_Categories_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CharonServ_Delete_Handler,
		},
		{
			MethodName: "Category",
			Handler:    _CharonServ_Category_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CharonServ_Update_Handler,
		},
		{
			MethodName: "CategoriesDelete",
			Handler:    _CharonServ_CategoriesDelete_Handler,
		},
		{
			MethodName: "ProductCreate",
			Handler:    _CharonServ_ProductCreate_Handler,
		},
		{
			MethodName: "ProductDelete",
			Handler:    _CharonServ_ProductDelete_Handler,
		},
		{
			MethodName: "ProductsDelete",
			Handler:    _CharonServ_ProductsDelete_Handler,
		},
		{
			MethodName: "Products",
			Handler:    _CharonServ_Products_Handler,
		},
		{
			MethodName: "Product",
			Handler:    _CharonServ_Product_Handler,
		},
		{
			MethodName: "ProductUpdate",
			Handler:    _CharonServ_ProductUpdate_Handler,
		},
		{
			MethodName: "SubProductCreate",
			Handler:    _CharonServ_SubProductCreate_Handler,
		},
		{
			MethodName: "SubProductDelete",
			Handler:    _CharonServ_SubProductDelete_Handler,
		},
		{
			MethodName: "SubProductsDelete",
			Handler:    _CharonServ_SubProductsDelete_Handler,
		},
		{
			MethodName: "SubProducts",
			Handler:    _CharonServ_SubProducts_Handler,
		},
		{
			MethodName: "SubProduct",
			Handler:    _CharonServ_SubProduct_Handler,
		},
		{
			MethodName: "SubProductUpdate",
			Handler:    _CharonServ_SubProductUpdate_Handler,
		},
		{
			MethodName: "AttributeCreate",
			Handler:    _CharonServ_AttributeCreate_Handler,
		},
		{
			MethodName: "AttributeDelete",
			Handler:    _CharonServ_AttributeDelete_Handler,
		},
		{
			MethodName: "Attributes",
			Handler:    _CharonServ_Attributes_Handler,
		},
		{
			MethodName: "Attribute",
			Handler:    _CharonServ_Attribute_Handler,
		},
		{
			MethodName: "AttributeUpdate",
			Handler:    _CharonServ_AttributeUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "charon/proto/service.proto",
}
