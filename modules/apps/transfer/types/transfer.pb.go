// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/transfer/v1/transfer.proto

package types

import (
	fmt "fmt"
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

// FungibleTokenPacketData defines a struct for the packet payload
// See FungibleTokenPacketData spec:
// https://github.com/cosmos/ics/tree/master/spec/ics-020-fungible-token-transfer#data-structures
type FungibleTokenPacketData struct {
	// the token denomination to be transferred
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// the token amount to be transferred
	Amount uint64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,string,omitempty"`
	// the sender address
	Sender string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	// the recipient address on the destination chain
	Receiver string `protobuf:"bytes,4,opt,name=receiver,proto3" json:"receiver,omitempty"`
}

func (m *FungibleTokenPacketData) Reset()         { *m = FungibleTokenPacketData{} }
func (m *FungibleTokenPacketData) String() string { return proto.CompactTextString(m) }
func (*FungibleTokenPacketData) ProtoMessage()    {}
func (*FungibleTokenPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_5041673e96e97901, []int{0}
}
func (m *FungibleTokenPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FungibleTokenPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FungibleTokenPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FungibleTokenPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FungibleTokenPacketData.Merge(m, src)
}
func (m *FungibleTokenPacketData) XXX_Size() int {
	return m.Size()
}
func (m *FungibleTokenPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_FungibleTokenPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_FungibleTokenPacketData proto.InternalMessageInfo

func (m *FungibleTokenPacketData) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *FungibleTokenPacketData) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *FungibleTokenPacketData) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *FungibleTokenPacketData) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

// DenomTrace contains the base denomination for ICS20 fungible tokens and the
// source tracing information path.
type DenomTrace struct {
	// path defines the chain of port/channel identifiers used for tracing the
	// source of the fungible token.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// base denomination of the relayed fungible token.
	BaseDenom string `protobuf:"bytes,2,opt,name=base_denom,json=baseDenom,proto3" json:"base_denom,omitempty"`
}

func (m *DenomTrace) Reset()         { *m = DenomTrace{} }
func (m *DenomTrace) String() string { return proto.CompactTextString(m) }
func (*DenomTrace) ProtoMessage()    {}
func (*DenomTrace) Descriptor() ([]byte, []int) {
	return fileDescriptor_5041673e96e97901, []int{1}
}
func (m *DenomTrace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DenomTrace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DenomTrace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DenomTrace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DenomTrace.Merge(m, src)
}
func (m *DenomTrace) XXX_Size() int {
	return m.Size()
}
func (m *DenomTrace) XXX_DiscardUnknown() {
	xxx_messageInfo_DenomTrace.DiscardUnknown(m)
}

var xxx_messageInfo_DenomTrace proto.InternalMessageInfo

func (m *DenomTrace) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *DenomTrace) GetBaseDenom() string {
	if m != nil {
		return m.BaseDenom
	}
	return ""
}

