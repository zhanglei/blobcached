// Code generated by protoc-gen-gogo.
// source: index.proto
// DO NOT EDIT!

/*
	Package cache is a generated protocol buffer package.

	It is generated from these files:
		index.proto

	It has these top-level messages:
		IndexMeta
		IndexItem
*/
package cache

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type IndexMeta struct {
	Term     int64 `protobuf:"varint,1,req,name=Term" json:"Term"`
	Head     int64 `protobuf:"varint,2,req,name=Head" json:"Head"`
	DataSize int64 `protobuf:"varint,3,req,name=DataSize" json:"DataSize"`
}

func (m *IndexMeta) Reset()                    { *m = IndexMeta{} }
func (m *IndexMeta) String() string            { return proto.CompactTextString(m) }
func (*IndexMeta) ProtoMessage()               {}
func (*IndexMeta) Descriptor() ([]byte, []int) { return fileDescriptorIndex, []int{0} }

type IndexItem struct {
	Term      int64  `protobuf:"varint,1,req,name=Term" json:"Term"`
	Offset    int64  `protobuf:"varint,2,req,name=Offset" json:"Offset"`
	ValueSize int32  `protobuf:"varint,3,req,name=ValueSize" json:"ValueSize"`
	Timestamp int64  `protobuf:"varint,4,req,name=Timestamp" json:"Timestamp"`
	TTL       uint32 `protobuf:"varint,5,req,name=TTL" json:"TTL"`
	Flags     uint32 `protobuf:"varint,6,req,name=Flags" json:"Flags"`
}

func (m *IndexItem) Reset()                    { *m = IndexItem{} }
func (m *IndexItem) String() string            { return proto.CompactTextString(m) }
func (*IndexItem) ProtoMessage()               {}
func (*IndexItem) Descriptor() ([]byte, []int) { return fileDescriptorIndex, []int{1} }

func init() {
	proto.RegisterType((*IndexMeta)(nil), "cache.IndexMeta")
	proto.RegisterType((*IndexItem)(nil), "cache.IndexItem")
}
func (m *IndexMeta) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *IndexMeta) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintIndex(data, i, uint64(m.Term))
	data[i] = 0x10
	i++
	i = encodeVarintIndex(data, i, uint64(m.Head))
	data[i] = 0x18
	i++
	i = encodeVarintIndex(data, i, uint64(m.DataSize))
	return i, nil
}

func (m *IndexItem) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *IndexItem) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintIndex(data, i, uint64(m.Term))
	data[i] = 0x10
	i++
	i = encodeVarintIndex(data, i, uint64(m.Offset))
	data[i] = 0x18
	i++
	i = encodeVarintIndex(data, i, uint64(m.ValueSize))
	data[i] = 0x20
	i++
	i = encodeVarintIndex(data, i, uint64(m.Timestamp))
	data[i] = 0x28
	i++
	i = encodeVarintIndex(data, i, uint64(m.TTL))
	data[i] = 0x30
	i++
	i = encodeVarintIndex(data, i, uint64(m.Flags))
	return i, nil
}

func encodeFixed64Index(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Index(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintIndex(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *IndexMeta) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovIndex(uint64(m.Term))
	n += 1 + sovIndex(uint64(m.Head))
	n += 1 + sovIndex(uint64(m.DataSize))
	return n
}

func (m *IndexItem) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovIndex(uint64(m.Term))
	n += 1 + sovIndex(uint64(m.Offset))
	n += 1 + sovIndex(uint64(m.ValueSize))
	n += 1 + sovIndex(uint64(m.Timestamp))
	n += 1 + sovIndex(uint64(m.TTL))
	n += 1 + sovIndex(uint64(m.Flags))
	return n
}

func sovIndex(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozIndex(x uint64) (n int) {
	return sovIndex(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IndexMeta) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIndex
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IndexMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Term |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Head", wireType)
			}
			m.Head = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Head |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataSize", wireType)
			}
			m.DataSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.DataSize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000004)
		default:
			iNdEx = preIndex
			skippy, err := skipIndex(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIndex
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Term")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Head")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("DataSize")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IndexItem) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIndex
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IndexItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Term |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Offset |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValueSize", wireType)
			}
			m.ValueSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ValueSize |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000004)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Timestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000008)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TTL", wireType)
			}
			m.TTL = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.TTL |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000010)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flags", wireType)
			}
			m.Flags = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Flags |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000020)
		default:
			iNdEx = preIndex
			skippy, err := skipIndex(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIndex
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Term")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Offset")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("ValueSize")
	}
	if hasFields[0]&uint64(0x00000008) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Timestamp")
	}
	if hasFields[0]&uint64(0x00000010) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("TTL")
	}
	if hasFields[0]&uint64(0x00000020) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Flags")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipIndex(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIndex
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowIndex
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowIndex
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthIndex
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowIndex
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipIndex(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthIndex = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIndex   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("index.proto", fileDescriptorIndex) }

var fileDescriptorIndex = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xcc, 0x4b, 0x49,
	0xad, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x4e, 0x4c, 0xce, 0x48, 0x95, 0xd2,
	0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7,
	0x07, 0xcb, 0x26, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa5, 0xe4, 0xcd, 0xc5,
	0xe9, 0x09, 0x32, 0xc4, 0x37, 0xb5, 0x24, 0x51, 0x48, 0x88, 0x8b, 0x25, 0x24, 0xb5, 0x28, 0x57,
	0x82, 0x51, 0x81, 0x49, 0x83, 0xd9, 0x89, 0xe5, 0xc4, 0x3d, 0x79, 0x06, 0x90, 0x98, 0x47, 0x6a,
	0x62, 0x8a, 0x04, 0x13, 0x92, 0x98, 0x18, 0x17, 0x87, 0x4b, 0x62, 0x49, 0x62, 0x70, 0x66, 0x55,
	0xaa, 0x04, 0x33, 0x42, 0x5c, 0xa9, 0x9f, 0x11, 0x6a, 0x9a, 0x67, 0x49, 0x6a, 0x2e, 0x56, 0xd3,
	0x44, 0xb8, 0xd8, 0xfc, 0xd3, 0xd2, 0x8a, 0x53, 0x4b, 0x50, 0xcc, 0x13, 0xe7, 0xe2, 0x0c, 0x4b,
	0xcc, 0x29, 0x4d, 0x85, 0x1b, 0xc8, 0x8a, 0x90, 0x08, 0xc9, 0xcc, 0x4d, 0x2d, 0x2e, 0x49, 0xcc,
	0x2d, 0x90, 0x60, 0x41, 0xd2, 0x21, 0xc8, 0xc5, 0x1c, 0x12, 0xe2, 0x23, 0xc1, 0xaa, 0xc0, 0xa4,
	0xc1, 0x0b, 0x15, 0x12, 0xe6, 0x62, 0x75, 0xcb, 0x49, 0x4c, 0x2f, 0x96, 0x60, 0x43, 0x08, 0x3a,
	0x89, 0x9c, 0x78, 0x28, 0xc7, 0x70, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e,
	0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x00, 0x08, 0x00, 0x00, 0xff, 0xff, 0x63, 0x2a, 0xe4, 0x44,
	0x38, 0x01, 0x00, 0x00,
}
