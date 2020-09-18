// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: eventTimer.proto

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

type TimerCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallbackUri string `protobuf:"bytes,1,opt,name=CallbackUri,proto3" json:"CallbackUri,omitempty"`
	SetTime     string `protobuf:"bytes,2,opt,name=SetTime,proto3" json:"SetTime,omitempty"`
	ExpireSec   int32  `protobuf:"varint,3,opt,name=ExpireSec,proto3" json:"ExpireSec,omitempty"`
	Data        string `protobuf:"bytes,4,opt,name=Data,proto3" json:"Data,omitempty"`
	RepeatCount int32  `protobuf:"varint,5,opt,name=RepeatCount,proto3" json:"RepeatCount,omitempty"`
}

func (x *TimerCreateRequest) Reset() {
	*x = TimerCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventTimer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimerCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimerCreateRequest) ProtoMessage() {}

func (x *TimerCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eventTimer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimerCreateRequest.ProtoReflect.Descriptor instead.
func (*TimerCreateRequest) Descriptor() ([]byte, []int) {
	return file_eventTimer_proto_rawDescGZIP(), []int{0}
}

func (x *TimerCreateRequest) GetCallbackUri() string {
	if x != nil {
		return x.CallbackUri
	}
	return ""
}

func (x *TimerCreateRequest) GetSetTime() string {
	if x != nil {
		return x.SetTime
	}
	return ""
}

func (x *TimerCreateRequest) GetExpireSec() int32 {
	if x != nil {
		return x.ExpireSec
	}
	return 0
}

func (x *TimerCreateRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *TimerCreateRequest) GetRepeatCount() int32 {
	if x != nil {
		return x.RepeatCount
	}
	return 0
}

type TimerCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallbackUri string `protobuf:"bytes,1,opt,name=CallbackUri,proto3" json:"CallbackUri,omitempty"`
	SetTime     string `protobuf:"bytes,2,opt,name=SetTime,proto3" json:"SetTime,omitempty"`
	ExpireSec   int32  `protobuf:"varint,3,opt,name=ExpireSec,proto3" json:"ExpireSec,omitempty"`
	Data        string `protobuf:"bytes,4,opt,name=Data,proto3" json:"Data,omitempty"`
	RepeatCount int32  `protobuf:"varint,5,opt,name=RepeatCount,proto3" json:"RepeatCount,omitempty"`
	TimerId     string `protobuf:"bytes,6,opt,name=TimerId,proto3" json:"TimerId,omitempty"` // CMD_CREATED, CMD_DELETE, CMD_EVENT
}

func (x *TimerCreateResponse) Reset() {
	*x = TimerCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventTimer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimerCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimerCreateResponse) ProtoMessage() {}

func (x *TimerCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eventTimer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimerCreateResponse.ProtoReflect.Descriptor instead.
func (*TimerCreateResponse) Descriptor() ([]byte, []int) {
	return file_eventTimer_proto_rawDescGZIP(), []int{1}
}

func (x *TimerCreateResponse) GetCallbackUri() string {
	if x != nil {
		return x.CallbackUri
	}
	return ""
}

func (x *TimerCreateResponse) GetSetTime() string {
	if x != nil {
		return x.SetTime
	}
	return ""
}

func (x *TimerCreateResponse) GetExpireSec() int32 {
	if x != nil {
		return x.ExpireSec
	}
	return 0
}

func (x *TimerCreateResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *TimerCreateResponse) GetRepeatCount() int32 {
	if x != nil {
		return x.RepeatCount
	}
	return 0
}

func (x *TimerCreateResponse) GetTimerId() string {
	if x != nil {
		return x.TimerId
	}
	return ""
}

type TimerDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimerId string `protobuf:"bytes,1,opt,name=TimerId,proto3" json:"TimerId,omitempty"` // CMD_CREATED, CMD_DELETE
}

func (x *TimerDeleteRequest) Reset() {
	*x = TimerDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventTimer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimerDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimerDeleteRequest) ProtoMessage() {}

func (x *TimerDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eventTimer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimerDeleteRequest.ProtoReflect.Descriptor instead.
func (*TimerDeleteRequest) Descriptor() ([]byte, []int) {
	return file_eventTimer_proto_rawDescGZIP(), []int{2}
}

func (x *TimerDeleteRequest) GetTimerId() string {
	if x != nil {
		return x.TimerId
	}
	return ""
}

type TimerDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *TimerDeleteResponse) Reset() {
	*x = TimerDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventTimer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimerDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimerDeleteResponse) ProtoMessage() {}

