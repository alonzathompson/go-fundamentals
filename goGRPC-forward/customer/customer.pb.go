// Code generated by protoc-gen-go. DO NOT EDIT.
// source: customer.proto

package customer

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Request message for creating a new customer
type CustomerRequest struct {
	Id                   int32                      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string                     `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string                     `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Addresses            []*CustomerRequest_Address `protobuf:"bytes,5,rep,name=addresses,proto3" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *CustomerRequest) Reset()         { *m = CustomerRequest{} }
func (m *CustomerRequest) String() string { return proto.CompactTextString(m) }
func (*CustomerRequest) ProtoMessage()    {}
func (*CustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{0}
}

func (m *CustomerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerRequest.Unmarshal(m, b)
}
func (m *CustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerRequest.Marshal(b, m, deterministic)
}
func (m *CustomerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerRequest.Merge(m, src)
}
func (m *CustomerRequest) XXX_Size() int {
	return xxx_messageInfo_CustomerRequest.Size(m)
}
func (m *CustomerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerRequest proto.InternalMessageInfo

func (m *CustomerRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CustomerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CustomerRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CustomerRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *CustomerRequest) GetAddresses() []*CustomerRequest_Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type CustomerRequest_Address struct {
	Street               string   `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`
	City                 string   `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	State                string   `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Zip                  string   `protobuf:"bytes,4,opt,name=zip,proto3" json:"zip,omitempty"`
	IsShippingAddress    bool     `protobuf:"varint,5,opt,name=isShippingAddress,proto3" json:"isShippingAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerRequest_Address) Reset()         { *m = CustomerRequest_Address{} }
func (m *CustomerRequest_Address) String() string { return proto.CompactTextString(m) }
func (*CustomerRequest_Address) ProtoMessage()    {}
func (*CustomerRequest_Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{0, 0}
}

func (m *CustomerRequest_Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerRequest_Address.Unmarshal(m, b)
}
func (m *CustomerRequest_Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerRequest_Address.Marshal(b, m, deterministic)
}
func (m *CustomerRequest_Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerRequest_Address.Merge(m, src)
}
func (m *CustomerRequest_Address) XXX_Size() int {
	return xxx_messageInfo_CustomerRequest_Address.Size(m)
}
func (m *CustomerRequest_Address) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerRequest_Address.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerRequest_Address proto.InternalMessageInfo

func (m *CustomerRequest_Address) GetStreet() string {
	if m != nil {
		return m.Street
	}
	return ""
}

func (m *CustomerRequest_Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *CustomerRequest_Address) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *CustomerRequest_Address) GetZip() string {
	if m != nil {
		return m.Zip
	}
	return ""
}

func (m *CustomerRequest_Address) GetIsShippingAddress() bool {
	if m != nil {
		return m.IsShippingAddress
	}
	return false
}

type CustomerResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerResponse) Reset()         { *m = CustomerResponse{} }
func (m *CustomerResponse) String() string { return proto.CompactTextString(m) }
func (*CustomerResponse) ProtoMessage()    {}
func (*CustomerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{1}
}

