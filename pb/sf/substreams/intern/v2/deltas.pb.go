// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: sf/substreams/intern/v2/deltas.proto

package pbssinternal

import (
	_ "github.com/streamingfast/substreams/pb/sf/substreams/index/v1"
	v1 "github.com/streamingfast/substreams/pb/sf/substreams/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Operation_Type int32

const (
	Operation_SET                     Operation_Type = 0
	Operation_SET_BYTES               Operation_Type = 1
	Operation_SET_IF_NOT_EXISTS       Operation_Type = 2
	Operation_SET_BYTES_IF_NOT_EXISTS Operation_Type = 3
	Operation_APPEND                  Operation_Type = 4
	Operation_DELETE_PREFIX           Operation_Type = 5
	Operation_SET_MAX_BIG_INT         Operation_Type = 6
	Operation_SET_MAX_INT64           Operation_Type = 7
	Operation_SET_MAX_FLOAT64         Operation_Type = 8
	Operation_SET_MAX_BIG_DECIMAL     Operation_Type = 9
	Operation_SET_MIN_BIG_INT         Operation_Type = 10
	Operation_SET_MIN_INT64           Operation_Type = 11
	Operation_SET_MIN_FLOAT64         Operation_Type = 12
	Operation_SET_MIN_BIG_DECIMAL     Operation_Type = 13
	Operation_SUM_BIG_INT             Operation_Type = 14
	Operation_SUM_INT64               Operation_Type = 15
	Operation_SUM_FLOAT64             Operation_Type = 16
	Operation_SUM_BIG_DECIMAL         Operation_Type = 17
	Operation_SET_SUM_INT64           Operation_Type = 18
	Operation_SET_SUM_FLOAT64         Operation_Type = 19
	Operation_SET_SUM_BIG_INT         Operation_Type = 20
	Operation_SET_SUM_BIG_DECIMAL     Operation_Type = 21
)

// Enum value maps for Operation_Type.
var (
	Operation_Type_name = map[int32]string{
		0:  "SET",
		1:  "SET_BYTES",
		2:  "SET_IF_NOT_EXISTS",
		3:  "SET_BYTES_IF_NOT_EXISTS",
		4:  "APPEND",
		5:  "DELETE_PREFIX",
		6:  "SET_MAX_BIG_INT",
		7:  "SET_MAX_INT64",
		8:  "SET_MAX_FLOAT64",
		9:  "SET_MAX_BIG_DECIMAL",
		10: "SET_MIN_BIG_INT",
		11: "SET_MIN_INT64",
		12: "SET_MIN_FLOAT64",
		13: "SET_MIN_BIG_DECIMAL",
		14: "SUM_BIG_INT",
		15: "SUM_INT64",
		16: "SUM_FLOAT64",
		17: "SUM_BIG_DECIMAL",
		18: "SET_SUM_INT64",
		19: "SET_SUM_FLOAT64",
		20: "SET_SUM_BIG_INT",
		21: "SET_SUM_BIG_DECIMAL",
	}
	Operation_Type_value = map[string]int32{
		"SET":                     0,
		"SET_BYTES":               1,
		"SET_IF_NOT_EXISTS":       2,
		"SET_BYTES_IF_NOT_EXISTS": 3,
		"APPEND":                  4,
		"DELETE_PREFIX":           5,
		"SET_MAX_BIG_INT":         6,
		"SET_MAX_INT64":           7,
		"SET_MAX_FLOAT64":         8,
		"SET_MAX_BIG_DECIMAL":     9,
		"SET_MIN_BIG_INT":         10,
		"SET_MIN_INT64":           11,
		"SET_MIN_FLOAT64":         12,
		"SET_MIN_BIG_DECIMAL":     13,
		"SUM_BIG_INT":             14,
		"SUM_INT64":               15,
		"SUM_FLOAT64":             16,
		"SUM_BIG_DECIMAL":         17,
		"SET_SUM_INT64":           18,
		"SET_SUM_FLOAT64":         19,
		"SET_SUM_BIG_INT":         20,
		"SET_SUM_BIG_DECIMAL":     21,
	}
)

func (x Operation_Type) Enum() *Operation_Type {
	p := new(Operation_Type)
	*p = x
	return p
}

func (x Operation_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operation_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_sf_substreams_intern_v2_deltas_proto_enumTypes[0].Descriptor()
}

func (Operation_Type) Type() protoreflect.EnumType {
	return &file_sf_substreams_intern_v2_deltas_proto_enumTypes[0]
}

func (x Operation_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operation_Type.Descriptor instead.
func (Operation_Type) EnumDescriptor() ([]byte, []int) {
	return file_sf_substreams_intern_v2_deltas_proto_rawDescGZIP(), []int{2, 0}
}

type ModuleOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModuleName string `protobuf:"bytes,1,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
	// Types that are assignable to Data:
	//
	//	*ModuleOutput_MapOutput
	//	*ModuleOutput_StoreDeltas
	Data               isModuleOutput_Data `protobuf_oneof:"data"`
	Logs               []string            `protobuf:"bytes,4,rep,name=logs,proto3" json:"logs,omitempty"`
	DebugLogsTruncated bool                `protobuf:"varint,5,opt,name=debug_logs_truncated,json=debugLogsTruncated,proto3" json:"debug_logs_truncated,omitempty"`
	Cached             bool                `protobuf:"varint,6,opt,name=cached,proto3" json:"cached,omitempty"`
}

func (x *ModuleOutput) Reset() {
	*x = ModuleOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_intern_v2_deltas_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModuleOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModuleOutput) ProtoMessage() {}

func (x *ModuleOutput) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_intern_v2_deltas_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModuleOutput.ProtoReflect.Descriptor instead.
func (*ModuleOutput) Descriptor() ([]byte, []int) {
	return file_sf_substreams_intern_v2_deltas_proto_rawDescGZIP(), []int{0}
}

func (x *ModuleOutput) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

func (m *ModuleOutput) GetData() isModuleOutput_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *ModuleOutput) GetMapOutput() *anypb.Any {
	if x, ok := x.GetData().(*ModuleOutput_MapOutput); ok {
		return x.MapOutput
	}
	return nil
}

func (x *ModuleOutput) GetStoreDeltas() *v1.StoreDeltas {
	if x, ok := x.GetData().(*ModuleOutput_StoreDeltas); ok {
		return x.StoreDeltas
	}
	return nil
}

func (x *ModuleOutput) GetLogs() []string {
	if x != nil {
		return x.Logs
	}
	return nil
}

func (x *ModuleOutput) GetDebugLogsTruncated() bool {
	if x != nil {
		return x.DebugLogsTruncated
	}
	return false
}

func (x *ModuleOutput) GetCached() bool {
	if x != nil {
		return x.Cached
	}
	return false
}

type isModuleOutput_Data interface {
	isModuleOutput_Data()
}

type ModuleOutput_MapOutput struct {
	MapOutput *anypb.Any `protobuf:"bytes,2,opt,name=map_output,json=mapOutput,proto3,oneof"`
}

type ModuleOutput_StoreDeltas struct {
	StoreDeltas *v1.StoreDeltas `protobuf:"bytes,3,opt,name=store_deltas,json=storeDeltas,proto3,oneof"`
}

func (*ModuleOutput_MapOutput) isModuleOutput_Data() {}

func (*ModuleOutput_StoreDeltas) isModuleOutput_Data() {}

type Operations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operations []*Operation `protobuf:"bytes,1,rep,name=operations,proto3" json:"operations,omitempty"`
}

func (x *Operations) Reset() {
	*x = Operations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_intern_v2_deltas_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operations) ProtoMessage() {}

func (x *Operations) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_intern_v2_deltas_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operations.ProtoReflect.Descriptor instead.
func (*Operations) Descriptor() ([]byte, []int) {
	return file_sf_substreams_intern_v2_deltas_proto_rawDescGZIP(), []int{1}
}

func (x *Operations) GetOperations() []*Operation {
	if x != nil {
		return x.Operations
	}
	return nil
}

type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  Operation_Type `protobuf:"varint,1,opt,name=type,proto3,enum=sf.substreams.internal.v2.Operation_Type" json:"type,omitempty"`
	Ord   uint64         `protobuf:"varint,2,opt,name=ord,proto3" json:"ord,omitempty"`
	Key   string         `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte         `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_intern_v2_deltas_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_intern_v2_deltas_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_sf_substreams_intern_v2_deltas_proto_rawDescGZIP(), []int{2}
}

func (x *Operation) GetType() Operation_Type {
	if x != nil {
		return x.Type
	}
	return Operation_SET
}

func (x *Operation) GetOrd() uint64 {
	if x != nil {
		return x.Ord
	}
	return 0
}

func (x *Operation) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Operation) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_sf_substreams_intern_v2_deltas_proto protoreflect.FileDescriptor

var file_sf_substreams_intern_v2_deltas_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x66, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x76,
	0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x73, 0x66,
	0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x65, 0x6c, 0x74, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x73, 0x66, 0x2f,
	0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90,
	0x02, 0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x35, 0x0a, 0x0a, 0x6d, 0x61, 0x70, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x09, 0x6d, 0x61,
	0x70, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x42, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x73, 0x48, 0x00, 0x52, 0x0b,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x6f, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x12,
	0x30, 0x0a, 0x14, 0x64, 0x65, 0x62, 0x75, 0x67, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x74, 0x72,
	0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x64,
	0x65, 0x62, 0x75, 0x67, 0x4c, 0x6f, 0x67, 0x73, 0x54, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x52, 0x0a, 0x0a, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x44, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x76, 0x32, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xc0, 0x04, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x29, 0x2e, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x6f, 0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xb9, 0x03, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x09, 0x53, 0x45, 0x54, 0x5f, 0x42, 0x59, 0x54, 0x45, 0x53, 0x10, 0x01, 0x12, 0x15, 0x0a,
	0x11, 0x53, 0x45, 0x54, 0x5f, 0x49, 0x46, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53,
	0x54, 0x53, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x45, 0x54, 0x5f, 0x42, 0x59, 0x54, 0x45,
	0x53, 0x5f, 0x49, 0x46, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10,
	0x03, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x50, 0x50, 0x45, 0x4e, 0x44, 0x10, 0x04, 0x12, 0x11, 0x0a,
	0x0d, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x50, 0x52, 0x45, 0x46, 0x49, 0x58, 0x10, 0x05,
	0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x54, 0x5f, 0x4d, 0x41, 0x58, 0x5f, 0x42, 0x49, 0x47, 0x5f,
	0x49, 0x4e, 0x54, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x45, 0x54, 0x5f, 0x4d, 0x41, 0x58,
	0x5f, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x54, 0x5f,
	0x4d, 0x41, 0x58, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x36, 0x34, 0x10, 0x08, 0x12, 0x17, 0x0a,
	0x13, 0x53, 0x45, 0x54, 0x5f, 0x4d, 0x41, 0x58, 0x5f, 0x42, 0x49, 0x47, 0x5f, 0x44, 0x45, 0x43,
	0x49, 0x4d, 0x41, 0x4c, 0x10, 0x09, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x54, 0x5f, 0x4d, 0x49,
	0x4e, 0x5f, 0x42, 0x49, 0x47, 0x5f, 0x49, 0x4e, 0x54, 0x10, 0x0a, 0x12, 0x11, 0x0a, 0x0d, 0x53,
	0x45, 0x54, 0x5f, 0x4d, 0x49, 0x4e, 0x5f, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x0b, 0x12, 0x13,
	0x0a, 0x0f, 0x53, 0x45, 0x54, 0x5f, 0x4d, 0x49, 0x4e, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x36,
	0x34, 0x10, 0x0c, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x45, 0x54, 0x5f, 0x4d, 0x49, 0x4e, 0x5f, 0x42,
	0x49, 0x47, 0x5f, 0x44, 0x45, 0x43, 0x49, 0x4d, 0x41, 0x4c, 0x10, 0x0d, 0x12, 0x0f, 0x0a, 0x0b,
	0x53, 0x55, 0x4d, 0x5f, 0x42, 0x49, 0x47, 0x5f, 0x49, 0x4e, 0x54, 0x10, 0x0e, 0x12, 0x0d, 0x0a,
	0x09, 0x53, 0x55, 0x4d, 0x5f, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x0f, 0x12, 0x0f, 0x0a, 0x0b,
	0x53, 0x55, 0x4d, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x36, 0x34, 0x10, 0x10, 0x12, 0x13, 0x0a,
	0x0f, 0x53, 0x55, 0x4d, 0x5f, 0x42, 0x49, 0x47, 0x5f, 0x44, 0x45, 0x43, 0x49, 0x4d, 0x41, 0x4c,
	0x10, 0x11, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x45, 0x54, 0x5f, 0x53, 0x55, 0x4d, 0x5f, 0x49, 0x4e,
	0x54, 0x36, 0x34, 0x10, 0x12, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x54, 0x5f, 0x53, 0x55, 0x4d,
	0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x36, 0x34, 0x10, 0x13, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45,
	0x54, 0x5f, 0x53, 0x55, 0x4d, 0x5f, 0x42, 0x49, 0x47, 0x5f, 0x49, 0x4e, 0x54, 0x10, 0x14, 0x12,
	0x17, 0x0a, 0x13, 0x53, 0x45, 0x54, 0x5f, 0x53, 0x55, 0x4d, 0x5f, 0x42, 0x49, 0x47, 0x5f, 0x44,
	0x45, 0x43, 0x49, 0x4d, 0x41, 0x4c, 0x10, 0x15, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x66, 0x61, 0x73, 0x74, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f,
	0x70, 0x62, 0x2f, 0x73, 0x66, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x2f, 0x76, 0x32, 0x3b, 0x70, 0x62, 0x73, 0x73, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sf_substreams_intern_v2_deltas_proto_rawDescOnce sync.Once
	file_sf_substreams_intern_v2_deltas_proto_rawDescData = file_sf_substreams_intern_v2_deltas_proto_rawDesc
)

func file_sf_substreams_intern_v2_deltas_proto_rawDescGZIP() []byte {
	file_sf_substreams_intern_v2_deltas_proto_rawDescOnce.Do(func() {
		file_sf_substreams_intern_v2_deltas_proto_rawDescData = protoimpl.X.CompressGZIP(file_sf_substreams_intern_v2_deltas_proto_rawDescData)
	})
	return file_sf_substreams_intern_v2_deltas_proto_rawDescData
}

var file_sf_substreams_intern_v2_deltas_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sf_substreams_intern_v2_deltas_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sf_substreams_intern_v2_deltas_proto_goTypes = []any{
	(Operation_Type)(0),    // 0: sf.substreams.internal.v2.Operation.Type
	(*ModuleOutput)(nil),   // 1: sf.substreams.internal.v2.ModuleOutput
	(*Operations)(nil),     // 2: sf.substreams.internal.v2.Operations
	(*Operation)(nil),      // 3: sf.substreams.internal.v2.Operation
	(*anypb.Any)(nil),      // 4: google.protobuf.Any
	(*v1.StoreDeltas)(nil), // 5: sf.substreams.v1.StoreDeltas
}
var file_sf_substreams_intern_v2_deltas_proto_depIdxs = []int32{
	4, // 0: sf.substreams.internal.v2.ModuleOutput.map_output:type_name -> google.protobuf.Any
	5, // 1: sf.substreams.internal.v2.ModuleOutput.store_deltas:type_name -> sf.substreams.v1.StoreDeltas
	3, // 2: sf.substreams.internal.v2.Operations.operations:type_name -> sf.substreams.internal.v2.Operation
	0, // 3: sf.substreams.internal.v2.Operation.type:type_name -> sf.substreams.internal.v2.Operation.Type
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_sf_substreams_intern_v2_deltas_proto_init() }
func file_sf_substreams_intern_v2_deltas_proto_init() {
	if File_sf_substreams_intern_v2_deltas_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sf_substreams_intern_v2_deltas_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ModuleOutput); i {
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
		file_sf_substreams_intern_v2_deltas_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Operations); i {
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
		file_sf_substreams_intern_v2_deltas_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Operation); i {
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
	file_sf_substreams_intern_v2_deltas_proto_msgTypes[0].OneofWrappers = []any{
		(*ModuleOutput_MapOutput)(nil),
		(*ModuleOutput_StoreDeltas)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sf_substreams_intern_v2_deltas_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sf_substreams_intern_v2_deltas_proto_goTypes,
		DependencyIndexes: file_sf_substreams_intern_v2_deltas_proto_depIdxs,
		EnumInfos:         file_sf_substreams_intern_v2_deltas_proto_enumTypes,
		MessageInfos:      file_sf_substreams_intern_v2_deltas_proto_msgTypes,
	}.Build()
	File_sf_substreams_intern_v2_deltas_proto = out.File
	file_sf_substreams_intern_v2_deltas_proto_rawDesc = nil
	file_sf_substreams_intern_v2_deltas_proto_goTypes = nil
	file_sf_substreams_intern_v2_deltas_proto_depIdxs = nil
}
