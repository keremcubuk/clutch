// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: config/service/topology/topology.proto

package topologyv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cache *Cache `protobuf:"bytes,1,opt,name=cache,proto3" json:"cache,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_topology_topology_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_topology_topology_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_config_service_topology_topology_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetCache() *Cache {
	if x != nil {
		return x.Cache
	}
	return nil
}

type Cache struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The time to live (ttl) for an item in cache, the default is 60 minutes.
	Ttl *duration.Duration `protobuf:"bytes,1,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *Cache) Reset() {
	*x = Cache{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_topology_topology_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cache) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cache) ProtoMessage() {}

func (x *Cache) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_topology_topology_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cache.ProtoReflect.Descriptor instead.
func (*Cache) Descriptor() ([]byte, []int) {
	return file_config_service_topology_topology_proto_rawDescGZIP(), []int{1}
}

func (x *Cache) GetTtl() *duration.Duration {
	if x != nil {
		return x.Ttl
	}
	return nil
}

var File_config_service_topology_topology_proto protoreflect.FileDescriptor

var file_config_service_topology_topology_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f,
	0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3e,
	0x0a, 0x05, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x05, 0x63, 0x61, 0x63, 0x68, 0x65, 0x22, 0x40,
	0x0a, 0x05, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x37, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x0a, 0xfa, 0x42, 0x07, 0xaa, 0x01, 0x04, 0x2a, 0x02, 0x08, 0x01, 0x52, 0x03, 0x74, 0x74, 0x6c,
	0x42, 0x0c, 0x5a, 0x0a, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_service_topology_topology_proto_rawDescOnce sync.Once
	file_config_service_topology_topology_proto_rawDescData = file_config_service_topology_topology_proto_rawDesc
)

func file_config_service_topology_topology_proto_rawDescGZIP() []byte {
	file_config_service_topology_topology_proto_rawDescOnce.Do(func() {
		file_config_service_topology_topology_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_service_topology_topology_proto_rawDescData)
	})
	return file_config_service_topology_topology_proto_rawDescData
}

var file_config_service_topology_topology_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_config_service_topology_topology_proto_goTypes = []interface{}{
	(*Config)(nil),            // 0: clutch.config.service.topology.v1.Config
	(*Cache)(nil),             // 1: clutch.config.service.topology.v1.Cache
	(*duration.Duration)(nil), // 2: google.protobuf.Duration
}
var file_config_service_topology_topology_proto_depIdxs = []int32{
	1, // 0: clutch.config.service.topology.v1.Config.cache:type_name -> clutch.config.service.topology.v1.Cache
	2, // 1: clutch.config.service.topology.v1.Cache.ttl:type_name -> google.protobuf.Duration
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_config_service_topology_topology_proto_init() }
func file_config_service_topology_topology_proto_init() {
	if File_config_service_topology_topology_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_service_topology_topology_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_config_service_topology_topology_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cache); i {
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
			RawDescriptor: file_config_service_topology_topology_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_service_topology_topology_proto_goTypes,
		DependencyIndexes: file_config_service_topology_topology_proto_depIdxs,
		MessageInfos:      file_config_service_topology_topology_proto_msgTypes,
	}.Build()
	File_config_service_topology_topology_proto = out.File
	file_config_service_topology_topology_proto_rawDesc = nil
	file_config_service_topology_topology_proto_goTypes = nil
	file_config_service_topology_topology_proto_depIdxs = nil
}