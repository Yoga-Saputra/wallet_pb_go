//
// Internal System gRPC Services
//
// Common task related messages.
// This proto contains everything related to common task services
// accepted payload, returned response and method.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.1
// source: v1/insys/common.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents the create wallet request and response standard message.
type CreateWallet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateWallet) Reset() {
	*x = CreateWallet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_insys_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWallet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWallet) ProtoMessage() {}

func (x *CreateWallet) ProtoReflect() protoreflect.Message {
	mi := &file_v1_insys_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWallet.ProtoReflect.Descriptor instead.
func (*CreateWallet) Descriptor() ([]byte, []int) {
	return file_v1_insys_common_proto_rawDescGZIP(), []int{0}
}

// * Accepted request payload for create wallet service.
type CreateWallet_Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BranchId uint32 `protobuf:"varint,1,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"` /// ID of branch.
	MemberId uint64 `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"` /// ID of member.
	PId      string `protobuf:"bytes,3,opt,name=p_id,json=pId,proto3" json:"p_id,omitempty"`                 /// pID of member.
	Currency string `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`                  /// Currency of member.
	Username string `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`                  /// Username of member.
}

func (x *CreateWallet_Req) Reset() {
	*x = CreateWallet_Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_insys_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWallet_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWallet_Req) ProtoMessage() {}

func (x *CreateWallet_Req) ProtoReflect() protoreflect.Message {
	mi := &file_v1_insys_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWallet_Req.ProtoReflect.Descriptor instead.
func (*CreateWallet_Req) Descriptor() ([]byte, []int) {
	return file_v1_insys_common_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CreateWallet_Req) GetBranchId() uint32 {
	if x != nil {
		return x.BranchId
	}
	return 0
}

func (x *CreateWallet_Req) GetMemberId() uint64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *CreateWallet_Req) GetPId() string {
	if x != nil {
		return x.PId
	}
	return ""
}

func (x *CreateWallet_Req) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *CreateWallet_Req) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

// * Response message that will be returned from create wallet service.
type CreateWallet_Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool                    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` /// Status of the response.
	Code    int32                   `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`       /// Code of the response.
	Error   *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`      /// Error info.
	Data    string                  `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`        /// Data of the response.
}

func (x *CreateWallet_Res) Reset() {
	*x = CreateWallet_Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_insys_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWallet_Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWallet_Res) ProtoMessage() {}

func (x *CreateWallet_Res) ProtoReflect() protoreflect.Message {
	mi := &file_v1_insys_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWallet_Res.ProtoReflect.Descriptor instead.
func (*CreateWallet_Res) Descriptor() ([]byte, []int) {
	return file_v1_insys_common_proto_rawDescGZIP(), []int{0, 1}
}

func (x *CreateWallet_Res) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateWallet_Res) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateWallet_Res) GetError() *wrapperspb.StringValue {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *CreateWallet_Res) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_v1_insys_common_proto protoreflect.FileDescriptor

var file_v1_insys_common_proto_rawDesc = []byte{
	0x0a, 0x15, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x79, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x69, 0x6e, 0x73, 0x79, 0x73, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x02, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x1a, 0x8a, 0x01, 0x0a, 0x03, 0x52, 0x65,
	0x71, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x11, 0x0a, 0x04, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x7b, 0x0a, 0x03, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x32, 0x5e, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x54, 0x0a,
	0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x21, 0x2e,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x73, 0x79, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x71,
	0x1a, 0x21, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x73,
	0x79, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x52, 0x65, 0x73, 0x42, 0x31, 0x5a, 0x16, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x6e, 0x73, 0x79, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xca, 0x02, 0x16,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x49, 0x6e, 0x53, 0x79, 0x73, 0x5c,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_insys_common_proto_rawDescOnce sync.Once
	file_v1_insys_common_proto_rawDescData = file_v1_insys_common_proto_rawDesc
)

func file_v1_insys_common_proto_rawDescGZIP() []byte {
	file_v1_insys_common_proto_rawDescOnce.Do(func() {
		file_v1_insys_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_insys_common_proto_rawDescData)
	})
	return file_v1_insys_common_proto_rawDescData
}

var file_v1_insys_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_insys_common_proto_goTypes = []any{
	(*CreateWallet)(nil),           // 0: wallet.v1.insys.CreateWallet
	(*CreateWallet_Req)(nil),       // 1: wallet.v1.insys.CreateWallet.Req
	(*CreateWallet_Res)(nil),       // 2: wallet.v1.insys.CreateWallet.Res
	(*wrapperspb.StringValue)(nil), // 3: google.protobuf.StringValue
}
var file_v1_insys_common_proto_depIdxs = []int32{
	3, // 0: wallet.v1.insys.CreateWallet.Res.error:type_name -> google.protobuf.StringValue
	1, // 1: wallet.v1.insys.Common.CreateWallet:input_type -> wallet.v1.insys.CreateWallet.Req
	2, // 2: wallet.v1.insys.Common.CreateWallet:output_type -> wallet.v1.insys.CreateWallet.Res
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_insys_common_proto_init() }
func file_v1_insys_common_proto_init() {
	if File_v1_insys_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_insys_common_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateWallet); i {
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
		file_v1_insys_common_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateWallet_Req); i {
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
		file_v1_insys_common_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateWallet_Res); i {
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
			RawDescriptor: file_v1_insys_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_insys_common_proto_goTypes,
		DependencyIndexes: file_v1_insys_common_proto_depIdxs,
		MessageInfos:      file_v1_insys_common_proto_msgTypes,
	}.Build()
	File_v1_insys_common_proto = out.File
	file_v1_insys_common_proto_rawDesc = nil
	file_v1_insys_common_proto_goTypes = nil
	file_v1_insys_common_proto_depIdxs = nil
}