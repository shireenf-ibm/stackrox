// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.3
// source: storage/cloud_source.proto

package storage

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

type CloudSource_Type int32

const (
	CloudSource_TYPE_UNSPECIFIED   CloudSource_Type = 0
	CloudSource_TYPE_PALADIN_CLOUD CloudSource_Type = 1
	CloudSource_TYPE_OCM           CloudSource_Type = 2
)

// Enum value maps for CloudSource_Type.
var (
	CloudSource_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_PALADIN_CLOUD",
		2: "TYPE_OCM",
	}
	CloudSource_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED":   0,
		"TYPE_PALADIN_CLOUD": 1,
		"TYPE_OCM":           2,
	}
)

func (x CloudSource_Type) Enum() *CloudSource_Type {
	p := new(CloudSource_Type)
	*p = x
	return p
}

func (x CloudSource_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CloudSource_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_cloud_source_proto_enumTypes[0].Descriptor()
}

func (CloudSource_Type) Type() protoreflect.EnumType {
	return &file_storage_cloud_source_proto_enumTypes[0]
}

func (x CloudSource_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CloudSource_Type.Descriptor instead.
func (CloudSource_Type) EnumDescriptor() ([]byte, []int) {
	return file_storage_cloud_source_proto_rawDescGZIP(), []int{0, 0}
}

type CloudSource struct {
	state               protoimpl.MessageState   `protogen:"open.v1"`
	Id                  string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk,type(uuid)"`                                    // @gotags: sql:"pk,type(uuid)"
	Name                string                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" sql:"unique" search:"Integration Name,hidden"`                                // @gotags: sql:"unique" search:"Integration Name,hidden"
	Type                CloudSource_Type         `protobuf:"varint,3,opt,name=type,proto3,enum=storage.CloudSource_Type" json:"type,omitempty" search:"Integration Type,hidden"` // @gotags: search:"Integration Type,hidden"
	Credentials         *CloudSource_Credentials `protobuf:"bytes,4,opt,name=credentials,proto3" json:"credentials,omitempty"`
	SkipTestIntegration bool                     `protobuf:"varint,5,opt,name=skip_test_integration,json=skipTestIntegration,proto3" json:"skip_test_integration,omitempty"`
	// Types that are valid to be assigned to Config:
	//
	//	*CloudSource_PaladinCloud
	//	*CloudSource_Ocm
	Config        isCloudSource_Config `protobuf_oneof:"Config"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CloudSource) Reset() {
	*x = CloudSource{}
	mi := &file_storage_cloud_source_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CloudSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudSource) ProtoMessage() {}

func (x *CloudSource) ProtoReflect() protoreflect.Message {
	mi := &file_storage_cloud_source_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudSource.ProtoReflect.Descriptor instead.
func (*CloudSource) Descriptor() ([]byte, []int) {
	return file_storage_cloud_source_proto_rawDescGZIP(), []int{0}
}

func (x *CloudSource) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CloudSource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CloudSource) GetType() CloudSource_Type {
	if x != nil {
		return x.Type
	}
	return CloudSource_TYPE_UNSPECIFIED
}

func (x *CloudSource) GetCredentials() *CloudSource_Credentials {
	if x != nil {
		return x.Credentials
	}
	return nil
}

func (x *CloudSource) GetSkipTestIntegration() bool {
	if x != nil {
		return x.SkipTestIntegration
	}
	return false
}

func (x *CloudSource) GetConfig() isCloudSource_Config {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *CloudSource) GetPaladinCloud() *PaladinCloudConfig {
	if x != nil {
		if x, ok := x.Config.(*CloudSource_PaladinCloud); ok {
			return x.PaladinCloud
		}
	}
	return nil
}

func (x *CloudSource) GetOcm() *OCMConfig {
	if x != nil {
		if x, ok := x.Config.(*CloudSource_Ocm); ok {
			return x.Ocm
		}
	}
	return nil
}

type isCloudSource_Config interface {
	isCloudSource_Config()
}

type CloudSource_PaladinCloud struct {
	PaladinCloud *PaladinCloudConfig `protobuf:"bytes,6,opt,name=paladin_cloud,json=paladinCloud,proto3,oneof"`
}

type CloudSource_Ocm struct {
	Ocm *OCMConfig `protobuf:"bytes,7,opt,name=ocm,proto3,oneof"`
}

func (*CloudSource_PaladinCloud) isCloudSource_Config() {}

func (*CloudSource_Ocm) isCloudSource_Config() {}

type PaladinCloudConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Endpoint      string                 `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty" validate:"nolocalendpoint"` // @gotags: validate:"nolocalendpoint"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaladinCloudConfig) Reset() {
	*x = PaladinCloudConfig{}
	mi := &file_storage_cloud_source_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaladinCloudConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaladinCloudConfig) ProtoMessage() {}

func (x *PaladinCloudConfig) ProtoReflect() protoreflect.Message {
	mi := &file_storage_cloud_source_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaladinCloudConfig.ProtoReflect.Descriptor instead.
func (*PaladinCloudConfig) Descriptor() ([]byte, []int) {
	return file_storage_cloud_source_proto_rawDescGZIP(), []int{1}
}

func (x *PaladinCloudConfig) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type OCMConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Endpoint      string                 `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty" validate:"nolocalendpoint"` // @gotags: validate:"nolocalendpoint"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OCMConfig) Reset() {
	*x = OCMConfig{}
	mi := &file_storage_cloud_source_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OCMConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OCMConfig) ProtoMessage() {}

func (x *OCMConfig) ProtoReflect() protoreflect.Message {
	mi := &file_storage_cloud_source_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OCMConfig.ProtoReflect.Descriptor instead.
func (*OCMConfig) Descriptor() ([]byte, []int) {
	return file_storage_cloud_source_proto_rawDescGZIP(), []int{2}
}

func (x *OCMConfig) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type CloudSource_Credentials struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Secret        string                 `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty" scrub:"always"`                                 // @gotags: scrub:"always"
	ClientId      string                 `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty" scrub:"always"`             // @gotags: scrub:"always"
	ClientSecret  string                 `protobuf:"bytes,3,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty" scrub:"always"` // @gotags: scrub:"always"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CloudSource_Credentials) Reset() {
	*x = CloudSource_Credentials{}
	mi := &file_storage_cloud_source_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CloudSource_Credentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudSource_Credentials) ProtoMessage() {}

func (x *CloudSource_Credentials) ProtoReflect() protoreflect.Message {
	mi := &file_storage_cloud_source_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudSource_Credentials.ProtoReflect.Descriptor instead.
func (*CloudSource_Credentials) Descriptor() ([]byte, []int) {
	return file_storage_cloud_source_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CloudSource_Credentials) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *CloudSource_Credentials) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *CloudSource_Credentials) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

