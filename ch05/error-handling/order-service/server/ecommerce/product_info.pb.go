// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product_info.proto

package ecommerce

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
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

type Order struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Items                []string `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price                float32  `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Destination          string   `protobuf:"bytes,5,opt,name=destination,proto3" json:"destination,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a4d768ec9cb4951, []int{0}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetItems() []string {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Order) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Order) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Order) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

type CombinedShipment struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	OrdersList           []*Order `protobuf:"bytes,3,rep,name=ordersList,proto3" json:"ordersList,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CombinedShipment) Reset()         { *m = CombinedShipment{} }
func (m *CombinedShipment) String() string { return proto.CompactTextString(m) }
func (*CombinedShipment) ProtoMessage()    {}
func (*CombinedShipment) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a4d768ec9cb4951, []int{1}
}

func (m *CombinedShipment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CombinedShipment.Unmarshal(m, b)
}
func (m *CombinedShipment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CombinedShipment.Marshal(b, m, deterministic)
}
func (m *CombinedShipment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CombinedShipment.Merge(m, src)
}
func (m *CombinedShipment) XXX_Size() int {
	return xxx_messageInfo_CombinedShipment.Size(m)
}
func (m *CombinedShipment) XXX_DiscardUnknown() {
	xxx_messageInfo_CombinedShipment.DiscardUnknown(m)
}

var xxx_messageInfo_CombinedShipment proto.InternalMessageInfo

func (m *CombinedShipment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CombinedShipment) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *CombinedShipment) GetOrdersList() []*Order {
	if m != nil {
		return m.OrdersList
	}
	return nil
}

func init() {
	proto.RegisterType((*Order)(nil), "ecommerce.Order")
	proto.RegisterType((*CombinedShipment)(nil), "ecommerce.CombinedShipment")
}

func init() { proto.RegisterFile("product_info.proto", fileDescriptor_9a4d768ec9cb4951) }

