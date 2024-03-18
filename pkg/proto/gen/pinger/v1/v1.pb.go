// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: pkg/proto/pinger/v1.proto

package v1

import (
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_pinger_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_pinger_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_pkg_proto_pinger_v1_proto_rawDescGZIP(), []int{0}
}

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request string `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_pinger_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_pinger_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_pinger_v1_proto_rawDescGZIP(), []int{1}
}

func (x *PingRequest) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response  string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	TimeStamp string `protobuf:"bytes,2,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_pinger_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_pinger_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_pinger_v1_proto_rawDescGZIP(), []int{2}
}

func (x *PingResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *PingResponse) GetTimeStamp() string {
	if x != nil {
		return x.TimeStamp
	}
	return ""
}

var File_pkg_proto_pinger_v1_proto protoreflect.FileDescriptor

var file_pkg_proto_pinger_v1_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x69, 0x6e, 0x67,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x69, 0x6e,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x27, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x49, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74,
	0x61, 0x6d, 0x70, 0x32, 0x81, 0x01, 0x0a, 0x0d, 0x50, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e,
	0x70, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x2e, 0x70,
	0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17,
	0x2e, 0x70, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x67, 0x65, 0x6e, 0x2f, 0x70,
	0x69, 0x6e, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_pinger_v1_proto_rawDescOnce sync.Once
	file_pkg_proto_pinger_v1_proto_rawDescData = file_pkg_proto_pinger_v1_proto_rawDesc
)

func file_pkg_proto_pinger_v1_proto_rawDescGZIP() []byte {
	file_pkg_proto_pinger_v1_proto_rawDescOnce.Do(func() {
		file_pkg_proto_pinger_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_pinger_v1_proto_rawDescData)
	})
	return file_pkg_proto_pinger_v1_proto_rawDescData
}

var file_pkg_proto_pinger_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_proto_pinger_v1_proto_goTypes = []interface{}{
	(*Empty)(nil),        // 0: pinger.v1.Empty
	(*PingRequest)(nil),  // 1: pinger.v1.PingRequest
	(*PingResponse)(nil), // 2: pinger.v1.PingResponse
}
var file_pkg_proto_pinger_v1_proto_depIdxs = []int32{
	1, // 0: pinger.v1.PingerService.Ping:input_type -> pinger.v1.PingRequest
	0, // 1: pinger.v1.PingerService.GetAllPing:input_type -> pinger.v1.Empty
	2, // 2: pinger.v1.PingerService.Ping:output_type -> pinger.v1.PingResponse
	2, // 3: pinger.v1.PingerService.GetAllPing:output_type -> pinger.v1.PingResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_proto_pinger_v1_proto_init() }
func file_pkg_proto_pinger_v1_proto_init() {
	if File_pkg_proto_pinger_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_pinger_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_pkg_proto_pinger_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_pkg_proto_pinger_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
			RawDescriptor: file_pkg_proto_pinger_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_pinger_v1_proto_goTypes,
		DependencyIndexes: file_pkg_proto_pinger_v1_proto_depIdxs,
		MessageInfos:      file_pkg_proto_pinger_v1_proto_msgTypes,
	}.Build()
	File_pkg_proto_pinger_v1_proto = out.File
	file_pkg_proto_pinger_v1_proto_rawDesc = nil
	file_pkg_proto_pinger_v1_proto_goTypes = nil
	file_pkg_proto_pinger_v1_proto_depIdxs = nil
}
