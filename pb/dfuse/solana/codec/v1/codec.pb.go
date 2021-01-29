// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dfuse/solana/codec/v1/codec.proto

package pbcodec

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Slot struct {
	Id                   string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Number               uint64         `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	PreviousId           string         `protobuf:"bytes,3,opt,name=previous_id,json=previousId,proto3" json:"previous_id,omitempty"`
	Version              uint32         `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	Block                *Block         `protobuf:"bytes,5,opt,name=block,proto3" json:"block,omitempty"`
	Transactions         []*Transaction `protobuf:"bytes,7,rep,name=transactions,proto3" json:"transactions,omitempty"`
	TransactionCount     uint32         `protobuf:"varint,8,opt,name=transaction_count,json=transactionCount,proto3" json:"transaction_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Slot) Reset()         { *m = Slot{} }
func (m *Slot) String() string { return proto.CompactTextString(m) }
func (*Slot) ProtoMessage()    {}
func (*Slot) Descriptor() ([]byte, []int) {
	return fileDescriptor_10cf9f0a2b7e0e4b, []int{0}
}

func (m *Slot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Slot.Unmarshal(m, b)
}
func (m *Slot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Slot.Marshal(b, m, deterministic)
}
func (m *Slot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Slot.Merge(m, src)
}
func (m *Slot) XXX_Size() int {
	return xxx_messageInfo_Slot.Size(m)
}
func (m *Slot) XXX_DiscardUnknown() {
	xxx_messageInfo_Slot.DiscardUnknown(m)
}

var xxx_messageInfo_Slot proto.InternalMessageInfo

func (m *Slot) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Slot) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Slot) GetPreviousId() string {
	if m != nil {
		return m.PreviousId
	}
	return ""
}

func (m *Slot) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Slot) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *Slot) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *Slot) GetTransactionCount() uint32 {
	if m != nil {
		return m.TransactionCount
	}
	return 0
}

type Block struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Number               uint64   `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Height               uint64   `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	PreviousId           string   `protobuf:"bytes,4,opt,name=previous_id,json=previousId,proto3" json:"previous_id,omitempty"`
	PreviousBlockSlot    uint64   `protobuf:"varint,5,opt,name=previous_block_slot,json=previousBlockSlot,proto3" json:"previous_block_slot,omitempty"`
	GenesisUnixTimestamp uint64   `protobuf:"varint,6,opt,name=genesis_unix_timestamp,json=genesisUnixTimestamp,proto3" json:"genesis_unix_timestamp,omitempty"`
	ClockUnixTimestamp   uint64   `protobuf:"varint,7,opt,name=clock_unix_timestamp,json=clockUnixTimestamp,proto3" json:"clock_unix_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_10cf9f0a2b7e0e4b, []int{1}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Block) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Block) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Block) GetPreviousId() string {
	if m != nil {
		return m.PreviousId
	}
	return ""
}

func (m *Block) GetPreviousBlockSlot() uint64 {
	if m != nil {
		return m.PreviousBlockSlot
	}
	return 0
}

func (m *Block) GetGenesisUnixTimestamp() uint64 {
	if m != nil {
		return m.GenesisUnixTimestamp
	}
	return 0
}

func (m *Block) GetClockUnixTimestamp() uint64 {
	if m != nil {
		return m.ClockUnixTimestamp
	}
	return 0
}

type Batch struct {
	Transactions         []*Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Batch) Reset()         { *m = Batch{} }
func (m *Batch) String() string { return proto.CompactTextString(m) }
func (*Batch) ProtoMessage()    {}
func (*Batch) Descriptor() ([]byte, []int) {
	return fileDescriptor_10cf9f0a2b7e0e4b, []int{2}
}

func (m *Batch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Batch.Unmarshal(m, b)
}
func (m *Batch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Batch.Marshal(b, m, deterministic)
}
func (m *Batch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Batch.Merge(m, src)
}
func (m *Batch) XXX_Size() int {
	return xxx_messageInfo_Batch.Size(m)
}
func (m *Batch) XXX_DiscardUnknown() {
	xxx_messageInfo_Batch.DiscardUnknown(m)
}

