// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: internalapi/common/boolean-operator.proto

package common

import (
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

type BooleanOperator int32

const (
	BooleanOperator_OR  BooleanOperator = 0
	BooleanOperator_AND BooleanOperator = 1
)

// Enum value maps for BooleanOperator.
var (
	BooleanOperator_name = map[int32]string{
		0: "OR",
		1: "AND",
	}
	BooleanOperator_value = map[string]int32{
		"OR":  0,
		"AND": 1,
	}
)

func (x BooleanOperator) Enum() *BooleanOperator {
	p := new(BooleanOperator)
	*p = x
	return p
}

func (x BooleanOperator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BooleanOperator) Descriptor() protoreflect.EnumDescriptor {
	return file_internalapi_common_boolean_operator_proto_enumTypes[0].Descriptor()
}

func (BooleanOperator) Type() protoreflect.EnumType {
	return &file_internalapi_common_boolean_operator_proto_enumTypes[0]
}

func (x BooleanOperator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BooleanOperator.Descriptor instead.
func (BooleanOperator) EnumDescriptor() ([]byte, []int) {
	return file_internalapi_common_boolean_operator_proto_rawDescGZIP(), []int{0}
}

var File_internalapi_common_boolean_operator_proto protoreflect.FileDescriptor

var file_internalapi_common_boolean_operator_proto_rawDesc = []byte{
	0x0a, 0x29, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2a, 0x22, 0x0a, 0x0f, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x07,
	0x0a, 0x03, 0x41, 0x4e, 0x44, 0x10, 0x01, 0x42, 0x1d, 0x5a, 0x1b, 0x2e, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internalapi_common_boolean_operator_proto_rawDescOnce sync.Once
	file_internalapi_common_boolean_operator_proto_rawDescData = file_internalapi_common_boolean_operator_proto_rawDesc
)

func file_internalapi_common_boolean_operator_proto_rawDescGZIP() []byte {
	file_internalapi_common_boolean_operator_proto_rawDescOnce.Do(func() {
		file_internalapi_common_boolean_operator_proto_rawDescData = protoimpl.X.CompressGZIP(file_internalapi_common_boolean_operator_proto_rawDescData)
	})
	return file_internalapi_common_boolean_operator_proto_rawDescData
}

var file_internalapi_common_boolean_operator_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internalapi_common_boolean_operator_proto_goTypes = []any{
	(BooleanOperator)(0), // 0: common.BooleanOperator
}
var file_internalapi_common_boolean_operator_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internalapi_common_boolean_operator_proto_init() }
func file_internalapi_common_boolean_operator_proto_init() {
	if File_internalapi_common_boolean_operator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internalapi_common_boolean_operator_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internalapi_common_boolean_operator_proto_goTypes,
		DependencyIndexes: file_internalapi_common_boolean_operator_proto_depIdxs,
		EnumInfos:         file_internalapi_common_boolean_operator_proto_enumTypes,
	}.Build()
	File_internalapi_common_boolean_operator_proto = out.File
	file_internalapi_common_boolean_operator_proto_rawDesc = nil
	file_internalapi_common_boolean_operator_proto_goTypes = nil
	file_internalapi_common_boolean_operator_proto_depIdxs = nil
}
