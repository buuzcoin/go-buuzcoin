// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blockchain.proto

package blockchain

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

type Blockchain struct {
	LastBlockHash        []byte   `protobuf:"bytes,1,opt,name=lastBlockHash,proto3" json:"lastBlockHash,omitempty"`
	StateMerkleRoot      []byte   `protobuf:"bytes,2,opt,name=stateMerkleRoot,proto3" json:"stateMerkleRoot,omitempty"`
	LastBlockIndex       uint64   `protobuf:"varint,3,opt,name=lastBlockIndex,proto3" json:"lastBlockIndex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blockchain) Reset()         { *m = Blockchain{} }
func (m *Blockchain) String() string { return proto.CompactTextString(m) }
func (*Blockchain) ProtoMessage()    {}
func (*Blockchain) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{0}
}

func (m *Blockchain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blockchain.Unmarshal(m, b)
}
func (m *Blockchain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blockchain.Marshal(b, m, deterministic)
}
func (m *Blockchain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blockchain.Merge(m, src)
}
func (m *Blockchain) XXX_Size() int {
	return xxx_messageInfo_Blockchain.Size(m)
}
func (m *Blockchain) XXX_DiscardUnknown() {
	xxx_messageInfo_Blockchain.DiscardUnknown(m)
}

var xxx_messageInfo_Blockchain proto.InternalMessageInfo

func (m *Blockchain) GetLastBlockHash() []byte {
	if m != nil {
		return m.LastBlockHash
	}
	return nil
}

func (m *Blockchain) GetStateMerkleRoot() []byte {
	if m != nil {
		return m.StateMerkleRoot
	}
	return nil
}

func (m *Blockchain) GetLastBlockIndex() uint64 {
	if m != nil {
		return m.LastBlockIndex
	}
	return 0
}

type AccountState struct {
	Balance uint64 `protobuf:"varint,1,opt,name=balance,proto3" json:"balance,omitempty"`
	// OutTxCounter is used for checking tx nonce, for valid outgoing transaction
	// tx.nonce == account.outTxCounter + 1, if transaction is accepted counter is increased
	OutTxCounter uint64 `protobuf:"varint,2,opt,name=outTxCounter,proto3" json:"outTxCounter,omitempty"`
	// Reserved for further use (32 bytes)
	Reserved []byte `protobuf:"bytes,3,opt,name=reserved,proto3" json:"reserved,omitempty"`
	// Reserved for further use (32 bytes)
	Reserved2            []byte   `protobuf:"bytes,4,opt,name=reserved2,proto3" json:"reserved2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountState) Reset()         { *m = AccountState{} }
func (m *AccountState) String() string { return proto.CompactTextString(m) }
func (*AccountState) ProtoMessage()    {}
func (*AccountState) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{1}
}

func (m *AccountState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountState.Unmarshal(m, b)
}
func (m *AccountState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountState.Marshal(b, m, deterministic)
}
func (m *AccountState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountState.Merge(m, src)
}
func (m *AccountState) XXX_Size() int {
	return xxx_messageInfo_AccountState.Size(m)
}
func (m *AccountState) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountState.DiscardUnknown(m)
}

var xxx_messageInfo_AccountState proto.InternalMessageInfo

func (m *AccountState) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *AccountState) GetOutTxCounter() uint64 {
	if m != nil {
		return m.OutTxCounter
	}
	return 0
}

func (m *AccountState) GetReserved() []byte {
	if m != nil {
		return m.Reserved
	}
	return nil
}

func (m *AccountState) GetReserved2() []byte {
	if m != nil {
		return m.Reserved2
	}
	return nil
}

