// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sifnode/dispensation/v1/types.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Distribution type enum
type DistributionType int32

const (
	// Unspecified distribution type
	DistributionType_DISTRIBUTION_TYPE_UNSPECIFIED DistributionType = 0
	// Airdrop distribution type
	DistributionType_DISTRIBUTION_TYPE_AIRDROP DistributionType = 1
)

var DistributionType_name = map[int32]string{
	0: "DISTRIBUTION_TYPE_UNSPECIFIED",
	1: "DISTRIBUTION_TYPE_AIRDROP",
}

var DistributionType_value = map[string]int32{
	"DISTRIBUTION_TYPE_UNSPECIFIED": 0,
	"DISTRIBUTION_TYPE_AIRDROP":     1,
}

func (x DistributionType) String() string {
	return proto.EnumName(DistributionType_name, int32(x))
}

func (DistributionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bfdf912039cd8799, []int{0}
}

// Claim status enum
type DistributionStatus int32

const (
	// Unspecified
	DistributionStatus_DISTRIBUTION_STATUS_UNSPECIFIED DistributionStatus = 0
	// Pending claim status
	DistributionStatus_DISTRIBUTION_STATUS_PENDING DistributionStatus = 1
	// Completed claim status
	DistributionStatus_DISTRIBUTION_STATUS_COMPLETED DistributionStatus = 2
)

var DistributionStatus_name = map[int32]string{
	0: "DISTRIBUTION_STATUS_UNSPECIFIED",
	1: "DISTRIBUTION_STATUS_PENDING",
	2: "DISTRIBUTION_STATUS_COMPLETED",
}

var DistributionStatus_value = map[string]int32{
	"DISTRIBUTION_STATUS_UNSPECIFIED": 0,
	"DISTRIBUTION_STATUS_PENDING":     1,
	"DISTRIBUTION_STATUS_COMPLETED":   2,
}

func (x DistributionStatus) String() string {
	return proto.EnumName(DistributionStatus_name, int32(x))
}

func (DistributionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bfdf912039cd8799, []int{1}
}

type GenesisState struct {
	DistributionRecords *DistributionRecords `protobuf:"bytes,1,opt,name=distribution_records,json=distributionRecords,proto3" json:"distribution_records,omitempty"`
	Distributions       *Distributions       `protobuf:"bytes,2,opt,name=distributions,proto3" json:"distributions,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfdf912039cd8799, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetDistributionRecords() *DistributionRecords {
	if m != nil {
		return m.DistributionRecords
	}
	return nil
}

func (m *GenesisState) GetDistributions() *Distributions {
	if m != nil {
		return m.Distributions
	}
	return nil
}

type DistributionRecord struct {
	DistributionStatus          DistributionStatus                       `protobuf:"varint,1,opt,name=distribution_status,json=distributionStatus,proto3,enum=sifnode.dispensation.v1.DistributionStatus" json:"distribution_status,omitempty"`
	DistributionName            string                                   `protobuf:"bytes,2,opt,name=distribution_name,json=distributionName,proto3" json:"distribution_name,omitempty"`
	RecipientAddress            string                                   `protobuf:"bytes,3,opt,name=recipient_address,json=recipientAddress,proto3" json:"recipient_address,omitempty"`
	Coins                       github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins" yaml:"coins"`
	DistributionStartHeight     int64                                    `protobuf:"varint,5,opt,name=distribution_start_height,json=distributionStartHeight,proto3" json:"distribution_start_height,omitempty"`
	DistributionCompletedHeight int64                                    `protobuf:"varint,6,opt,name=distribution_completed_height,json=distributionCompletedHeight,proto3" json:"distribution_completed_height,omitempty"`
}

func (m *DistributionRecord) Reset()         { *m = DistributionRecord{} }
func (m *DistributionRecord) String() string { return proto.CompactTextString(m) }
func (*DistributionRecord) ProtoMessage()    {}
func (*DistributionRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfdf912039cd8799, []int{1}
}
func (m *DistributionRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DistributionRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DistributionRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DistributionRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributionRecord.Merge(m, src)
}
func (m *DistributionRecord) XXX_Size() int {
	return m.Size()
}
func (m *DistributionRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributionRecord.DiscardUnknown(m)
}

var xxx_messageInfo_DistributionRecord proto.InternalMessageInfo

