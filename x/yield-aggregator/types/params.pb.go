// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: yield-aggregator/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

// Params defines the parameters for the module.
type Params struct {
	CommissionRate       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=commission_rate,json=commissionRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"commission_rate"`
	VaultCreationFee     types.Coin                             `protobuf:"bytes,2,opt,name=vault_creation_fee,json=vaultCreationFee,proto3" json:"vault_creation_fee"`
	VaultCreationDeposit types.Coin                             `protobuf:"bytes,3,opt,name=vault_creation_deposit,json=vaultCreationDeposit,proto3" json:"vault_creation_deposit"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5d4d0a492ea9705, []int{0}
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

func (m *Params) GetVaultCreationFee() types.Coin {
	if m != nil {
		return m.VaultCreationFee
	}
	return types.Coin{}
}

func (m *Params) GetVaultCreationDeposit() types.Coin {
	if m != nil {
		return m.VaultCreationDeposit
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*Params)(nil), "ununifi.chain.yieldaggregator.Params")
}

func init() { proto.RegisterFile("yield-aggregator/params.proto", fileDescriptor_b5d4d0a492ea9705) }

var fileDescriptor_b5d4d0a492ea9705 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xbd, 0x6a, 0x2a, 0x41,
	0x14, 0xde, 0xf5, 0x8a, 0x70, 0xf7, 0xc2, 0x4d, 0x58, 0x24, 0xa8, 0xe0, 0x28, 0x29, 0x82, 0x8d,
	0x33, 0x98, 0x74, 0x21, 0x95, 0x8a, 0x4d, 0x08, 0x04, 0xc1, 0x26, 0x8d, 0xcc, 0x8e, 0xc7, 0x75,
	0x88, 0x3b, 0xb3, 0xec, 0xcc, 0x4a, 0x7c, 0x8b, 0x94, 0x81, 0x34, 0x79, 0x88, 0x3c, 0x84, 0xa5,
	0xa4, 0x0a, 0x29, 0x24, 0xe8, 0x8b, 0x84, 0x9d, 0x1d, 0x30, 0x3f, 0x4d, 0xaa, 0x99, 0xc3, 0xf7,
	0x73, 0x3e, 0xbe, 0xe3, 0xd5, 0x97, 0x1c, 0xe6, 0x93, 0x36, 0x0d, 0xc3, 0x04, 0x42, 0xaa, 0x65,
	0x42, 0x62, 0x9a, 0xd0, 0x48, 0xe1, 0x38, 0x91, 0x5a, 0xfa, 0xf5, 0x54, 0xa4, 0x82, 0x4f, 0x39,
	0x66, 0x33, 0xca, 0x05, 0x36, 0xe4, 0x3d, 0xb7, 0x56, 0x0e, 0x65, 0x28, 0x0d, 0x93, 0x64, 0xbf,
	0x5c, 0x54, 0xab, 0x32, 0xa9, 0x22, 0xa9, 0xc6, 0x39, 0x90, 0x0f, 0x16, 0x42, 0xf9, 0x44, 0x02,
	0xaa, 0x80, 0x2c, 0x3a, 0x01, 0x68, 0xda, 0x21, 0x4c, 0x72, 0x91, 0xe3, 0xc7, 0x8f, 0x05, 0xaf,
	0x74, 0x6d, 0x02, 0xf8, 0xe0, 0x1d, 0x30, 0x19, 0x45, 0x5c, 0x29, 0x2e, 0xc5, 0x38, 0xa1, 0x1a,
	0x2a, 0x6e, 0xd3, 0x6d, 0xfd, 0xed, 0x5e, 0xac, 0x36, 0x0d, 0xe7, 0x6d, 0xd3, 0x38, 0x09, 0xb9,
	0x9e, 0xa5, 0x01, 0x66, 0x32, 0xb2, 0x4b, 0xec, 0xd3, 0x56, 0x93, 0x5b, 0xa2, 0x97, 0x31, 0x28,
	0xdc, 0x07, 0xf6, 0xf2, 0xdc, 0xf6, 0x6c, 0x86, 0x3e, 0xb0, 0xe1, 0xff, 0xbd, 0xe9, 0x90, 0x6a,
	0xf0, 0xaf, 0x3c, 0x7f, 0x41, 0xd3, 0xb9, 0x1e, 0xb3, 0x04, 0xa8, 0xce, 0x56, 0x4d, 0x01, 0x2a,
	0x85, 0xa6, 0xdb, 0xfa, 0x77, 0x5a, 0xc5, 0x56, 0x98, 0xc5, 0xc5, 0x36, 0x2e, 0xee, 0x49, 0x2e,
	0xba, 0xc5, 0x2c, 0xc4, 0xf0, 0xd0, 0x48, 0x7b, 0x56, 0x39, 0x00, 0xf0, 0x47, 0xde, 0xd1, 0x37,
	0xbb, 0x09, 0xc4, 0x52, 0x71, 0x5d, 0xf9, 0xf3, 0x3b, 0xcb, 0xf2, 0x17, 0xcb, 0x7e, 0x2e, 0x3e,
	0x2f, 0x3e, 0x3c, 0x35, 0x9c, 0xee, 0xe5, 0x6a, 0x8b, 0xdc, 0xf5, 0x16, 0xb9, 0xef, 0x5b, 0xe4,
	0xde, 0xef, 0x90, 0xb3, 0xde, 0x21, 0xe7, 0x75, 0x87, 0x9c, 0x9b, 0xce, 0xa7, 0x2e, 0x46, 0x62,
	0x24, 0xf8, 0x80, 0x13, 0x73, 0x32, 0x72, 0x47, 0x7e, 0x5c, 0xd8, 0x54, 0x13, 0x94, 0x4c, 0xe3,
	0x67, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x70, 0x6b, 0xc1, 0x71, 0x02, 0x02, 0x00, 0x00,
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
	{
		size, err := m.VaultCreationDeposit.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.VaultCreationFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.CommissionRate.Size()
		i -= size
		if _, err := m.CommissionRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CommissionRate.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.VaultCreationFee.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.VaultCreationDeposit.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommissionRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommissionRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VaultCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VaultCreationFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VaultCreationDeposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VaultCreationDeposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
