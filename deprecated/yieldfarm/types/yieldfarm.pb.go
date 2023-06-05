// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: yieldfarm/yieldfarm.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type FarmerInfo struct {
	Account string       `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Amount  []types.Coin `protobuf:"bytes,2,rep,name=amount,proto3" json:"amount"`
	Rewards []types.Coin `protobuf:"bytes,3,rep,name=rewards,proto3" json:"rewards"`
}

func (m *FarmerInfo) Reset()         { *m = FarmerInfo{} }
func (m *FarmerInfo) String() string { return proto.CompactTextString(m) }
func (*FarmerInfo) ProtoMessage()    {}
func (*FarmerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_24371ddc269ed4ea, []int{0}
}
func (m *FarmerInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FarmerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FarmerInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FarmerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FarmerInfo.Merge(m, src)
}
func (m *FarmerInfo) XXX_Size() int {
	return m.Size()
}
func (m *FarmerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FarmerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FarmerInfo proto.InternalMessageInfo

func (m *FarmerInfo) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *FarmerInfo) GetAmount() []types.Coin {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *FarmerInfo) GetRewards() []types.Coin {
	if m != nil {
		return m.Rewards
	}
	return nil
}

func init() {
	proto.RegisterType((*FarmerInfo)(nil), "ununifi.chain.yieldfarm.FarmerInfo")
}

func init() { proto.RegisterFile("yieldfarm/yieldfarm.proto", fileDescriptor_24371ddc269ed4ea) }

var fileDescriptor_24371ddc269ed4ea = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xac, 0xcc, 0x4c, 0xcd,
	0x49, 0x49, 0x4b, 0x2c, 0xca, 0xd5, 0x87, 0xb3, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xc4,
	0x4b, 0xf3, 0x4a, 0xf3, 0x32, 0xd3, 0x32, 0xf5, 0x92, 0x33, 0x12, 0x33, 0xf3, 0xf4, 0xe0, 0xd2,
	0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x35, 0xfa, 0x20, 0x16, 0x44, 0xb9, 0x94, 0x5c, 0x72,
	0x7e, 0x71, 0x6e, 0x7e, 0xb1, 0x7e, 0x52, 0x62, 0x71, 0xaa, 0x7e, 0x99, 0x61, 0x52, 0x6a, 0x49,
	0xa2, 0xa1, 0x7e, 0x72, 0x7e, 0x66, 0x1e, 0x44, 0x5e, 0x69, 0x16, 0x23, 0x17, 0x97, 0x5b, 0x62,
	0x51, 0x6e, 0x6a, 0x91, 0x67, 0x5e, 0x5a, 0xbe, 0x90, 0x04, 0x17, 0x7b, 0x62, 0x72, 0x72, 0x7e,
	0x69, 0x5e, 0x89, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0x2b, 0x64, 0xce, 0xc5, 0x96,
	0x98, 0x0b, 0x96, 0x60, 0x52, 0x60, 0xd6, 0xe0, 0x36, 0x92, 0xd4, 0x83, 0x98, 0xac, 0x07, 0x32,
	0x59, 0x0f, 0x6a, 0xb2, 0x9e, 0x73, 0x7e, 0x66, 0x9e, 0x13, 0xcb, 0x89, 0x7b, 0xf2, 0x0c, 0x41,
	0x50, 0xe5, 0x42, 0x96, 0x5c, 0xec, 0x45, 0xa9, 0xe5, 0x89, 0x45, 0x29, 0xc5, 0x12, 0xcc, 0xc4,
	0xe9, 0x84, 0xa9, 0x77, 0x72, 0x39, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f,
	0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28,
	0xad, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xd0, 0xbc, 0xd0, 0xbc,
	0x4c, 0xb7, 0x4c, 0x7d, 0x70, 0x80, 0xe8, 0x57, 0x20, 0x42, 0x4c, 0xbf, 0xa4, 0xb2, 0x20, 0xb5,
	0x38, 0x89, 0x0d, 0xec, 0x53, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x40, 0x4e, 0x6d, 0xe2,
	0x55, 0x01, 0x00, 0x00,
}

func (m *FarmerInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FarmerInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FarmerInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for iNdEx := len(m.Rewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintYieldfarm(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintYieldfarm(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintYieldfarm(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintYieldfarm(dAtA []byte, offset int, v uint64) int {
	offset -= sovYieldfarm(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FarmerInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovYieldfarm(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovYieldfarm(uint64(l))
		}
	}
	if len(m.Rewards) > 0 {
		for _, e := range m.Rewards {
			l = e.Size()
			n += 1 + l + sovYieldfarm(uint64(l))
		}
	}
	return n
}

func sovYieldfarm(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozYieldfarm(x uint64) (n int) {
	return sovYieldfarm(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FarmerInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowYieldfarm
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
			return fmt.Errorf("proto: FarmerInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FarmerInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowYieldfarm
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
				return ErrInvalidLengthYieldfarm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthYieldfarm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowYieldfarm
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
				return ErrInvalidLengthYieldfarm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthYieldfarm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowYieldfarm
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
				return ErrInvalidLengthYieldfarm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthYieldfarm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rewards = append(m.Rewards, types.Coin{})
			if err := m.Rewards[len(m.Rewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipYieldfarm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthYieldfarm
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
func skipYieldfarm(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowYieldfarm
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
					return 0, ErrIntOverflowYieldfarm
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
					return 0, ErrIntOverflowYieldfarm
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
				return 0, ErrInvalidLengthYieldfarm
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupYieldfarm
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthYieldfarm
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthYieldfarm        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowYieldfarm          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupYieldfarm = fmt.Errorf("proto: unexpected end of group")
)