func (m *DistributionRecord) GetDistributionStatus() DistributionStatus {
	if m != nil {
		return m.DistributionStatus
	}
	return DistributionStatus_DISTRIBUTION_STATUS_UNSPECIFIED
}

func (m *DistributionRecord) GetDistributionName() string {
	if m != nil {
		return m.DistributionName
	}
	return ""
}

func (m *DistributionRecord) GetRecipientAddress() string {
	if m != nil {
		return m.RecipientAddress
	}
	return ""
}

func (m *DistributionRecord) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

func (m *DistributionRecord) GetDistributionStartHeight() int64 {
	if m != nil {
		return m.DistributionStartHeight
	}
	return 0
}

func (m *DistributionRecord) GetDistributionCompletedHeight() int64 {
	if m != nil {
		return m.DistributionCompletedHeight
	}
	return 0
}

type DistributionRecords struct {
	DistributionRecords []*DistributionRecord `protobuf:"bytes,1,rep,name=distribution_records,json=distributionRecords,proto3" json:"distribution_records,omitempty"`
}

func (m *DistributionRecords) Reset()         { *m = DistributionRecords{} }
func (m *DistributionRecords) String() string { return proto.CompactTextString(m) }
func (*DistributionRecords) ProtoMessage()    {}
func (*DistributionRecords) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfdf912039cd8799, []int{2}
}
func (m *DistributionRecords) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DistributionRecords) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DistributionRecords.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DistributionRecords) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributionRecords.Merge(m, src)
}
func (m *DistributionRecords) XXX_Size() int {
	return m.Size()
}
func (m *DistributionRecords) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributionRecords.DiscardUnknown(m)
}

var xxx_messageInfo_DistributionRecords proto.InternalMessageInfo

func (m *DistributionRecords) GetDistributionRecords() []*DistributionRecord {
	if m != nil {
		return m.DistributionRecords
	}
	return nil
}

type Distributions struct {
	Distributions []*Distribution `protobuf:"bytes,1,rep,name=distributions,proto3" json:"distributions,omitempty"`
}

func (m *Distributions) Reset()         { *m = Distributions{} }
func (m *Distributions) String() string { return proto.CompactTextString(m) }
func (*Distributions) ProtoMessage()    {}
func (*Distributions) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfdf912039cd8799, []int{3}
}
func (m *Distributions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Distributions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Distributions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Distributions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Distributions.Merge(m, src)
}
func (m *Distributions) XXX_Size() int {
	return m.Size()
}
func (m *Distributions) XXX_DiscardUnknown() {
	xxx_messageInfo_Distributions.DiscardUnknown(m)
}

var xxx_messageInfo_Distributions proto.InternalMessageInfo

func (m *Distributions) GetDistributions() []*Distribution {
	if m != nil {
		return m.Distributions
	}
	return nil
}

type Distribution struct {
	DistributionType DistributionType `protobuf:"varint,1,opt,name=distribution_type,json=distributionType,proto3,enum=sifnode.dispensation.v1.DistributionType" json:"distribution_type,omitempty"`
	DistributionName string           `protobuf:"bytes,2,opt,name=distribution_name,json=distributionName,proto3" json:"distribution_name,omitempty"`
}

func (m *Distribution) Reset()         { *m = Distribution{} }
func (m *Distribution) String() string { return proto.CompactTextString(m) }
func (*Distribution) ProtoMessage()    {}
func (*Distribution) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfdf912039cd8799, []int{4}
}
func (m *Distribution) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Distribution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Distribution.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Distribution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Distribution.Merge(m, src)
}
func (m *Distribution) XXX_Size() int {
	return m.Size()
}
func (m *Distribution) XXX_DiscardUnknown() {
	xxx_messageInfo_Distribution.DiscardUnknown(m)
}

var xxx_messageInfo_Distribution proto.InternalMessageInfo

func (m *Distribution) GetDistributionType() DistributionType {
	if m != nil {
		return m.DistributionType
	}
	return DistributionType_DISTRIBUTION_TYPE_UNSPECIFIED
}

func (m *Distribution) GetDistributionName() string {
	if m != nil {
		return m.DistributionName
	}
	return ""
}

