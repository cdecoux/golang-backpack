// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: demo.proto

package demo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetRandomNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRandomNumberRequest) Reset() {
	*x = GetRandomNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRandomNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRandomNumberRequest) ProtoMessage() {}

func (x *GetRandomNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRandomNumberRequest.ProtoReflect.Descriptor instead.
func (*GetRandomNumberRequest) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{0}
}

type GetRandomNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RandomNumber int64                  `protobuf:"varint,1,opt,name=random_number,json=randomNumber,proto3" json:"random_number,omitempty"`
	Time         *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *GetRandomNumberResponse) Reset() {
	*x = GetRandomNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRandomNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRandomNumberResponse) ProtoMessage() {}

func (x *GetRandomNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRandomNumberResponse.ProtoReflect.Descriptor instead.
func (*GetRandomNumberResponse) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{1}
}

func (x *GetRandomNumberResponse) GetRandomNumber() int64 {
	if x != nil {
		return x.RandomNumber
	}
	return 0
}

func (x *GetRandomNumberResponse) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

var File_demo_proto protoreflect.FileDescriptor

var file_demo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x65,
	0x6d, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6e, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x61, 0x6e, 0x64,
	0x6f, 0x6d, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2e, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x32, 0xb9, 0x01,
	0x0a, 0x0c, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x50,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x1c, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x64,
	0x6f, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x57, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1c, 0x2e, 0x64, 0x65, 0x6d, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x30, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x64,
	0x65, 0x6d, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_demo_proto_rawDescOnce sync.Once
	file_demo_proto_rawDescData = file_demo_proto_rawDesc
)

func file_demo_proto_rawDescGZIP() []byte {
	file_demo_proto_rawDescOnce.Do(func() {
		file_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_demo_proto_rawDescData)
	})
	return file_demo_proto_rawDescData
}

var file_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_demo_proto_goTypes = []interface{}{
	(*GetRandomNumberRequest)(nil),  // 0: demo.GetRandomNumberRequest
	(*GetRandomNumberResponse)(nil), // 1: demo.GetRandomNumberResponse
	(*timestamppb.Timestamp)(nil),   // 2: google.protobuf.Timestamp
}
var file_demo_proto_depIdxs = []int32{
	2, // 0: demo.GetRandomNumberResponse.time:type_name -> google.protobuf.Timestamp
	0, // 1: demo.RandomNumber.GetRandomNumber:input_type -> demo.GetRandomNumberRequest
	0, // 2: demo.RandomNumber.GetRandomNumberStream:input_type -> demo.GetRandomNumberRequest
	1, // 3: demo.RandomNumber.GetRandomNumber:output_type -> demo.GetRandomNumberResponse
	0, // 4: demo.RandomNumber.GetRandomNumberStream:output_type -> demo.GetRandomNumberRequest
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_demo_proto_init() }
func file_demo_proto_init() {
	if File_demo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_demo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRandomNumberRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_demo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRandomNumberResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_demo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_demo_proto_goTypes,
		DependencyIndexes: file_demo_proto_depIdxs,
		MessageInfos:      file_demo_proto_msgTypes,
	}.Build()
	File_demo_proto = out.File
	file_demo_proto_rawDesc = nil
	file_demo_proto_goTypes = nil
	file_demo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RandomNumberClient is the client API for RandomNumber service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RandomNumberClient interface {
	GetRandomNumber(ctx context.Context, in *GetRandomNumberRequest, opts ...grpc.CallOption) (*GetRandomNumberResponse, error)
	GetRandomNumberStream(ctx context.Context, in *GetRandomNumberRequest, opts ...grpc.CallOption) (RandomNumber_GetRandomNumberStreamClient, error)
}

type randomNumberClient struct {
	cc grpc.ClientConnInterface
}

func NewRandomNumberClient(cc grpc.ClientConnInterface) RandomNumberClient {
	return &randomNumberClient{cc}
}

func (c *randomNumberClient) GetRandomNumber(ctx context.Context, in *GetRandomNumberRequest, opts ...grpc.CallOption) (*GetRandomNumberResponse, error) {
	out := new(GetRandomNumberResponse)
	err := c.cc.Invoke(ctx, "/demo.RandomNumber/GetRandomNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *randomNumberClient) GetRandomNumberStream(ctx context.Context, in *GetRandomNumberRequest, opts ...grpc.CallOption) (RandomNumber_GetRandomNumberStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RandomNumber_serviceDesc.Streams[0], "/demo.RandomNumber/GetRandomNumberStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &randomNumberGetRandomNumberStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RandomNumber_GetRandomNumberStreamClient interface {
	Recv() (*GetRandomNumberRequest, error)
	grpc.ClientStream
}

type randomNumberGetRandomNumberStreamClient struct {
	grpc.ClientStream
}

func (x *randomNumberGetRandomNumberStreamClient) Recv() (*GetRandomNumberRequest, error) {
	m := new(GetRandomNumberRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RandomNumberServer is the server API for RandomNumber service.
type RandomNumberServer interface {
	GetRandomNumber(context.Context, *GetRandomNumberRequest) (*GetRandomNumberResponse, error)
	GetRandomNumberStream(*GetRandomNumberRequest, RandomNumber_GetRandomNumberStreamServer) error
}

// UnimplementedRandomNumberServer can be embedded to have forward compatible implementations.
type UnimplementedRandomNumberServer struct {
}

func (*UnimplementedRandomNumberServer) GetRandomNumber(context.Context, *GetRandomNumberRequest) (*GetRandomNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRandomNumber not implemented")
}
func (*UnimplementedRandomNumberServer) GetRandomNumberStream(*GetRandomNumberRequest, RandomNumber_GetRandomNumberStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRandomNumberStream not implemented")
}

func RegisterRandomNumberServer(s *grpc.Server, srv RandomNumberServer) {
	s.RegisterService(&_RandomNumber_serviceDesc, srv)
}

func _RandomNumber_GetRandomNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRandomNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomNumberServer).GetRandomNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.RandomNumber/GetRandomNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomNumberServer).GetRandomNumber(ctx, req.(*GetRandomNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RandomNumber_GetRandomNumberStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRandomNumberRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RandomNumberServer).GetRandomNumberStream(m, &randomNumberGetRandomNumberStreamServer{stream})
}

type RandomNumber_GetRandomNumberStreamServer interface {
	Send(*GetRandomNumberRequest) error
	grpc.ServerStream
}

type randomNumberGetRandomNumberStreamServer struct {
	grpc.ServerStream
}

func (x *randomNumberGetRandomNumberStreamServer) Send(m *GetRandomNumberRequest) error {
	return x.ServerStream.SendMsg(m)
}

var _RandomNumber_serviceDesc = grpc.ServiceDesc{
	ServiceName: "demo.RandomNumber",
	HandlerType: (*RandomNumberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRandomNumber",
			Handler:    _RandomNumber_GetRandomNumber_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetRandomNumberStream",
			Handler:       _RandomNumber_GetRandomNumberStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "demo.proto",
}