var fileDescriptor_9a4d768ec9cb4951 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x50, 0xcd, 0x4a, 0x03, 0x31,
	0x18, 0x24, 0xbb, 0xb6, 0xd8, 0xaf, 0xfe, 0x94, 0x20, 0xb2, 0x54, 0x91, 0xa5, 0xa7, 0x3d, 0x6d,
	0x4b, 0xbd, 0x79, 0x12, 0xbc, 0xfa, 0x03, 0x5b, 0xf0, 0x2a, 0x69, 0xf2, 0x75, 0x1b, 0xe8, 0x26,
	0x21, 0xc9, 0xe2, 0x23, 0xf8, 0x94, 0xbe, 0x8b, 0x6c, 0xb6, 0x2d, 0xc5, 0x05, 0xb1, 0xc7, 0x99,
	0xcc, 0x4c, 0xe6, 0x1b, 0xa0, 0xc6, 0x6a, 0x51, 0x73, 0xff, 0x21, 0xd5, 0x4a, 0xe7, 0xc6, 0x6a,
	0xaf, 0xe9, 0x00, 0xb9, 0xae, 0x2a, 0xb4, 0x1c, 0xc7, 0x77, 0xa5, 0xd6, 0xe5, 0x06, 0xa7, 0xe1,
	0x61, 0x59, 0xaf, 0xa6, 0x9f, 0x96, 0x19, 0x83, 0xd6, 0xb5, 0xd2, 0xc9, 0x17, 0x81, 0xde, 0x9b,
	0x15, 0x68, 0xe9, 0x05, 0x44, 0x52, 0x24, 0x24, 0x25, 0xd9, 0xa0, 0x88, 0xa4, 0xa0, 0x57, 0xd0,
	0x93, 0x1e, 0x2b, 0x97, 0x44, 0x69, 0x9c, 0x0d, 0x8a, 0x16, 0xd0, 0x14, 0x86, 0x02, 0x1d, 0xb7,
	0xd2, 0x78, 0xa9, 0x55, 0x12, 0x07, 0xf9, 0x21, 0xd5, 0xf8, 0x8c, 0x95, 0x1c, 0x93, 0x93, 0x94,
	0x64, 0x51, 0xd1, 0x82, 0xad, 0xcf, 0x4b, 0xc5, 0x82, 0xaf, 0xb7, 0xf7, 0xed, 0xa8, 0xc9, 0x06,
	0x46, 0x4f, 0xba, 0x5a, 0x4a, 0x85, 0x62, 0xb1, 0x96, 0xa6, 0x42, 0xe5, 0x3b, 0x9d, 0xae, 0xa1,
	0xef, 0x3c, 0xf3, 0x75, 0x53, 0xaa, 0xe1, 0xb6, 0x88, 0xce, 0x00, 0x74, 0x73, 0x84, 0x7b, 0x96,
	0xce, 0x27, 0x71, 0x1a, 0x67, 0xc3, 0xf9, 0x28, 0xdf, 0xaf, 0x90, 0x87, 0x0b, 0x8b, 0x03, 0xcd,
	0xfc, 0x3b, 0x82, 0xcb, 0xc0, 0xbe, 0x30, 0xc5, 0x4a, 0x0c, 0xbf, 0x3d, 0xc0, 0x69, 0x89, 0xbe,
	0x5d, 0xe3, 0x36, 0x6f, 0x87, 0xcb, 0x77, 0xc3, 0xe5, 0x0b, 0x6f, 0xa5, 0x2a, 0xdf, 0xd9, 0xa6,
	0xc6, 0x71, 0x27, 0x9b, 0x3e, 0xc2, 0x99, 0x43, 0x66, 0xf9, 0x3a, 0x40, 0x77, 0xac, 0x7f, 0x46,
	0x9a, 0x84, 0xda, 0x08, 0xe6, 0x71, 0x9b, 0xd0, 0xd1, 0x8c, 0xff, 0xcc, 0xcc, 0x08, 0x7d, 0x85,
	0x73, 0x63, 0x35, 0x47, 0xe7, 0xfe, 0x55, 0xe2, 0xe6, 0xe0, 0x83, 0xdf, 0xcb, 0x67, 0x64, 0x46,
	0x9a, 0x3d, 0x98, 0x10, 0xed, 0x7d, 0x47, 0xb6, 0x59, 0xf6, 0x03, 0x7b, 0xff, 0x13, 0x00, 0x00,
	0xff, 0xff, 0xc5, 0xb3, 0xac, 0x05, 0x9f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrderManagementClient is the client API for OrderManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderManagementClient interface {
	GetOrder(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*Order, error)
	SearchOrders(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (OrderManagement_SearchOrdersClient, error)
	UpdateOrders(ctx context.Context, opts ...grpc.CallOption) (OrderManagement_UpdateOrdersClient, error)
	ProcessOrders(ctx context.Context, opts ...grpc.CallOption) (OrderManagement_ProcessOrdersClient, error)
	AddOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*wrappers.StringValue, error)
}

type orderManagementClient struct {
	cc *grpc.ClientConn
}

func NewOrderManagementClient(cc *grpc.ClientConn) OrderManagementClient {
	return &orderManagementClient{cc}
}

