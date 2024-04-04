// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: api/v1/notifier.proto

package notifier

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

type PlaceHolders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PlaceHolders) Reset() {
	*x = PlaceHolders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_notifier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaceHolders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceHolders) ProtoMessage() {}

func (x *PlaceHolders) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_notifier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceHolders.ProtoReflect.Descriptor instead.
func (*PlaceHolders) Descriptor() ([]byte, []int) {
	return file_api_v1_notifier_proto_rawDescGZIP(), []int{0}
}

func (x *PlaceHolders) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlaceHolders) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type NotifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message      string          `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Subject      string          `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Recipients   string          `protobuf:"bytes,3,opt,name=recipients,proto3" json:"recipients,omitempty"`
	Placeholders []*PlaceHolders `protobuf:"bytes,4,rep,name=placeholders,proto3" json:"placeholders,omitempty"`
}

func (x *NotifyRequest) Reset() {
	*x = NotifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_notifier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyRequest) ProtoMessage() {}

func (x *NotifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_notifier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyRequest.ProtoReflect.Descriptor instead.
func (*NotifyRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_notifier_proto_rawDescGZIP(), []int{1}
}

func (x *NotifyRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *NotifyRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *NotifyRequest) GetRecipients() string {
	if x != nil {
		return x.Recipients
	}
	return ""
}

func (x *NotifyRequest) GetPlaceholders() []*PlaceHolders {
	if x != nil {
		return x.Placeholders
	}
	return nil
}

type NotifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message      string          `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Subject      string          `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Recipients   string          `protobuf:"bytes,3,opt,name=recipients,proto3" json:"recipients,omitempty"`
	Placeholders []*PlaceHolders `protobuf:"bytes,4,rep,name=placeholders,proto3" json:"placeholders,omitempty"`
	Status       string          `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Error        string          `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *NotifyResponse) Reset() {
	*x = NotifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_notifier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyResponse) ProtoMessage() {}

func (x *NotifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_notifier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyResponse.ProtoReflect.Descriptor instead.
func (*NotifyResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_notifier_proto_rawDescGZIP(), []int{2}
}

func (x *NotifyResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *NotifyResponse) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *NotifyResponse) GetRecipients() string {
	if x != nil {
		return x.Recipients
	}
	return ""
}

func (x *NotifyResponse) GetPlaceholders() []*PlaceHolders {
	if x != nil {
		return x.Placeholders
	}
	return nil
}

func (x *NotifyResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *NotifyResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_api_v1_notifier_proto protoreflect.FileDescriptor

var file_api_v1_notifier_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x96, 0x01, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x31, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x0c, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x22, 0xc5, 0x01, 0x0a, 0x0e, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x31, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x48,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x32, 0x3e, 0x0a, 0x0f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12,
	0x0e, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0f, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x72, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x6e, 0x65, 0x74, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x2f, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_notifier_proto_rawDescOnce sync.Once
	file_api_v1_notifier_proto_rawDescData = file_api_v1_notifier_proto_rawDesc
)

func file_api_v1_notifier_proto_rawDescGZIP() []byte {
	file_api_v1_notifier_proto_rawDescOnce.Do(func() {
		file_api_v1_notifier_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_notifier_proto_rawDescData)
	})
	return file_api_v1_notifier_proto_rawDescData
}

var file_api_v1_notifier_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_v1_notifier_proto_goTypes = []interface{}{
	(*PlaceHolders)(nil),   // 0: PlaceHolders
	(*NotifyRequest)(nil),  // 1: NotifyRequest
	(*NotifyResponse)(nil), // 2: NotifyResponse
}
var file_api_v1_notifier_proto_depIdxs = []int32{
	0, // 0: NotifyRequest.placeholders:type_name -> PlaceHolders
	0, // 1: NotifyResponse.placeholders:type_name -> PlaceHolders
	1, // 2: NotifierService.Notify:input_type -> NotifyRequest
	2, // 3: NotifierService.Notify:output_type -> NotifyResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1_notifier_proto_init() }
func file_api_v1_notifier_proto_init() {
	if File_api_v1_notifier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_notifier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaceHolders); i {
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
		file_api_v1_notifier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyRequest); i {
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
		file_api_v1_notifier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyResponse); i {
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
			RawDescriptor: file_api_v1_notifier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_notifier_proto_goTypes,
		DependencyIndexes: file_api_v1_notifier_proto_depIdxs,
		MessageInfos:      file_api_v1_notifier_proto_msgTypes,
	}.Build()
	File_api_v1_notifier_proto = out.File
	file_api_v1_notifier_proto_rawDesc = nil
	file_api_v1_notifier_proto_goTypes = nil
	file_api_v1_notifier_proto_depIdxs = nil
}
