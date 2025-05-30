// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.3
// source: api/v1/cve_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

type SuppressCVERequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// These are (NVD) vulnerability identifiers, `cve` field of `storage.CVE`, and *not* the `id` field.
	// For example, CVE-2021-44832.
	Cves []string `protobuf:"bytes,1,rep,name=cves,proto3" json:"cves,omitempty"`
	// In JSON format, the Duration type is encoded as a string rather than an object,
	// where the string ends in the suffix "s" (indicating seconds) and is preceded by the number of seconds,
	// with nanoseconds expressed as fractional seconds.
	// For example, 3 seconds with 0 nanoseconds should be encoded in JSON format as "3s",
	// while 3 seconds and 1 nanosecond should be expressed in JSON format as "3.000000001s",
	// and 3 seconds and 1 microsecond should be expressed in JSON format as "3.000001s".
	Duration      *durationpb.Duration `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SuppressCVERequest) Reset() {
	*x = SuppressCVERequest{}
	mi := &file_api_v1_cve_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SuppressCVERequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuppressCVERequest) ProtoMessage() {}

func (x *SuppressCVERequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cve_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuppressCVERequest.ProtoReflect.Descriptor instead.
func (*SuppressCVERequest) Descriptor() ([]byte, []int) {
	return file_api_v1_cve_service_proto_rawDescGZIP(), []int{0}
}

func (x *SuppressCVERequest) GetCves() []string {
	if x != nil {
		return x.Cves
	}
	return nil
}

func (x *SuppressCVERequest) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

type UnsuppressCVERequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// These are (NVD) vulnerability identifiers, `cve` field of `storage.CVE`, and *not* the `id` field.
	// For example, CVE-2021-44832.
	Cves          []string `protobuf:"bytes,1,rep,name=cves,proto3" json:"cves,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UnsuppressCVERequest) Reset() {
	*x = UnsuppressCVERequest{}
	mi := &file_api_v1_cve_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnsuppressCVERequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsuppressCVERequest) ProtoMessage() {}

func (x *UnsuppressCVERequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cve_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsuppressCVERequest.ProtoReflect.Descriptor instead.
func (*UnsuppressCVERequest) Descriptor() ([]byte, []int) {
	return file_api_v1_cve_service_proto_rawDescGZIP(), []int{1}
}

func (x *UnsuppressCVERequest) GetCves() []string {
	if x != nil {
		return x.Cves
	}
	return nil
}

var File_api_v1_cve_service_proto protoreflect.FileDescriptor

const file_api_v1_cve_service_proto_rawDesc = "" +
	"\n" +
	"\x18api/v1/cve_service.proto\x12\x02v1\x1a\x12api/v1/empty.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1egoogle/protobuf/duration.proto\"e\n" +
	"\x12SuppressCVERequest\x12\x12\n" +
	"\x04cves\x18\x01 \x03(\tR\x04cves\x125\n" +
	"\bduration\x18\x03 \x01(\v2\x19.google.protobuf.DurationR\bdurationJ\x04\b\x02\x10\x03\"*\n" +
	"\x14UnsuppressCVERequest\x12\x12\n" +
	"\x04cves\x18\x01 \x03(\tR\x04cves2\xc3\x01\n" +
	"\x0fImageCVEService\x12T\n" +
	"\fSuppressCVEs\x12\x16.v1.SuppressCVERequest\x1a\t.v1.Empty\"!\x82\xd3\xe4\x93\x02\x1b:\x01*2\x16/v1/imagecves/suppress\x12Z\n" +
	"\x0eUnsuppressCVEs\x12\x18.v1.UnsuppressCVERequest\x1a\t.v1.Empty\"#\x82\xd3\xe4\x93\x02\x1d:\x01*2\x18/v1/imagecves/unsuppress2\xc0\x01\n" +
	"\x0eNodeCVEService\x12S\n" +
	"\fSuppressCVEs\x12\x16.v1.SuppressCVERequest\x1a\t.v1.Empty\" \x82\xd3\xe4\x93\x02\x1a:\x01*2\x15/v1/nodecves/suppress\x12Y\n" +
	"\x0eUnsuppressCVEs\x12\x18.v1.UnsuppressCVERequest\x1a\t.v1.Empty\"\"\x82\xd3\xe4\x93\x02\x1c:\x01*2\x17/v1/nodecves/unsuppress2\xc9\x01\n" +
	"\x11ClusterCVEService\x12V\n" +
	"\fSuppressCVEs\x12\x16.v1.SuppressCVERequest\x1a\t.v1.Empty\"#\x82\xd3\xe4\x93\x02\x1d:\x01*2\x18/v1/clustercves/suppress\x12\\\n" +
	"\x0eUnsuppressCVEs\x12\x18.v1.UnsuppressCVERequest\x1a\t.v1.Empty\"%\x82\xd3\xe4\x93\x02\x1f:\x01*2\x1a/v1/clustercves/unsuppressB'\n" +
	"\x18io.stackrox.proto.api.v1Z\v./api/v1;v1X\x01b\x06proto3"

var (
	file_api_v1_cve_service_proto_rawDescOnce sync.Once
	file_api_v1_cve_service_proto_rawDescData []byte
)

func file_api_v1_cve_service_proto_rawDescGZIP() []byte {
	file_api_v1_cve_service_proto_rawDescOnce.Do(func() {
		file_api_v1_cve_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1_cve_service_proto_rawDesc), len(file_api_v1_cve_service_proto_rawDesc)))
	})
	return file_api_v1_cve_service_proto_rawDescData
}

