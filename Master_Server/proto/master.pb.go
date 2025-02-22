// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.12.4
// source: master.proto

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

// Request to get chunk location from the Master Server
type GetChunkRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileName      string                 `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	ChunkIndex    int32                  `protobuf:"varint,2,opt,name=chunk_index,json=chunkIndex,proto3" json:"chunk_index,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetChunkRequest) Reset() {
	*x = GetChunkRequest{}
	mi := &file_master_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetChunkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChunkRequest) ProtoMessage() {}

func (x *GetChunkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChunkRequest.ProtoReflect.Descriptor instead.
func (*GetChunkRequest) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{0}
}

func (x *GetChunkRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *GetChunkRequest) GetChunkIndex() int32 {
	if x != nil {
		return x.ChunkIndex
	}
	return 0
}

// Response containing chunk locations
type GetChunkResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChunkId       string                 `protobuf:"bytes,1,opt,name=chunk_id,json=chunkId,proto3" json:"chunk_id,omitempty"`
	ChunkServers  []string               `protobuf:"bytes,2,rep,name=chunk_servers,json=chunkServers,proto3" json:"chunk_servers,omitempty"`
	Success       bool                   `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"` // ✅ Add this field
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetChunkResponse) Reset() {
	*x = GetChunkResponse{}
	mi := &file_master_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetChunkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChunkResponse) ProtoMessage() {}

func (x *GetChunkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChunkResponse.ProtoReflect.Descriptor instead.
func (*GetChunkResponse) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{1}
}

func (x *GetChunkResponse) GetChunkId() string {
	if x != nil {
		return x.ChunkId
	}
	return ""
}

func (x *GetChunkResponse) GetChunkServers() []string {
	if x != nil {
		return x.ChunkServers
	}
	return nil
}

func (x *GetChunkResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetChunkResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Chunk Server reporting new chunk
type ChunkReport struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChunkId       string                 `protobuf:"bytes,1,opt,name=chunk_id,json=chunkId,proto3" json:"chunk_id,omitempty"`
	ServerId      string                 `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	Version       int32                  `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChunkReport) Reset() {
	*x = ChunkReport{}
	mi := &file_master_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChunkReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkReport) ProtoMessage() {}

func (x *ChunkReport) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkReport.ProtoReflect.Descriptor instead.
func (*ChunkReport) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{2}
}

func (x *ChunkReport) GetChunkId() string {
	if x != nil {
		return x.ChunkId
	}
	return ""
}

func (x *ChunkReport) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *ChunkReport) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

// Response from the Master
type ChunkResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChunkResponse) Reset() {
	*x = ChunkResponse{}
	mi := &file_master_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChunkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkResponse) ProtoMessage() {}

func (x *ChunkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkResponse.ProtoReflect.Descriptor instead.
func (*ChunkResponse) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{3}
}

func (x *ChunkResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ChunkResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_master_proto protoreflect.FileDescriptor

var file_master_proto_rawDesc = string([]byte{
	0x0a, 0x0c, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x86, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x5f, 0x0a, 0x0b, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x43, 0x0a, 0x0d, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xd2, 0x01, 0x0a, 0x0d, 0x4d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37,
	0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_master_proto_rawDescOnce sync.Once
	file_master_proto_rawDescData []byte
)

func file_master_proto_rawDescGZIP() []byte {
	file_master_proto_rawDescOnce.Do(func() {
		file_master_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_master_proto_rawDesc), len(file_master_proto_rawDesc)))
	})
	return file_master_proto_rawDescData
}

var file_master_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_master_proto_goTypes = []any{
	(*GetChunkRequest)(nil),   // 0: proto.GetChunkRequest
	(*GetChunkResponse)(nil),  // 1: proto.GetChunkResponse
	(*ChunkReport)(nil),       // 2: proto.ChunkReport
	(*ChunkResponse)(nil),     // 3: proto.ChunkResponse
	(*HeartbeatRequest)(nil),  // 4: proto.HeartbeatRequest
	(*HeartbeatResponse)(nil), // 5: proto.HeartbeatResponse
}
var file_master_proto_depIdxs = []int32{
	0, // 0: proto.MasterService.GetChunkLocations:input_type -> proto.GetChunkRequest
	2, // 1: proto.MasterService.ReportChunk:input_type -> proto.ChunkReport
	4, // 2: proto.MasterService.SendHeartbeat:input_type -> proto.HeartbeatRequest
	1, // 3: proto.MasterService.GetChunkLocations:output_type -> proto.GetChunkResponse
	3, // 4: proto.MasterService.ReportChunk:output_type -> proto.ChunkResponse
	5, // 5: proto.MasterService.SendHeartbeat:output_type -> proto.HeartbeatResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_master_proto_init() }
func file_master_proto_init() {
	if File_master_proto != nil {
		return
	}
	file_heartbeat_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_master_proto_rawDesc), len(file_master_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_master_proto_goTypes,
		DependencyIndexes: file_master_proto_depIdxs,
		MessageInfos:      file_master_proto_msgTypes,
	}.Build()
	File_master_proto = out.File
	file_master_proto_goTypes = nil
	file_master_proto_depIdxs = nil
}
