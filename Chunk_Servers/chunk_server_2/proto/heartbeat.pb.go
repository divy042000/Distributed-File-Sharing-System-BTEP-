// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.2
// source: heartbeat.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Heartbeat Request from Chunk Server to Master Server
type HeartbeatRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ServerId      string                 `protobuf:"bytes,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`               // Unique ID of the chunk server
	StoragePath   string                 `protobuf:"bytes,2,opt,name=storage_path,json=storagePath,proto3" json:"storage_path,omitempty"`      // Path to chunk storage directory
	FreeSpace     int64                  `protobuf:"varint,3,opt,name=free_space,json=freeSpace,proto3" json:"free_space,omitempty"`           // Available disk space in MB
	CpuUsage      float32                `protobuf:"fixed32,4,opt,name=cpu_usage,json=cpuUsage,proto3" json:"cpu_usage,omitempty"`             // CPU usage percentage
	MemoryUsage   float32                `protobuf:"fixed32,5,opt,name=memory_usage,json=memoryUsage,proto3" json:"memory_usage,omitempty"`    // Memory usage percentage
	NetworkUsage  float32                `protobuf:"fixed32,6,opt,name=network_usage,json=networkUsage,proto3" json:"network_usage,omitempty"` // Network bandwidth usage
	Load          float32                `protobuf:"fixed32,7,opt,name=load,proto3" json:"load,omitempty"`                                     // System load average
	ChunkIds      []string               `protobuf:"bytes,8,rep,name=chunk_ids,json=chunkIds,proto3" json:"chunk_ids,omitempty"`               // List of stored chunks
	TotalSpace    int64                  `protobuf:"varint,9,opt,name=total_space,json=totalSpace,proto3" json:"total_space,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HeartbeatRequest) Reset() {
	*x = HeartbeatRequest{}
	mi := &file_heartbeat_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatRequest) ProtoMessage() {}

func (x *HeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heartbeat_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatRequest.ProtoReflect.Descriptor instead.
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_heartbeat_proto_rawDescGZIP(), []int{0}
}

func (x *HeartbeatRequest) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *HeartbeatRequest) GetStoragePath() string {
	if x != nil {
		return x.StoragePath
	}
	return ""
}

func (x *HeartbeatRequest) GetFreeSpace() int64 {
	if x != nil {
		return x.FreeSpace
	}
	return 0
}

func (x *HeartbeatRequest) GetCpuUsage() float32 {
	if x != nil {
		return x.CpuUsage
	}
	return 0
}

func (x *HeartbeatRequest) GetMemoryUsage() float32 {
	if x != nil {
		return x.MemoryUsage
	}
	return 0
}

func (x *HeartbeatRequest) GetNetworkUsage() float32 {
	if x != nil {
		return x.NetworkUsage
	}
	return 0
}

func (x *HeartbeatRequest) GetLoad() float32 {
	if x != nil {
		return x.Load
	}
	return 0
}

func (x *HeartbeatRequest) GetChunkIds() []string {
	if x != nil {
		return x.ChunkIds
	}
	return nil
}

func (x *HeartbeatRequest) GetTotalSpace() int64 {
	if x != nil {
		return x.TotalSpace
	}
	return 0
}

// Heartbeat Response from Master Server to Chunk Server
type HeartbeatResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // Indicates if heartbeat was received
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`  // Response message from Master
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HeartbeatResponse) Reset() {
	*x = HeartbeatResponse{}
	mi := &file_heartbeat_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResponse) ProtoMessage() {}

func (x *HeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_heartbeat_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_heartbeat_proto_rawDescGZIP(), []int{1}
}

func (x *HeartbeatResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *HeartbeatResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_heartbeat_proto protoreflect.FileDescriptor

var file_heartbeat_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x02, 0x0a, 0x10, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1d, 0x0a,
	0x0a, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x66, 0x72, 0x65, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x70, 0x75, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x08, 0x63, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x04, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49,
	0x64, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x70,
	0x61, 0x63, 0x65, 0x22, 0x47, 0x0a, 0x11, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x56, 0x0a, 0x10,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x42, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61,
	0x74, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_heartbeat_proto_rawDescOnce sync.Once
	file_heartbeat_proto_rawDescData []byte
)

func file_heartbeat_proto_rawDescGZIP() []byte {
	file_heartbeat_proto_rawDescOnce.Do(func() {
		file_heartbeat_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_heartbeat_proto_rawDesc), len(file_heartbeat_proto_rawDesc)))
	})
	return file_heartbeat_proto_rawDescData
}

var file_heartbeat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_heartbeat_proto_goTypes = []any{
	(*HeartbeatRequest)(nil),  // 0: proto.HeartbeatRequest
	(*HeartbeatResponse)(nil), // 1: proto.HeartbeatResponse
}
var file_heartbeat_proto_depIdxs = []int32{
	0, // 0: proto.HeartbeatService.SendHeartbeat:input_type -> proto.HeartbeatRequest
	1, // 1: proto.HeartbeatService.SendHeartbeat:output_type -> proto.HeartbeatResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_heartbeat_proto_init() }
func file_heartbeat_proto_init() {
	if File_heartbeat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_heartbeat_proto_rawDesc), len(file_heartbeat_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_heartbeat_proto_goTypes,
		DependencyIndexes: file_heartbeat_proto_depIdxs,
		MessageInfos:      file_heartbeat_proto_msgTypes,
	}.Build()
	File_heartbeat_proto = out.File
	file_heartbeat_proto_goTypes = nil
	file_heartbeat_proto_depIdxs = nil
}
