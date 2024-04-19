// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: skytree_backend.proto

package skytree_backend

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

type DepositPointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CrossChainTxId int64 `protobuf:"varint,1,opt,name=cross_chain_tx_id,json=crossChainTxId,proto3" json:"cross_chain_tx_id,omitempty"`
}

func (x *DepositPointRequest) Reset() {
	*x = DepositPointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skytree_backend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositPointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositPointRequest) ProtoMessage() {}

func (x *DepositPointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_skytree_backend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositPointRequest.ProtoReflect.Descriptor instead.
func (*DepositPointRequest) Descriptor() ([]byte, []int) {
	return file_skytree_backend_proto_rawDescGZIP(), []int{0}
}

func (x *DepositPointRequest) GetCrossChainTxId() int64 {
	if x != nil {
		return x.CrossChainTxId
	}
	return 0
}

type DepositPointResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DepositPointResponse) Reset() {
	*x = DepositPointResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skytree_backend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositPointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositPointResponse) ProtoMessage() {}

func (x *DepositPointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_skytree_backend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositPointResponse.ProtoReflect.Descriptor instead.
func (*DepositPointResponse) Descriptor() ([]byte, []int) {
	return file_skytree_backend_proto_rawDescGZIP(), []int{1}
}

func (x *DepositPointResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type ExchangePointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenAmount string `protobuf:"bytes,1,opt,name=token_amount,json=tokenAmount,proto3" json:"token_amount,omitempty"`
	User        string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	ClaimTx     string `protobuf:"bytes,3,opt,name=claim_tx,json=claimTx,proto3" json:"claim_tx,omitempty"`
	ClaimTime   int64  `protobuf:"varint,4,opt,name=claim_time,json=claimTime,proto3" json:"claim_time,omitempty"`
}

func (x *ExchangePointRequest) Reset() {
	*x = ExchangePointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skytree_backend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangePointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangePointRequest) ProtoMessage() {}

func (x *ExchangePointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_skytree_backend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangePointRequest.ProtoReflect.Descriptor instead.
func (*ExchangePointRequest) Descriptor() ([]byte, []int) {
	return file_skytree_backend_proto_rawDescGZIP(), []int{2}
}

func (x *ExchangePointRequest) GetTokenAmount() string {
	if x != nil {
		return x.TokenAmount
	}
	return ""
}

func (x *ExchangePointRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ExchangePointRequest) GetClaimTx() string {
	if x != nil {
		return x.ClaimTx
	}
	return ""
}

func (x *ExchangePointRequest) GetClaimTime() int64 {
	if x != nil {
		return x.ClaimTime
	}
	return 0
}

type ExchangePointResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ExchangePointResponse) Reset() {
	*x = ExchangePointResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skytree_backend_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangePointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangePointResponse) ProtoMessage() {}

func (x *ExchangePointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_skytree_backend_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangePointResponse.ProtoReflect.Descriptor instead.
func (*ExchangePointResponse) Descriptor() ([]byte, []int) {
	return file_skytree_backend_proto_rawDescGZIP(), []int{3}
}

func (x *ExchangePointResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ExchangePointResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type OnTokenExchangedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenAmount string  `protobuf:"bytes,1,opt,name=token_amount,json=tokenAmount,proto3" json:"token_amount,omitempty"`
	User        string  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	ClaimTx     string  `protobuf:"bytes,3,opt,name=claim_tx,json=claimTx,proto3" json:"claim_tx,omitempty"`
	ClaimTime   int64   `protobuf:"varint,4,opt,name=claim_time,json=claimTime,proto3" json:"claim_time,omitempty"`
	UserData    []int32 `protobuf:"varint,5,rep,packed,name=user_data,json=userData,proto3" json:"user_data,omitempty"`
}

func (x *OnTokenExchangedRequest) Reset() {
	*x = OnTokenExchangedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skytree_backend_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnTokenExchangedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnTokenExchangedRequest) ProtoMessage() {}

func (x *OnTokenExchangedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_skytree_backend_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnTokenExchangedRequest.ProtoReflect.Descriptor instead.
func (*OnTokenExchangedRequest) Descriptor() ([]byte, []int) {
	return file_skytree_backend_proto_rawDescGZIP(), []int{4}
}

func (x *OnTokenExchangedRequest) GetTokenAmount() string {
	if x != nil {
		return x.TokenAmount
	}
	return ""
}

func (x *OnTokenExchangedRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *OnTokenExchangedRequest) GetClaimTx() string {
	if x != nil {
		return x.ClaimTx
	}
	return ""
}

func (x *OnTokenExchangedRequest) GetClaimTime() int64 {
	if x != nil {
		return x.ClaimTime
	}
	return 0
}

func (x *OnTokenExchangedRequest) GetUserData() []int32 {
	if x != nil {
		return x.UserData
	}
	return nil
}

type OnTokenExchangedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *OnTokenExchangedResponse) Reset() {
	*x = OnTokenExchangedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skytree_backend_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnTokenExchangedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnTokenExchangedResponse) ProtoMessage() {}

func (x *OnTokenExchangedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_skytree_backend_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnTokenExchangedResponse.ProtoReflect.Descriptor instead.
func (*OnTokenExchangedResponse) Descriptor() ([]byte, []int) {
	return file_skytree_backend_proto_rawDescGZIP(), []int{5}
}

func (x *OnTokenExchangedResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *OnTokenExchangedResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_skytree_backend_proto protoreflect.FileDescriptor

var file_skytree_backend_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x6b, 0x79, 0x74, 0x72, 0x65, 0x65, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x6b, 0x79, 0x74, 0x72, 0x65, 0x65,
	0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x22, 0x40, 0x0a, 0x13,
	0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x11, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x5f, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x5f, 0x74, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x63, 0x72, 0x6f, 0x73, 0x73, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x78, 0x49, 0x64, 0x22, 0x2e,
	0x0a, 0x14, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x87,
	0x01, 0x0a, 0x14, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x19,
	0x0a, 0x08, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x74, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x54, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x61,
	0x69, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x6c, 0x61, 0x69, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x3d, 0x0a, 0x15, 0x45, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0xa7, 0x01, 0x0a, 0x17, 0x4f, 0x6e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c,
	0x61, 0x69, 0x6d, 0x5f, 0x74, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c,
	0x61, 0x69, 0x6d, 0x54, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x69, 0x6d,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x40, 0x0a, 0x18, 0x4f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x32, 0xd3, 0x02, 0x0a, 0x0d, 0x53, 0x6b, 0x79, 0x74, 0x65, 0x65, 0x42, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x12, 0x65, 0x0a, 0x0c, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x28, 0x2e, 0x73, 0x6b, 0x79, 0x74, 0x72, 0x65, 0x65, 0x5f,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x73, 0x6b, 0x79, 0x74, 0x72, 0x65, 0x65, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x0d,
	0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x29, 0x2e,
	0x73, 0x6b, 0x79, 0x74, 0x72, 0x65, 0x65, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x6b, 0x79, 0x74, 0x72,
	0x65, 0x65, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x71, 0x0a, 0x10, 0x4f, 0x6e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x2c, 0x2e, 0x73, 0x6b, 0x79,
	0x74, 0x72, 0x65, 0x65, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x73, 0x6b, 0x79, 0x74, 0x72,
	0x65, 0x65, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4f,
	0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2f, 0x3b,
	0x73, 0x6b, 0x79, 0x74, 0x72, 0x65, 0x65, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_skytree_backend_proto_rawDescOnce sync.Once
	file_skytree_backend_proto_rawDescData = file_skytree_backend_proto_rawDesc
)

func file_skytree_backend_proto_rawDescGZIP() []byte {
	file_skytree_backend_proto_rawDescOnce.Do(func() {
		file_skytree_backend_proto_rawDescData = protoimpl.X.CompressGZIP(file_skytree_backend_proto_rawDescData)
	})
	return file_skytree_backend_proto_rawDescData
}

var file_skytree_backend_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_skytree_backend_proto_goTypes = []interface{}{
	(*DepositPointRequest)(nil),      // 0: skytree_backend.api.DepositPointRequest
	(*DepositPointResponse)(nil),     // 1: skytree_backend.api.DepositPointResponse
	(*ExchangePointRequest)(nil),     // 2: skytree_backend.api.ExchangePointRequest
	(*ExchangePointResponse)(nil),    // 3: skytree_backend.api.ExchangePointResponse
	(*OnTokenExchangedRequest)(nil),  // 4: skytree_backend.api.OnTokenExchangedRequest
	(*OnTokenExchangedResponse)(nil), // 5: skytree_backend.api.OnTokenExchangedResponse
}
var file_skytree_backend_proto_depIdxs = []int32{
	0, // 0: skytree_backend.api.SkyteeBackend.DepositPoint:input_type -> skytree_backend.api.DepositPointRequest
	2, // 1: skytree_backend.api.SkyteeBackend.ExchangePoint:input_type -> skytree_backend.api.ExchangePointRequest
	4, // 2: skytree_backend.api.SkyteeBackend.OnTokenExchanged:input_type -> skytree_backend.api.OnTokenExchangedRequest
	1, // 3: skytree_backend.api.SkyteeBackend.DepositPoint:output_type -> skytree_backend.api.DepositPointResponse
	3, // 4: skytree_backend.api.SkyteeBackend.ExchangePoint:output_type -> skytree_backend.api.ExchangePointResponse
	5, // 5: skytree_backend.api.SkyteeBackend.OnTokenExchanged:output_type -> skytree_backend.api.OnTokenExchangedResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_skytree_backend_proto_init() }
func file_skytree_backend_proto_init() {
	if File_skytree_backend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_skytree_backend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositPointRequest); i {
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
		file_skytree_backend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositPointResponse); i {
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
		file_skytree_backend_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangePointRequest); i {
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
		file_skytree_backend_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangePointResponse); i {
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
		file_skytree_backend_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnTokenExchangedRequest); i {
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
		file_skytree_backend_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnTokenExchangedResponse); i {
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
			RawDescriptor: file_skytree_backend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_skytree_backend_proto_goTypes,
		DependencyIndexes: file_skytree_backend_proto_depIdxs,
		MessageInfos:      file_skytree_backend_proto_msgTypes,
	}.Build()
	File_skytree_backend_proto = out.File
	file_skytree_backend_proto_rawDesc = nil
	file_skytree_backend_proto_goTypes = nil
	file_skytree_backend_proto_depIdxs = nil
}
