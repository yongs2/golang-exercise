// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: eventTimerCallback.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TimerEventMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallbackUri string `protobuf:"bytes,2,opt,name=CallbackUri,proto3" json:"CallbackUri,omitempty"`
	SetTime     string `protobuf:"bytes,3,opt,name=SetTime,proto3" json:"SetTime,omitempty"`
	ExpireSec   int32  `protobuf:"varint,4,opt,name=ExpireSec,proto3" json:"ExpireSec,omitempty"`
	Data        string `protobuf:"bytes,5,opt,name=Data,proto3" json:"Data,omitempty"`
	RepeatCount int32  `protobuf:"varint,6,opt,name=RepeatCount,proto3" json:"RepeatCount,omitempty"`
	TimerId     string `protobuf:"bytes,7,opt,name=TimerId,proto3" json:"TimerId,omitempty"`     // CMD_CREATED, CMD_DELETE, CMD_EVENT
	EventTime   string `protobuf:"bytes,8,opt,name=EventTime,proto3" json:"EventTime,omitempty"` // CMD_EVENT
}

func (x *TimerEventMsg) Reset() {
	*x = TimerEventMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventTimerCallback_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimerEventMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimerEventMsg) ProtoMessage() {}

func (x *TimerEventMsg) ProtoReflect() protoreflect.Message {
	mi := &file_eventTimerCallback_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimerEventMsg.ProtoReflect.Descriptor instead.
func (*TimerEventMsg) Descriptor() ([]byte, []int) {
	return file_eventTimerCallback_proto_rawDescGZIP(), []int{0}
}

func (x *TimerEventMsg) GetCallbackUri() string {
	if x != nil {
		return x.CallbackUri
	}
	return ""
}

func (x *TimerEventMsg) GetSetTime() string {
	if x != nil {
		return x.SetTime
	}
	return ""
}

func (x *TimerEventMsg) GetExpireSec() int32 {
	if x != nil {
		return x.ExpireSec
	}
	return 0
}

func (x *TimerEventMsg) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *TimerEventMsg) GetRepeatCount() int32 {
	if x != nil {
		return x.RepeatCount
	}
	return 0
}

func (x *TimerEventMsg) GetTimerId() string {
	if x != nil {
		return x.TimerId
	}
	return ""
}

func (x *TimerEventMsg) GetEventTime() string {
	if x != nil {
		return x.EventTime
	}
	return ""
}

type TimerEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *TimerEventResponse) Reset() {
	*x = TimerEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventTimerCallback_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimerEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimerEventResponse) ProtoMessage() {}

func (x *TimerEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eventTimerCallback_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimerEventResponse.ProtoReflect.Descriptor instead.
func (*TimerEventResponse) Descriptor() ([]byte, []int) {
	return file_eventTimerCallback_proto_rawDescGZIP(), []int{1}
}

func (x *TimerEventResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_eventTimerCallback_proto protoreflect.FileDescriptor

var file_eventTimerCallback_proto_rawDesc = []byte{
	0x0a, 0x18, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x43, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x72, 0x70, 0x63,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x22, 0xd7, 0x01, 0x0a, 0x0d, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x4d, 0x73, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x55, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x55, 0x72, 0x69, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x53, 0x65, 0x63, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x53, 0x65, 0x63, 0x12, 0x12,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x12,
	0x54, 0x69, 0x6d, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x75, 0x0a, 0x12, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x12, 0x5f, 0x0a, 0x08, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x64, 0x12, 0x25, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x4d, 0x73, 0x67, 0x1a, 0x2a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x74,
	0x69, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x5d, 0x0a, 0x1c, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x69, 0x6d,
	0x65, 0x72, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x42, 0x17, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x43, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x22, 0x30, 0x36,
	0x2e, 0x67, 0x6f, 0x2d, 0x72, 0x70, 0x63, 0x2f, 0x30, 0x34, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x72,
	0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eventTimerCallback_proto_rawDescOnce sync.Once
	file_eventTimerCallback_proto_rawDescData = file_eventTimerCallback_proto_rawDesc
)

func file_eventTimerCallback_proto_rawDescGZIP() []byte {
	file_eventTimerCallback_proto_rawDescOnce.Do(func() {
		file_eventTimerCallback_proto_rawDescData = protoimpl.X.CompressGZIP(file_eventTimerCallback_proto_rawDescData)
	})
	return file_eventTimerCallback_proto_rawDescData
}

var file_eventTimerCallback_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eventTimerCallback_proto_goTypes = []interface{}{
	(*TimerEventMsg)(nil),      // 0: grpceventtimercallback.TimerEventMsg
	(*TimerEventResponse)(nil), // 1: grpceventtimercallback.TimerEventResponse
}
var file_eventTimerCallback_proto_depIdxs = []int32{
	0, // 0: grpceventtimercallback.EventTimerCallback.Occurred:input_type -> grpceventtimercallback.TimerEventMsg
	1, // 1: grpceventtimercallback.EventTimerCallback.Occurred:output_type -> grpceventtimercallback.TimerEventResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eventTimerCallback_proto_init() }
func file_eventTimerCallback_proto_init() {
	if File_eventTimerCallback_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eventTimerCallback_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimerEventMsg); i {
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
		file_eventTimerCallback_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimerEventResponse); i {
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
			RawDescriptor: file_eventTimerCallback_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eventTimerCallback_proto_goTypes,
		DependencyIndexes: file_eventTimerCallback_proto_depIdxs,
		MessageInfos:      file_eventTimerCallback_proto_msgTypes,
	}.Build()
	File_eventTimerCallback_proto = out.File
	file_eventTimerCallback_proto_rawDesc = nil
	file_eventTimerCallback_proto_goTypes = nil
	file_eventTimerCallback_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventTimerCallbackClient is the client API for EventTimerCallback service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventTimerCallbackClient interface {
	Occurred(ctx context.Context, in *TimerEventMsg, opts ...grpc.CallOption) (*TimerEventResponse, error)
}

type eventTimerCallbackClient struct {
	cc grpc.ClientConnInterface
}

func NewEventTimerCallbackClient(cc grpc.ClientConnInterface) EventTimerCallbackClient {
	return &eventTimerCallbackClient{cc}
}

func (c *eventTimerCallbackClient) Occurred(ctx context.Context, in *TimerEventMsg, opts ...grpc.CallOption) (*TimerEventResponse, error) {
	out := new(TimerEventResponse)
	err := c.cc.Invoke(ctx, "/grpceventtimercallback.EventTimerCallback/Occurred", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventTimerCallbackServer is the server API for EventTimerCallback service.
type EventTimerCallbackServer interface {
	Occurred(context.Context, *TimerEventMsg) (*TimerEventResponse, error)
}

// UnimplementedEventTimerCallbackServer can be embedded to have forward compatible implementations.
type UnimplementedEventTimerCallbackServer struct {
}

func (*UnimplementedEventTimerCallbackServer) Occurred(context.Context, *TimerEventMsg) (*TimerEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Occurred not implemented")
}

func RegisterEventTimerCallbackServer(s *grpc.Server, srv EventTimerCallbackServer) {
	s.RegisterService(&_EventTimerCallback_serviceDesc, srv)
}

func _EventTimerCallback_Occurred_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimerEventMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventTimerCallbackServer).Occurred(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpceventtimercallback.EventTimerCallback/Occurred",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventTimerCallbackServer).Occurred(ctx, req.(*TimerEventMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventTimerCallback_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpceventtimercallback.EventTimerCallback",
	HandlerType: (*EventTimerCallbackServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Occurred",
			Handler:    _EventTimerCallback_Occurred_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eventTimerCallback.proto",
}
