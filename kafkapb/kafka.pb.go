// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: kafkapb/kafka.proto

package kafkapb

import (
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

type Record struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Topic         string                 `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Partition     int32                  `protobuf:"varint,2,opt,name=partition,proto3" json:"partition,omitempty"`
	Key           []byte                 `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value         []byte                 `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Headers       []*Header              `protobuf:"bytes,6,rep,name=headers,proto3" json:"headers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Record) Reset() {
	*x = Record{}
	mi := &file_kafkapb_kafka_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_kafkapb_kafka_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_kafkapb_kafka_proto_rawDescGZIP(), []int{0}
}

func (x *Record) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Record) GetPartition() int32 {
	if x != nil {
		return x.Partition
	}
	return 0
}

func (x *Record) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Record) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Record) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Record) GetHeaders() []*Header {
	if x != nil {
		return x.Headers
	}
	return nil
}

type Header struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value         []byte                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Header) Reset() {
	*x = Header{}
	mi := &file_kafkapb_kafka_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_kafkapb_kafka_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_kafkapb_kafka_proto_rawDescGZIP(), []int{1}
}

func (x *Header) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Header) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_kafkapb_kafka_proto protoreflect.FileDescriptor

var file_kafkapb_kafka_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x70, 0x62, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x65, 0x64, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x01, 0x0a, 0x06,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x35, 0x0a, 0x07,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x64, 0x65, 0x76, 0x2e, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6b, 0x61,
	0x66, 0x6b, 0x61, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x22, 0x30, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x2a, 0x5a, 0x28, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kafkapb_kafka_proto_rawDescOnce sync.Once
	file_kafkapb_kafka_proto_rawDescData = file_kafkapb_kafka_proto_rawDesc
)

func file_kafkapb_kafka_proto_rawDescGZIP() []byte {
	file_kafkapb_kafka_proto_rawDescOnce.Do(func() {
		file_kafkapb_kafka_proto_rawDescData = protoimpl.X.CompressGZIP(file_kafkapb_kafka_proto_rawDescData)
	})
	return file_kafkapb_kafka_proto_rawDescData
}

var file_kafkapb_kafka_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_kafkapb_kafka_proto_goTypes = []any{
	(*Record)(nil),                // 0: dev.reduction.kafka.Record
	(*Header)(nil),                // 1: dev.reduction.kafka.Header
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_kafkapb_kafka_proto_depIdxs = []int32{
	2, // 0: dev.reduction.kafka.Record.timestamp:type_name -> google.protobuf.Timestamp
	1, // 1: dev.reduction.kafka.Record.headers:type_name -> dev.reduction.kafka.Header
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_kafkapb_kafka_proto_init() }
func file_kafkapb_kafka_proto_init() {
	if File_kafkapb_kafka_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kafkapb_kafka_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kafkapb_kafka_proto_goTypes,
		DependencyIndexes: file_kafkapb_kafka_proto_depIdxs,
		MessageInfos:      file_kafkapb_kafka_proto_msgTypes,
	}.Build()
	File_kafkapb_kafka_proto = out.File
	file_kafkapb_kafka_proto_rawDesc = nil
	file_kafkapb_kafka_proto_goTypes = nil
	file_kafkapb_kafka_proto_depIdxs = nil
}
