// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: dfuse/solana/serumhist/v1/serumhist.proto

package pbserumhist

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Side int32

const (
	Side_BID Side = 0
	Side_ASK Side = 1
)

// Enum value maps for Side.
var (
	Side_name = map[int32]string{
		0: "BID",
		1: "ASK",
	}
	Side_value = map[string]int32{
		"BID": 0,
		"ASK": 1,
	}
)

func (x Side) Enum() *Side {
	p := new(Side)
	*p = x
	return p
}

func (x Side) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Side) Descriptor() protoreflect.EnumDescriptor {
	return file_dfuse_solana_serumhist_v1_serumhist_proto_enumTypes[0].Descriptor()
}

func (Side) Type() protoreflect.EnumType {
	return &file_dfuse_solana_serumhist_v1_serumhist_proto_enumTypes[0]
}

func (x Side) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Side.Descriptor instead.
func (Side) EnumDescriptor() ([]byte, []int) {
	return file_dfuse_solana_serumhist_v1_serumhist_proto_rawDescGZIP(), []int{0}
}

type FeeTier int32

const (
	FeeTier_Base FeeTier = 0
	FeeTier_SRM2 FeeTier = 1
	FeeTier_SRM3 FeeTier = 2
	FeeTier_SRM4 FeeTier = 3
	FeeTier_SRM5 FeeTier = 4
	FeeTier_SRM6 FeeTier = 5
	FeeTier_MSRM FeeTier = 6
)

// Enum value maps for FeeTier.
var (
	FeeTier_name = map[int32]string{
		0: "Base",
		1: "SRM2",
		2: "SRM3",
		3: "SRM4",
		4: "SRM5",
		5: "SRM6",
		6: "MSRM",
	}
	FeeTier_value = map[string]int32{
		"Base": 0,
		"SRM2": 1,
		"SRM3": 2,
		"SRM4": 3,
		"SRM5": 4,
		"SRM6": 5,
		"MSRM": 6,
	}
)

func (x FeeTier) Enum() *FeeTier {
	p := new(FeeTier)
	*p = x
	return p
}

func (x FeeTier) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeeTier) Descriptor() protoreflect.EnumDescriptor {
	return file_dfuse_solana_serumhist_v1_serumhist_proto_enumTypes[1].Descriptor()
}

func (FeeTier) Type() protoreflect.EnumType {
	return &file_dfuse_solana_serumhist_v1_serumhist_proto_enumTypes[1]
}

func (x FeeTier) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeeTier.Descriptor instead.
func (FeeTier) EnumDescriptor() ([]byte, []int) {
	return file_dfuse_solana_serumhist_v1_serumhist_proto_rawDescGZIP(), []int{1}
}

type GetFillsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trader string `protobuf:"bytes,1,opt,name=trader,proto3" json:"trader,omitempty"`
	Market string `protobuf:"bytes,2,opt,name=market,proto3" json:"market,omitempty"`
}

func (x *GetFillsRequest) Reset() {
	*x = GetFillsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dfuse_solana_serumhist_v1_serumhist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFillsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFillsRequest) ProtoMessage() {}

func (x *GetFillsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dfuse_solana_serumhist_v1_serumhist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFillsRequest.ProtoReflect.Descriptor instead.
func (*GetFillsRequest) Descriptor() ([]byte, []int) {
	return file_dfuse_solana_serumhist_v1_serumhist_proto_rawDescGZIP(), []int{0}
}

func (x *GetFillsRequest) GetTrader() string {
	if x != nil {
		return x.Trader
	}
	return ""
}

func (x *GetFillsRequest) GetMarket() string {
	if x != nil {
		return x.Market
	}
	return ""
}

type FillsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fill []*Fill `protobuf:"bytes,1,rep,name=fill,proto3" json:"fill,omitempty"`
}

func (x *FillsResponse) Reset() {
	*x = FillsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dfuse_solana_serumhist_v1_serumhist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FillsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FillsResponse) ProtoMessage() {}

func (x *FillsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dfuse_solana_serumhist_v1_serumhist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FillsResponse.ProtoReflect.Descriptor instead.
func (*FillsResponse) Descriptor() ([]byte, []int) {
	return file_dfuse_solana_serumhist_v1_serumhist_proto_rawDescGZIP(), []int{1}
}

func (x *FillsResponse) GetFill() []*Fill {
	if x != nil {
		return x.Fill
	}
	return nil
}

type Checkpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastWrittenSlotNum uint64 `protobuf:"varint,1,opt,name=last_written_slot_num,json=lastWrittenSlotNum,proto3" json:"last_written_slot_num,omitempty"`
	LastWrittenLostId  string `protobuf:"bytes,2,opt,name=last_written_lost_id,json=lastWrittenLostId,proto3" json:"last_written_lost_id,omitempty"`
}

