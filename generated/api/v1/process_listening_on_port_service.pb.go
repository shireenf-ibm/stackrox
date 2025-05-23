// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.3
// source: api/v1/process_listening_on_port_service.proto

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

type GetProcessesListeningOnPortsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DeploymentId  string                 `protobuf:"bytes,1,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProcessesListeningOnPortsRequest) Reset() {
	*x = GetProcessesListeningOnPortsRequest{}
	mi := &file_api_v1_process_listening_on_port_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProcessesListeningOnPortsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProcessesListeningOnPortsRequest) ProtoMessage() {}

func (x *GetProcessesListeningOnPortsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_process_listening_on_port_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProcessesListeningOnPortsRequest.ProtoReflect.Descriptor instead.
func (*GetProcessesListeningOnPortsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_process_listening_on_port_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetProcessesListeningOnPortsRequest) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

type GetProcessesListeningOnPortsResponse struct {
	state              protoimpl.MessageState            `protogen:"open.v1"`
	ListeningEndpoints []*storage.ProcessListeningOnPort `protobuf:"bytes,1,rep,name=listening_endpoints,json=listeningEndpoints,proto3" json:"listening_endpoints,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *GetProcessesListeningOnPortsResponse) Reset() {
	*x = GetProcessesListeningOnPortsResponse{}
	mi := &file_api_v1_process_listening_on_port_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProcessesListeningOnPortsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProcessesListeningOnPortsResponse) ProtoMessage() {}

func (x *GetProcessesListeningOnPortsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_process_listening_on_port_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProcessesListeningOnPortsResponse.ProtoReflect.Descriptor instead.
func (*GetProcessesListeningOnPortsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_process_listening_on_port_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetProcessesListeningOnPortsResponse) GetListeningEndpoints() []*storage.ProcessListeningOnPort {
	if x != nil {
		return x.ListeningEndpoints
	}
	return nil
}

var File_api_v1_process_listening_on_port_service_proto protoreflect.FileDescriptor

const file_api_v1_process_listening_on_port_service_proto_rawDesc = "" +
	"\n" +
	".api/v1/process_listening_on_port_service.proto\x12\x02v1\x1a\x1cgoogle/api/annotations.proto\x1a'storage/process_listening_on_port.proto\"J\n" +
	"#GetProcessesListeningOnPortsRequest\x12#\n" +
	"\rdeployment_id\x18\x01 \x01(\tR\fdeploymentId\"x\n" +
	"$GetProcessesListeningOnPortsResponse\x12P\n" +
	"\x13listening_endpoints\x18\x01 \x03(\v2\x1f.storage.ProcessListeningOnPortR\x12listeningEndpoints2\xc4\x01\n" +
	"\x19ListeningEndpointsService\x12\xa6\x01\n" +
	"\x15GetListeningEndpoints\x12'.v1.GetProcessesListeningOnPortsRequest\x1a(.v1.GetProcessesListeningOnPortsResponse\":\x82\xd3\xe4\x93\x024\x122/v1/listening_endpoints/deployment/{deployment_id}B'\n" +
	"\x18io.stackrox.proto.api.v1Z\v./api/v1;v1X\x00b\x06proto3"

var (
	file_api_v1_process_listening_on_port_service_proto_rawDescOnce sync.Once
	file_api_v1_process_listening_on_port_service_proto_rawDescData []byte
)

func file_api_v1_process_listening_on_port_service_proto_rawDescGZIP() []byte {
	file_api_v1_process_listening_on_port_service_proto_rawDescOnce.Do(func() {
		file_api_v1_process_listening_on_port_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1_process_listening_on_port_service_proto_rawDesc), len(file_api_v1_process_listening_on_port_service_proto_rawDesc)))
	})
	return file_api_v1_process_listening_on_port_service_proto_rawDescData
}

var file_api_v1_process_listening_on_port_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_v1_process_listening_on_port_service_proto_goTypes = []any{
	(*GetProcessesListeningOnPortsRequest)(nil),  // 0: v1.GetProcessesListeningOnPortsRequest
	(*GetProcessesListeningOnPortsResponse)(nil), // 1: v1.GetProcessesListeningOnPortsResponse
	(*storage.ProcessListeningOnPort)(nil),       // 2: storage.ProcessListeningOnPort
}
var file_api_v1_process_listening_on_port_service_proto_depIdxs = []int32{
	2, // 0: v1.GetProcessesListeningOnPortsResponse.listening_endpoints:type_name -> storage.ProcessListeningOnPort
	0, // 1: v1.ListeningEndpointsService.GetListeningEndpoints:input_type -> v1.GetProcessesListeningOnPortsRequest
	1, // 2: v1.ListeningEndpointsService.GetListeningEndpoints:output_type -> v1.GetProcessesListeningOnPortsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_v1_process_listening_on_port_service_proto_init() }
func file_api_v1_process_listening_on_port_service_proto_init() {
	if File_api_v1_process_listening_on_port_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1_process_listening_on_port_service_proto_rawDesc), len(file_api_v1_process_listening_on_port_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_process_listening_on_port_service_proto_goTypes,
		DependencyIndexes: file_api_v1_process_listening_on_port_service_proto_depIdxs,
		MessageInfos:      file_api_v1_process_listening_on_port_service_proto_msgTypes,
	}.Build()
	File_api_v1_process_listening_on_port_service_proto = out.File
	file_api_v1_process_listening_on_port_service_proto_goTypes = nil
	file_api_v1_process_listening_on_port_service_proto_depIdxs = nil
}
