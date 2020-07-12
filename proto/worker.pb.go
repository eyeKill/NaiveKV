// Code generated by protoc-gen-go. DO NOT EDIT.
// source: worker.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type PutResponse struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=kv.proto.Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutResponse) Reset()         { *m = PutResponse{} }
func (m *PutResponse) String() string { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()    {}
func (*PutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4ff6184b07e587a, []int{0}
}

func (m *PutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutResponse.Unmarshal(m, b)
}
func (m *PutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutResponse.Marshal(b, m, deterministic)
}
func (m *PutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutResponse.Merge(m, src)
}
func (m *PutResponse) XXX_Size() int {
	return xxx_messageInfo_PutResponse.Size(m)
}
func (m *PutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutResponse proto.InternalMessageInfo

func (m *PutResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_OK
}

type GetResponse struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=kv.proto.Status" json:"status,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4ff6184b07e587a, []int{1}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_OK
}

func (m *GetResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type DeleteResponse struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=kv.proto.Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4ff6184b07e587a, []int{2}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_OK
}

func init() {
	proto.RegisterType((*PutResponse)(nil), "kv.proto.PutResponse")
	proto.RegisterType((*GetResponse)(nil), "kv.proto.GetResponse")
	proto.RegisterType((*DeleteResponse)(nil), "kv.proto.DeleteResponse")
}

func init() {
	proto.RegisterFile("worker.proto", fileDescriptor_e4ff6184b07e587a)
}

var fileDescriptor_e4ff6184b07e587a = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xcf, 0x2f, 0xca,
	0x4e, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0x2e, 0x83, 0xb0, 0xa4, 0x78,
	0x92, 0xf3, 0x73, 0x73, 0xf3, 0xf3, 0x20, 0x3c, 0x25, 0x73, 0x2e, 0xee, 0x80, 0xd2, 0x92, 0xa0,
	0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x0d, 0x2e, 0xb6, 0xe2, 0x92, 0xc4, 0x92, 0xd2,
	0x62, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x3e, 0x23, 0x01, 0x3d, 0x98, 0x3e, 0xbd, 0x60, 0xb0, 0x78,
	0x10, 0x54, 0x5e, 0xc9, 0x97, 0x8b, 0xdb, 0x3d, 0x95, 0x0c, 0x8d, 0x42, 0x22, 0x5c, 0xac, 0x65,
	0x89, 0x39, 0xa5, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x92, 0x15, 0x17,
	0x9f, 0x4b, 0x6a, 0x4e, 0x6a, 0x49, 0x2a, 0xe9, 0x26, 0x1a, 0x2d, 0x60, 0xe4, 0xe2, 0xf0, 0x0e,
	0x0b, 0x07, 0x7b, 0x57, 0xc8, 0x80, 0x8b, 0xb9, 0xa0, 0xb4, 0x44, 0x08, 0x49, 0xb5, 0x77, 0x58,
	0x40, 0x62, 0x66, 0x91, 0x94, 0x28, 0x42, 0x04, 0xc9, 0xc7, 0x4a, 0x0c, 0x42, 0xba, 0x5c, 0xcc,
	0xe9, 0xa9, 0x25, 0x42, 0xbc, 0x48, 0x3a, 0x52, 0x2b, 0x91, 0x95, 0x23, 0xf9, 0x53, 0x89, 0x41,
	0xc8, 0x98, 0x8b, 0x2d, 0x05, 0xec, 0x52, 0x74, 0x1d, 0x12, 0x08, 0x2e, 0xaa, 0x57, 0x94, 0x18,
	0x9c, 0x38, 0xa3, 0xd8, 0xf5, 0xac, 0xc1, 0x72, 0x49, 0x6c, 0x60, 0xca, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x9a, 0xf3, 0xba, 0x02, 0xa0, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KVWorkerClient is the client API for KVWorker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KVWorkerClient interface {
	Put(ctx context.Context, in *KVPair, opts ...grpc.CallOption) (*PutResponse, error)
	Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*GetResponse, error)
	Delete(ctx context.Context, in *Key, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type kVWorkerClient struct {
	cc grpc.ClientConnInterface
}

func NewKVWorkerClient(cc grpc.ClientConnInterface) KVWorkerClient {
	return &kVWorkerClient{cc}
}

func (c *kVWorkerClient) Put(ctx context.Context, in *KVPair, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/kv.proto.KVWorker/put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVWorkerClient) Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/kv.proto.KVWorker/get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVWorkerClient) Delete(ctx context.Context, in *Key, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/kv.proto.KVWorker/delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KVWorkerServer is the server API for KVWorker service.
type KVWorkerServer interface {
	Put(context.Context, *KVPair) (*PutResponse, error)
	Get(context.Context, *Key) (*GetResponse, error)
	Delete(context.Context, *Key) (*DeleteResponse, error)
}

// UnimplementedKVWorkerServer can be embedded to have forward compatible implementations.
type UnimplementedKVWorkerServer struct {
}

func (*UnimplementedKVWorkerServer) Put(ctx context.Context, req *KVPair) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (*UnimplementedKVWorkerServer) Get(ctx context.Context, req *Key) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedKVWorkerServer) Delete(ctx context.Context, req *Key) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterKVWorkerServer(s *grpc.Server, srv KVWorkerServer) {
	s.RegisterService(&_KVWorker_serviceDesc, srv)
}

func _KVWorker_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KVPair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVWorkerServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kv.proto.KVWorker/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVWorkerServer).Put(ctx, req.(*KVPair))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVWorker_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVWorkerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kv.proto.KVWorker/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVWorkerServer).Get(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVWorker_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVWorkerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kv.proto.KVWorker/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVWorkerServer).Delete(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

var _KVWorker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kv.proto.KVWorker",
	HandlerType: (*KVWorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "put",
			Handler:    _KVWorker_Put_Handler,
		},
		{
			MethodName: "get",
			Handler:    _KVWorker_Get_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _KVWorker_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "worker.proto",
}
