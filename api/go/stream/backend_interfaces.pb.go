// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: stream/backend_interfaces.proto

package stream

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type BackendInterfacesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Sender ID.
	SenderId string `protobuf:"bytes,1,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	// Receiver ID.
	ReceiverId string `protobuf:"bytes,2,opt,name=receiver_id,json=receiverId,proto3" json:"receiver_id,omitempty"`
	//   Timestamp.
	Time *timestamp.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	// Transaction ID.
	TransactionId uint32 `protobuf:"varint,4,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	// Message-type.
	MessageType string `protobuf:"bytes,5,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"`
	// Result code.
	ResultCode string `protobuf:"bytes,6,opt,name=result_code,json=resultCode,proto3" json:"result_code,omitempty"`
	// Request body.
	RequestBody string `protobuf:"bytes,7,opt,name=request_body,json=requestBody,proto3" json:"request_body,omitempty"`
	// Request error.
	RequestError string `protobuf:"bytes,8,opt,name=request_error,json=requestError,proto3" json:"request_error,omitempty"`
	// Response body.
	ResponseBody string `protobuf:"bytes,9,opt,name=response_body,json=responseBody,proto3" json:"response_body,omitempty"`
}

func (x *BackendInterfacesRequest) Reset() {
	*x = BackendInterfacesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_backend_interfaces_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackendInterfacesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackendInterfacesRequest) ProtoMessage() {}

func (x *BackendInterfacesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stream_backend_interfaces_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackendInterfacesRequest.ProtoReflect.Descriptor instead.
func (*BackendInterfacesRequest) Descriptor() ([]byte, []int) {
	return file_stream_backend_interfaces_proto_rawDescGZIP(), []int{0}
}

func (x *BackendInterfacesRequest) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *BackendInterfacesRequest) GetReceiverId() string {
	if x != nil {
		return x.ReceiverId
	}
	return ""
}

func (x *BackendInterfacesRequest) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *BackendInterfacesRequest) GetTransactionId() uint32 {
	if x != nil {
		return x.TransactionId
	}
	return 0
}

func (x *BackendInterfacesRequest) GetMessageType() string {
	if x != nil {
		return x.MessageType
	}
	return ""
}

func (x *BackendInterfacesRequest) GetResultCode() string {
	if x != nil {
		return x.ResultCode
	}
	return ""
}

func (x *BackendInterfacesRequest) GetRequestBody() string {
	if x != nil {
		return x.RequestBody
	}
	return ""
}

func (x *BackendInterfacesRequest) GetRequestError() string {
	if x != nil {
		return x.RequestError
	}
	return ""
}

func (x *BackendInterfacesRequest) GetResponseBody() string {
	if x != nil {
		return x.ResponseBody
	}
	return ""
}

var File_stream_backend_interfaces_proto protoreflect.FileDescriptor

var file_stream_backend_interfaces_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x02, 0x0a, 0x18, 0x42,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42,
	0x6f, 0x64, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x42, 0xaf, 0x01,
	0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x63, 0x68, 0x69, 0x72, 0x70, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x42, 0x16, 0x42, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x68, 0x69, 0x72, 0x70, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x63, 0x68, 0x69, 0x72,
	0x70, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x34,
	0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0xaa, 0x02, 0x11, 0x43, 0x68, 0x69, 0x72, 0x70, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0xca, 0x02, 0x11, 0x43, 0x68,
	0x69, 0x72, 0x70, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0xe2,
	0x02, 0x1d, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5c, 0x43, 0x68,
	0x69, 0x72, 0x70, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stream_backend_interfaces_proto_rawDescOnce sync.Once
	file_stream_backend_interfaces_proto_rawDescData = file_stream_backend_interfaces_proto_rawDesc
)

func file_stream_backend_interfaces_proto_rawDescGZIP() []byte {
	file_stream_backend_interfaces_proto_rawDescOnce.Do(func() {
		file_stream_backend_interfaces_proto_rawDescData = protoimpl.X.CompressGZIP(file_stream_backend_interfaces_proto_rawDescData)
	})
	return file_stream_backend_interfaces_proto_rawDescData
}

var file_stream_backend_interfaces_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_stream_backend_interfaces_proto_goTypes = []interface{}{
	(*BackendInterfacesRequest)(nil), // 0: stream.BackendInterfacesRequest
	(*timestamp.Timestamp)(nil),      // 1: google.protobuf.Timestamp
}
var file_stream_backend_interfaces_proto_depIdxs = []int32{
	1, // 0: stream.BackendInterfacesRequest.time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_stream_backend_interfaces_proto_init() }
func file_stream_backend_interfaces_proto_init() {
	if File_stream_backend_interfaces_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stream_backend_interfaces_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackendInterfacesRequest); i {
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
			RawDescriptor: file_stream_backend_interfaces_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stream_backend_interfaces_proto_goTypes,
		DependencyIndexes: file_stream_backend_interfaces_proto_depIdxs,
		MessageInfos:      file_stream_backend_interfaces_proto_msgTypes,
	}.Build()
	File_stream_backend_interfaces_proto = out.File
	file_stream_backend_interfaces_proto_rawDesc = nil
	file_stream_backend_interfaces_proto_goTypes = nil
	file_stream_backend_interfaces_proto_depIdxs = nil
}
