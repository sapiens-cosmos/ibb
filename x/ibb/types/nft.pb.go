// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibb/nft.proto

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

type Nft struct {
	Creator           string   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id                uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Collection        string   `protobuf:"bytes,3,opt,name=collection,proto3" json:"collection,omitempty"`
	OwnerAddress      string   `protobuf:"bytes,4,opt,name=ownerAddress,proto3" json:"ownerAddress,omitempty"`
	ImageUrl          string   `protobuf:"bytes,5,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	Name              string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	NftCreatorAddress string   `protobuf:"bytes,7,opt,name=nftCreatorAddress,proto3" json:"nftCreatorAddress,omitempty"`
	SelectedOffer     *Offer   `protobuf:"bytes,8,opt,name=selectedOffer,proto3" json:"selectedOffer,omitempty"`
	Offers            []*Offer `protobuf:"bytes,9,rep,name=offers,proto3" json:"offers,omitempty"`
}

func (m *Nft) Reset()         { *m = Nft{} }
func (m *Nft) String() string { return proto.CompactTextString(m) }
func (*Nft) ProtoMessage()    {}
func (*Nft) Descriptor() ([]byte, []int) {
	return fileDescriptor_c984c4bc68c11f1b, []int{0}
}
func (m *Nft) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Nft) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Nft.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Nft) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nft.Merge(m, src)
}
func (m *Nft) XXX_Size() int {
	return m.Size()
}
func (m *Nft) XXX_DiscardUnknown() {
	xxx_messageInfo_Nft.DiscardUnknown(m)
}

var xxx_messageInfo_Nft proto.InternalMessageInfo

func (m *Nft) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Nft) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Nft) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *Nft) GetOwnerAddress() string {
	if m != nil {
		return m.OwnerAddress
	}
	return ""
}

func (m *Nft) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *Nft) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Nft) GetNftCreatorAddress() string {
	if m != nil {
		return m.NftCreatorAddress
	}
	return ""
}

func (m *Nft) GetSelectedOffer() *Offer {
	if m != nil {
		return m.SelectedOffer
	}
	return nil
}

func (m *Nft) GetOffers() []*Offer {
	if m != nil {
		return m.Offers
	}
	return nil
}

type Offer struct {
	Denom           string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Amount          int32  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	PaybackAmount   int32  `protobuf:"varint,3,opt,name=paybackAmount,proto3" json:"paybackAmount,omitempty"`
	PaybackDuration int32  `protobuf:"varint,4,opt,name=paybackDuration,proto3" json:"paybackDuration,omitempty"`
	OfferStartAt    int64  `protobuf:"varint,5,opt,name=offerStartAt,proto3" json:"offerStartAt,omitempty"`
	NftId           int32  `protobuf:"varint,6,opt,name=nftId,proto3" json:"nftId,omitempty"`
	OfferCreator    string `protobuf:"bytes,7,opt,name=offerCreator,proto3" json:"offerCreator,omitempty"`
	Id              uint64 `protobuf:"varint,8,opt,name=id,proto3" json:"id,omitempty"`
	Interest        int32  `protobuf:"varint,9,opt,name=interest,proto3" json:"interest,omitempty"`
}

func (m *Offer) Reset()         { *m = Offer{} }
func (m *Offer) String() string { return proto.CompactTextString(m) }
func (*Offer) ProtoMessage()    {}
func (*Offer) Descriptor() ([]byte, []int) {
	return fileDescriptor_c984c4bc68c11f1b, []int{1}
}
func (m *Offer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Offer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Offer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Offer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Offer.Merge(m, src)
}
func (m *Offer) XXX_Size() int {
	return m.Size()
}
func (m *Offer) XXX_DiscardUnknown() {
	xxx_messageInfo_Offer.DiscardUnknown(m)
}

var xxx_messageInfo_Offer proto.InternalMessageInfo

func (m *Offer) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *Offer) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Offer) GetPaybackAmount() int32 {
	if m != nil {
		return m.PaybackAmount
	}
	return 0
}

func (m *Offer) GetPaybackDuration() int32 {
	if m != nil {
		return m.PaybackDuration
	}
	return 0
}

func (m *Offer) GetOfferStartAt() int64 {
	if m != nil {
		return m.OfferStartAt
	}
	return 0
}

func (m *Offer) GetNftId() int32 {
	if m != nil {
		return m.NftId
	}
	return 0
}

func (m *Offer) GetOfferCreator() string {
	if m != nil {
		return m.OfferCreator
	}
	return ""
}

func (m *Offer) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Offer) GetInterest() int32 {
	if m != nil {
		return m.Interest
	}
	return 0
}

func init() {
	proto.RegisterType((*Nft)(nil), "sapienscosmos.ibb.ibb.Nft")
	proto.RegisterType((*Offer)(nil), "sapienscosmos.ibb.ibb.Offer")
}

func init() { proto.RegisterFile("ibb/nft.proto", fileDescriptor_c984c4bc68c11f1b) }

var fileDescriptor_c984c4bc68c11f1b = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0xaa, 0xd3, 0x40,
	0x14, 0x86, 0x9b, 0xa4, 0xe9, 0x6d, 0x8f, 0x56, 0x71, 0xb8, 0xca, 0x70, 0x91, 0x50, 0x8a, 0x8b,
	0x0a, 0x9a, 0x80, 0xfa, 0x02, 0xbd, 0xd7, 0x8d, 0x1b, 0x85, 0x88, 0x1b, 0x77, 0x93, 0x64, 0x12,
	0x07, 0x9b, 0x99, 0x30, 0x33, 0x45, 0xef, 0x53, 0xe8, 0x63, 0x09, 0x6e, 0xee, 0xd2, 0xa5, 0xb4,
	0x2f, 0x22, 0x73, 0x32, 0x0d, 0x46, 0x05, 0x17, 0x09, 0xf3, 0xff, 0xe7, 0x3f, 0xc3, 0xe1, 0x9b,
	0x03, 0x4b, 0x51, 0x14, 0x99, 0xac, 0x6d, 0xda, 0x69, 0x65, 0x15, 0xb9, 0x6f, 0x58, 0x27, 0xb8,
	0x34, 0xa5, 0x32, 0xad, 0x32, 0xa9, 0x28, 0x0a, 0xf7, 0x5d, 0x9c, 0x37, 0xaa, 0x51, 0x98, 0xc8,
	0xdc, 0xa9, 0x0f, 0xaf, 0xbf, 0x87, 0x10, 0xbd, 0xae, 0x2d, 0xa1, 0x70, 0x56, 0x6a, 0xce, 0xac,
	0xd2, 0x34, 0x58, 0x05, 0x9b, 0x45, 0x7e, 0x92, 0xe4, 0x0e, 0x84, 0xa2, 0xa2, 0xe1, 0x2a, 0xd8,
	0x4c, 0xf3, 0x50, 0x54, 0x24, 0x01, 0x28, 0xd5, 0x6e, 0xc7, 0x4b, 0x2b, 0x94, 0xa4, 0x11, 0x86,
	0x7f, 0x73, 0xc8, 0x1a, 0x6e, 0xab, 0x4f, 0x92, 0xeb, 0x6d, 0x55, 0x69, 0x6e, 0x0c, 0x9d, 0x62,
	0x62, 0xe4, 0x91, 0x0b, 0x98, 0x8b, 0x96, 0x35, 0xfc, 0x9d, 0xde, 0xd1, 0x18, 0xeb, 0x83, 0x26,
	0x04, 0xa6, 0x92, 0xb5, 0x9c, 0xce, 0xd0, 0xc7, 0x33, 0x79, 0x02, 0xf7, 0x64, 0x6d, 0xaf, 0xfa,
	0x89, 0x4e, 0x17, 0x9f, 0x61, 0xe0, 0xef, 0x02, 0xb9, 0x84, 0xa5, 0xe1, 0x6e, 0x1c, 0x5e, 0xbd,
	0xa9, 0x6b, 0xae, 0xe9, 0x7c, 0x15, 0x6c, 0x6e, 0x3d, 0x7b, 0x98, 0xfe, 0x13, 0x4c, 0x8a, 0x99,
	0x7c, 0xdc, 0x42, 0x5e, 0xc0, 0x4c, 0xb9, 0x83, 0xa1, 0x8b, 0x55, 0xf4, 0xdf, 0x66, 0x9f, 0x5d,
	0x7f, 0x09, 0x21, 0xee, 0xfb, 0xcf, 0x21, 0xae, 0xb8, 0x54, 0xad, 0xa7, 0xd9, 0x0b, 0xf2, 0x00,
	0x66, 0xac, 0x55, 0x7b, 0x69, 0x91, 0x67, 0x9c, 0x7b, 0x45, 0x1e, 0xc1, 0xb2, 0x63, 0xd7, 0x05,
	0x2b, 0x3f, 0x6e, 0xfb, 0x72, 0x84, 0xe5, 0xb1, 0x49, 0x36, 0x70, 0xd7, 0x1b, 0x2f, 0xf7, 0x9a,
	0x21, 0xfe, 0x29, 0xe6, 0xfe, 0xb4, 0xf1, 0x0d, 0xdc, 0x18, 0x6f, 0x2d, 0xd3, 0x76, 0x6b, 0x91,
	0x71, 0x94, 0x8f, 0x3c, 0x37, 0xa1, 0xac, 0xed, 0xab, 0x0a, 0x41, 0xc7, 0x79, 0x2f, 0x86, 0x4e,
	0x8f, 0xd4, 0x43, 0x1e, 0x79, 0x7e, 0x23, 0xe6, 0xc3, 0x46, 0xb8, 0xd7, 0x94, 0x96, 0x6b, 0x6e,
	0x2c, 0x5d, 0xe0, 0x65, 0x83, 0xbe, 0xbc, 0xfa, 0x76, 0x48, 0x82, 0x9b, 0x43, 0x12, 0xfc, 0x3c,
	0x24, 0xc1, 0xd7, 0x63, 0x32, 0xb9, 0x39, 0x26, 0x93, 0x1f, 0xc7, 0x64, 0xf2, 0xfe, 0x71, 0x23,
	0xec, 0x87, 0x7d, 0x91, 0x96, 0xaa, 0xcd, 0x3c, 0xdb, 0xa7, 0x3d, 0xdc, 0xcc, 0xed, 0xf3, 0x67,
	0xfc, 0xdb, 0xeb, 0x8e, 0x9b, 0x62, 0x86, 0xbb, 0xfa, 0xfc, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x63, 0x11, 0x73, 0x4b, 0xe9, 0x02, 0x00, 0x00,
}

func (m *Nft) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Nft) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Nft) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Offers) > 0 {
		for iNdEx := len(m.Offers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Offers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNft(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.SelectedOffer != nil {
		{
			size, err := m.SelectedOffer.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNft(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if len(m.NftCreatorAddress) > 0 {
		i -= len(m.NftCreatorAddress)
		copy(dAtA[i:], m.NftCreatorAddress)
		i = encodeVarintNft(dAtA, i, uint64(len(m.NftCreatorAddress)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ImageUrl) > 0 {
		i -= len(m.ImageUrl)
		copy(dAtA[i:], m.ImageUrl)
		i = encodeVarintNft(dAtA, i, uint64(len(m.ImageUrl)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintNft(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Collection) > 0 {
		i -= len(m.Collection)
		copy(dAtA[i:], m.Collection)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Collection)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintNft(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Offer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Offer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Offer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Interest != 0 {
		i = encodeVarintNft(dAtA, i, uint64(m.Interest))
		i--
		dAtA[i] = 0x48
	}
	if m.Id != 0 {
		i = encodeVarintNft(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x40
	}
	if len(m.OfferCreator) > 0 {
		i -= len(m.OfferCreator)
		copy(dAtA[i:], m.OfferCreator)
		i = encodeVarintNft(dAtA, i, uint64(len(m.OfferCreator)))
		i--
		dAtA[i] = 0x3a
	}
	if m.NftId != 0 {
		i = encodeVarintNft(dAtA, i, uint64(m.NftId))
		i--
		dAtA[i] = 0x30
	}
	if m.OfferStartAt != 0 {
		i = encodeVarintNft(dAtA, i, uint64(m.OfferStartAt))
		i--
		dAtA[i] = 0x28
	}
	if m.PaybackDuration != 0 {
		i = encodeVarintNft(dAtA, i, uint64(m.PaybackDuration))
		i--
		dAtA[i] = 0x20
	}
	if m.PaybackAmount != 0 {
		i = encodeVarintNft(dAtA, i, uint64(m.PaybackAmount))
		i--
		dAtA[i] = 0x18
	}
	if m.Amount != 0 {
		i = encodeVarintNft(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNft(dAtA []byte, offset int, v uint64) int {
	offset -= sovNft(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Nft) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovNft(uint64(m.Id))
	}
	l = len(m.Collection)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.OwnerAddress)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.ImageUrl)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.NftCreatorAddress)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	if m.SelectedOffer != nil {
		l = m.SelectedOffer.Size()
		n += 1 + l + sovNft(uint64(l))
	}
	if len(m.Offers) > 0 {
		for _, e := range m.Offers {
			l = e.Size()
			n += 1 + l + sovNft(uint64(l))
		}
	}
	return n
}

func (m *Offer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovNft(uint64(m.Amount))
	}
	if m.PaybackAmount != 0 {
		n += 1 + sovNft(uint64(m.PaybackAmount))
	}
	if m.PaybackDuration != 0 {
		n += 1 + sovNft(uint64(m.PaybackDuration))
	}
	if m.OfferStartAt != 0 {
		n += 1 + sovNft(uint64(m.OfferStartAt))
	}
	if m.NftId != 0 {
		n += 1 + sovNft(uint64(m.NftId))
	}
	l = len(m.OfferCreator)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovNft(uint64(m.Id))
	}
	if m.Interest != 0 {
		n += 1 + sovNft(uint64(m.Interest))
	}
	return n
}

func sovNft(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNft(x uint64) (n int) {
	return sovNft(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Nft) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNft
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
			return fmt.Errorf("proto: Nft: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Nft: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
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
					return ErrIntOverflowNft
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collection", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Collection = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImageUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImageUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftCreatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftCreatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SelectedOffer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SelectedOffer == nil {
				m.SelectedOffer = &Offer{}
			}
			if err := m.SelectedOffer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Offers = append(m.Offers, &Offer{})
			if err := m.Offers[len(m.Offers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNft
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
func (m *Offer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNft
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
			return fmt.Errorf("proto: Offer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Offer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
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
					return ErrIntOverflowNft
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaybackAmount", wireType)
			}
			m.PaybackAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PaybackAmount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaybackDuration", wireType)
			}
			m.PaybackDuration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PaybackDuration |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OfferStartAt", wireType)
			}
			m.OfferStartAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OfferStartAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftId", wireType)
			}
			m.NftId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NftId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OfferCreator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OfferCreator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Interest", wireType)
			}
			m.Interest = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Interest |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNft
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
func skipNft(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNft
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
					return 0, ErrIntOverflowNft
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
					return 0, ErrIntOverflowNft
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
				return 0, ErrInvalidLengthNft
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNft
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNft
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNft        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNft          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNft = fmt.Errorf("proto: unexpected end of group")
)
