// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: aks.proto

package retropb

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

type AksPing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AksPing) Reset() {
	*x = AksPing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AksPing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AksPing) ProtoMessage() {}

func (x *AksPing) ProtoReflect() protoreflect.Message {
	mi := &file_aks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AksPing.ProtoReflect.Descriptor instead.
func (*AksPing) Descriptor() ([]byte, []int) {
	return file_aks_proto_rawDescGZIP(), []int{0}
}

type AksHelloConnect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Salt string `protobuf:"bytes,1,opt,name=salt,proto3" json:"salt,omitempty"`
}

func (x *AksHelloConnect) Reset() {
	*x = AksHelloConnect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AksHelloConnect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AksHelloConnect) ProtoMessage() {}

func (x *AksHelloConnect) ProtoReflect() protoreflect.Message {
	mi := &file_aks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AksHelloConnect.ProtoReflect.Descriptor instead.
func (*AksHelloConnect) Descriptor() ([]byte, []int) {
	return file_aks_proto_rawDescGZIP(), []int{1}
}

func (x *AksHelloConnect) GetSalt() string {
	if x != nil {
		return x.Salt
	}
	return ""
}

var File_aks_proto protoreflect.FileDescriptor

var file_aks_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6b, 0x72, 0x61,
	0x6c, 0x61, 0x6d, 0x6f, 0x75, 0x72, 0x65, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x6f, 0x22, 0x09, 0x0a,
	0x07, 0x41, 0x6b, 0x73, 0x50, 0x69, 0x6e, 0x67, 0x22, 0x25, 0x0a, 0x0f, 0x41, 0x6b, 0x73, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x61, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x61, 0x6c, 0x74, 0x42,
	0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x72,
	0x61, 0x6c, 0x61, 0x6d, 0x6f, 0x75, 0x72, 0x65, 0x2f, 0x72, 0x65, 0x74, 0x72, 0x6f, 0x70, 0x62,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65, 0x74, 0x72, 0x6f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aks_proto_rawDescOnce sync.Once
	file_aks_proto_rawDescData = file_aks_proto_rawDesc
)

func file_aks_proto_rawDescGZIP() []byte {
	file_aks_proto_rawDescOnce.Do(func() {
		file_aks_proto_rawDescData = protoimpl.X.CompressGZIP(file_aks_proto_rawDescData)
	})
	return file_aks_proto_rawDescData
}

var file_aks_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_aks_proto_goTypes = []interface{}{
	(*AksPing)(nil),         // 0: kralamoure.retro.AksPing
	(*AksHelloConnect)(nil), // 1: kralamoure.retro.AksHelloConnect
}
var file_aks_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_aks_proto_init() }
func file_aks_proto_init() {
	if File_aks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AksPing); i {
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
		file_aks_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AksHelloConnect); i {
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
			RawDescriptor: file_aks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aks_proto_goTypes,
		DependencyIndexes: file_aks_proto_depIdxs,
		MessageInfos:      file_aks_proto_msgTypes,
	}.Build()
	File_aks_proto = out.File
	file_aks_proto_rawDesc = nil
	file_aks_proto_goTypes = nil
	file_aks_proto_depIdxs = nil
}
