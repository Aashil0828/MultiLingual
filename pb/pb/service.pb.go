// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0--rc1
// source: service.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Language struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	LanguageName string `protobuf:"bytes,2,opt,name=LanguageName,proto3" json:"LanguageName,omitempty"`
}

func (x *Language) Reset() {
	*x = Language{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Language) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Language) ProtoMessage() {}

func (x *Language) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Language.ProtoReflect.Descriptor instead.
func (*Language) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *Language) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Language) GetLanguageName() string {
	if x != nil {
		return x.LanguageName
	}
	return ""
}

type TextContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=Label,proto3" json:"Label,omitempty"`
}

func (x *TextContent) Reset() {
	*x = TextContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextContent) ProtoMessage() {}

func (x *TextContent) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextContent.ProtoReflect.Descriptor instead.
func (*TextContent) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *TextContent) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TextContent) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type MultiLingualRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LanguageId uint32         `protobuf:"varint,1,opt,name=LanguageId,proto3" json:"LanguageId,omitempty"`
	Data       []*TextContent `protobuf:"bytes,2,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *MultiLingualRequest) Reset() {
	*x = MultiLingualRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiLingualRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiLingualRequest) ProtoMessage() {}

func (x *MultiLingualRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiLingualRequest.ProtoReflect.Descriptor instead.
func (*MultiLingualRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *MultiLingualRequest) GetLanguageId() uint32 {
	if x != nil {
		return x.LanguageId
	}
	return 0
}

func (x *MultiLingualRequest) GetData() []*TextContent {
	if x != nil {
		return x.Data
	}
	return nil
}

type MultiLingualResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32          `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Message string         `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	Data    []*TextContent `protobuf:"bytes,3,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *MultiLingualResponse) Reset() {
	*x = MultiLingualResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiLingualResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiLingualResponse) ProtoMessage() {}

func (x *MultiLingualResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiLingualResponse.ProtoReflect.Descriptor instead.
func (*MultiLingualResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *MultiLingualResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *MultiLingualResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MultiLingualResponse) GetData() []*TextContent {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetSupportedLanguagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32       `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Message string      `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	Data    []*Language `protobuf:"bytes,3,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *GetSupportedLanguagesResponse) Reset() {
	*x = GetSupportedLanguagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSupportedLanguagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSupportedLanguagesResponse) ProtoMessage() {}

func (x *GetSupportedLanguagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSupportedLanguagesResponse.ProtoReflect.Descriptor instead.
func (*GetSupportedLanguagesResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetSupportedLanguagesResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetSupportedLanguagesResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetSupportedLanguagesResponse) GetData() []*Language {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x08,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x0b,
	0x54, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x22, 0x57, 0x0a, 0x13, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x4c, 0x69, 0x6e, 0x67, 0x75, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x6a, 0x0a, 0x14, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x4c, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x70, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x32, 0xa0, 0x01, 0x0a, 0x13, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x4c, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x38, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x4c, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x4c, 0x69, 0x6e, 0x67, 0x75,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_service_proto_goTypes = []interface{}{
	(*Language)(nil),                      // 0: Language
	(*TextContent)(nil),                   // 1: TextContent
	(*MultiLingualRequest)(nil),           // 2: multiLingualRequest
	(*MultiLingualResponse)(nil),          // 3: multiLingualResponse
	(*GetSupportedLanguagesResponse)(nil), // 4: GetSupportedLanguagesResponse
	(*emptypb.Empty)(nil),                 // 5: google.protobuf.Empty
}
var file_service_proto_depIdxs = []int32{
	1, // 0: multiLingualRequest.Data:type_name -> TextContent
	1, // 1: multiLingualResponse.Data:type_name -> TextContent
	0, // 2: GetSupportedLanguagesResponse.Data:type_name -> Language
	2, // 3: multiLingualService.translate:input_type -> multiLingualRequest
	5, // 4: multiLingualService.GetSupportedLanguages:input_type -> google.protobuf.Empty
	3, // 5: multiLingualService.translate:output_type -> multiLingualResponse
	4, // 6: multiLingualService.GetSupportedLanguages:output_type -> GetSupportedLanguagesResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Language); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextContent); i {
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
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiLingualRequest); i {
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
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiLingualResponse); i {
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
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSupportedLanguagesResponse); i {
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
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