func init() {
	proto.RegisterEnum("sifnode.dispensation.v1.DistributionType", DistributionType_name, DistributionType_value)
	proto.RegisterEnum("sifnode.dispensation.v1.DistributionStatus", DistributionStatus_name, DistributionStatus_value)
	proto.RegisterType((*GenesisState)(nil), "sifnode.dispensation.v1.GenesisState")
	proto.RegisterType((*DistributionRecord)(nil), "sifnode.dispensation.v1.DistributionRecord")
	proto.RegisterType((*DistributionRecords)(nil), "sifnode.dispensation.v1.DistributionRecords")
	proto.RegisterType((*Distributions)(nil), "sifnode.dispensation.v1.Distributions")
	proto.RegisterType((*Distribution)(nil), "sifnode.dispensation.v1.Distribution")
}

func init() {
	proto.RegisterFile("sifnode/dispensation/v1/types.proto", fileDescriptor_bfdf912039cd8799)
}

var fileDescriptor_bfdf912039cd8799 = []byte{
	// 621 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdf, 0x4e, 0xd3, 0x50,
	0x18, 0x5f, 0x99, 0x90, 0x78, 0x00, 0x53, 0x0f, 0x44, 0x06, 0x84, 0x0e, 0x4b, 0x34, 0x13, 0xb4,
	0x75, 0x78, 0xc7, 0x95, 0x6c, 0xad, 0x58, 0xc5, 0xb1, 0xb4, 0xc5, 0x44, 0x43, 0x6c, 0xba, 0xf6,
	0xb0, 0x9d, 0x48, 0x7b, 0x6a, 0xcf, 0x19, 0x91, 0x0b, 0xdf, 0xc1, 0xc4, 0xb7, 0xf0, 0x1d, 0x8c,
	0xb7, 0x5c, 0x72, 0xe9, 0x15, 0x1a, 0xf6, 0x06, 0x3e, 0x81, 0xe9, 0x9f, 0x99, 0xb6, 0x1b, 0xc9,
	0xb8, 0xda, 0xf2, 0x7d, 0xbf, 0x3f, 0x3d, 0xe7, 0xfb, 0x7d, 0x07, 0x6c, 0x50, 0x7c, 0xec, 0x13,
	0x17, 0xc9, 0x2e, 0xa6, 0x01, 0xf2, 0xa9, 0xcd, 0x30, 0xf1, 0xe5, 0xd3, 0xba, 0xcc, 0xce, 0x02,
	0x44, 0xa5, 0x20, 0x24, 0x8c, 0xc0, 0xa5, 0x14, 0x24, 0x65, 0x41, 0xd2, 0x69, 0x7d, 0x65, 0xb1,
	0x4b, 0xba, 0x24, 0xc6, 0xc8, 0xd1, 0xbf, 0x04, 0xbe, 0x72, 0xcf, 0x21, 0xd4, 0x23, 0x54, 0xee,
	0xd8, 0x14, 0xc9, 0x0e, 0xc1, 0x7e, 0x52, 0x17, 0x7f, 0x70, 0x60, 0x6e, 0x0f, 0xf9, 0x88, 0x62,
	0x6a, 0x30, 0x9b, 0x21, 0x68, 0x81, 0x45, 0x17, 0x53, 0x16, 0xe2, 0x4e, 0x3f, 0x52, 0xb4, 0x42,
	0xe4, 0x90, 0xd0, 0xa5, 0x15, 0x6e, 0x9d, 0xab, 0xcd, 0x6e, 0x3f, 0x96, 0xae, 0xb1, 0x95, 0x94,
	0x0c, 0x49, 0x4f, 0x38, 0xfa, 0x82, 0x3b, 0x5a, 0x84, 0xfb, 0x60, 0x3e, 0x5b, 0xa6, 0x95, 0xa9,
	0x58, 0xf9, 0xe1, 0x44, 0xca, 0x54, 0xcf, 0x93, 0xc5, 0x9f, 0x65, 0x00, 0x47, 0xad, 0xe1, 0x11,
	0xc8, 0x79, 0x5b, 0x94, 0xd9, 0xac, 0x9f, 0x1c, 0xe2, 0xce, 0xf6, 0xd6, 0x44, 0x56, 0x46, 0x4c,
	0xd1, 0xa1, 0x3b, 0x52, 0x83, 0x5b, 0xe0, 0x6e, 0x4e, 0xdd, 0xb7, 0x3d, 0x14, 0x1f, 0xe3, 0xb6,
	0xce, 0x67, 0x1b, 0x2d, 0xdb, 0x43, 0x11, 0x38, 0x44, 0x0e, 0x0e, 0x30, 0xf2, 0x99, 0x65, 0xbb,
	0x6e, 0x88, 0x28, 0xad, 0x94, 0x13, 0xf0, 0xff, 0xc6, 0x6e, 0x52, 0x87, 0x9f, 0xc0, 0x74, 0x34,
	0x1c, 0x5a, 0xb9, 0xb5, 0x5e, 0xae, 0xcd, 0x6e, 0x2f, 0x4b, 0xc9, 0xd8, 0xa4, 0x68, 0x6c, 0xd2,
	0x69, 0xbd, 0x83, 0x98, 0x5d, 0x97, 0x9a, 0x04, 0xfb, 0x8d, 0xe7, 0xe7, 0x97, 0xd5, 0xd2, 0xdf,
	0xcb, 0xea, 0xdc, 0x99, 0xed, 0x9d, 0xec, 0x88, 0x31, 0x4b, 0xfc, 0xfe, 0xbb, 0x5a, 0xeb, 0x62,
	0xd6, 0xeb, 0x77, 0x24, 0x87, 0x78, 0x72, 0x3a, 0xf3, 0xe4, 0xe7, 0x09, 0x75, 0x3f, 0xa6, 0x09,
	0x8a, 0x04, 0xa8, 0x9e, 0x38, 0xc1, 0x1d, 0xb0, 0x5c, 0xbc, 0xaa, 0x90, 0x59, 0x3d, 0x84, 0xbb,
	0x3d, 0x56, 0x99, 0x5e, 0xe7, 0x6a, 0x65, 0x7d, 0xa9, 0x70, 0x07, 0x21, 0x7b, 0x19, 0xb7, 0x61,
	0x03, 0xac, 0xe5, 0xb8, 0x0e, 0xf1, 0x82, 0x13, 0xc4, 0x90, 0x3b, 0xe4, 0xcf, 0xc4, 0xfc, 0xd5,
	0x2c, 0xa8, 0x39, 0xc4, 0x24, 0x1a, 0x62, 0x1f, 0x2c, 0x8c, 0xc9, 0x0e, 0xfc, 0x70, 0x6d, 0x0e,
	0xa3, 0x8b, 0xd9, 0xba, 0x41, 0x0e, 0xc7, 0xc6, 0x50, 0x3c, 0x02, 0xf3, 0xb9, 0x60, 0xc1, 0xd7,
	0xc5, 0x5c, 0x26, 0x4e, 0x0f, 0x26, 0x73, 0x2a, 0xc4, 0xf2, 0x1b, 0x07, 0xe6, 0xb2, 0x7d, 0xf8,
	0xb6, 0x10, 0x99, 0x68, 0x10, 0x69, 0x1c, 0x1f, 0x4d, 0xe4, 0x60, 0x9e, 0x05, 0x28, 0x9f, 0xae,
	0xa8, 0x72, 0xa3, 0x28, 0x6e, 0x9a, 0x80, 0x2f, 0x4a, 0xc2, 0xfb, 0x60, 0x4d, 0xd1, 0x0c, 0x53,
	0xd7, 0x1a, 0x87, 0xa6, 0x76, 0xd0, 0xb2, 0xcc, 0x77, 0x6d, 0xd5, 0x3a, 0x6c, 0x19, 0x6d, 0xb5,
	0xa9, 0xbd, 0xd0, 0x54, 0x85, 0x2f, 0xc1, 0x35, 0xb0, 0x3c, 0x0a, 0xd9, 0xd5, 0x74, 0x45, 0x3f,
	0x68, 0xf3, 0xdc, 0xe6, 0x97, 0xfc, 0x06, 0xa6, 0x3b, 0xb2, 0x01, 0xaa, 0x39, 0x92, 0x61, 0xee,
	0x9a, 0x87, 0x46, 0x41, 0xb9, 0x0a, 0x56, 0xc7, 0x81, 0xda, 0x6a, 0x4b, 0xd1, 0x5a, 0x7b, 0x3c,
	0x37, 0xf2, 0x75, 0x29, 0xa0, 0x79, 0xf0, 0xa6, 0xbd, 0xaf, 0x9a, 0xaa, 0xc2, 0x4f, 0x35, 0x5e,
	0x9d, 0x5f, 0x09, 0xdc, 0xc5, 0x95, 0xc0, 0xfd, 0xb9, 0x12, 0xb8, 0xaf, 0x03, 0xa1, 0x74, 0x31,
	0x10, 0x4a, 0xbf, 0x06, 0x42, 0xe9, 0xfd, 0xd3, 0xcc, 0x2a, 0x18, 0xf8, 0xd8, 0xe9, 0xd9, 0xd8,
	0x97, 0x87, 0x6f, 0xeb, 0xe7, 0xfc, 0xeb, 0x1a, 0x2f, 0x46, 0x67, 0x26, 0x7e, 0x14, 0x9f, 0xfd,
	0x0b, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x8a, 0xdb, 0xee, 0x82, 0x05, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Distributions != nil {
		{
			size, err := m.Distributions.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.DistributionRecords != nil {
		{
			size, err := m.DistributionRecords.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DistributionRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DistributionRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DistributionRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DistributionCompletedHeight != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.DistributionCompletedHeight))
		i--
		dAtA[i] = 0x30
	}
	if m.DistributionStartHeight != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.DistributionStartHeight))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.RecipientAddress) > 0 {
		i -= len(m.RecipientAddress)
		copy(dAtA[i:], m.RecipientAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.RecipientAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DistributionName) > 0 {
		i -= len(m.DistributionName)
		copy(dAtA[i:], m.DistributionName)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.DistributionName)))
		i--
		dAtA[i] = 0x12
	}
	if m.DistributionStatus != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.DistributionStatus))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DistributionRecords) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DistributionRecords) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DistributionRecords) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DistributionRecords) > 0 {
		for iNdEx := len(m.DistributionRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DistributionRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Distributions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Distributions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Distributions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Distributions) > 0 {
		for iNdEx := len(m.Distributions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Distributions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Distribution) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Distribution) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Distribution) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DistributionName) > 0 {
		i -= len(m.DistributionName)
		copy(dAtA[i:], m.DistributionName)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.DistributionName)))
		i--
		dAtA[i] = 0x12
	}
	if m.DistributionType != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.DistributionType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DistributionRecords != nil {
		l = m.DistributionRecords.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Distributions != nil {
		l = m.Distributions.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *DistributionRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DistributionStatus != 0 {
		n += 1 + sovTypes(uint64(m.DistributionStatus))
	}
	l = len(m.DistributionName)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.RecipientAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.DistributionStartHeight != 0 {
		n += 1 + sovTypes(uint64(m.DistributionStartHeight))
	}
	if m.DistributionCompletedHeight != 0 {
		n += 1 + sovTypes(uint64(m.DistributionCompletedHeight))
	}
	return n
}

func (m *DistributionRecords) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DistributionRecords) > 0 {
		for _, e := range m.DistributionRecords {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *Distributions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Distributions) > 0 {
		for _, e := range m.Distributions {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *Distribution) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DistributionType != 0 {
		n += 1 + sovTypes(uint64(m.DistributionType))
	}
	l = len(m.DistributionName)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DistributionRecords == nil {
				m.DistributionRecords = &DistributionRecords{}
			}
			if err := m.DistributionRecords.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Distributions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Distributions == nil {
				m.Distributions = &Distributions{}
			}
			if err := m.Distributions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DistributionRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DistributionRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DistributionRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionStatus", wireType)
			}
			m.DistributionStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DistributionStatus |= DistributionStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DistributionName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecipientAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecipientAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionStartHeight", wireType)
			}
			m.DistributionStartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DistributionStartHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionCompletedHeight", wireType)
			}
			m.DistributionCompletedHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DistributionCompletedHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DistributionRecords) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DistributionRecords: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DistributionRecords: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DistributionRecords = append(m.DistributionRecords, &DistributionRecord{})
			if err := m.DistributionRecords[len(m.DistributionRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Distributions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Distributions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Distributions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Distributions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Distributions = append(m.Distributions, &Distribution{})
			if err := m.Distributions[len(m.Distributions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Distribution) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Distribution: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Distribution: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionType", wireType)
			}
			m.DistributionType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DistributionType |= DistributionType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DistributionName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
