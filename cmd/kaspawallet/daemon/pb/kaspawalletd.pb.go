// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: kaspawalletd.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBalanceRequest) Reset() {
	*x = GetBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kaspawalletd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceRequest) ProtoMessage() {}

func (x *GetBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kaspawalletd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceRequest.ProtoReflect.Descriptor instead.
func (*GetBalanceRequest) Descriptor() ([]byte, []int) {
	return file_kaspawalletd_proto_rawDescGZIP(), []int{0}
}

type GetBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Available uint64 `protobuf:"varint,1,opt,name=available,proto3" json:"available,omitempty"`
	Pending   uint64 `protobuf:"varint,2,opt,name=pending,proto3" json:"pending,omitempty"`
}

func (x *GetBalanceResponse) Reset() {
	*x = GetBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kaspawalletd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceResponse) ProtoMessage() {}

func (x *GetBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kaspawalletd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceResponse.ProtoReflect.Descriptor instead.
func (*GetBalanceResponse) Descriptor() ([]byte, []int) {
	return file_kaspawalletd_proto_rawDescGZIP(), []int{1}
}

func (x *GetBalanceResponse) GetAvailable() uint64 {
	if x != nil {
		return x.Available
	}
	return 0
}

func (x *GetBalanceResponse) GetPending() uint64 {
	if x != nil {
		return x.Pending
	}
	return 0
}

type CreateUnsignedTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Amount  uint64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *CreateUnsignedTransactionRequest) Reset() {
	*x = CreateUnsignedTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kaspawalletd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUnsignedTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUnsignedTransactionRequest) ProtoMessage() {}

func (x *CreateUnsignedTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kaspawalletd_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUnsignedTransactionRequest.ProtoReflect.Descriptor instead.
func (*CreateUnsignedTransactionRequest) Descriptor() ([]byte, []int) {
	return file_kaspawalletd_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUnsignedTransactionRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateUnsignedTransactionRequest) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type CreateUnsignedTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnsignedTransaction []byte `protobuf:"bytes,1,opt,name=unsignedTransaction,proto3" json:"unsignedTransaction,omitempty"`
}

func (x *CreateUnsignedTransactionResponse) Reset() {
	*x = CreateUnsignedTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kaspawalletd_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUnsignedTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUnsignedTransactionResponse) ProtoMessage() {}

func (x *CreateUnsignedTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kaspawalletd_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUnsignedTransactionResponse.ProtoReflect.Descriptor instead.
func (*CreateUnsignedTransactionResponse) Descriptor() ([]byte, []int) {
	return file_kaspawalletd_proto_rawDescGZIP(), []int{3}
}

func (x *CreateUnsignedTransactionResponse) GetUnsignedTransaction() []byte {
	if x != nil {
		return x.UnsignedTransaction
	}
	return nil
}

type GetReceiveAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetReceiveAddressRequest) Reset() {
	*x = GetReceiveAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kaspawalletd_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReceiveAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReceiveAddressRequest) ProtoMessage() {}

func (x *GetReceiveAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kaspawalletd_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReceiveAddressRequest.ProtoReflect.Descriptor instead.
func (*GetReceiveAddressRequest) Descriptor() ([]byte, []int) {
	return file_kaspawalletd_proto_rawDescGZIP(), []int{4}
}

type GetReceiveAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address        string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	DerivationPath string `protobuf:"bytes,2,opt,name=derivationPath,proto3" json:"derivationPath,omitempty"`
}

func (x *GetReceiveAddressResponse) Reset() {
	*x = GetReceiveAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kaspawalletd_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReceiveAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReceiveAddressResponse) ProtoMessage() {}