var file_api_v1_cve_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_v1_cve_service_proto_goTypes = []any{
	(*SuppressCVERequest)(nil),   // 0: v1.SuppressCVERequest
	(*UnsuppressCVERequest)(nil), // 1: v1.UnsuppressCVERequest
	(*durationpb.Duration)(nil),  // 2: google.protobuf.Duration
	(*Empty)(nil),                // 3: v1.Empty
}
var file_api_v1_cve_service_proto_depIdxs = []int32{
	2, // 0: v1.SuppressCVERequest.duration:type_name -> google.protobuf.Duration
	0, // 1: v1.ImageCVEService.SuppressCVEs:input_type -> v1.SuppressCVERequest
	1, // 2: v1.ImageCVEService.UnsuppressCVEs:input_type -> v1.UnsuppressCVERequest
	0, // 3: v1.NodeCVEService.SuppressCVEs:input_type -> v1.SuppressCVERequest
	1, // 4: v1.NodeCVEService.UnsuppressCVEs:input_type -> v1.UnsuppressCVERequest
	0, // 5: v1.ClusterCVEService.SuppressCVEs:input_type -> v1.SuppressCVERequest
	1, // 6: v1.ClusterCVEService.UnsuppressCVEs:input_type -> v1.UnsuppressCVERequest
	3, // 7: v1.ImageCVEService.SuppressCVEs:output_type -> v1.Empty
	3, // 8: v1.ImageCVEService.UnsuppressCVEs:output_type -> v1.Empty
	3, // 9: v1.NodeCVEService.SuppressCVEs:output_type -> v1.Empty
	3, // 10: v1.NodeCVEService.UnsuppressCVEs:output_type -> v1.Empty
	3, // 11: v1.ClusterCVEService.SuppressCVEs:output_type -> v1.Empty
	3, // 12: v1.ClusterCVEService.UnsuppressCVEs:output_type -> v1.Empty
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_v1_cve_service_proto_init() }
func file_api_v1_cve_service_proto_init() {
	if File_api_v1_cve_service_proto != nil {
		return
	}
	file_api_v1_empty_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1_cve_service_proto_rawDesc), len(file_api_v1_cve_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_api_v1_cve_service_proto_goTypes,
		DependencyIndexes: file_api_v1_cve_service_proto_depIdxs,
		MessageInfos:      file_api_v1_cve_service_proto_msgTypes,
	}.Build()
	File_api_v1_cve_service_proto = out.File
	file_api_v1_cve_service_proto_goTypes = nil
	file_api_v1_cve_service_proto_depIdxs = nil
}
