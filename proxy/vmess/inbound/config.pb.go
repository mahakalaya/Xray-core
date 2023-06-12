// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: proxy/vmess/inbound/config.proto

package inbound

import (
	protocol "github.com/xtls/xray-core/common/protocol"
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

type DetourConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To string `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *DetourConfig) Reset() {
	*x = DetourConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_vmess_inbound_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetourConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetourConfig) ProtoMessage() {}

func (x *DetourConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_vmess_inbound_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetourConfig.ProtoReflect.Descriptor instead.
func (*DetourConfig) Descriptor() ([]byte, []int) {
	return file_proxy_vmess_inbound_config_proto_rawDescGZIP(), []int{0}
}

func (x *DetourConfig) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

type DefaultConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlterId uint32 `protobuf:"varint,1,opt,name=alter_id,json=alterId,proto3" json:"alter_id,omitempty"`
	Level   uint32 `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *DefaultConfig) Reset() {
	*x = DefaultConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_vmess_inbound_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefaultConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefaultConfig) ProtoMessage() {}

func (x *DefaultConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_vmess_inbound_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefaultConfig.ProtoReflect.Descriptor instead.
func (*DefaultConfig) Descriptor() ([]byte, []int) {
	return file_proxy_vmess_inbound_config_proto_rawDescGZIP(), []int{1}
}

func (x *DefaultConfig) GetAlterId() uint32 {
	if x != nil {
		return x.AlterId
	}
	return 0
}

func (x *DefaultConfig) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User                 []*protocol.User `protobuf:"bytes,1,rep,name=user,proto3" json:"user,omitempty"`
	Default              *DefaultConfig   `protobuf:"bytes,2,opt,name=default,proto3" json:"default,omitempty"`
	Detour               *DetourConfig    `protobuf:"bytes,3,opt,name=detour,proto3" json:"detour,omitempty"`
	SecureEncryptionOnly bool             `protobuf:"varint,4,opt,name=secure_encryption_only,json=secureEncryptionOnly,proto3" json:"secure_encryption_only,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_vmess_inbound_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_vmess_inbound_config_proto_msgTypes[2]
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
	return file_proxy_vmess_inbound_config_proto_rawDescGZIP(), []int{2}
}

func (x *Config) GetUser() []*protocol.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Config) GetDefault() *DefaultConfig {
	if x != nil {
		return x.Default
	}
	return nil
}

func (x *Config) GetDetour() *DetourConfig {
	if x != nil {
		return x.Detour
	}
	return nil
}

func (x *Config) GetSecureEncryptionOnly() bool {
	if x != nil {
		return x.SecureEncryptionOnly
	}
	return false
}

var File_proxy_vmess_inbound_config_proto protoreflect.FileDescriptor

var file_proxy_vmess_inbound_config_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x6d, 0x65, 0x73, 0x73, 0x2f, 0x69, 0x6e,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x18, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76,
	0x6d, 0x65, 0x73, 0x73, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x1a, 0x1a, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x0c, 0x44, 0x65, 0x74, 0x6f,
	0x75, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x0d, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6c, 0x74,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x6c, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0xf1, 0x01, 0x0a, 0x06, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x76, 0x6d, 0x65, 0x73, 0x73, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x3e, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x6f,
	0x75, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x6d, 0x65, 0x73, 0x73, 0x2e, 0x69, 0x6e, 0x62, 0x6f,
	0x75, 0x6e, 0x64, 0x2e, 0x44, 0x65, 0x74, 0x6f, 0x75, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x06, 0x64, 0x65, 0x74, 0x6f, 0x75, 0x72, 0x12, 0x34, 0x0a, 0x16, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x65, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x6e,
	0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65,
	0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x6e, 0x6c, 0x79, 0x42, 0x6a,
	0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x76, 0x6d, 0x65, 0x73, 0x73, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x50, 0x01,
	0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x74, 0x6c,
	0x73, 0x2f, 0x78, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x76, 0x6d, 0x65, 0x73, 0x73, 0x2f, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0xaa,
	0x02, 0x18, 0x58, 0x72, 0x61, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x56, 0x6d, 0x65,
	0x73, 0x73, 0x2e, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proxy_vmess_inbound_config_proto_rawDescOnce sync.Once
	file_proxy_vmess_inbound_config_proto_rawDescData = file_proxy_vmess_inbound_config_proto_rawDesc
)

func file_proxy_vmess_inbound_config_proto_rawDescGZIP() []byte {
	file_proxy_vmess_inbound_config_proto_rawDescOnce.Do(func() {
		file_proxy_vmess_inbound_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_proxy_vmess_inbound_config_proto_rawDescData)
	})
	return file_proxy_vmess_inbound_config_proto_rawDescData
}

var file_proxy_vmess_inbound_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proxy_vmess_inbound_config_proto_goTypes = []interface{}{
	(*DetourConfig)(nil),  // 0: xray.proxy.vmess.inbound.DetourConfig
	(*DefaultConfig)(nil), // 1: xray.proxy.vmess.inbound.DefaultConfig
	(*Config)(nil),        // 2: xray.proxy.vmess.inbound.Config
	(*protocol.User)(nil), // 3: xray.common.protocol.User
}
var file_proxy_vmess_inbound_config_proto_depIdxs = []int32{
	3, // 0: xray.proxy.vmess.inbound.Config.user:type_name -> xray.common.protocol.User
	1, // 1: xray.proxy.vmess.inbound.Config.default:type_name -> xray.proxy.vmess.inbound.DefaultConfig
	0, // 2: xray.proxy.vmess.inbound.Config.detour:type_name -> xray.proxy.vmess.inbound.DetourConfig
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proxy_vmess_inbound_config_proto_init() }
func file_proxy_vmess_inbound_config_proto_init() {
	if File_proxy_vmess_inbound_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proxy_vmess_inbound_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetourConfig); i {
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
		file_proxy_vmess_inbound_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefaultConfig); i {
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
		file_proxy_vmess_inbound_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proxy_vmess_inbound_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proxy_vmess_inbound_config_proto_goTypes,
		DependencyIndexes: file_proxy_vmess_inbound_config_proto_depIdxs,
		MessageInfos:      file_proxy_vmess_inbound_config_proto_msgTypes,
	}.Build()
	File_proxy_vmess_inbound_config_proto = out.File
	file_proxy_vmess_inbound_config_proto_rawDesc = nil
	file_proxy_vmess_inbound_config_proto_goTypes = nil
	file_proxy_vmess_inbound_config_proto_depIdxs = nil
}