func (x *GetReceiveAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kaspawalletd_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReceiveAddressResponse.ProtoReflect.Descriptor instead.
func (*GetReceiveAddressResponse) Descriptor() ([]byte, []int) {
	return file_kaspawalletd_proto_rawDescGZIP(), []int{5}
}

func (x *GetReceiveAddressResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GetReceiveAddressResponse) GetDerivationPath() string {
	if x != nil {
		return x.DerivationPath
	}
	return ""
}

type BroadcastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction []byte `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *BroadcastRequest) Reset() {
	*x = BroadcastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kaspawalletd_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastRequest) ProtoMessage() {}

func (x *BroadcastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kaspawalletd_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastRequest.ProtoReflect.Descriptor instead.
func (*BroadcastRequest) Descriptor() ([]byte, []int) {
	return file_kaspawalletd_proto_rawDescGZIP(), []int{6}
}

func (x *BroadcastRequest) GetTransaction() []byte {
	if x != nil {
		return x.Transaction
	}
	return nil
}

type BroadcastResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxID string `protobuf:"bytes,1,opt,name=txID,proto3" json:"txID,omitempty"`
}

func (x *BroadcastResponse) Reset() {
	*x = BroadcastResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kaspawalletd_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastResponse) ProtoMessage() {}

func (x *BroadcastResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kaspawalletd_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastResponse.ProtoReflect.Descriptor instead.
func (*BroadcastResponse) Descriptor() ([]byte, []int) {
	return file_kaspawalletd_proto_rawDescGZIP(), []int{7}
}

func (x *BroadcastResponse) GetTxID() string {
	if x != nil {
		return x.TxID
	}
	return ""
}

type ShutdownRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ShutdownRequest) Reset() {
	*x = ShutdownRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kaspawalletd_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShutdownRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShutdownRequest) ProtoMessage() {}

func (x *ShutdownRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kaspawalletd_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShutdownRequest.ProtoReflect.Descriptor instead.
func (*ShutdownRequest) Descriptor() ([]byte, []int) {
	return file_kaspawalletd_proto_rawDescGZIP(), []int{8}
}

type ShutdownResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ShutdownResponse) Reset() {
	*x = ShutdownResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kaspawalletd_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShutdownResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShutdownResponse) ProtoMessage() {}

func (x *ShutdownResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kaspawalletd_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShutdownResponse.ProtoReflect.Descriptor instead.
func (*ShutdownResponse) Descriptor() ([]byte, []int) {
	return file_kaspawalletd_proto_rawDescGZIP(), []int{9}
}

var File_kaspawalletd_proto protoreflect.FileDescriptor

var file_kaspawalletd_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6b, 0x61, 0x73, 0x70, 0x61, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4c, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x54, 0x0a, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x55, 0x0a,
	0x21, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x75, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x13, 0x75, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x5d, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x64, 0x65, 0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x22,
	0x34, 0x0a, 0x10, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x27, 0x0a, 0x11, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x49, 0x44, 0x22, 0x11,
	0x0a, 0x0f, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xe4, 0x02, 0x0a, 0x0c, 0x6b, 0x61, 0x73, 0x70, 0x61, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x64, 0x12, 0x37, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x64, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x19, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x08, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x12,
	0x10, 0x2e, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x73, 0x70, 0x61,
	0x6e, 0x65, 0x74, 0x2f, 0x6b, 0x61, 0x73, 0x70, 0x61, 0x64, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x6b,
	0x61, 0x73, 0x70, 0x61, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x64, 0x61, 0x65, 0x6d, 0x6f,
	0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kaspawalletd_proto_rawDescOnce sync.Once
	file_kaspawalletd_proto_rawDescData = file_kaspawalletd_proto_rawDesc
)

func file_kaspawalletd_proto_rawDescGZIP() []byte {
	file_kaspawalletd_proto_rawDescOnce.Do(func() {
		file_kaspawalletd_proto_rawDescData = protoimpl.X.CompressGZIP(file_kaspawalletd_proto_rawDescData)
	})
	return file_kaspawalletd_proto_rawDescData
}

var file_kaspawalletd_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_kaspawalletd_proto_goTypes = []interface{}{
	(*GetBalanceRequest)(nil),                 // 0: GetBalanceRequest
	(*GetBalanceResponse)(nil),                // 1: GetBalanceResponse
	(*CreateUnsignedTransactionRequest)(nil),  // 2: CreateUnsignedTransactionRequest
	(*CreateUnsignedTransactionResponse)(nil), // 3: CreateUnsignedTransactionResponse
	(*GetReceiveAddressRequest)(nil),          // 4: GetReceiveAddressRequest
	(*GetReceiveAddressResponse)(nil),         // 5: GetReceiveAddressResponse
	(*BroadcastRequest)(nil),                  // 6: BroadcastRequest
	(*BroadcastResponse)(nil),                 // 7: BroadcastResponse
	(*ShutdownRequest)(nil),                   // 8: ShutdownRequest
	(*ShutdownResponse)(nil),                  // 9: ShutdownResponse
}
var file_kaspawalletd_proto_depIdxs = []int32{
	0, // 0: kaspawalletd.GetBalance:input_type -> GetBalanceRequest
	2, // 1: kaspawalletd.CreateUnsignedTransaction:input_type -> CreateUnsignedTransactionRequest
	4, // 2: kaspawalletd.GetReceiveAddress:input_type -> GetReceiveAddressRequest
	8, // 3: kaspawalletd.Shutdown:input_type -> ShutdownRequest
	6, // 4: kaspawalletd.Broadcast:input_type -> BroadcastRequest
	1, // 5: kaspawalletd.GetBalance:output_type -> GetBalanceResponse
	3, // 6: kaspawalletd.CreateUnsignedTransaction:output_type -> CreateUnsignedTransactionResponse
	5, // 7: kaspawalletd.GetReceiveAddress:output_type -> GetReceiveAddressResponse
	9, // 8: kaspawalletd.Shutdown:output_type -> ShutdownResponse
	7, // 9: kaspawalletd.Broadcast:output_type -> BroadcastResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kaspawalletd_proto_init() }
func file_kaspawalletd_proto_init() {
	if File_kaspawalletd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kaspawalletd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceRequest); i {
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
		file_kaspawalletd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceResponse); i {
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
		file_kaspawalletd_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUnsignedTransactionRequest); i {
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
		file_kaspawalletd_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUnsignedTransactionResponse); i {
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
		file_kaspawalletd_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReceiveAddressRequest); i {
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
		file_kaspawalletd_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReceiveAddressResponse); i {
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
		file_kaspawalletd_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastRequest); i {
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
		file_kaspawalletd_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastResponse); i {
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
		file_kaspawalletd_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShutdownRequest); i {
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
		file_kaspawalletd_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShutdownResponse); i {
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
			RawDescriptor: file_kaspawalletd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kaspawalletd_proto_goTypes,
		DependencyIndexes: file_kaspawalletd_proto_depIdxs,
		MessageInfos:      file_kaspawalletd_proto_msgTypes,
	}.Build()
	File_kaspawalletd_proto = out.File
	file_kaspawalletd_proto_rawDesc = nil
	file_kaspawalletd_proto_goTypes = nil
	file_kaspawalletd_proto_depIdxs = nil
}