var File_storage_cloud_source_proto protoreflect.FileDescriptor

const file_storage_cloud_source_proto_rawDesc = "" +
	"\n" +
	"\x1astorage/cloud_source.proto\x12\astorage\"\xfb\x03\n" +
	"\vCloudSource\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12-\n" +
	"\x04type\x18\x03 \x01(\x0e2\x19.storage.CloudSource.TypeR\x04type\x12B\n" +
	"\vcredentials\x18\x04 \x01(\v2 .storage.CloudSource.CredentialsR\vcredentials\x122\n" +
	"\x15skip_test_integration\x18\x05 \x01(\bR\x13skipTestIntegration\x12B\n" +
	"\rpaladin_cloud\x18\x06 \x01(\v2\x1b.storage.PaladinCloudConfigH\x00R\fpaladinCloud\x12&\n" +
	"\x03ocm\x18\a \x01(\v2\x12.storage.OCMConfigH\x00R\x03ocm\x1ag\n" +
	"\vCredentials\x12\x16\n" +
	"\x06secret\x18\x01 \x01(\tR\x06secret\x12\x1b\n" +
	"\tclient_id\x18\x02 \x01(\tR\bclientId\x12#\n" +
	"\rclient_secret\x18\x03 \x01(\tR\fclientSecret\"B\n" +
	"\x04Type\x12\x14\n" +
	"\x10TYPE_UNSPECIFIED\x10\x00\x12\x16\n" +
	"\x12TYPE_PALADIN_CLOUD\x10\x01\x12\f\n" +
	"\bTYPE_OCM\x10\x02B\b\n" +
	"\x06Config\"0\n" +
	"\x12PaladinCloudConfig\x12\x1a\n" +
	"\bendpoint\x18\x01 \x01(\tR\bendpoint\"'\n" +
	"\tOCMConfig\x12\x1a\n" +
	"\bendpoint\x18\x01 \x01(\tR\bendpointB.\n" +
	"\x19io.stackrox.proto.storageZ\x11./storage;storageb\x06proto3"

var (
	file_storage_cloud_source_proto_rawDescOnce sync.Once
	file_storage_cloud_source_proto_rawDescData []byte
)

func file_storage_cloud_source_proto_rawDescGZIP() []byte {
	file_storage_cloud_source_proto_rawDescOnce.Do(func() {
		file_storage_cloud_source_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_storage_cloud_source_proto_rawDesc), len(file_storage_cloud_source_proto_rawDesc)))
	})
	return file_storage_cloud_source_proto_rawDescData
}

var file_storage_cloud_source_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_storage_cloud_source_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_storage_cloud_source_proto_goTypes = []any{
	(CloudSource_Type)(0),           // 0: storage.CloudSource.Type
	(*CloudSource)(nil),             // 1: storage.CloudSource
	(*PaladinCloudConfig)(nil),      // 2: storage.PaladinCloudConfig
	(*OCMConfig)(nil),               // 3: storage.OCMConfig
	(*CloudSource_Credentials)(nil), // 4: storage.CloudSource.Credentials
}
var file_storage_cloud_source_proto_depIdxs = []int32{
	0, // 0: storage.CloudSource.type:type_name -> storage.CloudSource.Type
	4, // 1: storage.CloudSource.credentials:type_name -> storage.CloudSource.Credentials
	2, // 2: storage.CloudSource.paladin_cloud:type_name -> storage.PaladinCloudConfig
	3, // 3: storage.CloudSource.ocm:type_name -> storage.OCMConfig
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_storage_cloud_source_proto_init() }
func file_storage_cloud_source_proto_init() {
	if File_storage_cloud_source_proto != nil {
		return
	}
	file_storage_cloud_source_proto_msgTypes[0].OneofWrappers = []any{
		(*CloudSource_PaladinCloud)(nil),
		(*CloudSource_Ocm)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_storage_cloud_source_proto_rawDesc), len(file_storage_cloud_source_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_cloud_source_proto_goTypes,
		DependencyIndexes: file_storage_cloud_source_proto_depIdxs,
		EnumInfos:         file_storage_cloud_source_proto_enumTypes,
		MessageInfos:      file_storage_cloud_source_proto_msgTypes,
	}.Build()
	File_storage_cloud_source_proto = out.File
	file_storage_cloud_source_proto_goTypes = nil
	file_storage_cloud_source_proto_depIdxs = nil
}
