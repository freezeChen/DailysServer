// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto.proto

package proto

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

type Proto struct {
	Ver  int32  `protobuf:"varint,1,opt,name=ver,proto3" json:"ver"`
	Opr  int32  `protobuf:"varint,2,opt,name=opr,proto3" json:"opr"`
	Seq  int32  `protobuf:"varint,3,opt,name=seq,proto3" json:"seq"`
	Body []byte `protobuf:"bytes,4,opt,name=body,proto3" json:"body"`
}

func (m *Proto) Reset()         { *m = Proto{} }
func (m *Proto) String() string { return proto.CompactTextString(m) }
func (*Proto) ProtoMessage()    {}
func (*Proto) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{0}
}
func (m *Proto) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Proto.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proto.Merge(m, src)
}
func (m *Proto) XXX_Size() int {
	return m.Size()
}
func (m *Proto) XXX_DiscardUnknown() {
	xxx_messageInfo_Proto.DiscardUnknown(m)
}

var xxx_messageInfo_Proto proto.InternalMessageInfo

func (m *Proto) GetVer() int32 {
	if m != nil {
		return m.Ver
	}
	return 0
}

func (m *Proto) GetOpr() int32 {
	if m != nil {
		return m.Opr
	}
	return 0
}

func (m *Proto) GetSeq() int32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Proto) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type EmptyReply struct {
}

func (m *EmptyReply) Reset()         { *m = EmptyReply{} }
func (m *EmptyReply) String() string { return proto.CompactTextString(m) }
func (*EmptyReply) ProtoMessage()    {}
func (*EmptyReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{1}
}
func (m *EmptyReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EmptyReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EmptyReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EmptyReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyReply.Merge(m, src)
}
func (m *EmptyReply) XXX_Size() int {
	return m.Size()
}
func (m *EmptyReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyReply.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Proto)(nil), "proto.Proto")
	proto.RegisterType((*EmptyReply)(nil), "proto.EmptyReply")
}

func init() { proto.RegisterFile("proto.proto", fileDescriptor_2fcc84b9998d60d8) }

var fileDescriptor_2fcc84b9998d60d8 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x03, 0x93, 0x42, 0xac, 0x60, 0x4a, 0x4a, 0x37, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49,
	0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x3d, 0x3f, 0x3d, 0x5f, 0x1f, 0x2c, 0x9c, 0x54, 0x9a, 0x06, 0xe6,
	0x81, 0x39, 0x60, 0x16, 0x44, 0x97, 0x52, 0x19, 0x17, 0x6b, 0x00, 0x58, 0xbb, 0x24, 0x17, 0x73,
	0x59, 0x6a, 0x91, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xab, 0x13, 0xfb, 0xab, 0x7b, 0xf2, 0x20, 0x6e,
	0x10, 0x88, 0x00, 0x49, 0xe5, 0x17, 0x14, 0x49, 0x30, 0x21, 0xa4, 0xf2, 0x0b, 0x8a, 0x82, 0x40,
	0x04, 0x48, 0xaa, 0x38, 0xb5, 0x50, 0x82, 0x19, 0x21, 0x55, 0x9c, 0x5a, 0x18, 0x04, 0x22, 0x84,
	0x64, 0xb8, 0x58, 0x92, 0xf2, 0x53, 0x2a, 0x25, 0x58, 0x14, 0x18, 0x35, 0x78, 0x9c, 0x38, 0x5e,
	0xdd, 0x93, 0x07, 0xf3, 0x83, 0xc0, 0xa4, 0x12, 0x0f, 0x17, 0x97, 0x6b, 0x6e, 0x41, 0x49, 0x65,
	0x50, 0x6a, 0x41, 0x4e, 0xa5, 0x93, 0xc4, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e,
	0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31,
	0x24, 0xb1, 0x81, 0x9d, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x56, 0x32, 0x91, 0x9a, 0xeb,
	0x00, 0x00, 0x00,
}

func (m *Proto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proto) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proto) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintProto(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x22
	}
	if m.Seq != 0 {
		i = encodeVarintProto(dAtA, i, uint64(m.Seq))
		i--
		dAtA[i] = 0x18
	}
	if m.Opr != 0 {
		i = encodeVarintProto(dAtA, i, uint64(m.Opr))
		i--
		dAtA[i] = 0x10
	}
	if m.Ver != 0 {
		i = encodeVarintProto(dAtA, i, uint64(m.Ver))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EmptyReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EmptyReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EmptyReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintProto(dAtA []byte, offset int, v uint64) int {
	offset -= sovProto(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Proto) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ver != 0 {
		n += 1 + sovProto(uint64(m.Ver))
	}
	if m.Opr != 0 {
		n += 1 + sovProto(uint64(m.Opr))
	}
	if m.Seq != 0 {
		n += 1 + sovProto(uint64(m.Seq))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovProto(uint64(l))
	}
	return n
}

func (m *EmptyReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovProto(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProto(x uint64) (n int) {
	return sovProto(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Proto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProto
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
			return fmt.Errorf("proto: Proto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ver", wireType)
			}
			m.Ver = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ver |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Opr", wireType)
			}
			m.Opr = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Opr |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seq", wireType)
			}
			m.Seq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Seq |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = append(m.Body[:0], dAtA[iNdEx:postIndex]...)
			if m.Body == nil {
				m.Body = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProto
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProto
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
func (m *EmptyReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProto
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
			return fmt.Errorf("proto: EmptyReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmptyReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipProto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProto
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProto
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
func skipProto(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProto
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
					return 0, ErrIntOverflowProto
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
					return 0, ErrIntOverflowProto
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
				return 0, ErrInvalidLengthProto
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProto
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProto
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProto        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProto          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProto = fmt.Errorf("proto: unexpected end of group")
)
