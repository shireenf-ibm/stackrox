// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.3
// source: api/v1/rbac_service.proto

package v1

import (
	storage "github.com/stackrox/rox/generated/storage"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A list of k8s roles (free of scoped information)
// Next Tag: 2
type ListRolesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Roles         []*storage.K8SRole     `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRolesResponse) Reset() {
	*x = ListRolesResponse{}
	mi := &file_api_v1_rbac_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRolesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRolesResponse) ProtoMessage() {}

func (x *ListRolesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rbac_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRolesResponse.ProtoReflect.Descriptor instead.
func (*ListRolesResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_rbac_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListRolesResponse) GetRoles() []*storage.K8SRole {
	if x != nil {
		return x.Roles
	}
	return nil
}

type GetRoleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Role          *storage.K8SRole       `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRoleResponse) Reset() {
	*x = GetRoleResponse{}
	mi := &file_api_v1_rbac_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleResponse) ProtoMessage() {}

func (x *GetRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rbac_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleResponse.ProtoReflect.Descriptor instead.
func (*GetRoleResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_rbac_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetRoleResponse) GetRole() *storage.K8SRole {
	if x != nil {
		return x.Role
	}
	return nil
}

// A list of k8s role bindings (free of scoped information)
// Next Tag: 2
type ListRoleBindingsResponse struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Bindings      []*storage.K8SRoleBinding `protobuf:"bytes,1,rep,name=bindings,proto3" json:"bindings,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRoleBindingsResponse) Reset() {
	*x = ListRoleBindingsResponse{}
	mi := &file_api_v1_rbac_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRoleBindingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoleBindingsResponse) ProtoMessage() {}

func (x *ListRoleBindingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rbac_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoleBindingsResponse.ProtoReflect.Descriptor instead.
func (*ListRoleBindingsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_rbac_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListRoleBindingsResponse) GetBindings() []*storage.K8SRoleBinding {
	if x != nil {
		return x.Bindings
	}
	return nil
}

type GetRoleBindingResponse struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Binding       *storage.K8SRoleBinding `protobuf:"bytes,1,opt,name=binding,proto3" json:"binding,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRoleBindingResponse) Reset() {
	*x = GetRoleBindingResponse{}
	mi := &file_api_v1_rbac_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRoleBindingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleBindingResponse) ProtoMessage() {}

