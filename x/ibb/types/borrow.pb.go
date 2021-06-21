// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibb/borrow.proto

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

type Borrow struct {
	Creator     string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id          uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	BlockHeight int32  `protobuf:"varint,3,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	Asset       string `protobuf:"bytes,4,opt,name=asset,proto3" json:"asset,omitempty"`
	Amount      int32  `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Denom       string `protobuf:"bytes,6,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *Borrow) Reset()         { *m = Borrow{} }
func (m *Borrow) String() string { return proto.CompactTextString(m) }
func (*Borrow) ProtoMessage()    {}
func (*Borrow) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fcfa17d999a14cb, []int{0}
}
func (m *Borrow) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Borrow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Borrow.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Borrow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Borrow.Merge(m, src)
}
func (m *Borrow) XXX_Size() int {
	return m.Size()
}
func (m *Borrow) XXX_DiscardUnknown() {
	xxx_messageInfo_Borrow.DiscardUnknown(m)
}

var xxx_messageInfo_Borrow proto.InternalMessageInfo

func (m *Borrow) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Borrow) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Borrow) GetBlockHeight() int32 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *Borrow) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *Borrow) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Borrow) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func init() {
	proto.RegisterType((*Borrow)(nil), "sapienscosmos.ibb.ibb.Borrow")
}

func init() { proto.RegisterFile("ibb/borrow.proto", fileDescriptor_6fcfa17d999a14cb) }

var fileDescriptor_6fcfa17d999a14cb = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x31, 0x4b, 0xc4, 0x30,
	0x14, 0xc7, 0xfb, 0xea, 0xb5, 0x62, 0x04, 0x91, 0x70, 0x4a, 0x70, 0x08, 0xc5, 0xa9, 0x0e, 0xb6,
	0x83, 0xdf, 0xe0, 0x5c, 0x9c, 0x3b, 0xba, 0xf5, 0xb5, 0xa1, 0x17, 0xb4, 0x7d, 0x25, 0xc9, 0xa1,
	0x7e, 0x0b, 0x47, 0x3f, 0x92, 0xe3, 0x8d, 0x8e, 0xd2, 0x7e, 0x11, 0x69, 0x72, 0xc2, 0x0d, 0x09,
	0xf9, 0x85, 0xdf, 0xff, 0xc1, 0xff, 0xb1, 0x4b, 0x8d, 0x58, 0x22, 0x19, 0x43, 0x6f, 0xc5, 0x68,
	0xc8, 0x11, 0xbf, 0xb2, 0xf5, 0xa8, 0xd5, 0x60, 0x1b, 0xb2, 0x3d, 0xd9, 0x42, 0x23, 0x2e, 0xe7,
	0x66, 0xdd, 0x51, 0x47, 0xde, 0x28, 0x97, 0x57, 0x90, 0x6f, 0xbf, 0x80, 0xa5, 0x1b, 0x9f, 0xe6,
	0x82, 0x9d, 0x36, 0x46, 0xd5, 0x8e, 0x8c, 0x80, 0x0c, 0xf2, 0xb3, 0xea, 0x1f, 0xf9, 0x05, 0x8b,
	0x75, 0x2b, 0xe2, 0x0c, 0xf2, 0x55, 0x15, 0xeb, 0x96, 0x67, 0xec, 0x1c, 0x5f, 0xa9, 0x79, 0x79,
	0x52, 0xba, 0xdb, 0x3a, 0x71, 0x92, 0x41, 0x9e, 0x54, 0xc7, 0x5f, 0x7c, 0xcd, 0x92, 0xda, 0x5a,
	0xe5, 0xc4, 0xca, 0x4f, 0x0a, 0xc0, 0xaf, 0x59, 0x5a, 0xf7, 0xb4, 0x1b, 0x9c, 0x48, 0x7c, 0xe4,
	0x40, 0x8b, 0xdd, 0xaa, 0x81, 0x7a, 0x91, 0x06, 0xdb, 0xc3, 0xe6, 0xf1, 0x7b, 0x92, 0xb0, 0x9f,
	0x24, 0xfc, 0x4e, 0x12, 0x3e, 0x67, 0x19, 0xed, 0x67, 0x19, 0xfd, 0xcc, 0x32, 0x7a, 0xbe, 0xeb,
	0xb4, 0xdb, 0xee, 0xb0, 0x68, 0xa8, 0x2f, 0x0f, 0x65, 0xef, 0x43, 0xdb, 0x72, 0xd9, 0xc6, 0xbb,
	0xbf, 0xdd, 0xc7, 0xa8, 0x2c, 0xa6, 0xbe, 0xe6, 0xc3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdd,
	0x4e, 0xf1, 0xed, 0x27, 0x01, 0x00, 0x00,
}

func (m *Borrow) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Borrow) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Borrow) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintBorrow(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x32
	}
	if m.Amount != 0 {
		i = encodeVarintBorrow(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Asset) > 0 {
		i -= len(m.Asset)
		copy(dAtA[i:], m.Asset)
		i = encodeVarintBorrow(dAtA, i, uint64(len(m.Asset)))
		i--
		dAtA[i] = 0x22
	}
	if m.BlockHeight != 0 {
		i = encodeVarintBorrow(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x18
	}
	if m.Id != 0 {
		i = encodeVarintBorrow(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintBorrow(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBorrow(dAtA []byte, offset int, v uint64) int {
	offset -= sovBorrow(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Borrow) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovBorrow(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovBorrow(uint64(m.Id))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovBorrow(uint64(m.BlockHeight))
	}
	l = len(m.Asset)
	if l > 0 {
		n += 1 + l + sovBorrow(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovBorrow(uint64(m.Amount))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovBorrow(uint64(l))
	}
	return n
}

func sovBorrow(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBorrow(x uint64) (n int) {
	return sovBorrow(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Borrow) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBorrow
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
			return fmt.Errorf("proto: Borrow: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Borrow: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBorrow
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
				return ErrInvalidLengthBorrow
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBorrow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBorrow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBorrow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBorrow
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
				return ErrInvalidLengthBorrow
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBorrow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBorrow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBorrow
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
				return ErrInvalidLengthBorrow
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBorrow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBorrow(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBorrow
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
func skipBorrow(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBorrow
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
					return 0, ErrIntOverflowBorrow
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
					return 0, ErrIntOverflowBorrow
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
				return 0, ErrInvalidLengthBorrow
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBorrow
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBorrow
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBorrow        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBorrow          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBorrow = fmt.Errorf("proto: unexpected end of group")
)
