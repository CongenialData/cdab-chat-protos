// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: cdab-chat.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_cdab_chat_proto protoreflect.FileDescriptor

var file_cdab_chat_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x64, 0x61, 0x62, 0x2d, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x63, 0x64, 0x61, 0x62, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x1a, 0x0a, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xee, 0x01, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x2e, 0x63, 0x64, 0x61, 0x62, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x64, 0x61, 0x62, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4e, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x2e, 0x63, 0x64, 0x61, 0x62, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x64, 0x61, 0x62, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x04, 0x51, 0x75, 0x69, 0x7a, 0x12, 0x16,
	0x2e, 0x63, 0x64, 0x61, 0x62, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x51, 0x75, 0x69, 0x7a, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x64, 0x61, 0x62, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x51, 0x75, 0x69, 0x7a, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x6e, 0x67, 0x65, 0x6e, 0x69, 0x61, 0x6c, 0x44, 0x61, 0x74,
	0x61, 0x2f, 0x63, 0x64, 0x61, 0x62, 0x2d, 0x63, 0x68, 0x61, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_cdab_chat_proto_goTypes = []interface{}{
	(*ConnectionRequest)(nil),   // 0: cdab.chat.ConnectionRequest
	(*PostMessageRequest)(nil),  // 1: cdab.chat.PostMessageRequest
	(*QuizRequest)(nil),         // 2: cdab.chat.QuizRequest
	(*ConnectionResponse)(nil),  // 3: cdab.chat.ConnectionResponse
	(*PostMessageResponse)(nil), // 4: cdab.chat.PostMessageResponse
	(*QuizResponse)(nil),        // 5: cdab.chat.QuizResponse
}
var file_cdab_chat_proto_depIdxs = []int32{
	0, // 0: cdab.chat.ChatService.ClientConnect:input_type -> cdab.chat.ConnectionRequest
	1, // 1: cdab.chat.ChatService.PostMessage:input_type -> cdab.chat.PostMessageRequest
	2, // 2: cdab.chat.ChatService.Quiz:input_type -> cdab.chat.QuizRequest
	3, // 3: cdab.chat.ChatService.ClientConnect:output_type -> cdab.chat.ConnectionResponse
	4, // 4: cdab.chat.ChatService.PostMessage:output_type -> cdab.chat.PostMessageResponse
	5, // 5: cdab.chat.ChatService.Quiz:output_type -> cdab.chat.QuizResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cdab_chat_proto_init() }
func file_cdab_chat_proto_init() {
	if File_cdab_chat_proto != nil {
		return
	}
	file_chat_proto_init()
	file_quiz_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cdab_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cdab_chat_proto_goTypes,
		DependencyIndexes: file_cdab_chat_proto_depIdxs,
	}.Build()
	File_cdab_chat_proto = out.File
	file_cdab_chat_proto_rawDesc = nil
	file_cdab_chat_proto_goTypes = nil
	file_cdab_chat_proto_depIdxs = nil
}