func (x *TimerDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eventTimer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimerDeleteResponse.ProtoReflect.Descriptor instead.
func (*TimerDeleteResponse) Descriptor() ([]byte, []int) {
	return file_eventTimer_proto_rawDescGZIP(), []int{3}
}

func (x *TimerDeleteResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_eventTimer_proto protoreflect.FileDescriptor

var file_eventTimer_proto_rawDesc = []byte{
	0x0a, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x67, 0x72, 0x70, 0x63, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x74, 0x69, 0x6d,
	0x65, 0x72, 0x22, 0xa4, 0x01, 0x0a, 0x12, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x72, 0x69, 0x12, 0x18, 0x0a, 0x07, 0x53,
	0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x65,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x53,
	0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x53, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x52, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xbf, 0x01, 0x0a, 0x13, 0x54, 0x69,
	0x6d, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x72, 0x69,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x55, 0x72, 0x69, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x53, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x53, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x20, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x54,
	0x69, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x13, 0x54,
	0x69, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xb6, 0x01, 0x0a, 0x0a, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x53, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x74,
	0x69, 0x6d, 0x65, 0x72, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x4c, 0x0a, 0x13, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74,
	0x69, 0x6d, 0x65, 0x72, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x0f, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x22, 0x30,
	0x36, 0x2e, 0x67, 0x6f, 0x2d, 0x72, 0x70, 0x63, 0x2f, 0x30, 0x34, 0x2e, 0x74, 0x69, 0x6d, 0x65,
	0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eventTimer_proto_rawDescOnce sync.Once
	file_eventTimer_proto_rawDescData = file_eventTimer_proto_rawDesc
)

func file_eventTimer_proto_rawDescGZIP() []byte {
	file_eventTimer_proto_rawDescOnce.Do(func() {
		file_eventTimer_proto_rawDescData = protoimpl.X.CompressGZIP(file_eventTimer_proto_rawDescData)
	})
	return file_eventTimer_proto_rawDescData
}

var file_eventTimer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eventTimer_proto_goTypes = []interface{}{
	(*TimerCreateRequest)(nil),  // 0: grpceventtimer.TimerCreateRequest
	(*TimerCreateResponse)(nil), // 1: grpceventtimer.TimerCreateResponse
	(*TimerDeleteRequest)(nil),  // 2: grpceventtimer.TimerDeleteRequest
	(*TimerDeleteResponse)(nil), // 3: grpceventtimer.TimerDeleteResponse
}
var file_eventTimer_proto_depIdxs = []int32{
	0, // 0: grpceventtimer.EventTimer.Create:input_type -> grpceventtimer.TimerCreateRequest
	2, // 1: grpceventtimer.EventTimer.Delete:input_type -> grpceventtimer.TimerDeleteRequest
	1, // 2: grpceventtimer.EventTimer.Create:output_type -> grpceventtimer.TimerCreateResponse
	3, // 3: grpceventtimer.EventTimer.Delete:output_type -> grpceventtimer.TimerDeleteResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eventTimer_proto_init() }
func file_eventTimer_proto_init() {
	if File_eventTimer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eventTimer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimerCreateRequest); i {
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
		file_eventTimer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimerCreateResponse); i {
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
		file_eventTimer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimerDeleteRequest); i {
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
		file_eventTimer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimerDeleteResponse); i {
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
			RawDescriptor: file_eventTimer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eventTimer_proto_goTypes,
		DependencyIndexes: file_eventTimer_proto_depIdxs,
		MessageInfos:      file_eventTimer_proto_msgTypes,
	}.Build()
	File_eventTimer_proto = out.File
	file_eventTimer_proto_rawDesc = nil
	file_eventTimer_proto_goTypes = nil
	file_eventTimer_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventTimerClient is the client API for EventTimer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventTimerClient interface {
	Create(ctx context.Context, in *TimerCreateRequest, opts ...grpc.CallOption) (*TimerCreateResponse, error)
	Delete(ctx context.Context, in *TimerDeleteRequest, opts ...grpc.CallOption) (*TimerDeleteResponse, error)
}

type eventTimerClient struct {
	cc grpc.ClientConnInterface
}

func NewEventTimerClient(cc grpc.ClientConnInterface) EventTimerClient {
	return &eventTimerClient{cc}
}

func (c *eventTimerClient) Create(ctx context.Context, in *TimerCreateRequest, opts ...grpc.CallOption) (*TimerCreateResponse, error) {
	out := new(TimerCreateResponse)
	err := c.cc.Invoke(ctx, "/grpceventtimer.EventTimer/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventTimerClient) Delete(ctx context.Context, in *TimerDeleteRequest, opts ...grpc.CallOption) (*TimerDeleteResponse, error) {
	out := new(TimerDeleteResponse)
	err := c.cc.Invoke(ctx, "/grpceventtimer.EventTimer/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventTimerServer is the server API for EventTimer service.
type EventTimerServer interface {
	Create(context.Context, *TimerCreateRequest) (*TimerCreateResponse, error)
	Delete(context.Context, *TimerDeleteRequest) (*TimerDeleteResponse, error)
}

// UnimplementedEventTimerServer can be embedded to have forward compatible implementations.
type UnimplementedEventTimerServer struct {
}

func (*UnimplementedEventTimerServer) Create(context.Context, *TimerCreateRequest) (*TimerCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedEventTimerServer) Delete(context.Context, *TimerDeleteRequest) (*TimerDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterEventTimerServer(s *grpc.Server, srv EventTimerServer) {
	s.RegisterService(&_EventTimer_serviceDesc, srv)
}

func _EventTimer_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimerCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventTimerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpceventtimer.EventTimer/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventTimerServer).Create(ctx, req.(*TimerCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventTimer_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimerDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventTimerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpceventtimer.EventTimer/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventTimerServer).Delete(ctx, req.(*TimerDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventTimer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpceventtimer.EventTimer",
	HandlerType: (*EventTimerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _EventTimer_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _EventTimer_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eventTimer.proto",
}