type TX struct {
	Version      uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	From         []byte `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	Nonce        uint64 `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	To           []byte `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	Amount       uint64 `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Fee          uint64 `protobuf:"varint,6,opt,name=fee,proto3" json:"fee,omitempty"`
	OptionalData []byte `protobuf:"bytes,7,opt,name=optionalData,proto3" json:"optionalData,omitempty"`
	GasLimit     uint64 `protobuf:"varint,8,opt,name=gasLimit,proto3" json:"gasLimit,omitempty"`
	GasPrice     uint64 `protobuf:"varint,9,opt,name=gasPrice,proto3" json:"gasPrice,omitempty"`
	Hash         []byte `protobuf:"bytes,10,opt,name=hash,proto3" json:"hash,omitempty"`
	// Signature is calculated over hash of fields above, it must identify sender
	Signature            []byte   `protobuf:"bytes,11,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TX) Reset()         { *m = TX{} }
func (m *TX) String() string { return proto.CompactTextString(m) }
func (*TX) ProtoMessage()    {}
func (*TX) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{2}
}

func (m *TX) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TX.Unmarshal(m, b)
}
func (m *TX) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TX.Marshal(b, m, deterministic)
}
func (m *TX) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TX.Merge(m, src)
}
func (m *TX) XXX_Size() int {
	return xxx_messageInfo_TX.Size(m)
}
func (m *TX) XXX_DiscardUnknown() {
	xxx_messageInfo_TX.DiscardUnknown(m)
}

var xxx_messageInfo_TX proto.InternalMessageInfo

func (m *TX) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *TX) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *TX) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *TX) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *TX) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TX) GetFee() uint64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *TX) GetOptionalData() []byte {
	if m != nil {
		return m.OptionalData
	}
	return nil
}

func (m *TX) GetGasLimit() uint64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *TX) GetGasPrice() uint64 {
	if m != nil {
		return m.GasPrice
	}
	return 0
}

func (m *TX) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *TX) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Block struct {
	Version         uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Index           uint64 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Timestamp       int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	PrevBlockHash   []byte `protobuf:"bytes,4,opt,name=prevBlockHash,proto3" json:"prevBlockHash,omitempty"`
	TxMerkleRoot    []byte `protobuf:"bytes,5,opt,name=txMerkleRoot,proto3" json:"txMerkleRoot,omitempty"`
	StateMerkleRoot []byte `protobuf:"bytes,6,opt,name=stateMerkleRoot,proto3" json:"stateMerkleRoot,omitempty"`
	// Reserved for further use (32 bytes)
	Reserved       []byte `protobuf:"bytes,7,opt,name=reserved,proto3" json:"reserved,omitempty"`
	Beneficiary    []byte `protobuf:"bytes,8,opt,name=beneficiary,proto3" json:"beneficiary,omitempty"`
	AdditionalData []byte `protobuf:"bytes,9,opt,name=additionalData,proto3" json:"additionalData,omitempty"`
	// ProofData is binary representation of blockchain's consensus algorithm data
	ProofData []byte `protobuf:"bytes,10,opt,name=proofData,proto3" json:"proofData,omitempty"`
	// Signature is calculated over block header's hash,
	// public key should match beneficiary address
	Signature []byte   `protobuf:"bytes,11,opt,name=signature,proto3" json:"signature,omitempty"`
	TxHashes  [][]byte `protobuf:"bytes,12,rep,name=txHashes,proto3" json:"txHashes,omitempty"`
	// Transactions list is set for full-blocks and is optional for light-blocks
	Transactions         [][]byte `protobuf:"bytes,13,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{3}
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

func (m *Block) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Block) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Block) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Block) GetPrevBlockHash() []byte {
	if m != nil {
		return m.PrevBlockHash
	}
	return nil
}

func (m *Block) GetTxMerkleRoot() []byte {
	if m != nil {
		return m.TxMerkleRoot
	}
	return nil
}

func (m *Block) GetStateMerkleRoot() []byte {
	if m != nil {
		return m.StateMerkleRoot
	}
	return nil
}

func (m *Block) GetReserved() []byte {
	if m != nil {
		return m.Reserved
	}
	return nil
}

func (m *Block) GetBeneficiary() []byte {
	if m != nil {
		return m.Beneficiary
	}
	return nil
}

func (m *Block) GetAdditionalData() []byte {
	if m != nil {
		return m.AdditionalData
	}
	return nil
}

func (m *Block) GetProofData() []byte {
	if m != nil {
		return m.ProofData
	}
	return nil
}

func (m *Block) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Block) GetTxHashes() [][]byte {
	if m != nil {
		return m.TxHashes
	}
	return nil
}