func (x *Checkpoint) Reset() {
	*x = Checkpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dfuse_solana_serumhist_v1_serumhist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Checkpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Checkpoint) ProtoMessage() {}

func (x *Checkpoint) ProtoReflect() protoreflect.Message {
	mi := &file_dfuse_solana_serumhist_v1_serumhist_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Checkpoint.ProtoReflect.Descriptor instead.
func (*Checkpoint) Descriptor() ([]byte, []int) {
	return file_dfuse_solana_serumhist_v1_serumhist_proto_rawDescGZIP(), []int{2}
}

func (x *Checkpoint) GetLastWrittenSlotNum() uint64 {
	if x != nil {
		return x.LastWrittenSlotNum
	}
	return 0
}

func (x *Checkpoint) GetLastWrittenLostId() string {
	if x != nil {
		return x.LastWrittenLostId
	}
	return ""
}

type Fill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trader            string  `protobuf:"bytes,1,opt,name=trader,proto3" json:"trader,omitempty"`
	OrderId           string  `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Side              Side    `protobuf:"varint,3,opt,name=side,proto3,enum=dfuse.solana.serumhist.v1.Side" json:"side,omitempty"`
	Maker             bool    `protobuf:"varint,4,opt,name=maker,proto3" json:"maker,omitempty"`
	NativeQtyPaid     uint64  `protobuf:"varint,5,opt,name=native_qty_paid,json=nativeQtyPaid,proto3" json:"native_qty_paid,omitempty"`
	NativeQtyReceived uint64  `protobuf:"varint,6,opt,name=native_qty_received,json=nativeQtyReceived,proto3" json:"native_qty_received,omitempty"`
	NativeFeeOrRebate uint64  `protobuf:"varint,7,opt,name=native_fee_or_rebate,json=nativeFeeOrRebate,proto3" json:"native_fee_or_rebate,omitempty"`
	FeeTier           FeeTier `protobuf:"varint,8,opt,name=fee_tier,json=feeTier,proto3,enum=dfuse.solana.serumhist.v1.FeeTier" json:"fee_tier,omitempty"`
}

func (x *Fill) Reset() {
	*x = Fill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dfuse_solana_serumhist_v1_serumhist_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fill) ProtoMessage() {}

func (x *Fill) ProtoReflect() protoreflect.Message {
	mi := &file_dfuse_solana_serumhist_v1_serumhist_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fill.ProtoReflect.Descriptor instead.
func (*Fill) Descriptor() ([]byte, []int) {
	return file_dfuse_solana_serumhist_v1_serumhist_proto_rawDescGZIP(), []int{3}
}

func (x *Fill) GetTrader() string {
	if x != nil {
		return x.Trader
	}
	return ""
}

func (x *Fill) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Fill) GetSide() Side {
	if x != nil {
		return x.Side
	}
	return Side_BID
}

func (x *Fill) GetMaker() bool {
	if x != nil {
		return x.Maker
	}
	return false
}

func (x *Fill) GetNativeQtyPaid() uint64 {
	if x != nil {
		return x.NativeQtyPaid
	}
	return 0
}

func (x *Fill) GetNativeQtyReceived() uint64 {
	if x != nil {
		return x.NativeQtyReceived
	}
	return 0
}

func (x *Fill) GetNativeFeeOrRebate() uint64 {
	if x != nil {
		return x.NativeFeeOrRebate
	}
	return 0
}

func (x *Fill) GetFeeTier() FeeTier {
	if x != nil {
		return x.FeeTier
	}
	return FeeTier_Base
}

var File_dfuse_solana_serumhist_v1_serumhist_proto protoreflect.FileDescriptor

var file_dfuse_solana_serumhist_v1_serumhist_proto_rawDesc = []byte{
	0x0a, 0x29, 0x64, 0x66, 0x75, 0x73, 0x65, 0x2f, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2f, 0x73,
	0x65, 0x72, 0x75, 0x6d, 0x68, 0x69, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x75,
	0x6d, 0x68, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x64, 0x66, 0x75,
	0x73, 0x65, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x75, 0x6d, 0x68,
	0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x41, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x72, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x72, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x22, 0x44, 0x0a, 0x0d, 0x46, 0x69, 0x6c,
	0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x66, 0x69,
	0x6c, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x64, 0x66, 0x75, 0x73, 0x65,
	0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x75, 0x6d, 0x68, 0x69, 0x73,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x6c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x6c, 0x22,
	0x70, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x31, 0x0a,
	0x15, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x5f, 0x73, 0x6c,
	0x6f, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x6c, 0x61,
	0x73, 0x74, 0x57, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x53, 0x6c, 0x6f, 0x74, 0x4e, 0x75, 0x6d,
	0x12, 0x2f, 0x0a, 0x14, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e,
	0x5f, 0x6c, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x6c, 0x61, 0x73, 0x74, 0x57, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x4c, 0x6f, 0x73, 0x74, 0x49,
	0x64, 0x22, 0xcc, 0x02, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x72,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x72, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x33, 0x0a,
	0x04, 0x73, 0x69, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x64, 0x66,
	0x75, 0x73, 0x65, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x75, 0x6d,
	0x68, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x64, 0x65, 0x52, 0x04, 0x73, 0x69,
	0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x71, 0x74, 0x79, 0x5f, 0x70, 0x61, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0d, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x51, 0x74, 0x79, 0x50, 0x61, 0x69, 0x64,
	0x12, 0x2e, 0x0a, 0x13, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x71, 0x74, 0x79, 0x5f, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x6e,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x51, 0x74, 0x79, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64,
	0x12, 0x2f, 0x0a, 0x14, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x66, 0x65, 0x65, 0x5f, 0x6f,
	0x72, 0x5f, 0x72, 0x65, 0x62, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11,
	0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x46, 0x65, 0x65, 0x4f, 0x72, 0x52, 0x65, 0x62, 0x61, 0x74,
	0x65, 0x12, 0x3d, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x5f, 0x74, 0x69, 0x65, 0x72, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x64, 0x66, 0x75, 0x73, 0x65, 0x2e, 0x73, 0x6f, 0x6c, 0x61,
	0x6e, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x75, 0x6d, 0x68, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x65, 0x65, 0x54, 0x69, 0x65, 0x72, 0x52, 0x07, 0x66, 0x65, 0x65, 0x54, 0x69, 0x65, 0x72,
	0x2a, 0x18, 0x0a, 0x04, 0x53, 0x69, 0x64, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x49, 0x44, 0x10,
	0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x4b, 0x10, 0x01, 0x2a, 0x4f, 0x0a, 0x07, 0x46, 0x65,
	0x65, 0x54, 0x69, 0x65, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x53, 0x52, 0x4d, 0x32, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x52, 0x4d,
	0x33, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x52, 0x4d, 0x34, 0x10, 0x03, 0x12, 0x08, 0x0a,
	0x04, 0x53, 0x52, 0x4d, 0x35, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x52, 0x4d, 0x36, 0x10,
	0x05, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x53, 0x52, 0x4d, 0x10, 0x06, 0x32, 0x70, 0x0a, 0x0c, 0x53,
	0x65, 0x72, 0x75, 0x6d, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x60, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x46, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x2a, 0x2e, 0x64, 0x66, 0x75, 0x73, 0x65, 0x2e,
	0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x75, 0x6d, 0x68, 0x69, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x64, 0x66, 0x75, 0x73, 0x65, 0x2e, 0x73, 0x6f, 0x6c, 0x61,
	0x6e, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x75, 0x6d, 0x68, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4b, 0x5a,
	0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x66, 0x75, 0x73,
	0x65, 0x2d, 0x69, 0x6f, 0x2f, 0x64, 0x66, 0x75, 0x73, 0x65, 0x2d, 0x73, 0x6f, 0x6c, 0x61, 0x6e,
	0x61, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x66, 0x75, 0x73, 0x65, 0x2f, 0x73, 0x6f, 0x6c, 0x61, 0x6e,
	0x61, 0x2f, 0x73, 0x65, 0x72, 0x75, 0x6d, 0x68, 0x69, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x70,
	0x62, 0x73, 0x65, 0x72, 0x75, 0x6d, 0x68, 0x69, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_dfuse_solana_serumhist_v1_serumhist_proto_rawDescOnce sync.Once
	file_dfuse_solana_serumhist_v1_serumhist_proto_rawDescData = file_dfuse_solana_serumhist_v1_serumhist_proto_rawDesc
)

func file_dfuse_solana_serumhist_v1_serumhist_proto_rawDescGZIP() []byte {
	file_dfuse_solana_serumhist_v1_serumhist_proto_rawDescOnce.Do(func() {
		file_dfuse_solana_serumhist_v1_serumhist_proto_rawDescData = protoimpl.X.CompressGZIP(file_dfuse_solana_serumhist_v1_serumhist_proto_rawDescData)
	})
	return file_dfuse_solana_serumhist_v1_serumhist_proto_rawDescData
}

var file_dfuse_solana_serumhist_v1_serumhist_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_dfuse_solana_serumhist_v1_serumhist_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_dfuse_solana_serumhist_v1_serumhist_proto_goTypes = []interface{}{
	(Side)(0),               // 0: dfuse.solana.serumhist.v1.Side
	(FeeTier)(0),            // 1: dfuse.solana.serumhist.v1.FeeTier
	(*GetFillsRequest)(nil), // 2: dfuse.solana.serumhist.v1.GetFillsRequest
	(*FillsResponse)(nil),   // 3: dfuse.solana.serumhist.v1.FillsResponse
	(*Checkpoint)(nil),      // 4: dfuse.solana.serumhist.v1.Checkpoint
	(*Fill)(nil),            // 5: dfuse.solana.serumhist.v1.Fill
}
var file_dfuse_solana_serumhist_v1_serumhist_proto_depIdxs = []int32{
	5, // 0: dfuse.solana.serumhist.v1.FillsResponse.fill:type_name -> dfuse.solana.serumhist.v1.Fill
	0, // 1: dfuse.solana.serumhist.v1.Fill.side:type_name -> dfuse.solana.serumhist.v1.Side
	1, // 2: dfuse.solana.serumhist.v1.Fill.fee_tier:type_name -> dfuse.solana.serumhist.v1.FeeTier
	2, // 3: dfuse.solana.serumhist.v1.SerumHistory.GetFills:input_type -> dfuse.solana.serumhist.v1.GetFillsRequest
	3, // 4: dfuse.solana.serumhist.v1.SerumHistory.GetFills:output_type -> dfuse.solana.serumhist.v1.FillsResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_dfuse_solana_serumhist_v1_serumhist_proto_init() }
func file_dfuse_solana_serumhist_v1_serumhist_proto_init() {
	if File_dfuse_solana_serumhist_v1_serumhist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dfuse_solana_serumhist_v1_serumhist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFillsRequest); i {
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
		file_dfuse_solana_serumhist_v1_serumhist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FillsResponse); i {
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
		file_dfuse_solana_serumhist_v1_serumhist_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Checkpoint); i {
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
		file_dfuse_solana_serumhist_v1_serumhist_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fill); i {
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
			RawDescriptor: file_dfuse_solana_serumhist_v1_serumhist_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dfuse_solana_serumhist_v1_serumhist_proto_goTypes,
		DependencyIndexes: file_dfuse_solana_serumhist_v1_serumhist_proto_depIdxs,
		EnumInfos:         file_dfuse_solana_serumhist_v1_serumhist_proto_enumTypes,
		MessageInfos:      file_dfuse_solana_serumhist_v1_serumhist_proto_msgTypes,
	}.Build()
	File_dfuse_solana_serumhist_v1_serumhist_proto = out.File
	file_dfuse_solana_serumhist_v1_serumhist_proto_rawDesc = nil
	file_dfuse_solana_serumhist_v1_serumhist_proto_goTypes = nil
	file_dfuse_solana_serumhist_v1_serumhist_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SerumHistoryClient is the client API for SerumHistory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SerumHistoryClient interface {
	GetFills(ctx context.Context, in *GetFillsRequest, opts ...grpc.CallOption) (*FillsResponse, error)
}

type serumHistoryClient struct {
	cc grpc.ClientConnInterface
}

func NewSerumHistoryClient(cc grpc.ClientConnInterface) SerumHistoryClient {
	return &serumHistoryClient{cc}
}

func (c *serumHistoryClient) GetFills(ctx context.Context, in *GetFillsRequest, opts ...grpc.CallOption) (*FillsResponse, error) {
	out := new(FillsResponse)
	err := c.cc.Invoke(ctx, "/dfuse.solana.serumhist.v1.SerumHistory/GetFills", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SerumHistoryServer is the server API for SerumHistory service.
type SerumHistoryServer interface {
	GetFills(context.Context, *GetFillsRequest) (*FillsResponse, error)
}

// UnimplementedSerumHistoryServer can be embedded to have forward compatible implementations.
type UnimplementedSerumHistoryServer struct {
}

func (*UnimplementedSerumHistoryServer) GetFills(context.Context, *GetFillsRequest) (*FillsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFills not implemented")
}

func RegisterSerumHistoryServer(s *grpc.Server, srv SerumHistoryServer) {
	s.RegisterService(&_SerumHistory_serviceDesc, srv)
}

func _SerumHistory_GetFills_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFillsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SerumHistoryServer).GetFills(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfuse.solana.serumhist.v1.SerumHistory/GetFills",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SerumHistoryServer).GetFills(ctx, req.(*GetFillsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SerumHistory_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dfuse.solana.serumhist.v1.SerumHistory",
	HandlerType: (*SerumHistoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFills",
			Handler:    _SerumHistory_GetFills_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dfuse/solana/serumhist/v1/serumhist.proto",
}
