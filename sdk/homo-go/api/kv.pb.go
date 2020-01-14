// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/kv.proto

package api

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	types "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// KV kv message
type KV struct {
	// key is the key, in bytes, to put into the key-value store.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// value is the value, in bytes, to associate with the key in the key-value store.
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KV) Reset()         { *m = KV{} }
func (m *KV) String() string { return proto.CompactTextString(m) }
func (*KV) ProtoMessage()    {}
func (*KV) Descriptor() ([]byte, []int) {
	return fileDescriptor_1604c0fd7bb99b7b, []int{0}
}

func (m *KV) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KV.Unmarshal(m, b)
}
func (m *KV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KV.Marshal(b, m, deterministic)
}
func (m *KV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KV.Merge(m, src)
}
func (m *KV) XXX_Size() int {
	return xxx_messageInfo_KV.Size(m)
}
func (m *KV) XXX_DiscardUnknown() {
	xxx_messageInfo_KV.DiscardUnknown(m)
}

var xxx_messageInfo_KV proto.InternalMessageInfo

func (m *KV) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KV) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// KVs kvs message
type KVs struct {
	Kvs                  []*KV    `protobuf:"bytes,1,rep,name=kvs,proto3" json:"kvs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KVs) Reset()         { *m = KVs{} }
func (m *KVs) String() string { return proto.CompactTextString(m) }
func (*KVs) ProtoMessage()    {}
func (*KVs) Descriptor() ([]byte, []int) {
	return fileDescriptor_1604c0fd7bb99b7b, []int{1}
}

func (m *KVs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVs.Unmarshal(m, b)
}
func (m *KVs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVs.Marshal(b, m, deterministic)
}
func (m *KVs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVs.Merge(m, src)
}
func (m *KVs) XXX_Size() int {
	return xxx_messageInfo_KVs.Size(m)
}
func (m *KVs) XXX_DiscardUnknown() {
	xxx_messageInfo_KVs.DiscardUnknown(m)
}

var xxx_messageInfo_KVs proto.InternalMessageInfo

func (m *KVs) GetKvs() []*KV {
	if m != nil {
		return m.Kvs
	}
	return nil
}

func init() {
	proto.RegisterType((*KV)(nil), "api.KV")
	proto.RegisterType((*KVs)(nil), "api.KVs")
}

func init() { proto.RegisterFile("api/kv.proto", fileDescriptor_1604c0fd7bb99b7b) }

var fileDescriptor_1604c0fd7bb99b7b = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x2c, 0xc8, 0xd4,
	0xcf, 0x2e, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x94, 0x92, 0x4e,
	0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x0b, 0x25, 0x95, 0xa6, 0xe9, 0xa7, 0xe6, 0x16, 0x94,
	0x54, 0x42, 0x54, 0x48, 0xe9, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea,
	0xa7, 0xe7, 0xa7, 0xe7, 0x23, 0x54, 0x81, 0x78, 0x60, 0x0e, 0x98, 0x05, 0x51, 0xae, 0xa4, 0xc3,
	0xc5, 0xe4, 0x1d, 0x26, 0x24, 0xc0, 0xc5, 0x9c, 0x9d, 0x5a, 0x29, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0x13, 0x04, 0x62, 0x0a, 0x89, 0x70, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x81, 0xc5,
	0x20, 0x1c, 0x25, 0x05, 0x2e, 0x66, 0xef, 0xb0, 0x62, 0x21, 0x49, 0x2e, 0xe6, 0xec, 0xb2, 0x62,
	0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x76, 0xbd, 0xc4, 0x82, 0x4c, 0x3d, 0xef, 0xb0, 0x20,
	0x90, 0x98, 0xd1, 0x74, 0x46, 0x2e, 0x4e, 0xef, 0xb0, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54,
	0x21, 0x0d, 0x2e, 0xe6, 0xe0, 0xd4, 0x12, 0x21, 0x98, 0x12, 0x29, 0x31, 0x3d, 0x88, 0xd3, 0xf5,
	0x60, 0x8e, 0xd2, 0x73, 0x05, 0x39, 0x5d, 0x89, 0x01, 0x64, 0xa4, 0x3b, 0xb2, 0x4a, 0x18, 0x43,
	0x89, 0x01, 0x64, 0x88, 0x4b, 0x6a, 0x0e, 0x31, 0x86, 0x48, 0x73, 0xb1, 0xf8, 0x64, 0x16, 0x23,
	0x99, 0xc2, 0x01, 0x65, 0x14, 0x2b, 0x31, 0x38, 0x29, 0x9c, 0x78, 0x28, 0xc7, 0xf0, 0xe3, 0xa1,
	0x1c, 0xe3, 0x8a, 0x47, 0x72, 0x8c, 0x3b, 0x1e, 0xc9, 0x31, 0x1e, 0x78, 0x24, 0xc7, 0x78, 0xe2,
	0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x26, 0xb1, 0x81, 0x0d, 0x34,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xa6, 0xe9, 0xe9, 0x73, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KVServiceClient is the client API for KVService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KVServiceClient interface {
	Set(ctx context.Context, in *KV, opts ...grpc.CallOption) (*types.Empty, error)
	Get(ctx context.Context, in *KV, opts ...grpc.CallOption) (*KV, error)
	Del(ctx context.Context, in *KV, opts ...grpc.CallOption) (*types.Empty, error)
	List(ctx context.Context, in *KV, opts ...grpc.CallOption) (*KVs, error)
}

type kVServiceClient struct {
	cc *grpc.ClientConn
}

func NewKVServiceClient(cc *grpc.ClientConn) KVServiceClient {
	return &kVServiceClient{cc}
}

func (c *kVServiceClient) Set(ctx context.Context, in *KV, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/api.KVService/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVServiceClient) Get(ctx context.Context, in *KV, opts ...grpc.CallOption) (*KV, error) {
	out := new(KV)
	err := c.cc.Invoke(ctx, "/api.KVService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVServiceClient) Del(ctx context.Context, in *KV, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/api.KVService/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVServiceClient) List(ctx context.Context, in *KV, opts ...grpc.CallOption) (*KVs, error) {
	out := new(KVs)
	err := c.cc.Invoke(ctx, "/api.KVService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KVServiceServer is the server API for KVService service.
type KVServiceServer interface {
	Set(context.Context, *KV) (*types.Empty, error)
	Get(context.Context, *KV) (*KV, error)
	Del(context.Context, *KV) (*types.Empty, error)
	List(context.Context, *KV) (*KVs, error)
}

func RegisterKVServiceServer(s *grpc.Server, srv KVServiceServer) {
	s.RegisterService(&_KVService_serviceDesc, srv)
}

func _KVService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KV)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.KVService/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServiceServer).Set(ctx, req.(*KV))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KV)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.KVService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServiceServer).Get(ctx, req.(*KV))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVService_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KV)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServiceServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.KVService/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServiceServer).Del(ctx, req.(*KV))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KV)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.KVService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServiceServer).List(ctx, req.(*KV))
	}
	return interceptor(ctx, in, info, handler)
}

var _KVService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.KVService",
	HandlerType: (*KVServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _KVService_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _KVService_Get_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _KVService_Del_Handler,
		},
		{
			MethodName: "List",
			Handler:    _KVService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/kv.proto",
}