func (m *Block) GetTransactions() [][]byte {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func init() {
	proto.RegisterType((*Blockchain)(nil), "blockchain.Blockchain")
	proto.RegisterType((*AccountState)(nil), "blockchain.AccountState")
	proto.RegisterType((*TX)(nil), "blockchain.TX")
	proto.RegisterType((*Block)(nil), "blockchain.Block")
}

func init() { proto.RegisterFile("blockchain.proto", fileDescriptor_e9ac6287ce250c9a) }

var fileDescriptor_e9ac6287ce250c9a = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0xc7, 0xd5, 0x7c, 0x74, 0xb7, 0xb3, 0xe9, 0xb2, 0xb2, 0x56, 0xc8, 0x42, 0x1c, 0xaa, 0x08,
	0xa1, 0x9e, 0x38, 0xc0, 0x13, 0xf0, 0x71, 0x00, 0x09, 0x24, 0x14, 0xf6, 0xc0, 0xd5, 0x4d, 0xa7,
	0x5b, 0x6b, 0x13, 0x3b, 0xb2, 0xdd, 0xaa, 0xdc, 0x11, 0x67, 0xde, 0x83, 0x97, 0x44, 0x33, 0x69,
	0xf3, 0x51, 0xad, 0xf6, 0xe6, 0xff, 0xcf, 0x13, 0xcf, 0xcc, 0x7f, 0x26, 0x70, 0xb3, 0xaa, 0x6c,
	0xf9, 0x50, 0x6e, 0x95, 0x36, 0x6f, 0x1a, 0x67, 0x83, 0x15, 0xd0, 0x93, 0xfc, 0xf7, 0x04, 0xe0,
	0x43, 0x27, 0xc5, 0x2b, 0x98, 0x57, 0xca, 0x07, 0x26, 0x9f, 0x95, 0xdf, 0xca, 0xc9, 0x62, 0xb2,
	0xcc, 0x8a, 0x31, 0x14, 0x4b, 0x78, 0xe6, 0x83, 0x0a, 0xf8, 0x0d, 0xdd, 0x43, 0x85, 0x85, 0xb5,
	0x41, 0x46, 0x1c, 0x77, 0x8e, 0xc5, 0x6b, 0xb8, 0xee, 0x3e, 0xfd, 0x62, 0xd6, 0x78, 0x90, 0xf1,
	0x62, 0xb2, 0x4c, 0x8a, 0x33, 0x9a, 0xff, 0x99, 0x40, 0xf6, 0xbe, 0x2c, 0xed, 0xce, 0x84, 0x1f,
	0xf4, 0x84, 0x90, 0x70, 0xb1, 0x52, 0x95, 0x32, 0x25, 0x72, 0x09, 0x49, 0x71, 0x92, 0x22, 0x87,
	0xcc, 0xee, 0xc2, 0xdd, 0xe1, 0x23, 0x05, 0xa3, 0xe3, 0xcc, 0x49, 0x31, 0x62, 0xe2, 0x05, 0x5c,
	0x3a, 0xf4, 0xe8, 0xf6, 0xb8, 0xe6, 0x84, 0x59, 0xd1, 0x69, 0xf1, 0x12, 0x66, 0xa7, 0xf3, 0x5b,
	0x99, 0xf0, 0x65, 0x0f, 0xf2, 0xbf, 0x11, 0x44, 0x77, 0x3f, 0x29, 0xfd, 0x1e, 0x9d, 0xd7, 0xd6,
	0x70, 0xfa, 0x79, 0x71, 0x92, 0x42, 0x40, 0xb2, 0x71, 0xb6, 0x3e, 0x36, 0xcc, 0x67, 0x71, 0x0b,
	0xa9, 0xb1, 0x54, 0x6a, 0xdb, 0x5c, 0x2b, 0xc4, 0x35, 0x44, 0xc1, 0x1e, 0x33, 0x44, 0xc1, 0x8a,
	0xe7, 0x30, 0x55, 0x35, 0x15, 0x28, 0x53, 0x0e, 0x3b, 0x2a, 0x71, 0x03, 0xf1, 0x06, 0x51, 0x4e,
	0x19, 0xd2, 0x91, 0x5b, 0x6c, 0x82, 0xb6, 0x46, 0x55, 0x9f, 0x54, 0x50, 0xf2, 0x82, 0xdf, 0x18,
	0x31, 0x6a, 0xf1, 0x5e, 0xf9, 0xaf, 0xba, 0xd6, 0x41, 0x5e, 0xf2, 0xa7, 0x9d, 0x3e, 0xde, 0x7d,
	0x77, 0xba, 0x44, 0x39, 0xeb, 0xee, 0x58, 0x53, 0xfd, 0x5b, 0x1a, 0x2c, 0xb4, 0xf5, 0xd3, 0x99,
	0x2c, 0xf1, 0xfa, 0xde, 0xa8, 0xb0, 0x73, 0x28, 0xaf, 0x5a, 0x4b, 0x3a, 0x90, 0xff, 0x8b, 0x21,
	0xe5, 0x51, 0x3d, 0xe1, 0xca, 0x2d, 0xa4, 0x9a, 0xc7, 0xdb, 0x4e, 0xa3, 0x15, 0xf4, 0x6e, 0xd0,
	0x35, 0xfa, 0xa0, 0xea, 0x86, 0xbd, 0x89, 0x8b, 0x1e, 0xd0, 0xae, 0x35, 0x0e, 0xf7, 0xfd, 0xae,
	0xb5, 0x56, 0x8d, 0x21, 0x79, 0x11, 0x0e, 0x83, 0x45, 0x4b, 0x5b, 0x2f, 0x86, 0xec, 0xb1, 0x7d,
	0x9c, 0x3e, 0xbe, 0x8f, 0xc3, 0xc5, 0xb8, 0x38, 0x5b, 0x8c, 0x05, 0x5c, 0xad, 0xd0, 0xe0, 0x46,
	0x97, 0x5a, 0xb9, 0x5f, 0x6c, 0x6a, 0x56, 0x0c, 0x11, 0x6d, 0xb3, 0x5a, 0xaf, 0xf5, 0x60, 0x32,
	0x33, 0x0e, 0x3a, 0xa3, 0xd4, 0x77, 0xe3, 0xac, 0xdd, 0x70, 0x48, 0x6b, 0x74, 0x0f, 0x9e, 0x76,
	0x9b, 0x2a, 0x0c, 0x07, 0xea, 0x1c, 0xbd, 0xcc, 0x16, 0x31, 0x55, 0x78, 0xd2, 0xec, 0x85, 0x53,
	0xc6, 0xab, 0x92, 0x92, 0x79, 0x39, 0xe7, 0xfb, 0x11, 0x5b, 0x4d, 0xf9, 0x1f, 0x7f, 0xf7, 0x3f,
	0x00, 0x00, 0xff, 0xff, 0x82, 0x99, 0x49, 0x9a, 0xf7, 0x03, 0x00, 0x00,
}