// Params defines the set of IBC transfer parameters.
// NOTE: To prevent a single token from being transferred, set the
// TransfersEnabled parameter to true and then set the bank module's SendEnabled
// parameter for the denomination to false.
type Params struct {
	// send_enabled enables or disables all cross-chain token transfers from this
	// chain.
	SendEnabled bool `protobuf:"varint,1,opt,name=send_enabled,json=sendEnabled,proto3" json:"send_enabled,omitempty" yaml:"send_enabled"`
	// receive_enabled enables or disables all cross-chain token transfers to this
	// chain.
	ReceiveEnabled bool `protobuf:"varint,2,opt,name=receive_enabled,json=receiveEnabled,proto3" json:"receive_enabled,omitempty" yaml:"receive_enabled"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_5041673e96e97901, []int{2}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetSendEnabled() bool {
	if m != nil {
		return m.SendEnabled
	}
	return false
}

func (m *Params) GetReceiveEnabled() bool {
	if m != nil {
		return m.ReceiveEnabled
	}
	return false
}

func init() {
	proto.RegisterType((*FungibleTokenPacketData)(nil), "ibc.applications.transfer.v1.FungibleTokenPacketData")
	proto.RegisterType((*DenomTrace)(nil), "ibc.applications.transfer.v1.DenomTrace")
	proto.RegisterType((*Params)(nil), "ibc.applications.transfer.v1.Params")
}

func init() {
	proto.RegisterFile("ibc/applications/transfer/v1/transfer.proto", fileDescriptor_5041673e96e97901)
}

var fileDescriptor_5041673e96e97901 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xc1, 0xaa, 0xda, 0x40,
	0x14, 0x86, 0x8d, 0xb5, 0xa2, 0xd3, 0xd2, 0xc2, 0x54, 0x34, 0x48, 0x1b, 0x25, 0x2b, 0xa1, 0x34,
	0x41, 0xda, 0x95, 0x9b, 0x82, 0xb5, 0x5d, 0x16, 0x09, 0xae, 0xba, 0x91, 0x99, 0xc9, 0x69, 0x1c,
	0x4c, 0x66, 0xc2, 0xcc, 0x44, 0x90, 0x3e, 0x41, 0x77, 0x7d, 0xac, 0x2e, 0x5d, 0xde, 0x95, 0x5c,
	0xf4, 0x0d, 0x7c, 0x82, 0x4b, 0x26, 0x21, 0xc8, 0xdd, 0x9d, 0xef, 0x9c, 0xff, 0x3f, 0xe7, 0xc0,
	0x8f, 0x3e, 0x72, 0xca, 0x42, 0x92, 0xe7, 0x29, 0x67, 0xc4, 0x70, 0x29, 0x74, 0x68, 0x14, 0x11,
	0xfa, 0x37, 0xa8, 0xf0, 0x30, 0x6f, 0xea, 0x20, 0x57, 0xd2, 0x48, 0xfc, 0x9e, 0x53, 0x16, 0xdc,
	0x8b, 0x83, 0x46, 0x70, 0x98, 0x8f, 0x07, 0x89, 0x4c, 0xa4, 0x15, 0x86, 0x65, 0x55, 0x79, 0xfc,
	0x3f, 0x68, 0xf4, 0xa3, 0x10, 0x09, 0xa7, 0x29, 0x6c, 0xe4, 0x1e, 0xc4, 0x9a, 0xb0, 0x3d, 0x98,
	0x15, 0x31, 0x04, 0x0f, 0xd0, 0xcb, 0x18, 0x84, 0xcc, 0x5c, 0x67, 0xea, 0xcc, 0xfa, 0x51, 0x05,
	0x78, 0x88, 0xba, 0x24, 0x93, 0x85, 0x30, 0x6e, 0x7b, 0xea, 0xcc, 0x3a, 0x51, 0x4d, 0x65, 0x5f,
	0x83, 0x88, 0x41, 0xb9, 0x2f, 0xac, 0xbc, 0x26, 0x3c, 0x46, 0x3d, 0x05, 0x0c, 0xf8, 0x01, 0x94,
	0xdb, 0xb1, 0x93, 0x86, 0xfd, 0xaf, 0x08, 0xad, 0xca, 0xa5, 0x1b, 0x45, 0x18, 0x60, 0x8c, 0x3a,
	0x39, 0x31, 0xbb, 0xfa, 0x9c, 0xad, 0xf1, 0x07, 0x84, 0x28, 0xd1, 0xb0, 0xad, 0x1e, 0x69, 0xdb,
	0x49, 0xbf, 0xec, 0x58, 0x9f, 0xff, 0xd7, 0x41, 0xdd, 0x35, 0x51, 0x24, 0xd3, 0x78, 0x81, 0x5e,
	0x97, 0x17, 0xb7, 0x20, 0x08, 0x4d, 0x21, 0xb6, 0x5b, 0x7a, 0xcb, 0xd1, 0xed, 0x3c, 0x79, 0x77,
	0x24, 0x59, 0xba, 0xf0, 0xef, 0xa7, 0x7e, 0xf4, 0xaa, 0xc4, 0xef, 0x15, 0xe1, 0x6f, 0xe8, 0x6d,
	0xfd, 0x53, 0x63, 0x6f, 0x5b, 0xfb, 0xf8, 0x76, 0x9e, 0x0c, 0x2b, 0xfb, 0x33, 0x81, 0x1f, 0xbd,
	0xa9, 0x3b, 0xf5, 0x92, 0xe5, 0xcf, 0xff, 0x17, 0xcf, 0x39, 0x5d, 0x3c, 0xe7, 0xf1, 0xe2, 0x39,
	0xff, 0xae, 0x5e, 0xeb, 0x74, 0xf5, 0x5a, 0x0f, 0x57, 0xaf, 0xf5, 0xeb, 0x4b, 0xc2, 0xcd, 0xae,
	0xa0, 0x01, 0x93, 0x59, 0xc8, 0xa4, 0xce, 0xa4, 0x0e, 0x39, 0x65, 0x9f, 0x12, 0x19, 0x66, 0x32,
	0x2e, 0x52, 0xd0, 0x65, 0xc2, 0x77, 0xc9, 0x9a, 0x63, 0x0e, 0x9a, 0x76, 0x6d, 0x40, 0x9f, 0x9f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xbc, 0x5f, 0xc9, 0x03, 0x02, 0x00, 0x00,
}

func (m *FungibleTokenPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FungibleTokenPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FungibleTokenPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintTransfer(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTransfer(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Amount != 0 {
		i = encodeVarintTransfer(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTransfer(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DenomTrace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DenomTrace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DenomTrace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BaseDenom) > 0 {
		i -= len(m.BaseDenom)
		copy(dAtA[i:], m.BaseDenom)
		i = encodeVarintTransfer(dAtA, i, uint64(len(m.BaseDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintTransfer(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ReceiveEnabled {
		i--
		if m.ReceiveEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.SendEnabled {
		i--
		if m.SendEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTransfer(dAtA []byte, offset int, v uint64) int {
	offset -= sovTransfer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FungibleTokenPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTransfer(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovTransfer(uint64(m.Amount))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTransfer(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovTransfer(uint64(l))
	}
	return n
}

func (m *DenomTrace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovTransfer(uint64(l))
	}
	l = len(m.BaseDenom)
	if l > 0 {
		n += 1 + l + sovTransfer(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SendEnabled {
		n += 2
	}
	if m.ReceiveEnabled {
		n += 2
	}
	return n
}

func sovTransfer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTransfer(x uint64) (n int) {
	return sovTransfer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FungibleTokenPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransfer
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
			return fmt.Errorf("proto: FungibleTokenPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FungibleTokenPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
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
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
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
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
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
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransfer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTransfer
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
func (m *DenomTrace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransfer
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
			return fmt.Errorf("proto: DenomTrace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DenomTrace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
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
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
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
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransfer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTransfer
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransfer
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SendEnabled = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceiveEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ReceiveEnabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTransfer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTransfer
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
func skipTransfer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTransfer
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
					return 0, ErrIntOverflowTransfer
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
					return 0, ErrIntOverflowTransfer
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
				return 0, ErrInvalidLengthTransfer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTransfer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTransfer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTransfer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTransfer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTransfer = fmt.Errorf("proto: unexpected end of group")
)