func (c *orderManagementClient) GetOrder(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/ecommerce.OrderManagement/getOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderManagementClient) SearchOrders(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (OrderManagement_SearchOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OrderManagement_serviceDesc.Streams[0], "/ecommerce.OrderManagement/searchOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderManagementSearchOrdersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderManagement_SearchOrdersClient interface {
	Recv() (*Order, error)
	grpc.ClientStream
}

type orderManagementSearchOrdersClient struct {
	grpc.ClientStream
}

func (x *orderManagementSearchOrdersClient) Recv() (*Order, error) {
	m := new(Order)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderManagementClient) UpdateOrders(ctx context.Context, opts ...grpc.CallOption) (OrderManagement_UpdateOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OrderManagement_serviceDesc.Streams[1], "/ecommerce.OrderManagement/updateOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderManagementUpdateOrdersClient{stream}
	return x, nil
}

type OrderManagement_UpdateOrdersClient interface {
	Send(*Order) error
	CloseAndRecv() (*wrappers.StringValue, error)
	grpc.ClientStream
}

type orderManagementUpdateOrdersClient struct {
	grpc.ClientStream
}

func (x *orderManagementUpdateOrdersClient) Send(m *Order) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderManagementUpdateOrdersClient) CloseAndRecv() (*wrappers.StringValue, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(wrappers.StringValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderManagementClient) ProcessOrders(ctx context.Context, opts ...grpc.CallOption) (OrderManagement_ProcessOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OrderManagement_serviceDesc.Streams[2], "/ecommerce.OrderManagement/processOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderManagementProcessOrdersClient{stream}
	return x, nil
}

type OrderManagement_ProcessOrdersClient interface {
	Send(*wrappers.StringValue) error
	Recv() (*CombinedShipment, error)
	grpc.ClientStream
}

type orderManagementProcessOrdersClient struct {
	grpc.ClientStream
}

func (x *orderManagementProcessOrdersClient) Send(m *wrappers.StringValue) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderManagementProcessOrdersClient) Recv() (*CombinedShipment, error) {
	m := new(CombinedShipment)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderManagementClient) AddOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*wrappers.StringValue, error) {
	out := new(wrappers.StringValue)
	err := c.cc.Invoke(ctx, "/ecommerce.OrderManagement/addOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderManagementServer is the server API for OrderManagement service.
type OrderManagementServer interface {
	GetOrder(context.Context, *wrappers.StringValue) (*Order, error)
	SearchOrders(*wrappers.StringValue, OrderManagement_SearchOrdersServer) error
	UpdateOrders(OrderManagement_UpdateOrdersServer) error
	ProcessOrders(OrderManagement_ProcessOrdersServer) error
	AddOrder(context.Context, *Order) (*wrappers.StringValue, error)
}

func RegisterOrderManagementServer(s *grpc.Server, srv OrderManagementServer) {
	s.RegisterService(&_OrderManagement_serviceDesc, srv)
}

func _OrderManagement_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.OrderManagement/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServer).GetOrder(ctx, req.(*wrappers.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderManagement_SearchOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(wrappers.StringValue)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderManagementServer).SearchOrders(m, &orderManagementSearchOrdersServer{stream})
}

type OrderManagement_SearchOrdersServer interface {
	Send(*Order) error
	grpc.ServerStream
}

type orderManagementSearchOrdersServer struct {
	grpc.ServerStream
}

func (x *orderManagementSearchOrdersServer) Send(m *Order) error {
	return x.ServerStream.SendMsg(m)
}

func _OrderManagement_UpdateOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderManagementServer).UpdateOrders(&orderManagementUpdateOrdersServer{stream})
}

type OrderManagement_UpdateOrdersServer interface {
	SendAndClose(*wrappers.StringValue) error
	Recv() (*Order, error)
	grpc.ServerStream
}

type orderManagementUpdateOrdersServer struct {
	grpc.ServerStream
}

func (x *orderManagementUpdateOrdersServer) SendAndClose(m *wrappers.StringValue) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderManagementUpdateOrdersServer) Recv() (*Order, error) {
	m := new(Order)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _OrderManagement_ProcessOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderManagementServer).ProcessOrders(&orderManagementProcessOrdersServer{stream})
}

type OrderManagement_ProcessOrdersServer interface {
	Send(*CombinedShipment) error
	Recv() (*wrappers.StringValue, error)
	grpc.ServerStream
}

type orderManagementProcessOrdersServer struct {
	grpc.ServerStream
}

func (x *orderManagementProcessOrdersServer) Send(m *CombinedShipment) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderManagementProcessOrdersServer) Recv() (*wrappers.StringValue, error) {
	m := new(wrappers.StringValue)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _OrderManagement_AddOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServer).AddOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.OrderManagement/AddOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServer).AddOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderManagement_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ecommerce.OrderManagement",
	HandlerType: (*OrderManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getOrder",
			Handler:    _OrderManagement_GetOrder_Handler,
		},
		{
			MethodName: "addOrder",
			Handler:    _OrderManagement_AddOrder_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "searchOrders",
			Handler:       _OrderManagement_SearchOrders_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "updateOrders",
			Handler:       _OrderManagement_UpdateOrders_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "processOrders",
			Handler:       _OrderManagement_ProcessOrders_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "product_info.proto",
}