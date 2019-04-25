// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Product struct {
	Id                   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string              `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Details              map[string]*any.Any `protobuf:"bytes,4,rep,name=details,proto3" json:"details,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Price                float64             `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	Quantity             int32               `protobuf:"varint,6,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{0}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Product) GetDetails() map[string]*any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Product) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Product) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type ProductId struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductId) Reset()         { *m = ProductId{} }
func (m *ProductId) String() string { return proto.CompactTextString(m) }
func (*ProductId) ProtoMessage()    {}
func (*ProductId) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{1}
}

func (m *ProductId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductId.Unmarshal(m, b)
}
func (m *ProductId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductId.Marshal(b, m, deterministic)
}
func (m *ProductId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductId.Merge(m, src)
}
func (m *ProductId) XXX_Size() int {
	return xxx_messageInfo_ProductId.Size(m)
}
func (m *ProductId) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductId.DiscardUnknown(m)
}

var xxx_messageInfo_ProductId proto.InternalMessageInfo

func (m *ProductId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ValidateQuantity struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Quantity             int32    `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateQuantity) Reset()         { *m = ValidateQuantity{} }
func (m *ValidateQuantity) String() string { return proto.CompactTextString(m) }
func (*ValidateQuantity) ProtoMessage()    {}
func (*ValidateQuantity) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{2}
}

func (m *ValidateQuantity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateQuantity.Unmarshal(m, b)
}
func (m *ValidateQuantity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateQuantity.Marshal(b, m, deterministic)
}
func (m *ValidateQuantity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateQuantity.Merge(m, src)
}
func (m *ValidateQuantity) XXX_Size() int {
	return xxx_messageInfo_ValidateQuantity.Size(m)
}
func (m *ValidateQuantity) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateQuantity.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateQuantity proto.InternalMessageInfo

func (m *ValidateQuantity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ValidateQuantity) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func init() {
	proto.RegisterType((*Product)(nil), "pb.Product")
	proto.RegisterMapType((map[string]*any.Any)(nil), "pb.Product.DetailsEntry")
	proto.RegisterType((*ProductId)(nil), "pb.ProductId")
	proto.RegisterType((*ValidateQuantity)(nil), "pb.ValidateQuantity")
}

func init() { proto.RegisterFile("product.proto", fileDescriptor_f0fd8b59378f44a5) }

var fileDescriptor_f0fd8b59378f44a5 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x37, 0xd3, 0x4d, 0xba, 0xfb, 0xe2, 0x2e, 0xe5, 0x51, 0x24, 0xa6, 0x97, 0x10, 0x04,
	0x83, 0x87, 0x54, 0xe3, 0x45, 0x7a, 0x10, 0x8a, 0xad, 0xd0, 0x93, 0x35, 0xa2, 0xf7, 0x49, 0x67,
	0x2c, 0x83, 0xe9, 0x64, 0x4c, 0x27, 0x85, 0x7c, 0x02, 0xbf, 0x8a, 0x1f, 0x53, 0x92, 0x49, 0x25,
	0xa6, 0x94, 0xbd, 0xcd, 0x7b, 0xff, 0xff, 0x7b, 0xf3, 0x9b, 0x3f, 0x03, 0x0f, 0xaa, 0x2c, 0x58,
	0xb5, 0xd3, 0xb1, 0x2a, 0x0b, 0x5d, 0x20, 0x51, 0x99, 0x3f, 0xdb, 0x17, 0xc5, 0x3e, 0xe7, 0xf3,
	0xb6, 0x93, 0x55, 0x3f, 0xe6, 0xfc, 0xa0, 0x74, 0x6d, 0x0c, 0xfe, 0x8b, 0xa1, 0x48, 0x65, 0x27,
	0x85, 0xbf, 0x09, 0x8c, 0xb7, 0x66, 0x1b, 0x3e, 0x02, 0x11, 0xcc, 0xb3, 0x02, 0x2b, 0xba, 0x4f,
	0x89, 0x60, 0x88, 0x70, 0x2b, 0xe9, 0x81, 0x7b, 0xa4, 0xed, 0xb4, 0x67, 0x0c, 0xc0, 0x65, 0xfc,
	0xb8, 0x2b, 0x85, 0xd2, 0xa2, 0x90, 0xde, 0xa8, 0x95, 0xfa, 0x2d, 0x4c, 0x60, 0xcc, 0xb8, 0xa6,
	0x22, 0x3f, 0x7a, 0xb7, 0xc1, 0x28, 0x72, 0x13, 0x2f, 0x56, 0x59, 0xdc, 0xdd, 0x11, 0xaf, 0x8c,
	0xb4, 0x96, 0xba, 0xac, 0xd3, 0xb3, 0x11, 0xa7, 0x60, 0xab, 0x52, 0xec, 0xb8, 0x67, 0x07, 0x56,
	0x64, 0xa5, 0xa6, 0x40, 0x1f, 0xee, 0x7e, 0x55, 0x54, 0x6a, 0xa1, 0x6b, 0xcf, 0x09, 0xac, 0xc8,
	0x4e, 0xff, 0xd5, 0xfe, 0x16, 0x9e, 0xf5, 0x57, 0xe1, 0x04, 0x46, 0x3f, 0x79, 0xdd, 0xc1, 0x37,
	0x47, 0x7c, 0x0d, 0xf6, 0x89, 0xe6, 0x95, 0xc1, 0x77, 0x93, 0x69, 0x6c, 0x42, 0x88, 0xcf, 0x21,
	0xc4, 0x4b, 0x59, 0xa7, 0xc6, 0xb2, 0x20, 0xef, 0xad, 0x70, 0x06, 0xf7, 0x1d, 0xe4, 0x86, 0x0d,
	0xa3, 0x08, 0x3f, 0xc0, 0xe4, 0x3b, 0xcd, 0x05, 0xa3, 0x9a, 0x7f, 0xe9, 0x10, 0x2e, 0xe2, 0xea,
	0xe3, 0x92, 0xff, 0x71, 0x93, 0x3f, 0x04, 0x1e, 0xbb, 0xed, 0x5f, 0x79, 0x79, 0x6a, 0x5e, 0xf7,
	0x12, 0x9c, 0x8f, 0x25, 0xa7, 0x9a, 0xa3, 0xdb, 0x0b, 0xc8, 0xef, 0x17, 0xe1, 0x0d, 0xbe, 0x82,
	0xf1, 0x27, 0x21, 0xd9, 0x67, 0xc9, 0xf1, 0xa1, 0xa7, 0x6c, 0xd8, 0xd0, 0x98, 0x18, 0xe3, 0x32,
	0xcf, 0xf1, 0xf9, 0xc5, 0x53, 0xd7, 0xcd, 0x67, 0x18, 0x4c, 0xbc, 0xb1, 0x1a, 0x84, 0x6f, 0x8a,
	0x3d, 0x85, 0xf0, 0x16, 0x9c, 0x15, 0xcf, 0xb9, 0xbe, 0x20, 0xb8, 0x72, 0x4f, 0x78, 0x83, 0x0b,
	0xb8, 0x3b, 0xc7, 0x85, 0xd3, 0x66, 0x68, 0x18, 0xde, 0xf5, 0xd9, 0xcc, 0x69, 0x3b, 0xef, 0xfe,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x94, 0xca, 0x7c, 0xb3, 0xe5, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	Create(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Product, error)
	FindOne(ctx context.Context, in *ProductId, opts ...grpc.CallOption) (*Product, error)
	FindAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (ProductService_FindAllClient, error)
	Update(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Product, error)
	Delete(ctx context.Context, in *ProductId, opts ...grpc.CallOption) (*empty.Empty, error)
	Validate(ctx context.Context, in *ValidateQuantity, opts ...grpc.CallOption) (*empty.Empty, error)
}

type productServiceClient struct {
	cc *grpc.ClientConn
}

func NewProductServiceClient(cc *grpc.ClientConn) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) Create(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/pb.ProductService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) FindOne(ctx context.Context, in *ProductId, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/pb.ProductService/FindOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) FindAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (ProductService_FindAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProductService_serviceDesc.Streams[0], "/pb.ProductService/FindAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &productServiceFindAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProductService_FindAllClient interface {
	Recv() (*Product, error)
	grpc.ClientStream
}

type productServiceFindAllClient struct {
	grpc.ClientStream
}

func (x *productServiceFindAllClient) Recv() (*Product, error) {
	m := new(Product)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *productServiceClient) Update(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/pb.ProductService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) Delete(ctx context.Context, in *ProductId, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.ProductService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) Validate(ctx context.Context, in *ValidateQuantity, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.ProductService/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	Create(context.Context, *Product) (*Product, error)
	FindOne(context.Context, *ProductId) (*Product, error)
	FindAll(*empty.Empty, ProductService_FindAllServer) error
	Update(context.Context, *Product) (*Product, error)
	Delete(context.Context, *ProductId) (*empty.Empty, error)
	Validate(context.Context, *ValidateQuantity) (*empty.Empty, error)
}

// UnimplementedProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (*UnimplementedProductServiceServer) Create(ctx context.Context, req *Product) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedProductServiceServer) FindOne(ctx context.Context, req *ProductId) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOne not implemented")
}
func (*UnimplementedProductServiceServer) FindAll(req *empty.Empty, srv ProductService_FindAllServer) error {
	return status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}
func (*UnimplementedProductServiceServer) Update(ctx context.Context, req *Product) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedProductServiceServer) Delete(ctx context.Context, req *ProductId) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedProductServiceServer) Validate(ctx context.Context, req *ValidateQuantity) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProductService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).Create(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_FindOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).FindOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProductService/FindOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).FindOne(ctx, req.(*ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_FindAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProductServiceServer).FindAll(m, &productServiceFindAllServer{stream})
}

type ProductService_FindAllServer interface {
	Send(*Product) error
	grpc.ServerStream
}

type productServiceFindAllServer struct {
	grpc.ServerStream
}

func (x *productServiceFindAllServer) Send(m *Product) error {
	return x.ServerStream.SendMsg(m)
}

func _ProductService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProductService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).Update(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProductService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).Delete(ctx, req.(*ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateQuantity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProductService/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).Validate(ctx, req.(*ValidateQuantity))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ProductService_Create_Handler,
		},
		{
			MethodName: "FindOne",
			Handler:    _ProductService_FindOne_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ProductService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ProductService_Delete_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _ProductService_Validate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FindAll",
			Handler:       _ProductService_FindAll_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "product.proto",
}