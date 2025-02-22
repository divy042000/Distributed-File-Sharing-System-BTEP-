// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.12.4
// source: chunk.proto

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

// Request for writing a chunk
type WriteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChunkId       string                 `protobuf:"bytes,1,opt,name=chunk_id,json=chunkId,proto3" json:"chunk_id,omitempty"`
	Data          []byte                 `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Version       int32                  `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	mi := &file_chunk_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chunk_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_chunk_proto_rawDescGZIP(), []int{0}
}

func (x *WriteRequest) GetChunkId() string {
	if x != nil {
		return x.ChunkId
	}
	return ""
}

func (x *WriteRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *WriteRequest) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

// Response after writing a chunk
type WriteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WriteResponse) Reset() {
	*x = WriteResponse{}
	mi := &file_chunk_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteResponse) ProtoMessage() {}

func (x *WriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chunk_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteResponse.ProtoReflect.Descriptor instead.
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return file_chunk_proto_rawDescGZIP(), []int{1}
}

func (x *WriteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *WriteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Request for reading a chunk
type ReadRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChunkId       string                 `protobuf:"bytes,1,opt,name=chunk_id,json=chunkId,proto3" json:"chunk_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	mi := &file_chunk_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chunk_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_chunk_proto_rawDescGZIP(), []int{2}
}

func (x *ReadRequest) GetChunkId() string {
	if x != nil {
		return x.ChunkId
	}
	return ""
}

// Response after reading a chunk
type ReadResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Data          []byte                 `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadResponse) Reset() {
	*x = ReadResponse{}
	mi := &file_chunk_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse) ProtoMessage() {}

func (x *ReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chunk_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse.ProtoReflect.Descriptor instead.
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return file_chunk_proto_rawDescGZIP(), []int{3}
}

func (x *ReadResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ReadResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// Request for chunk replication
type ReplicationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChunkId       string                 `protobuf:"bytes,1,opt,name=chunk_id,json=chunkId,proto3" json:"chunk_id,omitempty"`
	Data          []byte                 `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Version       int32                  `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"` // Ensures consistent versioning across replicas
	Followers     []string               `protobuf:"bytes,4,rep,name=followers,proto3" json:"followers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReplicationRequest) Reset() {
	*x = ReplicationRequest{}
	mi := &file_chunk_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReplicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplicationRequest) ProtoMessage() {}

func (x *ReplicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chunk_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplicationRequest.ProtoReflect.Descriptor instead.
func (*ReplicationRequest) Descriptor() ([]byte, []int) {
	return file_chunk_proto_rawDescGZIP(), []int{4}
}

func (x *ReplicationRequest) GetChunkId() string {
	if x != nil {
		return x.ChunkId
	}
	return ""
}

func (x *ReplicationRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ReplicationRequest) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ReplicationRequest) GetFollowers() []string {
	if x != nil {
		return x.Followers
	}
	return nil
}

// Response after chunk replication
type ReplicationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReplicationResponse) Reset() {
	*x = ReplicationResponse{}
	mi := &file_chunk_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReplicationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplicationResponse) ProtoMessage() {}

func (x *ReplicationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chunk_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplicationResponse.ProtoReflect.Descriptor instead.
func (*ReplicationResponse) Descriptor() ([]byte, []int) {
	return file_chunk_proto_rawDescGZIP(), []int{5}
}

func (x *ReplicationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ReplicationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_chunk_proto protoreflect.FileDescriptor

var file_chunk_proto_rawDesc = string([]byte{
	0x0a, 0x0b, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x0c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x43, 0x0a,
	0x0d, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x28, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x0c,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x7b, 0x0a, 0x12, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x22, 0x49, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0xc1, 0x01, 0x0a, 0x0c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x57, 0x72, 0x69, 0x74, 0x65, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09,
	0x52, 0x65, 0x61, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_chunk_proto_rawDescOnce sync.Once
	file_chunk_proto_rawDescData []byte
)

func file_chunk_proto_rawDescGZIP() []byte {
	file_chunk_proto_rawDescOnce.Do(func() {
		file_chunk_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_chunk_proto_rawDesc), len(file_chunk_proto_rawDesc)))
	})
	return file_chunk_proto_rawDescData
}

var file_chunk_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_chunk_proto_goTypes = []any{
	(*WriteRequest)(nil),        // 0: proto.WriteRequest
	(*WriteResponse)(nil),       // 1: proto.WriteResponse
	(*ReadRequest)(nil),         // 2: proto.ReadRequest
	(*ReadResponse)(nil),        // 3: proto.ReadResponse
	(*ReplicationRequest)(nil),  // 4: proto.ReplicationRequest
	(*ReplicationResponse)(nil), // 5: proto.ReplicationResponse
}
var file_chunk_proto_depIdxs = []int32{
	0, // 0: proto.ChunkService.WriteChunk:input_type -> proto.WriteRequest
	2, // 1: proto.ChunkService.ReadChunk:input_type -> proto.ReadRequest
	4, // 2: proto.ChunkService.SendChunk:input_type -> proto.ReplicationRequest
	1, // 3: proto.ChunkService.WriteChunk:output_type -> proto.WriteResponse
	3, // 4: proto.ChunkService.ReadChunk:output_type -> proto.ReadResponse
	5, // 5: proto.ChunkService.SendChunk:output_type -> proto.ReplicationResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chunk_proto_init() }
func file_chunk_proto_init() {
	if File_chunk_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_chunk_proto_rawDesc), len(file_chunk_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chunk_proto_goTypes,
		DependencyIndexes: file_chunk_proto_depIdxs,
		MessageInfos:      file_chunk_proto_msgTypes,
	}.Build()
	File_chunk_proto = out.File
	file_chunk_proto_goTypes = nil
	file_chunk_proto_depIdxs = nil
}
