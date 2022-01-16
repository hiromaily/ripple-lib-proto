// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rippleapi/address.proto

package rippleapi

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ResponseGenerateAddress struct {
	XAddress       string `protobuf:"bytes,1,opt,name=xAddress,proto3" json:"xAddress,omitempty"`
	ClassicAddress string `protobuf:"bytes,2,opt,name=classicAddress,proto3" json:"classicAddress,omitempty"`
	Address        string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Secret         string `protobuf:"bytes,4,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (m *ResponseGenerateAddress) Reset()      { *m = ResponseGenerateAddress{} }
func (*ResponseGenerateAddress) ProtoMessage() {}
func (*ResponseGenerateAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_1abc8906907d13d8, []int{0}
}
func (m *ResponseGenerateAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseGenerateAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseGenerateAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseGenerateAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseGenerateAddress.Merge(m, src)
}
func (m *ResponseGenerateAddress) XXX_Size() int {
	return m.Size()
}
func (m *ResponseGenerateAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseGenerateAddress.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseGenerateAddress proto.InternalMessageInfo

func (m *ResponseGenerateAddress) GetXAddress() string {
	if m != nil {
		return m.XAddress
	}
	return ""
}

func (m *ResponseGenerateAddress) GetClassicAddress() string {
	if m != nil {
		return m.ClassicAddress
	}
	return ""
}

func (m *ResponseGenerateAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ResponseGenerateAddress) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

type ResponseGenerateXAddress struct {
	XAddress string `protobuf:"bytes,1,opt,name=xAddress,proto3" json:"xAddress,omitempty"`
	Secret   string `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (m *ResponseGenerateXAddress) Reset()      { *m = ResponseGenerateXAddress{} }
func (*ResponseGenerateXAddress) ProtoMessage() {}
func (*ResponseGenerateXAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_1abc8906907d13d8, []int{1}
}
func (m *ResponseGenerateXAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseGenerateXAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseGenerateXAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseGenerateXAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseGenerateXAddress.Merge(m, src)
}
func (m *ResponseGenerateXAddress) XXX_Size() int {
	return m.Size()
}
func (m *ResponseGenerateXAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseGenerateXAddress.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseGenerateXAddress proto.InternalMessageInfo

func (m *ResponseGenerateXAddress) GetXAddress() string {
	if m != nil {
		return m.XAddress
	}
	return ""
}

func (m *ResponseGenerateXAddress) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

type RequestIsValidAddress struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *RequestIsValidAddress) Reset()      { *m = RequestIsValidAddress{} }
func (*RequestIsValidAddress) ProtoMessage() {}
func (*RequestIsValidAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_1abc8906907d13d8, []int{2}
}
func (m *RequestIsValidAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestIsValidAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestIsValidAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestIsValidAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestIsValidAddress.Merge(m, src)
}
func (m *RequestIsValidAddress) XXX_Size() int {
	return m.Size()
}
func (m *RequestIsValidAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestIsValidAddress.DiscardUnknown(m)
}

var xxx_messageInfo_RequestIsValidAddress proto.InternalMessageInfo

func (m *RequestIsValidAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type ResponseIsValidAddress struct {
	IsValid bool `protobuf:"varint,1,opt,name=isValid,proto3" json:"isValid,omitempty"`
}

func (m *ResponseIsValidAddress) Reset()      { *m = ResponseIsValidAddress{} }
func (*ResponseIsValidAddress) ProtoMessage() {}
func (*ResponseIsValidAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_1abc8906907d13d8, []int{3}
}
func (m *ResponseIsValidAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseIsValidAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseIsValidAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseIsValidAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseIsValidAddress.Merge(m, src)
}
func (m *ResponseIsValidAddress) XXX_Size() int {
	return m.Size()
}
func (m *ResponseIsValidAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseIsValidAddress.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseIsValidAddress proto.InternalMessageInfo

func (m *ResponseIsValidAddress) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

func init() {
	proto.RegisterType((*ResponseGenerateAddress)(nil), "rippleapi.address.ResponseGenerateAddress")
	proto.RegisterType((*ResponseGenerateXAddress)(nil), "rippleapi.address.ResponseGenerateXAddress")
	proto.RegisterType((*RequestIsValidAddress)(nil), "rippleapi.address.RequestIsValidAddress")
	proto.RegisterType((*ResponseIsValidAddress)(nil), "rippleapi.address.ResponseIsValidAddress")
}

func init() { proto.RegisterFile("rippleapi/address.proto", fileDescriptor_1abc8906907d13d8) }

var fileDescriptor_1abc8906907d13d8 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xcd, 0x4e, 0xea, 0x40,
	0x18, 0x9d, 0xe1, 0xde, 0x70, 0xe1, 0x5b, 0x70, 0xb1, 0x89, 0xd0, 0xd4, 0x64, 0x34, 0x5d, 0x18,
	0xd4, 0x38, 0x44, 0x7c, 0x02, 0x48, 0x8c, 0x61, 0x63, 0x4c, 0x17, 0x8a, 0xee, 0xda, 0x32, 0x96,
	0x49, 0x0a, 0xad, 0x9d, 0x92, 0xc8, 0xce, 0x37, 0xd0, 0xc7, 0x70, 0xe7, 0x6b, 0xb8, 0x64, 0xc9,
	0x52, 0xca, 0xc6, 0x25, 0x8f, 0x60, 0xec, 0x0f, 0x84, 0x16, 0x0c, 0xbb, 0x9e, 0xf9, 0xce, 0x77,
	0xbe, 0xd3, 0x73, 0xa0, 0xea, 0x71, 0xd7, 0xb5, 0x99, 0xee, 0xf2, 0xba, 0xde, 0xed, 0x7a, 0x4c,
	0x08, 0xea, 0x7a, 0x8e, 0xef, 0x48, 0x3b, 0x8b, 0x01, 0x8d, 0x07, 0xca, 0x9e, 0xe5, 0x38, 0x96,
	0xcd, 0xea, 0x21, 0xc1, 0x18, 0x3e, 0xd4, 0x59, 0xdf, 0xf5, 0x47, 0x11, 0x5f, 0x7d, 0xc1, 0x50,
	0xd5, 0x98, 0x70, 0x9d, 0x81, 0x60, 0x97, 0x6c, 0xc0, 0x3c, 0xdd, 0x67, 0xcd, 0x68, 0x51, 0x52,
	0xa0, 0xf0, 0x14, 0x7f, 0xcb, 0xf8, 0x00, 0xd7, 0x8a, 0xda, 0x02, 0x4b, 0x87, 0x50, 0x32, 0x6d,
	0x5d, 0x08, 0x6e, 0x26, 0x8c, 0x5c, 0xc8, 0x48, 0xbd, 0x4a, 0x32, 0xfc, 0x8b, 0x7d, 0xc8, 0x7f,
	0x42, 0x42, 0x02, 0xa5, 0x0a, 0xe4, 0x05, 0x33, 0x3d, 0xe6, 0xcb, 0x7f, 0xc3, 0x41, 0x8c, 0xd4,
	0x2b, 0x90, 0xd3, 0x86, 0x3a, 0xdb, 0x38, 0x5a, 0xea, 0xe5, 0x56, 0xf4, 0xce, 0x60, 0x57, 0x63,
	0x8f, 0x43, 0x26, 0xfc, 0xb6, 0xb8, 0xd1, 0x6d, 0xde, 0x5d, 0x63, 0x0d, 0xaf, 0x58, 0x53, 0x1b,
	0x50, 0x49, 0x2c, 0x64, 0x77, 0x78, 0xf4, 0x12, 0xee, 0x14, 0xb4, 0x04, 0x36, 0xde, 0x73, 0x50,
	0xd6, 0xc2, 0xec, 0x63, 0x6e, 0xf3, 0xba, 0x2d, 0xdd, 0xc2, 0xff, 0x74, 0xa8, 0x15, 0x1a, 0xd5,
	0x41, 0x93, 0x3a, 0xe8, 0xc5, 0x4f, 0x1d, 0xca, 0x31, 0xcd, 0x34, 0x47, 0x37, 0x14, 0xa3, 0x22,
	0xe9, 0x0e, 0xca, 0x99, 0x70, 0x36, 0x29, 0x9f, 0x6c, 0xa1, 0xdc, 0x59, 0x4a, 0x5b, 0x50, 0x4a,
	0xfd, 0x74, 0x6d, 0xad, 0xc0, 0x9a, 0x48, 0x95, 0xa3, 0x5f, 0x4e, 0xad, 0x52, 0x55, 0xd4, 0xf2,
	0xc6, 0x53, 0x82, 0x26, 0x53, 0x82, 0xe6, 0x53, 0x82, 0x9f, 0x03, 0x82, 0xdf, 0x02, 0x82, 0x3f,
	0x02, 0x82, 0xc7, 0x01, 0xc1, 0x9f, 0x01, 0xc1, 0x5f, 0x01, 0x41, 0xf3, 0x80, 0xe0, 0xd7, 0x19,
	0x41, 0xe3, 0x19, 0x41, 0x93, 0x19, 0x41, 0xb0, 0x6f, 0x3a, 0x7d, 0x6a, 0x71, 0xbf, 0x37, 0x34,
	0x68, 0x8f, 0x7b, 0x4e, 0x5f, 0xe7, 0xf6, 0x28, 0x3e, 0x7b, 0x6a, 0x73, 0xa3, 0x95, 0x69, 0xe1,
	0xbe, 0xb8, 0x30, 0x65, 0xe4, 0xc3, 0x6c, 0xce, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x78, 0x24,
	0x9b, 0x73, 0x40, 0x03, 0x00, 0x00,
}

func (this *ResponseGenerateAddress) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ResponseGenerateAddress)
	if !ok {
		that2, ok := that.(ResponseGenerateAddress)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.XAddress != that1.XAddress {
		return false
	}
	if this.ClassicAddress != that1.ClassicAddress {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if this.Secret != that1.Secret {
		return false
	}
	return true
}
func (this *ResponseGenerateXAddress) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ResponseGenerateXAddress)
	if !ok {
		that2, ok := that.(ResponseGenerateXAddress)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.XAddress != that1.XAddress {
		return false
	}
	if this.Secret != that1.Secret {
		return false
	}
	return true
}
func (this *RequestIsValidAddress) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RequestIsValidAddress)
	if !ok {
		that2, ok := that.(RequestIsValidAddress)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	return true
}
func (this *ResponseIsValidAddress) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ResponseIsValidAddress)
	if !ok {
		that2, ok := that.(ResponseIsValidAddress)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.IsValid != that1.IsValid {
		return false
	}
	return true
}
func (this *ResponseGenerateAddress) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&rippleapi.ResponseGenerateAddress{")
	s = append(s, "XAddress: "+fmt.Sprintf("%#v", this.XAddress)+",\n")
	s = append(s, "ClassicAddress: "+fmt.Sprintf("%#v", this.ClassicAddress)+",\n")
	s = append(s, "Address: "+fmt.Sprintf("%#v", this.Address)+",\n")
	s = append(s, "Secret: "+fmt.Sprintf("%#v", this.Secret)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ResponseGenerateXAddress) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&rippleapi.ResponseGenerateXAddress{")
	s = append(s, "XAddress: "+fmt.Sprintf("%#v", this.XAddress)+",\n")
	s = append(s, "Secret: "+fmt.Sprintf("%#v", this.Secret)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *RequestIsValidAddress) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&rippleapi.RequestIsValidAddress{")
	s = append(s, "Address: "+fmt.Sprintf("%#v", this.Address)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ResponseIsValidAddress) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&rippleapi.ResponseIsValidAddress{")
	s = append(s, "IsValid: "+fmt.Sprintf("%#v", this.IsValid)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAddress(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RippleAddressAPIClient is the client API for RippleAddressAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RippleAddressAPIClient interface {
	// https://xrpl.org/rippleapi-reference.html#generateaddress
	GenerateAddress(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*ResponseGenerateAddress, error)
	// https://xrpl.org/rippleapi-reference.html#generatexaddress
	GenerateXAddress(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*ResponseGenerateXAddress, error)
	// https://xrpl.org/rippleapi-reference.html#isvalidaddress
	IsValidAddress(ctx context.Context, in *RequestIsValidAddress, opts ...grpc.CallOption) (*ResponseIsValidAddress, error)
}

type rippleAddressAPIClient struct {
	cc *grpc.ClientConn
}

func NewRippleAddressAPIClient(cc *grpc.ClientConn) RippleAddressAPIClient {
	return &rippleAddressAPIClient{cc}
}

func (c *rippleAddressAPIClient) GenerateAddress(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*ResponseGenerateAddress, error) {
	out := new(ResponseGenerateAddress)
	err := c.cc.Invoke(ctx, "/rippleapi.address.RippleAddressAPI/GenerateAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rippleAddressAPIClient) GenerateXAddress(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*ResponseGenerateXAddress, error) {
	out := new(ResponseGenerateXAddress)
	err := c.cc.Invoke(ctx, "/rippleapi.address.RippleAddressAPI/GenerateXAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rippleAddressAPIClient) IsValidAddress(ctx context.Context, in *RequestIsValidAddress, opts ...grpc.CallOption) (*ResponseIsValidAddress, error) {
	out := new(ResponseIsValidAddress)
	err := c.cc.Invoke(ctx, "/rippleapi.address.RippleAddressAPI/IsValidAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RippleAddressAPIServer is the server API for RippleAddressAPI service.
type RippleAddressAPIServer interface {
	// https://xrpl.org/rippleapi-reference.html#generateaddress
	GenerateAddress(context.Context, *types.Empty) (*ResponseGenerateAddress, error)
	// https://xrpl.org/rippleapi-reference.html#generatexaddress
	GenerateXAddress(context.Context, *types.Empty) (*ResponseGenerateXAddress, error)
	// https://xrpl.org/rippleapi-reference.html#isvalidaddress
	IsValidAddress(context.Context, *RequestIsValidAddress) (*ResponseIsValidAddress, error)
}

// UnimplementedRippleAddressAPIServer can be embedded to have forward compatible implementations.
type UnimplementedRippleAddressAPIServer struct {
}

func (*UnimplementedRippleAddressAPIServer) GenerateAddress(ctx context.Context, req *types.Empty) (*ResponseGenerateAddress, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateAddress not implemented")
}
func (*UnimplementedRippleAddressAPIServer) GenerateXAddress(ctx context.Context, req *types.Empty) (*ResponseGenerateXAddress, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateXAddress not implemented")
}
func (*UnimplementedRippleAddressAPIServer) IsValidAddress(ctx context.Context, req *RequestIsValidAddress) (*ResponseIsValidAddress, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsValidAddress not implemented")
}

func RegisterRippleAddressAPIServer(s *grpc.Server, srv RippleAddressAPIServer) {
	s.RegisterService(&_RippleAddressAPI_serviceDesc, srv)
}

func _RippleAddressAPI_GenerateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RippleAddressAPIServer).GenerateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rippleapi.address.RippleAddressAPI/GenerateAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RippleAddressAPIServer).GenerateAddress(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RippleAddressAPI_GenerateXAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RippleAddressAPIServer).GenerateXAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rippleapi.address.RippleAddressAPI/GenerateXAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RippleAddressAPIServer).GenerateXAddress(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RippleAddressAPI_IsValidAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestIsValidAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RippleAddressAPIServer).IsValidAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rippleapi.address.RippleAddressAPI/IsValidAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RippleAddressAPIServer).IsValidAddress(ctx, req.(*RequestIsValidAddress))
	}
	return interceptor(ctx, in, info, handler)
}

var _RippleAddressAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rippleapi.address.RippleAddressAPI",
	HandlerType: (*RippleAddressAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateAddress",
			Handler:    _RippleAddressAPI_GenerateAddress_Handler,
		},
		{
			MethodName: "GenerateXAddress",
			Handler:    _RippleAddressAPI_GenerateXAddress_Handler,
		},
		{
			MethodName: "IsValidAddress",
			Handler:    _RippleAddressAPI_IsValidAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rippleapi/address.proto",
}

func (m *ResponseGenerateAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseGenerateAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseGenerateAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Secret) > 0 {
		i -= len(m.Secret)
		copy(dAtA[i:], m.Secret)
		i = encodeVarintAddress(dAtA, i, uint64(len(m.Secret)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAddress(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ClassicAddress) > 0 {
		i -= len(m.ClassicAddress)
		copy(dAtA[i:], m.ClassicAddress)
		i = encodeVarintAddress(dAtA, i, uint64(len(m.ClassicAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.XAddress) > 0 {
		i -= len(m.XAddress)
		copy(dAtA[i:], m.XAddress)
		i = encodeVarintAddress(dAtA, i, uint64(len(m.XAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResponseGenerateXAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseGenerateXAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseGenerateXAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Secret) > 0 {
		i -= len(m.Secret)
		copy(dAtA[i:], m.Secret)
		i = encodeVarintAddress(dAtA, i, uint64(len(m.Secret)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.XAddress) > 0 {
		i -= len(m.XAddress)
		copy(dAtA[i:], m.XAddress)
		i = encodeVarintAddress(dAtA, i, uint64(len(m.XAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RequestIsValidAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestIsValidAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestIsValidAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAddress(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResponseIsValidAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseIsValidAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseIsValidAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsValid {
		i--
		if m.IsValid {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAddress(dAtA []byte, offset int, v uint64) int {
	offset -= sovAddress(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ResponseGenerateAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.XAddress)
	if l > 0 {
		n += 1 + l + sovAddress(uint64(l))
	}
	l = len(m.ClassicAddress)
	if l > 0 {
		n += 1 + l + sovAddress(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAddress(uint64(l))
	}
	l = len(m.Secret)
	if l > 0 {
		n += 1 + l + sovAddress(uint64(l))
	}
	return n
}

func (m *ResponseGenerateXAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.XAddress)
	if l > 0 {
		n += 1 + l + sovAddress(uint64(l))
	}
	l = len(m.Secret)
	if l > 0 {
		n += 1 + l + sovAddress(uint64(l))
	}
	return n
}

func (m *RequestIsValidAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAddress(uint64(l))
	}
	return n
}

func (m *ResponseIsValidAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsValid {
		n += 2
	}
	return n
}

func sovAddress(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAddress(x uint64) (n int) {
	return sovAddress(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ResponseGenerateAddress) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ResponseGenerateAddress{`,
		`XAddress:` + fmt.Sprintf("%v", this.XAddress) + `,`,
		`ClassicAddress:` + fmt.Sprintf("%v", this.ClassicAddress) + `,`,
		`Address:` + fmt.Sprintf("%v", this.Address) + `,`,
		`Secret:` + fmt.Sprintf("%v", this.Secret) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ResponseGenerateXAddress) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ResponseGenerateXAddress{`,
		`XAddress:` + fmt.Sprintf("%v", this.XAddress) + `,`,
		`Secret:` + fmt.Sprintf("%v", this.Secret) + `,`,
		`}`,
	}, "")
	return s
}
func (this *RequestIsValidAddress) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RequestIsValidAddress{`,
		`Address:` + fmt.Sprintf("%v", this.Address) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ResponseIsValidAddress) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ResponseIsValidAddress{`,
		`IsValid:` + fmt.Sprintf("%v", this.IsValid) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAddress(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ResponseGenerateAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAddress
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResponseGenerateAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseGenerateAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field XAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.XAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassicAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassicAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secret", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Secret = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAddress(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAddress
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResponseGenerateXAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAddress
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResponseGenerateXAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseGenerateXAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field XAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.XAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secret", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Secret = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAddress(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAddress
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RequestIsValidAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAddress
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RequestIsValidAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestIsValidAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAddress(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAddress
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResponseIsValidAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAddress
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResponseIsValidAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseIsValidAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsValid", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsValid = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipAddress(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAddress
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAddress(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAddress
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthAddress
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAddress
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAddress
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAddress        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAddress          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAddress = fmt.Errorf("proto: unexpected end of group")
)
