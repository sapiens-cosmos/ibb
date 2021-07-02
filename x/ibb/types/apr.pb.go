// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibb/apr.proto

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

type Apr struct {
	Creator     string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id          uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	BlockHeight int32  `protobuf:"varint,3,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	DepositApy  int32  `protobuf:"varint,4,opt,name=depositApy,proto3" json:"depositApy,omitempty"`
	BorrowApy   int32  `protobuf:"varint,5,opt,name=borrowApy,proto3" json:"borrowApy,omitempty"`
}

func (m *Apr) Reset()         { *m = Apr{} }
func (m *Apr) String() string { return proto.CompactTextString(m) }
func (*Apr) ProtoMessage()    {}
func (*Apr) Descriptor() ([]byte, []int) {
	return fileDescriptor_40bf6b2d94260981, []int{0}
}
func (m *Apr) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Apr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Apr.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Apr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Apr.Merge(m, src)
}
func (m *Apr) XXX_Size() int {
	return m.Size()
}
func (m *Apr) XXX_DiscardUnknown() {
	xxx_messageInfo_Apr.DiscardUnknown(m)
}

var xxx_messageInfo_Apr proto.InternalMessageInfo

func (m *Apr) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Apr) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Apr) GetBlockHeight() int32 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *Apr) GetDepositApy() int32 {
	if m != nil {
		return m.DepositApy
	}
	return 0
}

func (m *Apr) GetBorrowApy() int32 {
	if m != nil {
		return m.BorrowApy
	}
	return 0
}

func init() {
	proto.RegisterType((*Apr)(nil), "sapienscosmos.ibb.ibb.Apr")
}

func init() { proto.RegisterFile("ibb/apr.proto", fileDescriptor_40bf6b2d94260981) }

var fileDescriptor_40bf6b2d94260981 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xe3, 0xb4, 0x05, 0xf5, 0x10, 0x0c, 0x16, 0x48, 0x16, 0x42, 0x56, 0xc4, 0x14, 0x06,
	0x92, 0x81, 0x27, 0x28, 0x2c, 0xcc, 0x19, 0xd9, 0x72, 0x89, 0x95, 0x9e, 0xa0, 0x9c, 0x65, 0x1b,
	0x41, 0x9f, 0x02, 0x1e, 0x8b, 0xb1, 0x23, 0x23, 0x4a, 0x5e, 0x04, 0xc5, 0x05, 0x91, 0xc1, 0xd6,
	0xdd, 0xff, 0x7f, 0xcb, 0x77, 0x70, 0x4c, 0x88, 0x65, 0x6d, 0x5d, 0x61, 0x1d, 0x07, 0x96, 0x67,
	0xbe, 0xb6, 0x64, 0x9e, 0x7d, 0xc3, 0x7e, 0xc3, 0xbe, 0x20, 0xc4, 0xf1, 0x9d, 0x9f, 0x76, 0xdc,
	0x71, 0x24, 0xca, 0x71, 0xda, 0xc3, 0x97, 0xef, 0x02, 0x66, 0x2b, 0xeb, 0xa4, 0x82, 0xc3, 0xc6,
	0x99, 0x3a, 0xb0, 0x53, 0x22, 0x13, 0xf9, 0xb2, 0xfa, 0x5b, 0xe5, 0x09, 0xa4, 0xd4, 0xaa, 0x34,
	0x13, 0xf9, 0xbc, 0x4a, 0xa9, 0x95, 0x19, 0x1c, 0xe1, 0x13, 0x37, 0x8f, 0xf7, 0x86, 0xba, 0x75,
	0x50, 0xb3, 0x4c, 0xe4, 0x8b, 0x6a, 0x1a, 0x49, 0x0d, 0xd0, 0x1a, 0xcb, 0x9e, 0xc2, 0xca, 0x6e,
	0xd5, 0x3c, 0x02, 0x93, 0x44, 0x5e, 0xc0, 0x12, 0xd9, 0x39, 0x7e, 0x1d, 0xeb, 0x45, 0xac, 0xff,
	0x83, 0xdb, 0xbb, 0xcf, 0x5e, 0x8b, 0x5d, 0xaf, 0xc5, 0x77, 0xaf, 0xc5, 0xc7, 0xa0, 0x93, 0xdd,
	0xa0, 0x93, 0xaf, 0x41, 0x27, 0x0f, 0x57, 0x1d, 0x85, 0xf5, 0x0b, 0x16, 0x0d, 0x6f, 0xca, 0x5f,
	0xc7, 0xeb, 0xbd, 0x64, 0x39, 0x5e, 0xe0, 0x2d, 0xfe, 0x61, 0x6b, 0x8d, 0xc7, 0x83, 0x68, 0x77,
	0xf3, 0x13, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xa2, 0xff, 0x83, 0x1b, 0x01, 0x00, 0x00,
}

func (m *Apr) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Apr) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Apr) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BorrowApy != 0 {
		i = encodeVarintApr(dAtA, i, uint64(m.BorrowApy))
		i--
		dAtA[i] = 0x28
	}
	if m.DepositApy != 0 {
		i = encodeVarintApr(dAtA, i, uint64(m.DepositApy))
		i--
		dAtA[i] = 0x20
	}
	if m.BlockHeight != 0 {
		i = encodeVarintApr(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x18
	}
	if m.Id != 0 {
		i = encodeVarintApr(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintApr(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintApr(dAtA []byte, offset int, v uint64) int {
	offset -= sovApr(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Apr) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovApr(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovApr(uint64(m.Id))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovApr(uint64(m.BlockHeight))
	}
	if m.DepositApy != 0 {
		n += 1 + sovApr(uint64(m.DepositApy))
	}
	if m.BorrowApy != 0 {
		n += 1 + sovApr(uint64(m.BorrowApy))
	}
	return n
}

func sovApr(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApr(x uint64) (n int) {
	return sovApr(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Apr) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApr
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
			return fmt.Errorf("proto: Apr: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Apr: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApr
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
				return ErrInvalidLengthApr
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApr
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
					return ErrIntOverflowApr
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
					return ErrIntOverflowApr
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositApy", wireType)
			}
			m.DepositApy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DepositApy |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BorrowApy", wireType)
			}
			m.BorrowApy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BorrowApy |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApr(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApr
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
func skipApr(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApr
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
					return 0, ErrIntOverflowApr
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
					return 0, ErrIntOverflowApr
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
				return 0, ErrInvalidLengthApr
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApr
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApr
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApr        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApr          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApr = fmt.Errorf("proto: unexpected end of group")
)