func (m *CustomerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerResponse.Unmarshal(m, b)
}
func (m *CustomerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerResponse.Marshal(b, m, deterministic)
}
func (m *CustomerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerResponse.Merge(m, src)
}
func (m *CustomerResponse) XXX_Size() int {
	return xxx_messageInfo_CustomerResponse.Size(m)
}
func (m *CustomerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerResponse proto.InternalMessageInfo

func (m *CustomerResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CustomerResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type CustomerFilter struct {
	Keyword              string   `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerFilter) Reset()         { *m = CustomerFilter{} }
func (m *CustomerFilter) String() string { return proto.CompactTextString(m) }
func (*CustomerFilter) ProtoMessage()    {}
func (*CustomerFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{2}
}

func (m *CustomerFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerFilter.Unmarshal(m, b)
}
func (m *CustomerFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerFilter.Marshal(b, m, deterministic)
}
func (m *CustomerFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerFilter.Merge(m, src)
}
func (m *CustomerFilter) XXX_Size() int {
	return xxx_messageInfo_CustomerFilter.Size(m)
}
func (m *CustomerFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerFilter.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerFilter proto.InternalMessageInfo

func (m *CustomerFilter) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

func init() {
	proto.RegisterType((*CustomerRequest)(nil), "customer.CustomerRequest")
	proto.RegisterType((*CustomerRequest_Address)(nil), "customer.CustomerRequest.Address")
	proto.RegisterType((*CustomerResponse)(nil), "customer.CustomerResponse")
	proto.RegisterType((*CustomerFilter)(nil), "customer.CustomerFilter")
}

func init() { proto.RegisterFile("customer.proto", fileDescriptor_9efa92dae3d6ec46) }

var fileDescriptor_9efa92dae3d6ec46 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x4f, 0x4e, 0xf3, 0x30,
	0x10, 0xc5, 0xeb, 0xb4, 0xe9, 0x9f, 0xf9, 0x3e, 0xa5, 0xc5, 0x42, 0xc8, 0x64, 0x15, 0xb2, 0x8a,
	0x10, 0xaa, 0x50, 0xd9, 0x22, 0x21, 0x54, 0xa9, 0x15, 0x2c, 0xc3, 0x09, 0x42, 0x32, 0xa2, 0x16,
	0x6d, 0x1c, 0x6c, 0x57, 0x55, 0xb9, 0x02, 0x97, 0xe1, 0x5c, 0x9c, 0x02, 0xd9, 0x89, 0x5b, 0x89,
	0xd2, 0x15, 0xbb, 0xf9, 0x3d, 0xcf, 0x3c, 0x3f, 0x8f, 0x0c, 0x41, 0xbe, 0x56, 0x5a, 0xac, 0x50,
	0x8e, 0x2b, 0x29, 0xb4, 0xa0, 0x7d, 0xc7, 0xf1, 0xa7, 0x07, 0xc3, 0x69, 0x03, 0x29, 0xbe, 0xad,
	0x51, 0x69, 0x1a, 0x80, 0xc7, 0x0b, 0x46, 0x22, 0x92, 0xf8, 0xa9, 0xc7, 0x0b, 0x4a, 0xa1, 0x53,
	0x66, 0x2b, 0x64, 0x5e, 0x44, 0x92, 0x41, 0x6a, 0x6b, 0x7a, 0x0a, 0x3e, 0xae, 0x32, 0xbe, 0x64,
	0x6d, 0x2b, 0xd6, 0x60, 0xd4, 0x6a, 0x21, 0x4a, 0x64, 0x9d, 0x5a, 0xb5, 0x40, 0xef, 0x60, 0x90,
	0x15, 0x85, 0x44, 0xa5, 0x50, 0x31, 0x3f, 0x6a, 0x27, 0xff, 0x26, 0x17, 0xe3, 0x5d, 0xa2, 0x1f,
	0xb7, 0x8f, 0xef, 0xeb, 0xd6, 0x74, 0x3f, 0x13, 0x7e, 0x10, 0xe8, 0x35, 0x32, 0x3d, 0x83, 0xae,
	0xd2, 0x12, 0x51, 0xdb, 0x80, 0x83, 0xb4, 0x21, 0x13, 0x32, 0xe7, 0x7a, 0xeb, 0x42, 0x9a, 0xda,
	0xc4, 0x51, 0x3a, 0xd3, 0xe8, 0x42, 0x5a, 0xa0, 0x23, 0x68, 0xbf, 0xf3, 0xaa, 0x89, 0x68, 0x4a,
	0x7a, 0x05, 0x27, 0x5c, 0x3d, 0x2d, 0x78, 0x55, 0xf1, 0xf2, 0xa5, 0xb9, 0x88, 0xf9, 0x11, 0x49,
	0xfa, 0xe9, 0xe1, 0x41, 0x7c, 0x0b, 0xa3, 0x7d, 0x66, 0x55, 0x89, 0x52, 0xe1, 0xc1, 0xca, 0x18,
	0xf4, 0xd4, 0x3a, 0xcf, 0x8d, 0x8f, 0x67, 0x7d, 0x1c, 0xc6, 0x97, 0x10, 0xb8, 0xe9, 0x19, 0x5f,
	0x6a, 0x94, 0xa6, 0xf7, 0x15, 0xb7, 0x1b, 0x21, 0x8b, 0xe6, 0x49, 0x0e, 0x27, 0x5f, 0x04, 0xfa,
	0xae, 0x99, 0xce, 0xe1, 0xff, 0x1c, 0xb5, 0x43, 0x45, 0xd9, 0xe1, 0x0a, 0x6b, 0xc3, 0xf0, 0xfc,
	0xe8, 0x72, 0xe3, 0xd6, 0x35, 0xa1, 0x0f, 0x10, 0x4c, 0x25, 0x66, 0x1a, 0x77, 0xd6, 0xc7, 0x07,
	0xc2, 0xf0, 0xb7, 0xa3, 0xfa, 0xd1, 0x71, 0x8b, 0x3e, 0xc2, 0x70, 0x26, 0xe4, 0x26, 0x93, 0xc5,
	0x9f, 0xbd, 0x9e, 0xbb, 0xf6, 0x6b, 0xde, 0x7c, 0x07, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x98, 0x3a,
	0xb3, 0xac, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomerClient is the client API for Customer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerClient interface {
	// Get all Customers with filter - A server-to-client streaming RPC.
	GetCustomers(ctx context.Context, in *CustomerFilter, opts ...grpc.CallOption) (Customer_GetCustomersClient, error)
	// Create a new Customer - A simple RPC
	CreateCustomer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*CustomerResponse, error)
	//forward request
	ForwardCustomer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*CustomerResponse, error)
}

type customerClient struct {
	cc *grpc.ClientConn
}

func NewCustomerClient(cc *grpc.ClientConn) CustomerClient {
	return &customerClient{cc}
}

func (c *customerClient) GetCustomers(ctx context.Context, in *CustomerFilter, opts ...grpc.CallOption) (Customer_GetCustomersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Customer_serviceDesc.Streams[0], "/customer.Customer/GetCustomers", opts...)
	if err != nil {
		return nil, err
	}
	x := &customerGetCustomersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Customer_GetCustomersClient interface {
	Recv() (*CustomerRequest, error)
	grpc.ClientStream
}

type customerGetCustomersClient struct {
	grpc.ClientStream
}

func (x *customerGetCustomersClient) Recv() (*CustomerRequest, error) {
	m := new(CustomerRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *customerClient) CreateCustomer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*CustomerResponse, error) {
	out := new(CustomerResponse)
	err := c.cc.Invoke(ctx, "/customer.Customer/CreateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) ForwardCustomer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*CustomerResponse, error) {
	out := new(CustomerResponse)
	err := c.cc.Invoke(ctx, "/customer.Customer/ForwardCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServer is the server API for Customer service.
type CustomerServer interface {
	// Get all Customers with filter - A server-to-client streaming RPC.
	GetCustomers(*CustomerFilter, Customer_GetCustomersServer) error
	// Create a new Customer - A simple RPC
	CreateCustomer(context.Context, *CustomerRequest) (*CustomerResponse, error)
	//forward request
	ForwardCustomer(context.Context, *CustomerRequest) (*CustomerResponse, error)
}

func RegisterCustomerServer(s *grpc.Server, srv CustomerServer) {
	s.RegisterService(&_Customer_serviceDesc, srv)
}

func _Customer_GetCustomers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CustomerFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CustomerServer).GetCustomers(m, &customerGetCustomersServer{stream})
}

type Customer_GetCustomersServer interface {
	Send(*CustomerRequest) error
	grpc.ServerStream
}

type customerGetCustomersServer struct {
	grpc.ServerStream
}

func (x *customerGetCustomersServer) Send(m *CustomerRequest) error {
	return x.ServerStream.SendMsg(m)
}

func _Customer_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.Customer/CreateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).CreateCustomer(ctx, req.(*CustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_ForwardCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).ForwardCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.Customer/ForwardCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).ForwardCustomer(ctx, req.(*CustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Customer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "customer.Customer",
	HandlerType: (*CustomerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCustomer",
			Handler:    _Customer_CreateCustomer_Handler,
		},
		{
			MethodName: "ForwardCustomer",
			Handler:    _Customer_ForwardCustomer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCustomers",
			Handler:       _Customer_GetCustomers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "customer.proto",
}
