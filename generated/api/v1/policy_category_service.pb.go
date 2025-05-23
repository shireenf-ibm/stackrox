// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.3
// source: api/v1/policy_category_service.proto

package v1

import (
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

type PolicyCategory struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsDefault     bool                   `protobuf:"varint,3,opt,name=isDefault,proto3" json:"isDefault,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PolicyCategory) Reset() {
	*x = PolicyCategory{}
	mi := &file_api_v1_policy_category_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PolicyCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyCategory) ProtoMessage() {}

func (x *PolicyCategory) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_policy_category_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyCategory.ProtoReflect.Descriptor instead.
func (*PolicyCategory) Descriptor() ([]byte, []int) {
	return file_api_v1_policy_category_service_proto_rawDescGZIP(), []int{0}
}

func (x *PolicyCategory) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PolicyCategory) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PolicyCategory) GetIsDefault() bool {
	if x != nil {
		return x.IsDefault
	}
	return false
}

type PostPolicyCategoryRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	PolicyCategory *PolicyCategory        `protobuf:"bytes,1,opt,name=policyCategory,proto3" json:"policyCategory,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *PostPolicyCategoryRequest) Reset() {
	*x = PostPolicyCategoryRequest{}
	mi := &file_api_v1_policy_category_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostPolicyCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostPolicyCategoryRequest) ProtoMessage() {}

func (x *PostPolicyCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_policy_category_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostPolicyCategoryRequest.ProtoReflect.Descriptor instead.
func (*PostPolicyCategoryRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_policy_category_service_proto_rawDescGZIP(), []int{1}
}

func (x *PostPolicyCategoryRequest) GetPolicyCategory() *PolicyCategory {
	if x != nil {
		return x.PolicyCategory
	}
	return nil
}

type GetPolicyCategoriesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Categories    []*PolicyCategory      `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPolicyCategoriesResponse) Reset() {
	*x = GetPolicyCategoriesResponse{}
	mi := &file_api_v1_policy_category_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPolicyCategoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPolicyCategoriesResponse) ProtoMessage() {}

func (x *GetPolicyCategoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_policy_category_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPolicyCategoriesResponse.ProtoReflect.Descriptor instead.
func (*GetPolicyCategoriesResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_policy_category_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetPolicyCategoriesResponse) GetCategories() []*PolicyCategory {
	if x != nil {
		return x.Categories
	}
	return nil
}

type RenamePolicyCategoryRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NewCategoryName string                 `protobuf:"bytes,2,opt,name=new_category_name,json=newCategoryName,proto3" json:"new_category_name,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *RenamePolicyCategoryRequest) Reset() {
	*x = RenamePolicyCategoryRequest{}
	mi := &file_api_v1_policy_category_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RenamePolicyCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenamePolicyCategoryRequest) ProtoMessage() {}

func (x *RenamePolicyCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_policy_category_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenamePolicyCategoryRequest.ProtoReflect.Descriptor instead.
func (*RenamePolicyCategoryRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_policy_category_service_proto_rawDescGZIP(), []int{3}
}

func (x *RenamePolicyCategoryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RenamePolicyCategoryRequest) GetNewCategoryName() string {
	if x != nil {
		return x.NewCategoryName
	}
	return ""
}

type DeletePolicyCategoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeletePolicyCategoryRequest) Reset() {
	*x = DeletePolicyCategoryRequest{}
	mi := &file_api_v1_policy_category_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePolicyCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePolicyCategoryRequest) ProtoMessage() {}