var xxx_messageInfo_Batch proto.InternalMessageInfo

func (m *Batch) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type Transaction struct {
	// The transaction ID corresponds to the _first_
	// signature. Additional signatures are in `additional_signatures`.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// slot_num could be zero for non-executed transactions
	SlotNum uint64 `protobuf:"varint,2,opt,name=slot_num,json=slotNum,proto3" json:"slot_num,omitempty"`
	// slot_hash could be empty for non-executed transactions
	SlotHash string `protobuf:"bytes,3,opt,name=slot_hash,json=slotHash,proto3" json:"slot_hash,omitempty"`
	// Index from within a single Slot, deterministically ordered to the
	// best of our ability using the transaction ID as a sort key for
	// the batch of transactions executed in parallel.
	Index                uint64         `protobuf:"varint,4,opt,name=index,proto3" json:"index,omitempty"`
	AdditionalSignatures []string       `protobuf:"bytes,5,rep,name=additional_signatures,json=additionalSignatures,proto3" json:"additional_signatures,omitempty"`
	Header               *MessageHeader `protobuf:"bytes,6,opt,name=header,proto3" json:"header,omitempty"`
	// From the original Message object
	AccountKeys []string `protobuf:"bytes,7,rep,name=account_keys,json=accountKeys,proto3" json:"account_keys,omitempty"`
	// From the original Message object
	RecentBlockhash string `protobuf:"bytes,8,opt,name=recent_blockhash,json=recentBlockhash,proto3" json:"recent_blockhash,omitempty"`
	// What follows Once executed these can be set:
	LogMessages []string `protobuf:"bytes,12,rep,name=log_messages,json=logMessages,proto3" json:"log_messages,omitempty"`
	// Instructions, containing both top-level and nested transactions
	Instructions         []*Instruction `protobuf:"bytes,13,rep,name=instructions,proto3" json:"instructions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_10cf9f0a2b7e0e4b, []int{3}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Transaction) GetSlotNum() uint64 {
	if m != nil {
		return m.SlotNum
	}
	return 0
}

func (m *Transaction) GetSlotHash() string {
	if m != nil {
		return m.SlotHash
	}
	return ""
}

func (m *Transaction) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Transaction) GetAdditionalSignatures() []string {
	if m != nil {
		return m.AdditionalSignatures
	}
	return nil
}

func (m *Transaction) GetHeader() *MessageHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Transaction) GetAccountKeys() []string {
	if m != nil {
		return m.AccountKeys
	}
	return nil
}

func (m *Transaction) GetRecentBlockhash() string {
	if m != nil {
		return m.RecentBlockhash
	}
	return ""
}

func (m *Transaction) GetLogMessages() []string {
	if m != nil {
		return m.LogMessages
	}
	return nil
}

func (m *Transaction) GetInstructions() []*Instruction {
	if m != nil {
		return m.Instructions
	}
	return nil
}

type MessageHeader struct {
	NumRequiredSignatures       uint32   `protobuf:"varint,1,opt,name=num_required_signatures,json=numRequiredSignatures,proto3" json:"num_required_signatures,omitempty"`
	NumReadonlySignedAccounts   uint32   `protobuf:"varint,2,opt,name=num_readonly_signed_accounts,json=numReadonlySignedAccounts,proto3" json:"num_readonly_signed_accounts,omitempty"`
	NumReadonlyUnsignedAccounts uint32   `protobuf:"varint,3,opt,name=num_readonly_unsigned_accounts,json=numReadonlyUnsignedAccounts,proto3" json:"num_readonly_unsigned_accounts,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *MessageHeader) Reset()         { *m = MessageHeader{} }
func (m *MessageHeader) String() string { return proto.CompactTextString(m) }
func (*MessageHeader) ProtoMessage()    {}
func (*MessageHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_10cf9f0a2b7e0e4b, []int{4}
}

func (m *MessageHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageHeader.Unmarshal(m, b)
}
func (m *MessageHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageHeader.Marshal(b, m, deterministic)
}
func (m *MessageHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageHeader.Merge(m, src)
}
func (m *MessageHeader) XXX_Size() int {
	return xxx_messageInfo_MessageHeader.Size(m)
}
func (m *MessageHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageHeader.DiscardUnknown(m)
}

var xxx_messageInfo_MessageHeader proto.InternalMessageInfo

func (m *MessageHeader) GetNumRequiredSignatures() uint32 {
	if m != nil {
		return m.NumRequiredSignatures
	}
	return 0
}

func (m *MessageHeader) GetNumReadonlySignedAccounts() uint32 {
	if m != nil {
		return m.NumReadonlySignedAccounts
	}
	return 0
}

func (m *MessageHeader) GetNumReadonlyUnsignedAccounts() uint32 {
	if m != nil {
		return m.NumReadonlyUnsignedAccounts
	}
	return 0
}

type Instruction struct {
	ProgramId            string           `protobuf:"bytes,3,opt,name=program_id,json=programId,proto3" json:"program_id,omitempty"`
	AccountKeys          []string         `protobuf:"bytes,4,rep,name=account_keys,json=accountKeys,proto3" json:"account_keys,omitempty"`
	Data                 []byte           `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	Ordinal              uint32           `protobuf:"varint,6,opt,name=ordinal,proto3" json:"ordinal,omitempty"`
	ParentOrdinal        uint32           `protobuf:"varint,7,opt,name=parent_ordinal,json=parentOrdinal,proto3" json:"parent_ordinal,omitempty"`
	Depth                uint32           `protobuf:"varint,8,opt,name=depth,proto3" json:"depth,omitempty"`
	BalanceChanges       []*BalanceChange `protobuf:"bytes,9,rep,name=balance_changes,json=balanceChanges,proto3" json:"balance_changes,omitempty"`
	AccountChanges       []*AccountChange `protobuf:"bytes,10,rep,name=account_changes,json=accountChanges,proto3" json:"account_changes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Instruction) Reset()         { *m = Instruction{} }
func (m *Instruction) String() string { return proto.CompactTextString(m) }
func (*Instruction) ProtoMessage()    {}
func (*Instruction) Descriptor() ([]byte, []int) {
	return fileDescriptor_10cf9f0a2b7e0e4b, []int{5}
}

func (m *Instruction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instruction.Unmarshal(m, b)
}
func (m *Instruction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instruction.Marshal(b, m, deterministic)
}
func (m *Instruction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instruction.Merge(m, src)
}
func (m *Instruction) XXX_Size() int {
	return xxx_messageInfo_Instruction.Size(m)
}
func (m *Instruction) XXX_DiscardUnknown() {
	xxx_messageInfo_Instruction.DiscardUnknown(m)
}

var xxx_messageInfo_Instruction proto.InternalMessageInfo

func (m *Instruction) GetProgramId() string {
	if m != nil {
		return m.ProgramId
	}
	return ""
}

func (m *Instruction) GetAccountKeys() []string {
	if m != nil {
		return m.AccountKeys
	}
	return nil
}

func (m *Instruction) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Instruction) GetOrdinal() uint32 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

func (m *Instruction) GetParentOrdinal() uint32 {
	if m != nil {
		return m.ParentOrdinal
	}
	return 0
}

func (m *Instruction) GetDepth() uint32 {
	if m != nil {
		return m.Depth
	}
	return 0
}

func (m *Instruction) GetBalanceChanges() []*BalanceChange {
	if m != nil {
		return m.BalanceChanges
	}
	return nil
}

func (m *Instruction) GetAccountChanges() []*AccountChange {
	if m != nil {
		return m.AccountChanges
	}
	return nil
}

type BalanceChange struct {
	Pubkey               string   `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	PrevLamports         uint64   `protobuf:"varint,2,opt,name=prev_lamports,json=prevLamports,proto3" json:"prev_lamports,omitempty"`
	NewLamports          uint64   `protobuf:"varint,3,opt,name=new_lamports,json=newLamports,proto3" json:"new_lamports,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BalanceChange) Reset()         { *m = BalanceChange{} }
func (m *BalanceChange) String() string { return proto.CompactTextString(m) }
func (*BalanceChange) ProtoMessage()    {}
func (*BalanceChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_10cf9f0a2b7e0e4b, []int{6}
}

func (m *BalanceChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceChange.Unmarshal(m, b)
}
func (m *BalanceChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceChange.Marshal(b, m, deterministic)
}
func (m *BalanceChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceChange.Merge(m, src)
}
func (m *BalanceChange) XXX_Size() int {
	return xxx_messageInfo_BalanceChange.Size(m)
}
func (m *BalanceChange) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceChange.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceChange proto.InternalMessageInfo

func (m *BalanceChange) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func (m *BalanceChange) GetPrevLamports() uint64 {
	if m != nil {
		return m.PrevLamports
	}
	return 0
}

func (m *BalanceChange) GetNewLamports() uint64 {
	if m != nil {
		return m.NewLamports
	}
	return 0
}

type AccountChange struct {
	Pubkey               string   `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	PrevData             []byte   `protobuf:"bytes,2,opt,name=prev_data,json=prevData,proto3" json:"prev_data,omitempty"`
	NewData              []byte   `protobuf:"bytes,3,opt,name=new_data,json=newData,proto3" json:"new_data,omitempty"`
	NewDataLength        uint64   `protobuf:"varint,4,opt,name=new_data_length,json=newDataLength,proto3" json:"new_data_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountChange) Reset()         { *m = AccountChange{} }
func (m *AccountChange) String() string { return proto.CompactTextString(m) }
func (*AccountChange) ProtoMessage()    {}
func (*AccountChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_10cf9f0a2b7e0e4b, []int{7}
}

func (m *AccountChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountChange.Unmarshal(m, b)
}
func (m *AccountChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountChange.Marshal(b, m, deterministic)
}
func (m *AccountChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountChange.Merge(m, src)
}
func (m *AccountChange) XXX_Size() int {
	return xxx_messageInfo_AccountChange.Size(m)
}
func (m *AccountChange) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountChange.DiscardUnknown(m)
}

var xxx_messageInfo_AccountChange proto.InternalMessageInfo

func (m *AccountChange) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func (m *AccountChange) GetPrevData() []byte {
	if m != nil {
		return m.PrevData
	}
	return nil
}

func (m *AccountChange) GetNewData() []byte {
	if m != nil {
		return m.NewData
	}
	return nil
}

func (m *AccountChange) GetNewDataLength() uint64 {
	if m != nil {
		return m.NewDataLength
	}
	return 0
}

func init() {
	proto.RegisterType((*Slot)(nil), "dfuse.solana.codec.v1.Slot")
	proto.RegisterType((*Block)(nil), "dfuse.solana.codec.v1.Block")
	proto.RegisterType((*Batch)(nil), "dfuse.solana.codec.v1.Batch")
	proto.RegisterType((*Transaction)(nil), "dfuse.solana.codec.v1.Transaction")
	proto.RegisterType((*MessageHeader)(nil), "dfuse.solana.codec.v1.MessageHeader")
	proto.RegisterType((*Instruction)(nil), "dfuse.solana.codec.v1.Instruction")
	proto.RegisterType((*BalanceChange)(nil), "dfuse.solana.codec.v1.BalanceChange")
	proto.RegisterType((*AccountChange)(nil), "dfuse.solana.codec.v1.AccountChange")
}

func init() {
	proto.RegisterFile("dfuse/solana/codec/v1/codec.proto", fileDescriptor_10cf9f0a2b7e0e4b)
}

var fileDescriptor_10cf9f0a2b7e0e4b = []byte{
	// 874 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcf, 0x8f, 0xdb, 0x44,
	0x14, 0x56, 0x12, 0x67, 0x93, 0xbc, 0xc4, 0xdd, 0x76, 0xc8, 0x2e, 0x5e, 0x6d, 0x81, 0x34, 0xfc,
	0x50, 0x10, 0x22, 0xa1, 0x5b, 0xc4, 0x05, 0x24, 0xd4, 0x2c, 0x42, 0x5d, 0xd1, 0x52, 0x69, 0xb6,
	0xbd, 0x70, 0xb1, 0xc6, 0xf6, 0x60, 0x8f, 0xd6, 0x9e, 0x31, 0x1e, 0x7b, 0x7f, 0xdc, 0xb8, 0x71,
	0xe7, 0xc2, 0x5f, 0xc5, 0x9d, 0x3f, 0x07, 0xcd, 0xf3, 0xb8, 0x71, 0x36, 0xdd, 0x0a, 0x71, 0x9b,
	0xf7, 0xbe, 0xef, 0xbd, 0x99, 0xf7, 0x7d, 0xe3, 0x31, 0x3c, 0x8a, 0x7e, 0xad, 0x34, 0x5f, 0x69,
	0x95, 0x32, 0xc9, 0x56, 0xa1, 0x8a, 0x78, 0xb8, 0xba, 0x7c, 0x5c, 0x2f, 0x96, 0x79, 0xa1, 0x4a,
	0x45, 0x0e, 0x90, 0xb2, 0xac, 0x29, 0xcb, 0x1a, 0xb9, 0x7c, 0x3c, 0xff, 0xb3, 0x0b, 0xce, 0x79,
	0xaa, 0x4a, 0x72, 0x0f, 0xba, 0x22, 0xf2, 0x3a, 0xb3, 0xce, 0x62, 0x44, 0xbb, 0x22, 0x22, 0x87,
	0xb0, 0x27, 0xab, 0x2c, 0xe0, 0x85, 0xd7, 0x9d, 0x75, 0x16, 0x0e, 0xb5, 0x11, 0xf9, 0x08, 0xc6,
	0x79, 0xc1, 0x2f, 0x85, 0xaa, 0xb4, 0x2f, 0x22, 0xaf, 0x87, 0x05, 0xd0, 0xa4, 0xce, 0x22, 0xe2,
	0xc1, 0xe0, 0x92, 0x17, 0x5a, 0x28, 0xe9, 0x39, 0xb3, 0xce, 0xc2, 0xa5, 0x4d, 0x48, 0x4e, 0xa0,
	0x1f, 0xa4, 0x2a, 0xbc, 0xf0, 0xfa, 0xb3, 0xce, 0x62, 0x7c, 0xf2, 0x70, 0xf9, 0xd6, 0x23, 0x2d,
	0xd7, 0x86, 0x43, 0x6b, 0x2a, 0xf9, 0x11, 0x26, 0x65, 0xc1, 0xa4, 0x66, 0x61, 0x29, 0x94, 0xd4,
	0xde, 0x60, 0xd6, 0x5b, 0x8c, 0x4f, 0xe6, 0x77, 0x94, 0xbe, 0xda, 0x50, 0xe9, 0x56, 0x1d, 0xf9,
	0x02, 0x1e, 0xb4, 0x62, 0x3f, 0x54, 0x95, 0x2c, 0xbd, 0x21, 0x9e, 0xef, 0x7e, 0x0b, 0x38, 0x35,
	0xf9, 0xf9, 0xef, 0x5d, 0xe8, 0xe3, 0x29, 0xfe, 0xb3, 0x2a, 0x87, 0xb0, 0x97, 0x70, 0x11, 0x27,
	0x25, 0x0a, 0xe2, 0x50, 0x1b, 0xdd, 0x56, 0xcb, 0xd9, 0x51, 0x6b, 0x09, 0xef, 0xbd, 0x21, 0xe0,
	0xc4, 0xbe, 0x4e, 0x55, 0x89, 0x0a, 0x39, 0xf4, 0x41, 0x03, 0xe1, 0x61, 0xd0, 0xa6, 0xaf, 0xe1,
	0x30, 0xe6, 0x92, 0x6b, 0xa1, 0xfd, 0x4a, 0x8a, 0x6b, 0xbf, 0x14, 0x19, 0xd7, 0x25, 0xcb, 0x72,
	0x6f, 0x0f, 0x4b, 0xa6, 0x16, 0x7d, 0x2d, 0xc5, 0xf5, 0xab, 0x06, 0x23, 0x5f, 0xc1, 0x34, 0xc4,
	0xe6, 0xb7, 0x6a, 0x06, 0x58, 0x43, 0x10, 0xdb, 0xaa, 0x98, 0xbf, 0x84, 0xfe, 0x9a, 0x95, 0x61,
	0xb2, 0x63, 0x40, 0xe7, 0xff, 0x19, 0x30, 0xff, 0xab, 0x07, 0xe3, 0x16, 0xba, 0xa3, 0xec, 0x11,
	0x0c, 0xcd, 0xe4, 0xbe, 0xac, 0x32, 0xab, 0xed, 0xc0, 0xc4, 0x3f, 0x57, 0x19, 0x39, 0x86, 0x11,
	0x42, 0x09, 0xd3, 0x89, 0xbd, 0x70, 0xc8, 0x7d, 0xc6, 0x74, 0x42, 0xa6, 0xd0, 0x17, 0x32, 0xe2,
	0xd7, 0xa8, 0xad, 0x43, 0xeb, 0x80, 0x3c, 0x81, 0x03, 0x16, 0x45, 0xc2, 0xec, 0xc4, 0x52, 0x5f,
	0x8b, 0x58, 0xb2, 0xb2, 0x2a, 0xb8, 0xf6, 0xfa, 0xb3, 0xde, 0x62, 0x44, 0xa7, 0x1b, 0xf0, 0xfc,
	0x0d, 0x46, 0xbe, 0x33, 0x26, 0xb2, 0x88, 0x17, 0xa8, 0xe5, 0xf8, 0xe4, 0x93, 0x3b, 0x86, 0x7c,
	0xc1, 0xb5, 0x66, 0x31, 0x7f, 0x86, 0x5c, 0x6a, 0x6b, 0xc8, 0x23, 0x98, 0xb0, 0x10, 0xef, 0x95,
	0x7f, 0xc1, 0x6f, 0xea, 0x9b, 0x3a, 0xa2, 0x63, 0x9b, 0xfb, 0x89, 0xdf, 0x68, 0xf2, 0x39, 0xdc,
	0x2f, 0x78, 0xc8, 0x65, 0x59, 0x5b, 0x8d, 0xf3, 0x0c, 0x71, 0x9e, 0xfd, 0x3a, 0xbf, 0x6e, 0xd2,
	0xa6, 0x5b, 0xaa, 0x62, 0x3f, 0xab, 0xb7, 0xd2, 0xde, 0xa4, 0xee, 0x96, 0xaa, 0xd8, 0xee, 0xae,
	0x8d, 0x33, 0x42, 0xea, 0xb2, 0xa8, 0xac, 0x33, 0xee, 0x3b, 0x9d, 0x39, 0xdb, 0x50, 0xe9, 0x56,
	0xdd, 0xfc, 0xef, 0x0e, 0xb8, 0x5b, 0x23, 0x91, 0x6f, 0xe0, 0x7d, 0x59, 0x65, 0x7e, 0xc1, 0x7f,
	0xab, 0x44, 0xc1, 0xa3, 0xb6, 0x7e, 0x1d, 0xfc, 0x64, 0x0e, 0x64, 0x95, 0x51, 0x8b, 0xb6, 0x04,
	0xfc, 0x1e, 0x1e, 0xd6, 0x75, 0x2c, 0x52, 0x32, 0xbd, 0xc1, 0x3a, 0x1e, 0xf9, 0x56, 0x02, 0x8d,
	0xbe, 0xba, 0xf4, 0x08, 0x8b, 0x6b, 0xca, 0x39, 0x32, 0x9e, 0x5a, 0x02, 0x39, 0x85, 0x0f, 0xb7,
	0x1a, 0x54, 0xf2, 0x76, 0x8b, 0x1e, 0xb6, 0x38, 0x6e, 0xb5, 0x78, 0x6d, 0x39, 0x4d, 0x93, 0xf9,
	0x3f, 0x5d, 0x18, 0xb7, 0xa6, 0x25, 0x1f, 0x00, 0xe4, 0x85, 0x8a, 0x0b, 0x96, 0x6d, 0x1e, 0xac,
	0x91, 0xcd, 0x9c, 0x45, 0x3b, 0xbe, 0x39, 0xbb, 0xbe, 0x11, 0x70, 0x22, 0x56, 0x32, 0xfc, 0x2a,
	0x27, 0x14, 0xd7, 0xe6, 0x99, 0x53, 0x45, 0x24, 0x24, 0x4b, 0xf1, 0xb6, 0xb8, 0xb4, 0x09, 0xc9,
	0xa7, 0x70, 0x2f, 0x67, 0x85, 0x71, 0xb9, 0x21, 0x0c, 0x90, 0xe0, 0xd6, 0xd9, 0x97, 0x96, 0x36,
	0x85, 0x7e, 0xc4, 0xf3, 0x32, 0xb1, 0xaf, 0x50, 0x1d, 0x90, 0x17, 0xb0, 0x1f, 0xb0, 0x94, 0xc9,
	0x90, 0xfb, 0x61, 0xc2, 0xa4, 0xb1, 0x7e, 0x84, 0xbe, 0xde, 0x75, 0x19, 0xd7, 0x35, 0xfb, 0x14,
	0xc9, 0xf4, 0x5e, 0xd0, 0x0e, 0xb5, 0x69, 0xd7, 0x0c, 0xd7, 0xb4, 0x83, 0x77, 0xb6, 0xb3, 0x2a,
	0x36, 0xed, 0x58, 0x3b, 0xd4, 0x73, 0x05, 0xee, 0xd6, 0x7e, 0xe6, 0xdd, 0xcb, 0xab, 0xe0, 0x82,
	0xdf, 0xd8, 0x2f, 0xd9, 0x46, 0xe4, 0x63, 0x70, 0xcd, 0xdb, 0xe5, 0xa7, 0x2c, 0xcb, 0x55, 0x61,
	0xad, 0x77, 0xe8, 0xc4, 0x24, 0x9f, 0xdb, 0x9c, 0x51, 0x5e, 0xf2, 0xab, 0x0d, 0xa7, 0x7e, 0x3a,
	0xc7, 0x92, 0x5f, 0x35, 0x94, 0xf9, 0x1f, 0x1d, 0x70, 0xb7, 0x8e, 0x74, 0xe7, 0x8e, 0xc7, 0x30,
	0xc2, 0x1d, 0xd1, 0xa8, 0x2e, 0x1a, 0x35, 0x34, 0x89, 0x1f, 0x8c, 0x59, 0x47, 0x30, 0x34, 0x3b,
	0x21, 0xd6, 0x43, 0x6c, 0x20, 0xf9, 0x15, 0x42, 0x9f, 0xc1, 0x7e, 0x03, 0xf9, 0x29, 0x97, 0x71,
	0x99, 0xd8, 0x97, 0xc4, 0xb5, 0x8c, 0xe7, 0x98, 0x5c, 0x9f, 0xfe, 0xf2, 0x34, 0x16, 0x65, 0x52,
	0x05, 0xcb, 0x50, 0x65, 0x2b, 0x14, 0xef, 0x4b, 0xa1, 0xec, 0xc2, 0xfe, 0x78, 0xf3, 0x60, 0xf5,
	0xd6, 0x1f, 0xf1, 0xb7, 0x79, 0x80, 0xcb, 0x60, 0x0f, 0xff, 0xc5, 0x4f, 0xfe, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0x42, 0x9f, 0xef, 0x78, 0xb0, 0x07, 0x00, 0x00,
}
