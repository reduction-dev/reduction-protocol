// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: jobconfigpb/jobconfig.proto

package jobconfigpb

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

// JobConfig is the configuration needed to run a job
type JobConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Job           *Job                   `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	Sources       []*Source              `protobuf:"bytes,2,rep,name=sources,proto3" json:"sources,omitempty"`
	Sinks         []*Sink                `protobuf:"bytes,3,rep,name=sinks,proto3" json:"sinks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobConfig) Reset() {
	*x = JobConfig{}
	mi := &file_jobconfigpb_jobconfig_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobConfig) ProtoMessage() {}

func (x *JobConfig) ProtoReflect() protoreflect.Message {
	mi := &file_jobconfigpb_jobconfig_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobConfig.ProtoReflect.Descriptor instead.
func (*JobConfig) Descriptor() ([]byte, []int) {
	return file_jobconfigpb_jobconfig_proto_rawDescGZIP(), []int{0}
}

func (x *JobConfig) GetJob() *Job {
	if x != nil {
		return x.Job
	}
	return nil
}

func (x *JobConfig) GetSources() []*Source {
	if x != nil {
		return x.Sources
	}
	return nil
}

func (x *JobConfig) GetSinks() []*Sink {
	if x != nil {
		return x.Sinks
	}
	return nil
}

// Job contains job's parameters
type Job struct {
	state                    protoimpl.MessageState `protogen:"open.v1"`
	WorkerCount              *Int32Var              `protobuf:"bytes,1,opt,name=worker_count,json=workerCount,proto3" json:"worker_count,omitempty"`
	KeyGroupCount            int32                  `protobuf:"varint,2,opt,name=key_group_count,json=keyGroupCount,proto3" json:"key_group_count,omitempty"`
	WorkingStorageLocation   *StringVar             `protobuf:"bytes,3,opt,name=working_storage_location,json=workingStorageLocation,proto3" json:"working_storage_location,omitempty"`
	SavepointStorageLocation *StringVar             `protobuf:"bytes,4,opt,name=savepoint_storage_location,json=savepointStorageLocation,proto3" json:"savepoint_storage_location,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *Job) Reset() {
	*x = Job{}
	mi := &file_jobconfigpb_jobconfig_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_jobconfigpb_jobconfig_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_jobconfigpb_jobconfig_proto_rawDescGZIP(), []int{1}
}

func (x *Job) GetWorkerCount() *Int32Var {
	if x != nil {
		return x.WorkerCount
	}
	return nil
}

func (x *Job) GetKeyGroupCount() int32 {
	if x != nil {
		return x.KeyGroupCount
	}
	return 0
}

func (x *Job) GetWorkingStorageLocation() *StringVar {
	if x != nil {
		return x.WorkingStorageLocation
	}
	return nil
}

func (x *Job) GetSavepointStorageLocation() *StringVar {
	if x != nil {
		return x.SavepointStorageLocation
	}
	return nil
}

// Source represents any type of input source
type Source struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Id    string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Config:
	//
	//	*Source_Stdio
	//	*Source_Kinesis
	//	*Source_HttpApi
	//	*Source_Embedded
	//	*Source_Kafka
	Config        isSource_Config `protobuf_oneof:"config"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Source) Reset() {
	*x = Source{}
	mi := &file_jobconfigpb_jobconfig_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Source) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Source) ProtoMessage() {}

func (x *Source) ProtoReflect() protoreflect.Message {
	mi := &file_jobconfigpb_jobconfig_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Source.ProtoReflect.Descriptor instead.
func (*Source) Descriptor() ([]byte, []int) {
	return file_jobconfigpb_jobconfig_proto_rawDescGZIP(), []int{2}
}

func (x *Source) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Source) GetConfig() isSource_Config {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Source) GetStdio() *StdioSource {
	if x != nil {
		if x, ok := x.Config.(*Source_Stdio); ok {
			return x.Stdio
		}
	}
	return nil
}

func (x *Source) GetKinesis() *KinesisSource {
	if x != nil {
		if x, ok := x.Config.(*Source_Kinesis); ok {
			return x.Kinesis
		}
	}
	return nil
}

func (x *Source) GetHttpApi() *HTTPAPISource {
	if x != nil {
		if x, ok := x.Config.(*Source_HttpApi); ok {
			return x.HttpApi
		}
	}
	return nil
}

func (x *Source) GetEmbedded() *EmbeddedSource {
	if x != nil {
		if x, ok := x.Config.(*Source_Embedded); ok {
			return x.Embedded
		}
	}
	return nil
}

func (x *Source) GetKafka() *KafkaSource {
	if x != nil {
		if x, ok := x.Config.(*Source_Kafka); ok {
			return x.Kafka
		}
	}
	return nil
}