func (x *GetRoleBindingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rbac_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleBindingResponse.ProtoReflect.Descriptor instead.
func (*GetRoleBindingResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_rbac_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetRoleBindingResponse) GetBinding() *storage.K8SRoleBinding {
	if x != nil {
		return x.Binding
	}
	return nil
}

// A list of k8s subjects (users and groups only, for service accounts, try the service account service)
// Next Tag: 2
type ListSubjectsResponse struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	SubjectAndRoles []*SubjectAndRoles     `protobuf:"bytes,1,rep,name=subject_and_roles,json=subjectAndRoles,proto3" json:"subject_and_roles,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ListSubjectsResponse) Reset() {
	*x = ListSubjectsResponse{}
	mi := &file_api_v1_rbac_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSubjectsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubjectsResponse) ProtoMessage() {}

func (x *ListSubjectsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rbac_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubjectsResponse.ProtoReflect.Descriptor instead.
func (*ListSubjectsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_rbac_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListSubjectsResponse) GetSubjectAndRoles() []*SubjectAndRoles {
	if x != nil {
		return x.SubjectAndRoles
	}
	return nil
}

type SubjectAndRoles struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Subject       *storage.Subject       `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Roles         []*storage.K8SRole     `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubjectAndRoles) Reset() {
	*x = SubjectAndRoles{}
	mi := &file_api_v1_rbac_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubjectAndRoles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubjectAndRoles) ProtoMessage() {}

func (x *SubjectAndRoles) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rbac_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubjectAndRoles.ProtoReflect.Descriptor instead.
func (*SubjectAndRoles) Descriptor() ([]byte, []int) {
	return file_api_v1_rbac_service_proto_rawDescGZIP(), []int{5}
}

func (x *SubjectAndRoles) GetSubject() *storage.Subject {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *SubjectAndRoles) GetRoles() []*storage.K8SRole {
	if x != nil {
		return x.Roles
	}
	return nil
}

type GetSubjectResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Subject       *storage.Subject       `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	ClusterRoles  []*storage.K8SRole     `protobuf:"bytes,2,rep,name=cluster_roles,json=clusterRoles,proto3" json:"cluster_roles,omitempty"`
	ScopedRoles   []*ScopedRoles         `protobuf:"bytes,3,rep,name=scoped_roles,json=scopedRoles,proto3" json:"scoped_roles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSubjectResponse) Reset() {
	*x = GetSubjectResponse{}
	mi := &file_api_v1_rbac_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSubjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubjectResponse) ProtoMessage() {}

func (x *GetSubjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rbac_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubjectResponse.ProtoReflect.Descriptor instead.
func (*GetSubjectResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_rbac_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetSubjectResponse) GetSubject() *storage.Subject {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *GetSubjectResponse) GetClusterRoles() []*storage.K8SRole {
	if x != nil {
		return x.ClusterRoles
	}
	return nil
}

func (x *GetSubjectResponse) GetScopedRoles() []*ScopedRoles {
	if x != nil {
		return x.ScopedRoles
	}
	return nil
}

type ScopedRoles struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Namespace     string                 `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Roles         []*storage.K8SRole     `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ScopedRoles) Reset() {
	*x = ScopedRoles{}
	mi := &file_api_v1_rbac_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScopedRoles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScopedRoles) ProtoMessage() {}

func (x *ScopedRoles) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rbac_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScopedRoles.ProtoReflect.Descriptor instead.
func (*ScopedRoles) Descriptor() ([]byte, []int) {
	return file_api_v1_rbac_service_proto_rawDescGZIP(), []int{7}
}

func (x *ScopedRoles) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ScopedRoles) GetRoles() []*storage.K8SRole {
	if x != nil {
		return x.Roles
	}
	return nil
}

var File_api_v1_rbac_service_proto protoreflect.FileDescriptor

const file_api_v1_rbac_service_proto_rawDesc = "" +
	"\n" +
	"\x19api/v1/rbac_service.proto\x12\x02v1\x1a\x13api/v1/common.proto\x1a\x1bapi/v1/search_service.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x12storage/rbac.proto\";\n" +
	"\x11ListRolesResponse\x12&\n" +
	"\x05roles\x18\x01 \x03(\v2\x10.storage.K8sRoleR\x05roles\"7\n" +
	"\x0fGetRoleResponse\x12$\n" +
	"\x04role\x18\x01 \x01(\v2\x10.storage.K8sRoleR\x04role\"O\n" +
	"\x18ListRoleBindingsResponse\x123\n" +
	"\bbindings\x18\x01 \x03(\v2\x17.storage.K8sRoleBindingR\bbindings\"K\n" +
	"\x16GetRoleBindingResponse\x121\n" +
	"\abinding\x18\x01 \x01(\v2\x17.storage.K8sRoleBindingR\abinding\"W\n" +
	"\x14ListSubjectsResponse\x12?\n" +
	"\x11subject_and_roles\x18\x01 \x03(\v2\x13.v1.SubjectAndRolesR\x0fsubjectAndRoles\"e\n" +
	"\x0fSubjectAndRoles\x12*\n" +
	"\asubject\x18\x01 \x01(\v2\x10.storage.SubjectR\asubject\x12&\n" +
	"\x05roles\x18\x02 \x03(\v2\x10.storage.K8sRoleR\x05roles\"\xab\x01\n" +
	"\x12GetSubjectResponse\x12*\n" +
	"\asubject\x18\x01 \x01(\v2\x10.storage.SubjectR\asubject\x125\n" +
	"\rcluster_roles\x18\x02 \x03(\v2\x10.storage.K8sRoleR\fclusterRoles\x122\n" +
	"\fscoped_roles\x18\x03 \x03(\v2\x0f.v1.ScopedRolesR\vscopedRoles\"S\n" +
	"\vScopedRoles\x12\x1c\n" +
	"\tnamespace\x18\x01 \x01(\tR\tnamespace\x12&\n" +
	"\x05roles\x18\x02 \x03(\v2\x10.storage.K8sRoleR\x05roles2\x8b\x04\n" +
	"\vRbacService\x12M\n" +
	"\aGetRole\x12\x10.v1.ResourceByID\x1a\x13.v1.GetRoleResponse\"\x1b\x82\xd3\xe4\x93\x02\x15\x12\x13/v1/rbac/roles/{id}\x12H\n" +
	"\tListRoles\x12\f.v1.RawQuery\x1a\x15.v1.ListRolesResponse\"\x16\x82\xd3\xe4\x93\x02\x10\x12\x0e/v1/rbac/roles\x12^\n" +
	"\x0eGetRoleBinding\x12\x10.v1.ResourceByID\x1a\x1a.v1.GetRoleBindingResponse\"\x1e\x82\xd3\xe4\x93\x02\x18\x12\x16/v1/rbac/bindings/{id}\x12Y\n" +
	"\x10ListRoleBindings\x12\f.v1.RawQuery\x1a\x1c.v1.ListRoleBindingsResponse\"\x19\x82\xd3\xe4\x93\x02\x13\x12\x11/v1/rbac/bindings\x12U\n" +
	"\n" +
	"GetSubject\x12\x10.v1.ResourceByID\x1a\x16.v1.GetSubjectResponse\"\x1d\x82\xd3\xe4\x93\x02\x17\x12\x15/v1/rbac/subject/{id}\x12Q\n" +
	"\fListSubjects\x12\f.v1.RawQuery\x1a\x18.v1.ListSubjectsResponse\"\x19\x82\xd3\xe4\x93\x02\x13\x12\x11/v1/rbac/subjectsB'\n" +
	"\x18io.stackrox.proto.api.v1Z\v./api/v1;v1X\x02b\x06proto3"

var (
	file_api_v1_rbac_service_proto_rawDescOnce sync.Once
	file_api_v1_rbac_service_proto_rawDescData []byte
)

func file_api_v1_rbac_service_proto_rawDescGZIP() []byte {
	file_api_v1_rbac_service_proto_rawDescOnce.Do(func() {
		file_api_v1_rbac_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1_rbac_service_proto_rawDesc), len(file_api_v1_rbac_service_proto_rawDesc)))
	})
	return file_api_v1_rbac_service_proto_rawDescData
}

var file_api_v1_rbac_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_v1_rbac_service_proto_goTypes = []any{
	(*ListRolesResponse)(nil),        // 0: v1.ListRolesResponse
	(*GetRoleResponse)(nil),          // 1: v1.GetRoleResponse
	(*ListRoleBindingsResponse)(nil), // 2: v1.ListRoleBindingsResponse
	(*GetRoleBindingResponse)(nil),   // 3: v1.GetRoleBindingResponse
	(*ListSubjectsResponse)(nil),     // 4: v1.ListSubjectsResponse
	(*SubjectAndRoles)(nil),          // 5: v1.SubjectAndRoles
	(*GetSubjectResponse)(nil),       // 6: v1.GetSubjectResponse
	(*ScopedRoles)(nil),              // 7: v1.ScopedRoles
	(*storage.K8SRole)(nil),          // 8: storage.K8sRole
	(*storage.K8SRoleBinding)(nil),   // 9: storage.K8sRoleBinding
	(*storage.Subject)(nil),          // 10: storage.Subject
	(*ResourceByID)(nil),             // 11: v1.ResourceByID
	(*RawQuery)(nil),                 // 12: v1.RawQuery
}
var file_api_v1_rbac_service_proto_depIdxs = []int32{
	8,  // 0: v1.ListRolesResponse.roles:type_name -> storage.K8sRole
	8,  // 1: v1.GetRoleResponse.role:type_name -> storage.K8sRole
	9,  // 2: v1.ListRoleBindingsResponse.bindings:type_name -> storage.K8sRoleBinding
	9,  // 3: v1.GetRoleBindingResponse.binding:type_name -> storage.K8sRoleBinding
	5,  // 4: v1.ListSubjectsResponse.subject_and_roles:type_name -> v1.SubjectAndRoles
	10, // 5: v1.SubjectAndRoles.subject:type_name -> storage.Subject
	8,  // 6: v1.SubjectAndRoles.roles:type_name -> storage.K8sRole
	10, // 7: v1.GetSubjectResponse.subject:type_name -> storage.Subject
	8,  // 8: v1.GetSubjectResponse.cluster_roles:type_name -> storage.K8sRole
	7,  // 9: v1.GetSubjectResponse.scoped_roles:type_name -> v1.ScopedRoles
	8,  // 10: v1.ScopedRoles.roles:type_name -> storage.K8sRole
	11, // 11: v1.RbacService.GetRole:input_type -> v1.ResourceByID
	12, // 12: v1.RbacService.ListRoles:input_type -> v1.RawQuery
	11, // 13: v1.RbacService.GetRoleBinding:input_type -> v1.ResourceByID
	12, // 14: v1.RbacService.ListRoleBindings:input_type -> v1.RawQuery
	11, // 15: v1.RbacService.GetSubject:input_type -> v1.ResourceByID
	12, // 16: v1.RbacService.ListSubjects:input_type -> v1.RawQuery
	1,  // 17: v1.RbacService.GetRole:output_type -> v1.GetRoleResponse
	0,  // 18: v1.RbacService.ListRoles:output_type -> v1.ListRolesResponse
	3,  // 19: v1.RbacService.GetRoleBinding:output_type -> v1.GetRoleBindingResponse
	2,  // 20: v1.RbacService.ListRoleBindings:output_type -> v1.ListRoleBindingsResponse
	6,  // 21: v1.RbacService.GetSubject:output_type -> v1.GetSubjectResponse
	4,  // 22: v1.RbacService.ListSubjects:output_type -> v1.ListSubjectsResponse
	17, // [17:23] is the sub-list for method output_type
	11, // [11:17] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_api_v1_rbac_service_proto_init() }
func file_api_v1_rbac_service_proto_init() {
	if File_api_v1_rbac_service_proto != nil {
		return
	}
	file_api_v1_common_proto_init()
	file_api_v1_search_service_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1_rbac_service_proto_rawDesc), len(file_api_v1_rbac_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_rbac_service_proto_goTypes,
		DependencyIndexes: file_api_v1_rbac_service_proto_depIdxs,
		MessageInfos:      file_api_v1_rbac_service_proto_msgTypes,
	}.Build()
	File_api_v1_rbac_service_proto = out.File
	file_api_v1_rbac_service_proto_goTypes = nil
	file_api_v1_rbac_service_proto_depIdxs = nil
}