func (x *DeletePolicyCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_policy_category_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePolicyCategoryRequest.ProtoReflect.Descriptor instead.
func (*DeletePolicyCategoryRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_policy_category_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeletePolicyCategoryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_api_v1_policy_category_service_proto protoreflect.FileDescriptor

const file_api_v1_policy_category_service_proto_rawDesc = "" +
	"\n" +
	"$api/v1/policy_category_service.proto\x12\x02v1\x1a\x13api/v1/common.proto\x1a\x12api/v1/empty.proto\x1a\x1bapi/v1/search_service.proto\x1a\x1cgoogle/api/annotations.proto\"R\n" +
	"\x0ePolicyCategory\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x1c\n" +
	"\tisDefault\x18\x03 \x01(\bR\tisDefault\"W\n" +
	"\x19PostPolicyCategoryRequest\x12:\n" +
	"\x0epolicyCategory\x18\x01 \x01(\v2\x12.v1.PolicyCategoryR\x0epolicyCategory\"Q\n" +
	"\x1bGetPolicyCategoriesResponse\x122\n" +
	"\n" +
	"categories\x18\x01 \x03(\v2\x12.v1.PolicyCategoryR\n" +
	"categories\"Y\n" +
	"\x1bRenamePolicyCategoryRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12*\n" +
	"\x11new_category_name\x18\x02 \x01(\tR\x0fnewCategoryName\"-\n" +
	"\x1bDeletePolicyCategoryRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id2\xa5\x04\n" +
	"\x15PolicyCategoryService\x12\\\n" +
	"\x11GetPolicyCategory\x12\x10.v1.ResourceByID\x1a\x12.v1.PolicyCategory\"!\x82\xd3\xe4\x93\x02\x1b\x12\x19/v1/policycategories/{id}\x12b\n" +
	"\x13GetPolicyCategories\x12\f.v1.RawQuery\x1a\x1f.v1.GetPolicyCategoriesResponse\"\x1c\x82\xd3\xe4\x93\x02\x16\x12\x14/v1/policycategories\x12u\n" +
	"\x12PostPolicyCategory\x12\x1d.v1.PostPolicyCategoryRequest\x1a\x12.v1.PolicyCategory\",\x82\xd3\xe4\x93\x02&:\x0epolicyCategory\"\x14/v1/policycategories\x12l\n" +
	"\x14RenamePolicyCategory\x12\x1f.v1.RenamePolicyCategoryRequest\x1a\x12.v1.PolicyCategory\"\x1f\x82\xd3\xe4\x93\x02\x19:\x01*\x1a\x14/v1/policycategories\x12e\n" +
	"\x14DeletePolicyCategory\x12\x1f.v1.DeletePolicyCategoryRequest\x1a\t.v1.Empty\"!\x82\xd3\xe4\x93\x02\x1b*\x19/v1/policycategories/{id}B'\n" +
	"\x18io.stackrox.proto.api.v1Z\v./api/v1;v1X\x03b\x06proto3"

var (
	file_api_v1_policy_category_service_proto_rawDescOnce sync.Once
	file_api_v1_policy_category_service_proto_rawDescData []byte
)

func file_api_v1_policy_category_service_proto_rawDescGZIP() []byte {
	file_api_v1_policy_category_service_proto_rawDescOnce.Do(func() {
		file_api_v1_policy_category_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1_policy_category_service_proto_rawDesc), len(file_api_v1_policy_category_service_proto_rawDesc)))
	})
	return file_api_v1_policy_category_service_proto_rawDescData
}

var file_api_v1_policy_category_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v1_policy_category_service_proto_goTypes = []any{
	(*PolicyCategory)(nil),              // 0: v1.PolicyCategory
	(*PostPolicyCategoryRequest)(nil),   // 1: v1.PostPolicyCategoryRequest
	(*GetPolicyCategoriesResponse)(nil), // 2: v1.GetPolicyCategoriesResponse
	(*RenamePolicyCategoryRequest)(nil), // 3: v1.RenamePolicyCategoryRequest
	(*DeletePolicyCategoryRequest)(nil), // 4: v1.DeletePolicyCategoryRequest
	(*ResourceByID)(nil),                // 5: v1.ResourceByID
	(*RawQuery)(nil),                    // 6: v1.RawQuery
	(*Empty)(nil),                       // 7: v1.Empty
}
var file_api_v1_policy_category_service_proto_depIdxs = []int32{
	0, // 0: v1.PostPolicyCategoryRequest.policyCategory:type_name -> v1.PolicyCategory
	0, // 1: v1.GetPolicyCategoriesResponse.categories:type_name -> v1.PolicyCategory
	5, // 2: v1.PolicyCategoryService.GetPolicyCategory:input_type -> v1.ResourceByID
	6, // 3: v1.PolicyCategoryService.GetPolicyCategories:input_type -> v1.RawQuery
	1, // 4: v1.PolicyCategoryService.PostPolicyCategory:input_type -> v1.PostPolicyCategoryRequest
	3, // 5: v1.PolicyCategoryService.RenamePolicyCategory:input_type -> v1.RenamePolicyCategoryRequest
	4, // 6: v1.PolicyCategoryService.DeletePolicyCategory:input_type -> v1.DeletePolicyCategoryRequest
	0, // 7: v1.PolicyCategoryService.GetPolicyCategory:output_type -> v1.PolicyCategory
	2, // 8: v1.PolicyCategoryService.GetPolicyCategories:output_type -> v1.GetPolicyCategoriesResponse
	0, // 9: v1.PolicyCategoryService.PostPolicyCategory:output_type -> v1.PolicyCategory
	0, // 10: v1.PolicyCategoryService.RenamePolicyCategory:output_type -> v1.PolicyCategory
	7, // 11: v1.PolicyCategoryService.DeletePolicyCategory:output_type -> v1.Empty
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1_policy_category_service_proto_init() }
func file_api_v1_policy_category_service_proto_init() {
	if File_api_v1_policy_category_service_proto != nil {
		return
	}
	file_api_v1_common_proto_init()
	file_api_v1_empty_proto_init()
	file_api_v1_search_service_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1_policy_category_service_proto_rawDesc), len(file_api_v1_policy_category_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_policy_category_service_proto_goTypes,
		DependencyIndexes: file_api_v1_policy_category_service_proto_depIdxs,
		MessageInfos:      file_api_v1_policy_category_service_proto_msgTypes,
	}.Build()
	File_api_v1_policy_category_service_proto = out.File
	file_api_v1_policy_category_service_proto_goTypes = nil
	file_api_v1_policy_category_service_proto_depIdxs = nil
}