type isSource_Config interface {
	isSource_Config()
}

type Source_Stdio struct {
	Stdio *StdioSource `protobuf:"bytes,2,opt,name=stdio,proto3,oneof"`
}

type Source_Kinesis struct {
	Kinesis *KinesisSource `protobuf:"bytes,3,opt,name=kinesis,proto3,oneof"`
}

type Source_HttpApi struct {
	HttpApi *HTTPAPISource `protobuf:"bytes,4,opt,name=http_api,json=httpApi,proto3,oneof"`
}

type Source_Embedded struct {
	Embedded *EmbeddedSource `protobuf:"bytes,5,opt,name=embedded,proto3,oneof"`
}

type Source_Kafka struct {
	Kafka *KafkaSource `protobuf:"bytes,6,opt,name=kafka,proto3,oneof"`
}

func (*Source_Stdio) isSource_Config() {}

func (*Source_Kinesis) isSource_Config() {}

func (*Source_HttpApi) isSource_Config() {}

func (*Source_Embedded) isSource_Config() {}

func (*Source_Kafka) isSource_Config() {}

// Sink represents any type of output sink
type Sink struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Id    string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Config:
	//
	//	*Sink_Stdio
	//	*Sink_HttpApi
	//	*Sink_Memory
	//	*Sink_Kafka
	Config        isSink_Config `protobuf_oneof:"config"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Sink) Reset() {
	*x = Sink{}
	mi := &file_jobconfigpb_jobconfig_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Sink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sink) ProtoMessage() {}

func (x *Sink) ProtoReflect() protoreflect.Message {
	mi := &file_jobconfigpb_jobconfig_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sink.ProtoReflect.Descriptor instead.
func (*Sink) Descriptor() ([]byte, []int) {
	return file_jobconfigpb_jobconfig_proto_rawDescGZIP(), []int{3}
}

func (x *Sink) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Sink) GetConfig() isSink_Config {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Sink) GetStdio() *StdioSink {
	if x != nil {
		if x, ok := x.Config.(*Sink_Stdio); ok {
			return x.Stdio
		}
	}
	return nil
}

func (x *Sink) GetHttpApi() *HTTPAPISink {
	if x != nil {
		if x, ok := x.Config.(*Sink_HttpApi); ok {
			return x.HttpApi
		}
	}
	return nil
}

func (x *Sink) GetMemory() *MemorySink {
	if x != nil {
		if x, ok := x.Config.(*Sink_Memory); ok {
			return x.Memory
		}
	}
	return nil
}

func (x *Sink) GetKafka() *KafkaSink {
	if x != nil {
		if x, ok := x.Config.(*Sink_Kafka); ok {
			return x.Kafka
		}
	}
	return nil
}

type isSink_Config interface {
	isSink_Config()
}

type Sink_Stdio struct {
	Stdio *StdioSink `protobuf:"bytes,2,opt,name=stdio,proto3,oneof"`
}

type Sink_HttpApi struct {
	HttpApi *HTTPAPISink `protobuf:"bytes,3,opt,name=http_api,json=httpApi,proto3,oneof"`
}

type Sink_Memory struct {
	Memory *MemorySink `protobuf:"bytes,4,opt,name=memory,proto3,oneof"`
}

type Sink_Kafka struct {
	Kafka *KafkaSink `protobuf:"bytes,5,opt,name=kafka,proto3,oneof"`
}

func (*Sink_Stdio) isSink_Config() {}

func (*Sink_HttpApi) isSink_Config() {}

func (*Sink_Memory) isSink_Config() {}

func (*Sink_Kafka) isSink_Config() {}

var File_jobconfigpb_jobconfig_proto protoreflect.FileDescriptor

var file_jobconfigpb_jobconfig_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6a, 0x6f, 0x62, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x70, 0x62, 0x2f, 0x6a, 0x6f,
	0x62, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x64,
	0x65, 0x76, 0x2e, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6a, 0x6f, 0x62,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x1c, 0x6a, 0x6f, 0x62, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6a, 0x6f, 0x62, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x70,
	0x62, 0x2f, 0x73, 0x74, 0x64, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6a,
	0x6f, 0x62, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x70, 0x62, 0x2f, 0x6b, 0x69, 0x6e, 0x65, 0x73,
	0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6a, 0x6f, 0x62, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x70, 0x62, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6a, 0x6f, 0x62, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x70, 0x62,
	0x2f, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x18, 0x6a, 0x6f, 0x62, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6a, 0x6f, 0x62, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x70, 0x62, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x2e, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x64, 0x65, 0x76, 0x2e, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6a, 0x6f,
	0x62, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x03, 0x6a, 0x6f, 0x62,
	0x12, 0x39, 0x0a, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x6a, 0x6f, 0x62, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x05, 0x73,
	0x69, 0x6e, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x65, 0x76,
	0x2e, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6a, 0x6f, 0x62, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x73, 0x69, 0x6e, 0x6b, 0x73,
	0x22, 0xb3, 0x02, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x44, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6a,
	0x6f, 0x62, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x72, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26,
	0x0a, 0x0f, 0x6b, 0x65, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6b, 0x65, 0x79, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x5c, 0x0a, 0x18, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72,
	0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6a, 0x6f, 0x62, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x72, 0x52, 0x16, 0x77, 0x6f,
	0x72, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x60, 0x0a, 0x1a, 0x73, 0x61, 0x76, 0x65, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72,
	0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6a, 0x6f, 0x62, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x72, 0x52, 0x18, 0x73, 0x61,
	0x76, 0x65, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xee, 0x02, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x3c, 0x0a, 0x05, 0x73, 0x74, 0x64, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x6a, 0x6f, 0x62, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x74, 0x64, 0x69, 0x6f,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x05, 0x73, 0x74, 0x64, 0x69, 0x6f, 0x12,
	0x42, 0x0a, 0x07, 0x6b, 0x69, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x6a, 0x6f, 0x62, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4b, 0x69, 0x6e, 0x65, 0x73,
	0x69, 0x73, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x07, 0x6b, 0x69, 0x6e, 0x65,
	0x73, 0x69, 0x73, 0x12, 0x43, 0x0a, 0x08, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x61, 0x70, 0x69, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x65, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6a, 0x6f, 0x62, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x48, 0x54, 0x54, 0x50, 0x41, 0x50, 0x49, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52,
	0x07, 0x68, 0x74, 0x74, 0x70, 0x41, 0x70, 0x69, 0x12, 0x45, 0x0a, 0x08, 0x65, 0x6d, 0x62, 0x65,
	0x64, 0x64, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x64, 0x65, 0x76,
	0x2e, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6a, 0x6f, 0x62, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x08, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x12,
	0x3c, 0x0a, 0x05, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6a,
	0x6f, 0x62, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x05, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x42, 0x08, 0x0a,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x9a, 0x02, 0x0a, 0x04, 0x53, 0x69, 0x6e, 0x6b,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x3a, 0x0a, 0x05, 0x73, 0x74, 0x64, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x6a, 0x6f, 0x62, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x74, 0x64, 0x69, 0x6f, 0x53,
	0x69, 0x6e, 0x6b, 0x48, 0x00, 0x52, 0x05, 0x73, 0x74, 0x64, 0x69, 0x6f, 0x12, 0x41, 0x0a, 0x08,
	0x68, 0x74, 0x74, 0x70, 0x5f, 0x61, 0x70, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6a,
	0x6f, 0x62, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x41, 0x50, 0x49,
	0x53, 0x69, 0x6e, 0x6b, 0x48, 0x00, 0x52, 0x07, 0x68, 0x74, 0x74, 0x70, 0x41, 0x70, 0x69, 0x12,
	0x3d, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x6a, 0x6f, 0x62, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x53, 0x69, 0x6e, 0x6b, 0x48, 0x00, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x3a,
	0x0a, 0x05, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x64, 0x65, 0x76, 0x2e, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6a, 0x6f,
	0x62, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x53, 0x69, 0x6e,
	0x6b, 0x48, 0x00, 0x52, 0x05, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x42, 0x08, 0x0a, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x42, 0x2e, 0x5a, 0x2c, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6a, 0x6f, 0x62, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_jobconfigpb_jobconfig_proto_rawDescOnce sync.Once
	file_jobconfigpb_jobconfig_proto_rawDescData = file_jobconfigpb_jobconfig_proto_rawDesc
)

func file_jobconfigpb_jobconfig_proto_rawDescGZIP() []byte {
	file_jobconfigpb_jobconfig_proto_rawDescOnce.Do(func() {
		file_jobconfigpb_jobconfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_jobconfigpb_jobconfig_proto_rawDescData)
	})
	return file_jobconfigpb_jobconfig_proto_rawDescData
}

var file_jobconfigpb_jobconfig_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_jobconfigpb_jobconfig_proto_goTypes = []any{
	(*JobConfig)(nil),      // 0: dev.reduction.jobconfig.JobConfig
	(*Job)(nil),            // 1: dev.reduction.jobconfig.Job
	(*Source)(nil),         // 2: dev.reduction.jobconfig.Source
	(*Sink)(nil),           // 3: dev.reduction.jobconfig.Sink
	(*Int32Var)(nil),       // 4: dev.reduction.jobconfig.Int32Var
	(*StringVar)(nil),      // 5: dev.reduction.jobconfig.StringVar
	(*StdioSource)(nil),    // 6: dev.reduction.jobconfig.StdioSource
	(*KinesisSource)(nil),  // 7: dev.reduction.jobconfig.KinesisSource
	(*HTTPAPISource)(nil),  // 8: dev.reduction.jobconfig.HTTPAPISource
	(*EmbeddedSource)(nil), // 9: dev.reduction.jobconfig.EmbeddedSource
	(*KafkaSource)(nil),    // 10: dev.reduction.jobconfig.KafkaSource
	(*StdioSink)(nil),      // 11: dev.reduction.jobconfig.StdioSink
	(*HTTPAPISink)(nil),    // 12: dev.reduction.jobconfig.HTTPAPISink
	(*MemorySink)(nil),     // 13: dev.reduction.jobconfig.MemorySink
	(*KafkaSink)(nil),      // 14: dev.reduction.jobconfig.KafkaSink
}
var file_jobconfigpb_jobconfig_proto_depIdxs = []int32{
	1,  // 0: dev.reduction.jobconfig.JobConfig.job:type_name -> dev.reduction.jobconfig.Job
	2,  // 1: dev.reduction.jobconfig.JobConfig.sources:type_name -> dev.reduction.jobconfig.Source
	3,  // 2: dev.reduction.jobconfig.JobConfig.sinks:type_name -> dev.reduction.jobconfig.Sink
	4,  // 3: dev.reduction.jobconfig.Job.worker_count:type_name -> dev.reduction.jobconfig.Int32Var
	5,  // 4: dev.reduction.jobconfig.Job.working_storage_location:type_name -> dev.reduction.jobconfig.StringVar
	5,  // 5: dev.reduction.jobconfig.Job.savepoint_storage_location:type_name -> dev.reduction.jobconfig.StringVar
	6,  // 6: dev.reduction.jobconfig.Source.stdio:type_name -> dev.reduction.jobconfig.StdioSource
	7,  // 7: dev.reduction.jobconfig.Source.kinesis:type_name -> dev.reduction.jobconfig.KinesisSource
	8,  // 8: dev.reduction.jobconfig.Source.http_api:type_name -> dev.reduction.jobconfig.HTTPAPISource
	9,  // 9: dev.reduction.jobconfig.Source.embedded:type_name -> dev.reduction.jobconfig.EmbeddedSource
	10, // 10: dev.reduction.jobconfig.Source.kafka:type_name -> dev.reduction.jobconfig.KafkaSource
	11, // 11: dev.reduction.jobconfig.Sink.stdio:type_name -> dev.reduction.jobconfig.StdioSink
	12, // 12: dev.reduction.jobconfig.Sink.http_api:type_name -> dev.reduction.jobconfig.HTTPAPISink
	13, // 13: dev.reduction.jobconfig.Sink.memory:type_name -> dev.reduction.jobconfig.MemorySink
	14, // 14: dev.reduction.jobconfig.Sink.kafka:type_name -> dev.reduction.jobconfig.KafkaSink
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_jobconfigpb_jobconfig_proto_init() }
func file_jobconfigpb_jobconfig_proto_init() {
	if File_jobconfigpb_jobconfig_proto != nil {
		return
	}
	file_jobconfigpb_resolvable_proto_init()
	file_jobconfigpb_stdio_proto_init()
	file_jobconfigpb_kinesis_proto_init()
	file_jobconfigpb_httpapi_proto_init()
	file_jobconfigpb_embedded_proto_init()
	file_jobconfigpb_memory_proto_init()
	file_jobconfigpb_kafka_proto_init()
	file_jobconfigpb_jobconfig_proto_msgTypes[2].OneofWrappers = []any{
		(*Source_Stdio)(nil),
		(*Source_Kinesis)(nil),
		(*Source_HttpApi)(nil),
		(*Source_Embedded)(nil),
		(*Source_Kafka)(nil),
	}
	file_jobconfigpb_jobconfig_proto_msgTypes[3].OneofWrappers = []any{
		(*Sink_Stdio)(nil),
		(*Sink_HttpApi)(nil),
		(*Sink_Memory)(nil),
		(*Sink_Kafka)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_jobconfigpb_jobconfig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_jobconfigpb_jobconfig_proto_goTypes,
		DependencyIndexes: file_jobconfigpb_jobconfig_proto_depIdxs,
		MessageInfos:      file_jobconfigpb_jobconfig_proto_msgTypes,
	}.Build()
	File_jobconfigpb_jobconfig_proto = out.File
	file_jobconfigpb_jobconfig_proto_rawDesc = nil
	file_jobconfigpb_jobconfig_proto_goTypes = nil
	file_jobconfigpb_jobconfig_proto_depIdxs = nil
}
