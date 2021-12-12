// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: downloader/downloader.proto

package downloader

import (
	types "github.com/ledgerwatch/erigon-lib/gointerfaces/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DownloadItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path        string      `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	TorrentHash *types.H160 `protobuf:"bytes,2,opt,name=torrent_hash,json=torrentHash,proto3" json:"torrent_hash,omitempty"` // single hash will be resolved as magnet link
}

func (x *DownloadItem) Reset() {
	*x = DownloadItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_downloader_downloader_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadItem) ProtoMessage() {}

func (x *DownloadItem) ProtoReflect() protoreflect.Message {
	mi := &file_downloader_downloader_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadItem.ProtoReflect.Descriptor instead.
func (*DownloadItem) Descriptor() ([]byte, []int) {
	return file_downloader_downloader_proto_rawDescGZIP(), []int{0}
}

func (x *DownloadItem) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *DownloadItem) GetTorrentHash() *types.H160 {
	if x != nil {
		return x.TorrentHash
	}
	return nil
}

type DownloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*DownloadItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"` // single hash will be resolved as magnet link
}

func (x *DownloadRequest) Reset() {
	*x = DownloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_downloader_downloader_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadRequest) ProtoMessage() {}

func (x *DownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_downloader_downloader_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadRequest.ProtoReflect.Descriptor instead.
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return file_downloader_downloader_proto_rawDescGZIP(), []int{1}
}

func (x *DownloadRequest) GetItems() []*DownloadItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type SnapshotsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SnapshotsRequest) Reset() {
	*x = SnapshotsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_downloader_downloader_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotsRequest) ProtoMessage() {}

func (x *SnapshotsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_downloader_downloader_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotsRequest.ProtoReflect.Descriptor instead.
func (*SnapshotsRequest) Descriptor() ([]byte, []int) {
	return file_downloader_downloader_proto_rawDescGZIP(), []int{2}
}

type SnapshotInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path      string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Readiness int32  `protobuf:"varint,2,opt,name=readiness,proto3" json:"readiness,omitempty"`
}

func (x *SnapshotInfo) Reset() {
	*x = SnapshotInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_downloader_downloader_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotInfo) ProtoMessage() {}

func (x *SnapshotInfo) ProtoReflect() protoreflect.Message {
	mi := &file_downloader_downloader_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotInfo.ProtoReflect.Descriptor instead.
func (*SnapshotInfo) Descriptor() ([]byte, []int) {
	return file_downloader_downloader_proto_rawDescGZIP(), []int{3}
}

func (x *SnapshotInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *SnapshotInfo) GetReadiness() int32 {
	if x != nil {
		return x.Readiness
	}
	return 0
}

type SnapshotsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info []*SnapshotInfo `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
}

func (x *SnapshotsReply) Reset() {
	*x = SnapshotsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_downloader_downloader_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotsReply) ProtoMessage() {}

func (x *SnapshotsReply) ProtoReflect() protoreflect.Message {
	mi := &file_downloader_downloader_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotsReply.ProtoReflect.Descriptor instead.
func (*SnapshotsReply) Descriptor() ([]byte, []int) {
	return file_downloader_downloader_proto_rawDescGZIP(), []int{4}
}

func (x *SnapshotsReply) GetInfo() []*SnapshotInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_downloader_downloader_proto protoreflect.FileDescriptor

var file_downloader_downloader_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2f, 0x64, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x64,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x0c, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x2e, 0x0a,
	0x0c, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x48, 0x31, 0x36, 0x30,
	0x52, 0x0b, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x22, 0x41, 0x0a,
	0x0f, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0x12, 0x0a, 0x10, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x40, 0x0a, 0x0c, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65, 0x61,
	0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x22, 0x3e, 0x0a, 0x0e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2c, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x65, 0x72, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x98, 0x01, 0x0a, 0x0a, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x08, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x1b, 0x2e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x09, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x73, 0x12, 0x1c, 0x2e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x72, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72,
	0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x42, 0x19, 0x5a, 0x17, 0x2e, 0x2f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65,
	0x72, 0x3b, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_downloader_downloader_proto_rawDescOnce sync.Once
	file_downloader_downloader_proto_rawDescData = file_downloader_downloader_proto_rawDesc
)

func file_downloader_downloader_proto_rawDescGZIP() []byte {
	file_downloader_downloader_proto_rawDescOnce.Do(func() {
		file_downloader_downloader_proto_rawDescData = protoimpl.X.CompressGZIP(file_downloader_downloader_proto_rawDescData)
	})
	return file_downloader_downloader_proto_rawDescData
}

var file_downloader_downloader_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_downloader_downloader_proto_goTypes = []interface{}{
	(*DownloadItem)(nil),     // 0: downloader.DownloadItem
	(*DownloadRequest)(nil),  // 1: downloader.DownloadRequest
	(*SnapshotsRequest)(nil), // 2: downloader.SnapshotsRequest
	(*SnapshotInfo)(nil),     // 3: downloader.SnapshotInfo
	(*SnapshotsReply)(nil),   // 4: downloader.SnapshotsReply
	(*types.H160)(nil),       // 5: types.H160
	(*emptypb.Empty)(nil),    // 6: google.protobuf.Empty
}
var file_downloader_downloader_proto_depIdxs = []int32{
	5, // 0: downloader.DownloadItem.torrent_hash:type_name -> types.H160
	0, // 1: downloader.DownloadRequest.items:type_name -> downloader.DownloadItem
	3, // 2: downloader.SnapshotsReply.info:type_name -> downloader.SnapshotInfo
	1, // 3: downloader.Downloader.Download:input_type -> downloader.DownloadRequest
	2, // 4: downloader.Downloader.Snapshots:input_type -> downloader.SnapshotsRequest
	6, // 5: downloader.Downloader.Download:output_type -> google.protobuf.Empty
	4, // 6: downloader.Downloader.Snapshots:output_type -> downloader.SnapshotsReply
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_downloader_downloader_proto_init() }
func file_downloader_downloader_proto_init() {
	if File_downloader_downloader_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_downloader_downloader_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadItem); i {
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
		file_downloader_downloader_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadRequest); i {
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
		file_downloader_downloader_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotsRequest); i {
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
		file_downloader_downloader_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotInfo); i {
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
		file_downloader_downloader_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotsReply); i {
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
			RawDescriptor: file_downloader_downloader_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_downloader_downloader_proto_goTypes,
		DependencyIndexes: file_downloader_downloader_proto_depIdxs,
		MessageInfos:      file_downloader_downloader_proto_msgTypes,
	}.Build()
	File_downloader_downloader_proto = out.File
	file_downloader_downloader_proto_rawDesc = nil
	file_downloader_downloader_proto_goTypes = nil
	file_downloader_downloader_proto_depIdxs = nil
}