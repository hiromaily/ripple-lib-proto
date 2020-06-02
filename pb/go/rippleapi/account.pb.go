// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rippleapi/account.proto

package rippleapi

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type RequestGetAccountInfo struct {
	Address       string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	LedgerVersion uint64 `protobuf:"varint,2,opt,name=ledgerVersion,proto3" json:"ledgerVersion,omitempty"`
}

func (m *RequestGetAccountInfo) Reset()      { *m = RequestGetAccountInfo{} }
func (*RequestGetAccountInfo) ProtoMessage() {}
func (*RequestGetAccountInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d4e74ae4f9bff72, []int{0}
}
func (m *RequestGetAccountInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestGetAccountInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestGetAccountInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestGetAccountInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestGetAccountInfo.Merge(m, src)
}
func (m *RequestGetAccountInfo) XXX_Size() int {
	return m.Size()
}
func (m *RequestGetAccountInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestGetAccountInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RequestGetAccountInfo proto.InternalMessageInfo

func (m *RequestGetAccountInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RequestGetAccountInfo) GetLedgerVersion() uint64 {
	if m != nil {
		return m.LedgerVersion
	}
	return 0
}

type ResponseGetAccountInfo struct {
	Sequence                                  string `protobuf:"bytes,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	XrpBalance                                string `protobuf:"bytes,2,opt,name=xrpBalance,proto3" json:"xrpBalance,omitempty"`
	OwnerCount                                uint64 `protobuf:"varint,3,opt,name=ownerCount,proto3" json:"ownerCount,omitempty"`
	PreviousAffectingTransactionID            string `protobuf:"bytes,4,opt,name=previousAffectingTransactionID,proto3" json:"previousAffectingTransactionID,omitempty"`
	PreviousAffectingTransactionLedgerVersion uint64 `protobuf:"varint,5,opt,name=previousAffectingTransactionLedgerVersion,proto3" json:"previousAffectingTransactionLedgerVersion,omitempty"`
}

func (m *ResponseGetAccountInfo) Reset()      { *m = ResponseGetAccountInfo{} }
func (*ResponseGetAccountInfo) ProtoMessage() {}
func (*ResponseGetAccountInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d4e74ae4f9bff72, []int{1}
}
func (m *ResponseGetAccountInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseGetAccountInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseGetAccountInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseGetAccountInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseGetAccountInfo.Merge(m, src)
}
func (m *ResponseGetAccountInfo) XXX_Size() int {
	return m.Size()
}
func (m *ResponseGetAccountInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseGetAccountInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseGetAccountInfo proto.InternalMessageInfo

func (m *ResponseGetAccountInfo) GetSequence() string {
	if m != nil {
		return m.Sequence
	}
	return ""
}

func (m *ResponseGetAccountInfo) GetXrpBalance() string {
	if m != nil {
		return m.XrpBalance
	}
	return ""
}

func (m *ResponseGetAccountInfo) GetOwnerCount() uint64 {
	if m != nil {
		return m.OwnerCount
	}
	return 0
}

func (m *ResponseGetAccountInfo) GetPreviousAffectingTransactionID() string {
	if m != nil {
		return m.PreviousAffectingTransactionID
	}
	return ""
}

func (m *ResponseGetAccountInfo) GetPreviousAffectingTransactionLedgerVersion() uint64 {
	if m != nil {
		return m.PreviousAffectingTransactionLedgerVersion
	}
	return 0
}

func init() {
	proto.RegisterType((*RequestGetAccountInfo)(nil), "rippleapi.account.RequestGetAccountInfo")
	proto.RegisterType((*ResponseGetAccountInfo)(nil), "rippleapi.account.ResponseGetAccountInfo")
}

func init() { proto.RegisterFile("rippleapi/account.proto", fileDescriptor_3d4e74ae4f9bff72) }

var fileDescriptor_3d4e74ae4f9bff72 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0xcd, 0xf4, 0xeb, 0xf7, 0x7d, 0x76, 0x40, 0xd1, 0x80, 0x1a, 0xba, 0xb8, 0x96, 0xe2, 0xa2,
	0x5d, 0x18, 0x41, 0x9f, 0xa0, 0x55, 0x94, 0x82, 0x0b, 0x09, 0xa2, 0x20, 0x6e, 0xa6, 0xe9, 0x6d,
	0x3a, 0x90, 0xce, 0xc4, 0x99, 0xc4, 0x1f, 0xdc, 0xf8, 0x08, 0xae, 0x7d, 0x02, 0x1f, 0xc5, 0x65,
	0x97, 0x5d, 0xda, 0x74, 0xe3, 0xb2, 0x8f, 0x20, 0x89, 0xb1, 0xf4, 0x0f, 0x75, 0x79, 0xcf, 0xb9,
	0x39, 0xe7, 0xe6, 0x9c, 0xa1, 0x9b, 0x8a, 0x07, 0x81, 0x8f, 0x2c, 0xe0, 0xbb, 0xcc, 0x75, 0x65,
	0x24, 0x42, 0x3b, 0x50, 0x32, 0x94, 0xe6, 0xda, 0x98, 0xb0, 0x33, 0xa2, 0x7c, 0x41, 0xd7, 0x1d,
	0xbc, 0x8e, 0x50, 0x87, 0xc7, 0x18, 0xd6, 0x3e, 0xc1, 0x86, 0x68, 0x4b, 0xd3, 0xa2, 0xff, 0x59,
	0xab, 0xa5, 0x50, 0x6b, 0x8b, 0x94, 0x48, 0xa5, 0xe0, 0x7c, 0x8d, 0xe6, 0x36, 0x5d, 0xf6, 0xb1,
	0xe5, 0xa1, 0x3a, 0x47, 0xa5, 0xb9, 0x14, 0x56, 0xae, 0x44, 0x2a, 0x79, 0x67, 0x1a, 0x2c, 0x3f,
	0xe7, 0xe8, 0x86, 0x83, 0x3a, 0x90, 0x42, 0xe3, 0x8c, 0x74, 0x91, 0x2e, 0xe9, 0xc4, 0x53, 0xb8,
	0x98, 0x69, 0x8f, 0x67, 0x13, 0x28, 0xbd, 0x53, 0x41, 0x9d, 0xf9, 0x2c, 0x61, 0x73, 0x29, 0x3b,
	0x81, 0x24, 0xbc, 0xbc, 0x15, 0xa8, 0x0e, 0x12, 0x35, 0xeb, 0x4f, 0xea, 0x3c, 0x81, 0x98, 0x47,
	0x14, 0x02, 0x85, 0x37, 0x5c, 0x46, 0xba, 0xd6, 0x6e, 0xa3, 0x1b, 0x72, 0xe1, 0x9d, 0x29, 0x26,
	0x34, 0x73, 0x43, 0x2e, 0x45, 0xe3, 0xd0, 0xca, 0xa7, 0x9a, 0x3f, 0x6c, 0x99, 0x57, 0xb4, 0xfa,
	0xdd, 0xc6, 0xc9, 0x54, 0x00, 0x7f, 0xd3, 0x33, 0x7e, 0xff, 0xc1, 0xde, 0x03, 0x5d, 0x75, 0xd2,
	0x2a, 0xb2, 0x58, 0x6a, 0xa7, 0x0d, 0xd3, 0xa3, 0x2b, 0x33, 0x39, 0x55, 0xec, 0xb9, 0xbe, 0xec,
	0x85, 0x65, 0x15, 0xab, 0x0b, 0x37, 0x17, 0x85, 0x5f, 0x36, 0xea, 0xaa, 0x37, 0x00, 0xa3, 0x3f,
	0x00, 0x63, 0x34, 0x00, 0xf2, 0x18, 0x03, 0x79, 0x89, 0x81, 0xbc, 0xc6, 0x40, 0x7a, 0x31, 0x90,
	0xb7, 0x18, 0xc8, 0x7b, 0x0c, 0xc6, 0x28, 0x06, 0xf2, 0x34, 0x04, 0xa3, 0x37, 0x04, 0xa3, 0x3f,
	0x04, 0x83, 0x6e, 0xb9, 0xb2, 0x6b, 0x7b, 0x3c, 0xec, 0x44, 0x4d, 0xbb, 0xc3, 0x95, 0xec, 0x32,
	0xee, 0xdf, 0x67, 0xb6, 0x3b, 0x3e, 0x6f, 0xd6, 0xe7, 0xfe, 0xe8, 0xb2, 0x30, 0x3e, 0xaa, 0xf9,
	0x2f, 0x7d, 0x80, 0xfb, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x08, 0x7f, 0xfe, 0x9b, 0x02,
	0x00, 0x00,
}

func (this *RequestGetAccountInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RequestGetAccountInfo)
	if !ok {
		that2, ok := that.(RequestGetAccountInfo)
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
	if this.LedgerVersion != that1.LedgerVersion {
		return false
	}
	return true
}
func (this *ResponseGetAccountInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ResponseGetAccountInfo)
	if !ok {
		that2, ok := that.(ResponseGetAccountInfo)
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
	if this.Sequence != that1.Sequence {
		return false
	}
	if this.XrpBalance != that1.XrpBalance {
		return false
	}
	if this.OwnerCount != that1.OwnerCount {
		return false
	}
	if this.PreviousAffectingTransactionID != that1.PreviousAffectingTransactionID {
		return false
	}
	if this.PreviousAffectingTransactionLedgerVersion != that1.PreviousAffectingTransactionLedgerVersion {
		return false
	}
	return true
}
func (this *RequestGetAccountInfo) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&rippleapi.RequestGetAccountInfo{")
	s = append(s, "Address: "+fmt.Sprintf("%#v", this.Address)+",\n")
	s = append(s, "LedgerVersion: "+fmt.Sprintf("%#v", this.LedgerVersion)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ResponseGetAccountInfo) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&rippleapi.ResponseGetAccountInfo{")
	s = append(s, "Sequence: "+fmt.Sprintf("%#v", this.Sequence)+",\n")
	s = append(s, "XrpBalance: "+fmt.Sprintf("%#v", this.XrpBalance)+",\n")
	s = append(s, "OwnerCount: "+fmt.Sprintf("%#v", this.OwnerCount)+",\n")
	s = append(s, "PreviousAffectingTransactionID: "+fmt.Sprintf("%#v", this.PreviousAffectingTransactionID)+",\n")
	s = append(s, "PreviousAffectingTransactionLedgerVersion: "+fmt.Sprintf("%#v", this.PreviousAffectingTransactionLedgerVersion)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAccount(v interface{}, typ string) string {
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

// RippleAccountAPIClient is the client API for RippleAccountAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RippleAccountAPIClient interface {
	// https://xrpl.org/rippleapi-reference.html#getaccountinfo
	GetAccountInfo(ctx context.Context, in *RequestGetAccountInfo, opts ...grpc.CallOption) (*ResponseGetAccountInfo, error)
}

type rippleAccountAPIClient struct {
	cc *grpc.ClientConn
}

func NewRippleAccountAPIClient(cc *grpc.ClientConn) RippleAccountAPIClient {
	return &rippleAccountAPIClient{cc}
}

func (c *rippleAccountAPIClient) GetAccountInfo(ctx context.Context, in *RequestGetAccountInfo, opts ...grpc.CallOption) (*ResponseGetAccountInfo, error) {
	out := new(ResponseGetAccountInfo)
	err := c.cc.Invoke(ctx, "/rippleapi.account.RippleAccountAPI/GetAccountInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RippleAccountAPIServer is the server API for RippleAccountAPI service.
type RippleAccountAPIServer interface {
	// https://xrpl.org/rippleapi-reference.html#getaccountinfo
	GetAccountInfo(context.Context, *RequestGetAccountInfo) (*ResponseGetAccountInfo, error)
}

// UnimplementedRippleAccountAPIServer can be embedded to have forward compatible implementations.
type UnimplementedRippleAccountAPIServer struct {
}

func (*UnimplementedRippleAccountAPIServer) GetAccountInfo(ctx context.Context, req *RequestGetAccountInfo) (*ResponseGetAccountInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountInfo not implemented")
}

func RegisterRippleAccountAPIServer(s *grpc.Server, srv RippleAccountAPIServer) {
	s.RegisterService(&_RippleAccountAPI_serviceDesc, srv)
}

func _RippleAccountAPI_GetAccountInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetAccountInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RippleAccountAPIServer).GetAccountInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rippleapi.account.RippleAccountAPI/GetAccountInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RippleAccountAPIServer).GetAccountInfo(ctx, req.(*RequestGetAccountInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _RippleAccountAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rippleapi.account.RippleAccountAPI",
	HandlerType: (*RippleAccountAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccountInfo",
			Handler:    _RippleAccountAPI_GetAccountInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rippleapi/account.proto",
}

func (m *RequestGetAccountInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestGetAccountInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestGetAccountInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LedgerVersion != 0 {
		i = encodeVarintAccount(dAtA, i, uint64(m.LedgerVersion))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResponseGetAccountInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseGetAccountInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseGetAccountInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PreviousAffectingTransactionLedgerVersion != 0 {
		i = encodeVarintAccount(dAtA, i, uint64(m.PreviousAffectingTransactionLedgerVersion))
		i--
		dAtA[i] = 0x28
	}
	if len(m.PreviousAffectingTransactionID) > 0 {
		i -= len(m.PreviousAffectingTransactionID)
		copy(dAtA[i:], m.PreviousAffectingTransactionID)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.PreviousAffectingTransactionID)))
		i--
		dAtA[i] = 0x22
	}
	if m.OwnerCount != 0 {
		i = encodeVarintAccount(dAtA, i, uint64(m.OwnerCount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.XrpBalance) > 0 {
		i -= len(m.XrpBalance)
		copy(dAtA[i:], m.XrpBalance)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.XrpBalance)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sequence) > 0 {
		i -= len(m.Sequence)
		copy(dAtA[i:], m.Sequence)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Sequence)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RequestGetAccountInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	if m.LedgerVersion != 0 {
		n += 1 + sovAccount(uint64(m.LedgerVersion))
	}
	return n
}

func (m *ResponseGetAccountInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sequence)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	l = len(m.XrpBalance)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	if m.OwnerCount != 0 {
		n += 1 + sovAccount(uint64(m.OwnerCount))
	}
	l = len(m.PreviousAffectingTransactionID)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	if m.PreviousAffectingTransactionLedgerVersion != 0 {
		n += 1 + sovAccount(uint64(m.PreviousAffectingTransactionLedgerVersion))
	}
	return n
}

func sovAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccount(x uint64) (n int) {
	return sovAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *RequestGetAccountInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RequestGetAccountInfo{`,
		`Address:` + fmt.Sprintf("%v", this.Address) + `,`,
		`LedgerVersion:` + fmt.Sprintf("%v", this.LedgerVersion) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ResponseGetAccountInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ResponseGetAccountInfo{`,
		`Sequence:` + fmt.Sprintf("%v", this.Sequence) + `,`,
		`XrpBalance:` + fmt.Sprintf("%v", this.XrpBalance) + `,`,
		`OwnerCount:` + fmt.Sprintf("%v", this.OwnerCount) + `,`,
		`PreviousAffectingTransactionID:` + fmt.Sprintf("%v", this.PreviousAffectingTransactionID) + `,`,
		`PreviousAffectingTransactionLedgerVersion:` + fmt.Sprintf("%v", this.PreviousAffectingTransactionLedgerVersion) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAccount(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *RequestGetAccountInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
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
			return fmt.Errorf("proto: RequestGetAccountInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestGetAccountInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LedgerVersion", wireType)
			}
			m.LedgerVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LedgerVersion |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAccount
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
func (m *ResponseGetAccountInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
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
			return fmt.Errorf("proto: ResponseGetAccountInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseGetAccountInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sequence = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field XrpBalance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.XrpBalance = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerCount", wireType)
			}
			m.OwnerCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OwnerCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousAffectingTransactionID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreviousAffectingTransactionID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousAffectingTransactionLedgerVersion", wireType)
			}
			m.PreviousAffectingTransactionLedgerVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PreviousAffectingTransactionLedgerVersion |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAccount
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
func skipAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccount
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
					return 0, ErrIntOverflowAccount
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
					return 0, ErrIntOverflowAccount
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
				return 0, ErrInvalidLengthAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccount = fmt.Errorf("proto: unexpected end of group")
)