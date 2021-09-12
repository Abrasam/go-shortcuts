// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/appengine/v1beta/network_settings.proto

package appengine

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// If unspecified, INGRESS_TRAFFIC_ALLOWED_ALL will be used.
type NetworkSettings_IngressTrafficAllowed int32

const (
	// Unspecified
	NetworkSettings_INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED NetworkSettings_IngressTrafficAllowed = 0
	// Allow HTTP traffic from public and private sources.
	NetworkSettings_INGRESS_TRAFFIC_ALLOWED_ALL NetworkSettings_IngressTrafficAllowed = 1
	// Allow HTTP traffic from only private VPC sources.
	NetworkSettings_INGRESS_TRAFFIC_ALLOWED_INTERNAL_ONLY NetworkSettings_IngressTrafficAllowed = 2
	// Allow HTTP traffic from private VPC sources and through load balancers.
	NetworkSettings_INGRESS_TRAFFIC_ALLOWED_INTERNAL_AND_LB NetworkSettings_IngressTrafficAllowed = 3
)

// Enum value maps for NetworkSettings_IngressTrafficAllowed.
var (
	NetworkSettings_IngressTrafficAllowed_name = map[int32]string{
		0: "INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED",
		1: "INGRESS_TRAFFIC_ALLOWED_ALL",
		2: "INGRESS_TRAFFIC_ALLOWED_INTERNAL_ONLY",
		3: "INGRESS_TRAFFIC_ALLOWED_INTERNAL_AND_LB",
	}
	NetworkSettings_IngressTrafficAllowed_value = map[string]int32{
		"INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED":     0,
		"INGRESS_TRAFFIC_ALLOWED_ALL":             1,
		"INGRESS_TRAFFIC_ALLOWED_INTERNAL_ONLY":   2,
		"INGRESS_TRAFFIC_ALLOWED_INTERNAL_AND_LB": 3,
	}
)

func (x NetworkSettings_IngressTrafficAllowed) Enum() *NetworkSettings_IngressTrafficAllowed {
	p := new(NetworkSettings_IngressTrafficAllowed)
	*p = x
	return p
}

func (x NetworkSettings_IngressTrafficAllowed) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NetworkSettings_IngressTrafficAllowed) Descriptor() protoreflect.EnumDescriptor {
	return file_google_appengine_v1beta_network_settings_proto_enumTypes[0].Descriptor()
}

func (NetworkSettings_IngressTrafficAllowed) Type() protoreflect.EnumType {
	return &file_google_appengine_v1beta_network_settings_proto_enumTypes[0]
}

func (x NetworkSettings_IngressTrafficAllowed) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NetworkSettings_IngressTrafficAllowed.Descriptor instead.
func (NetworkSettings_IngressTrafficAllowed) EnumDescriptor() ([]byte, []int) {
	return file_google_appengine_v1beta_network_settings_proto_rawDescGZIP(), []int{0, 0}
}

// A NetworkSettings resource is a container for ingress settings for a version
// or service.
type NetworkSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ingress settings for version or service.
	IngressTrafficAllowed NetworkSettings_IngressTrafficAllowed `protobuf:"varint,1,opt,name=ingress_traffic_allowed,json=ingressTrafficAllowed,proto3,enum=google.appengine.v1beta.NetworkSettings_IngressTrafficAllowed" json:"ingress_traffic_allowed,omitempty"`
}

func (x *NetworkSettings) Reset() {
	*x = NetworkSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_v1beta_network_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkSettings) ProtoMessage() {}

func (x *NetworkSettings) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_v1beta_network_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkSettings.ProtoReflect.Descriptor instead.
func (*NetworkSettings) Descriptor() ([]byte, []int) {
	return file_google_appengine_v1beta_network_settings_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkSettings) GetIngressTrafficAllowed() NetworkSettings_IngressTrafficAllowed {
	if x != nil {
		return x.IngressTrafficAllowed
	}
	return NetworkSettings_INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED
}

var File_google_appengine_v1beta_network_settings_proto protoreflect.FileDescriptor

var file_google_appengine_v1beta_network_settings_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x02, 0x0a, 0x0f, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x76, 0x0a, 0x17, 0x69,
	0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x54, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x52, 0x15, 0x69, 0x6e,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x41, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x64, 0x22, 0xb9, 0x01, 0x0a, 0x15, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x54,
	0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x12, 0x27, 0x0a,
	0x23, 0x49, 0x4e, 0x47, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49, 0x43,
	0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x49, 0x4e, 0x47, 0x52, 0x45, 0x53,
	0x53, 0x5f, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49, 0x43, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45,
	0x44, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x29, 0x0a, 0x25, 0x49, 0x4e, 0x47, 0x52, 0x45,
	0x53, 0x53, 0x5f, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49, 0x43, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57,
	0x45, 0x44, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x4f, 0x4e, 0x4c, 0x59,
	0x10, 0x02, 0x12, 0x2b, 0x0a, 0x27, 0x49, 0x4e, 0x47, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x52,
	0x41, 0x46, 0x46, 0x49, 0x43, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x5f, 0x49, 0x4e,
	0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x4c, 0x42, 0x10, 0x03, 0x42,
	0xda, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x42,
	0x14, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x3b,
	0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0xaa, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0xca, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0xea, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x70, 0x70, 0x45, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_appengine_v1beta_network_settings_proto_rawDescOnce sync.Once
	file_google_appengine_v1beta_network_settings_proto_rawDescData = file_google_appengine_v1beta_network_settings_proto_rawDesc
)

func file_google_appengine_v1beta_network_settings_proto_rawDescGZIP() []byte {
	file_google_appengine_v1beta_network_settings_proto_rawDescOnce.Do(func() {
		file_google_appengine_v1beta_network_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_appengine_v1beta_network_settings_proto_rawDescData)
	})
	return file_google_appengine_v1beta_network_settings_proto_rawDescData
}

var file_google_appengine_v1beta_network_settings_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_appengine_v1beta_network_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_appengine_v1beta_network_settings_proto_goTypes = []interface{}{
	(NetworkSettings_IngressTrafficAllowed)(0), // 0: google.appengine.v1beta.NetworkSettings.IngressTrafficAllowed
	(*NetworkSettings)(nil),                    // 1: google.appengine.v1beta.NetworkSettings
}
var file_google_appengine_v1beta_network_settings_proto_depIdxs = []int32{
	0, // 0: google.appengine.v1beta.NetworkSettings.ingress_traffic_allowed:type_name -> google.appengine.v1beta.NetworkSettings.IngressTrafficAllowed
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_appengine_v1beta_network_settings_proto_init() }
func file_google_appengine_v1beta_network_settings_proto_init() {
	if File_google_appengine_v1beta_network_settings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_appengine_v1beta_network_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkSettings); i {
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
			RawDescriptor: file_google_appengine_v1beta_network_settings_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_appengine_v1beta_network_settings_proto_goTypes,
		DependencyIndexes: file_google_appengine_v1beta_network_settings_proto_depIdxs,
		EnumInfos:         file_google_appengine_v1beta_network_settings_proto_enumTypes,
		MessageInfos:      file_google_appengine_v1beta_network_settings_proto_msgTypes,
	}.Build()
	File_google_appengine_v1beta_network_settings_proto = out.File
	file_google_appengine_v1beta_network_settings_proto_rawDesc = nil
	file_google_appengine_v1beta_network_settings_proto_goTypes = nil
	file_google_appengine_v1beta_network_settings_proto_depIdxs = nil
}
