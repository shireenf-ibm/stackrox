// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.3
// source: storage/service_account.proto

package storage

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// Any properties of an individual service account.
// (regardless of time, scope, or context)
// ////////////////////////////////////////
type ServiceAccount struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Id               string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk,type(uuid)"`                                                                                             // @gotags: sql:"pk,type(uuid)"
	Name             string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" search:"Service Account,store"`                                                                                         // @gotags: search:"Service Account,store"
	Namespace        string                 `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty" search:"Namespace,store"`                                                                               // @gotags: search:"Namespace,store"
	ClusterName      string                 `protobuf:"bytes,4,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty" search:"Cluster,store"`                                                        // @gotags: search:"Cluster,store"
	ClusterId        string                 `protobuf:"bytes,5,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty" search:"Cluster ID,store,hidden" sql:"type(uuid)"`                                                              // @gotags: search:"Cluster ID,store,hidden" sql:"type(uuid)"
	Labels           map[string]string      `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value" search:"Service Account Label"`           // @gotags: search:"Service Account Label"
	Annotations      map[string]string      `protobuf:"bytes,7,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value" search:"Service Account Annotation"` // @gotags: search:"Service Account Annotation"
	CreatedAt        *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	AutomountToken   bool                   `protobuf:"varint,9,opt,name=automount_token,json=automountToken,proto3" json:"automount_token,omitempty"`
	Secrets          []string               `protobuf:"bytes,10,rep,name=secrets,proto3" json:"secrets,omitempty"`
	ImagePullSecrets []string               `protobuf:"bytes,11,rep,name=image_pull_secrets,json=imagePullSecrets,proto3" json:"image_pull_secrets,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *ServiceAccount) Reset() {
	*x = ServiceAccount{}
	mi := &file_storage_service_account_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServiceAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceAccount) ProtoMessage() {}

func (x *ServiceAccount) ProtoReflect() protoreflect.Message {
	mi := &file_storage_service_account_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceAccount.ProtoReflect.Descriptor instead.
func (*ServiceAccount) Descriptor() ([]byte, []int) {
	return file_storage_service_account_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceAccount) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ServiceAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceAccount) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ServiceAccount) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *ServiceAccount) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *ServiceAccount) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *ServiceAccount) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *ServiceAccount) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ServiceAccount) GetAutomountToken() bool {
	if x != nil {
		return x.AutomountToken
	}
	return false
}

func (x *ServiceAccount) GetSecrets() []string {
	if x != nil {
		return x.Secrets
	}
	return nil
}

func (x *ServiceAccount) GetImagePullSecrets() []string {
	if x != nil {
		return x.ImagePullSecrets
	}
	return nil
}

var File_storage_service_account_proto protoreflect.FileDescriptor

const file_storage_service_account_proto_rawDesc = "" +
	"\n" +
	"\x1dstorage/service_account.proto\x12\astorage\x1a\x1fgoogle/protobuf/timestamp.proto\"\xc4\x04\n" +
	"\x0eServiceAccount\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x1c\n" +
	"\tnamespace\x18\x03 \x01(\tR\tnamespace\x12!\n" +
	"\fcluster_name\x18\x04 \x01(\tR\vclusterName\x12\x1d\n" +
	"\n" +
	"cluster_id\x18\x05 \x01(\tR\tclusterId\x12;\n" +
	"\x06labels\x18\x06 \x03(\v2#.storage.ServiceAccount.LabelsEntryR\x06labels\x12J\n" +
	"\vannotations\x18\a \x03(\v2(.storage.ServiceAccount.AnnotationsEntryR\vannotations\x129\n" +
	"\n" +
	"created_at\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12'\n" +
	"\x0fautomount_token\x18\t \x01(\bR\x0eautomountToken\x12\x18\n" +
	"\asecrets\x18\n" +
	" \x03(\tR\asecrets\x12,\n" +
	"\x12image_pull_secrets\x18\v \x03(\tR\x10imagePullSecrets\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\x1a>\n" +
	"\x10AnnotationsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B.\n" +
	"\x19io.stackrox.proto.storageZ\x11./storage;storageb\x06proto3"

var (
	file_storage_service_account_proto_rawDescOnce sync.Once
	file_storage_service_account_proto_rawDescData []byte
)

func file_storage_service_account_proto_rawDescGZIP() []byte {
	file_storage_service_account_proto_rawDescOnce.Do(func() {
		file_storage_service_account_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_storage_service_account_proto_rawDesc), len(file_storage_service_account_proto_rawDesc)))
	})
	return file_storage_service_account_proto_rawDescData
}

var file_storage_service_account_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_storage_service_account_proto_goTypes = []any{
	(*ServiceAccount)(nil),        // 0: storage.ServiceAccount
	nil,                           // 1: storage.ServiceAccount.LabelsEntry
	nil,                           // 2: storage.ServiceAccount.AnnotationsEntry
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_storage_service_account_proto_depIdxs = []int32{
	1, // 0: storage.ServiceAccount.labels:type_name -> storage.ServiceAccount.LabelsEntry
	2, // 1: storage.ServiceAccount.annotations:type_name -> storage.ServiceAccount.AnnotationsEntry
	3, // 2: storage.ServiceAccount.created_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_storage_service_account_proto_init() }
func file_storage_service_account_proto_init() {
	if File_storage_service_account_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_storage_service_account_proto_rawDesc), len(file_storage_service_account_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_service_account_proto_goTypes,
		DependencyIndexes: file_storage_service_account_proto_depIdxs,
		MessageInfos:      file_storage_service_account_proto_msgTypes,
	}.Build()
	File_storage_service_account_proto = out.File
	file_storage_service_account_proto_goTypes = nil
	file_storage_service_account_proto_depIdxs = nil
}
