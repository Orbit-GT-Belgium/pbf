// Code generated by protoc-gen-gogo.
// source: fileformat.proto
// DO NOT EDIT!

/*
	Package OSMPBF is a generated protocol buffer package.

	It is generated from these files:
		fileformat.proto

	It has these top-level messages:
		Blob
		BlobHeader
*/
package OSMPBF

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import github_com_golang_protobuf_proto "github.com/golang/protobuf/proto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Blob struct {
	Raw     []byte `protobuf:"bytes,1,opt,name=raw" json:"raw,omitempty"`
	RawSize *int32 `protobuf:"varint,2,opt,name=raw_size,json=rawSize" json:"raw_size,omitempty"`
	// Possible compressed versions of the data.
	ZlibData []byte `protobuf:"bytes,3,opt,name=zlib_data,json=zlibData" json:"zlib_data,omitempty"`
	// PROPOSED feature for LZMA compressed data. SUPPORT IS NOT REQUIRED.
	LzmaData []byte `protobuf:"bytes,4,opt,name=lzma_data,json=lzmaData" json:"lzma_data,omitempty"`
	// Formerly used for bzip2 compressed data. Depreciated in 2010.
	OBSOLETEBzip2Data []byte `protobuf:"bytes,5,opt,name=OBSOLETE_bzip2_data,json=oBSOLETEBzip2Data" json:"OBSOLETE_bzip2_data,omitempty"`
	XXX_unrecognized  []byte `json:"-"`
}

func (m *Blob) Reset()                    { *m = Blob{} }
func (m *Blob) String() string            { return proto.CompactTextString(m) }
func (*Blob) ProtoMessage()               {}
func (*Blob) Descriptor() ([]byte, []int) { return fileDescriptorFileformat, []int{0} }

func (m *Blob) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

func (m *Blob) GetRawSize() int32 {
	if m != nil && m.RawSize != nil {
		return *m.RawSize
	}
	return 0
}

func (m *Blob) GetZlibData() []byte {
	if m != nil {
		return m.ZlibData
	}
	return nil
}

func (m *Blob) GetLzmaData() []byte {
	if m != nil {
		return m.LzmaData
	}
	return nil
}

func (m *Blob) GetOBSOLETEBzip2Data() []byte {
	if m != nil {
		return m.OBSOLETEBzip2Data
	}
	return nil
}

type BlobHeader struct {
	Type             *string `protobuf:"bytes,1,req,name=type" json:"type,omitempty"`
	Indexdata        []byte  `protobuf:"bytes,2,opt,name=indexdata" json:"indexdata,omitempty"`
	Datasize         *int32  `protobuf:"varint,3,req,name=datasize" json:"datasize,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BlobHeader) Reset()                    { *m = BlobHeader{} }
func (m *BlobHeader) String() string            { return proto.CompactTextString(m) }
func (*BlobHeader) ProtoMessage()               {}
func (*BlobHeader) Descriptor() ([]byte, []int) { return fileDescriptorFileformat, []int{1} }

func (m *BlobHeader) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *BlobHeader) GetIndexdata() []byte {
	if m != nil {
		return m.Indexdata
	}
	return nil
}

func (m *BlobHeader) GetDatasize() int32 {
	if m != nil && m.Datasize != nil {
		return *m.Datasize
	}
	return 0
}

func init() {
	proto.RegisterType((*Blob)(nil), "OSMPBF.Blob")
	proto.RegisterType((*BlobHeader)(nil), "OSMPBF.BlobHeader")
}
func (m *Blob) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Blob) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Raw != nil {
		data[i] = 0xa
		i++
		i = encodeVarintFileformat(data, i, uint64(len(m.Raw)))
		i += copy(data[i:], m.Raw)
	}
	if m.RawSize != nil {
		data[i] = 0x10
		i++
		i = encodeVarintFileformat(data, i, uint64(*m.RawSize))
	}
	if m.ZlibData != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintFileformat(data, i, uint64(len(m.ZlibData)))
		i += copy(data[i:], m.ZlibData)
	}
	if m.LzmaData != nil {
		data[i] = 0x22
		i++
		i = encodeVarintFileformat(data, i, uint64(len(m.LzmaData)))
		i += copy(data[i:], m.LzmaData)
	}
	if m.OBSOLETEBzip2Data != nil {
		data[i] = 0x2a
		i++
		i = encodeVarintFileformat(data, i, uint64(len(m.OBSOLETEBzip2Data)))
		i += copy(data[i:], m.OBSOLETEBzip2Data)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *BlobHeader) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *BlobHeader) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		data[i] = 0xa
		i++
		i = encodeVarintFileformat(data, i, uint64(len(*m.Type)))
		i += copy(data[i:], *m.Type)
	}
	if m.Indexdata != nil {
		data[i] = 0x12
		i++
		i = encodeVarintFileformat(data, i, uint64(len(m.Indexdata)))
		i += copy(data[i:], m.Indexdata)
	}
	if m.Datasize == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		data[i] = 0x18
		i++
		i = encodeVarintFileformat(data, i, uint64(*m.Datasize))
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Fileformat(data []byte, offset int, v uint64) int {
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
func encodeFixed32Fileformat(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintFileformat(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Blob) Size() (n int) {
	var l int
	_ = l
	if m.Raw != nil {
		l = len(m.Raw)
		n += 1 + l + sovFileformat(uint64(l))
	}
	if m.RawSize != nil {
		n += 1 + sovFileformat(uint64(*m.RawSize))
	}
	if m.ZlibData != nil {
		l = len(m.ZlibData)
		n += 1 + l + sovFileformat(uint64(l))
	}
	if m.LzmaData != nil {
		l = len(m.LzmaData)
		n += 1 + l + sovFileformat(uint64(l))
	}
	if m.OBSOLETEBzip2Data != nil {
		l = len(m.OBSOLETEBzip2Data)
		n += 1 + l + sovFileformat(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BlobHeader) Size() (n int) {
	var l int
	_ = l
	if m.Type != nil {
		l = len(*m.Type)
		n += 1 + l + sovFileformat(uint64(l))
	}
	if m.Indexdata != nil {
		l = len(m.Indexdata)
		n += 1 + l + sovFileformat(uint64(l))
	}
	if m.Datasize != nil {
		n += 1 + sovFileformat(uint64(*m.Datasize))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFileformat(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFileformat(x uint64) (n int) {
	return sovFileformat(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Blob) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFileformat
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
			return fmt.Errorf("proto: Blob: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Blob: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Raw", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileformat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFileformat
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Raw = append(m.Raw[:0], data[iNdEx:postIndex]...)
			if m.Raw == nil {
				m.Raw = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawSize", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileformat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.RawSize = &v
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZlibData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileformat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFileformat
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ZlibData = append(m.ZlibData[:0], data[iNdEx:postIndex]...)
			if m.ZlibData == nil {
				m.ZlibData = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LzmaData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileformat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFileformat
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LzmaData = append(m.LzmaData[:0], data[iNdEx:postIndex]...)
			if m.LzmaData == nil {
				m.LzmaData = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OBSOLETEBzip2Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileformat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFileformat
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OBSOLETEBzip2Data = append(m.OBSOLETEBzip2Data[:0], data[iNdEx:postIndex]...)
			if m.OBSOLETEBzip2Data == nil {
				m.OBSOLETEBzip2Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFileformat(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFileformat
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BlobHeader) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFileformat
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
			return fmt.Errorf("proto: BlobHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlobHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileformat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFileformat
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.Type = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Indexdata", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileformat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFileformat
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Indexdata = append(m.Indexdata[:0], data[iNdEx:postIndex]...)
			if m.Indexdata == nil {
				m.Indexdata = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Datasize", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileformat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Datasize = &v
			hasFields[0] |= uint64(0x00000002)
		default:
			iNdEx = preIndex
			skippy, err := skipFileformat(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFileformat
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFileformat(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFileformat
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
					return 0, ErrIntOverflowFileformat
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
					return 0, ErrIntOverflowFileformat
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
				return 0, ErrInvalidLengthFileformat
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFileformat
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
				next, err := skipFileformat(data[start:])
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
	ErrInvalidLengthFileformat = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFileformat   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorFileformat = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xcb, 0xcc, 0x49,
	0x4d, 0xcb, 0x2f, 0xca, 0x4d, 0x2c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xf3, 0x0f,
	0xf6, 0x0d, 0x70, 0x72, 0x53, 0x5a, 0xc8, 0xc8, 0xc5, 0xe2, 0x94, 0x93, 0x9f, 0x24, 0x24, 0xc0,
	0xc5, 0x5c, 0x94, 0x58, 0x2e, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x13, 0x04, 0x62, 0x0a, 0x49, 0x72,
	0x71, 0x00, 0xa9, 0xf8, 0xe2, 0xcc, 0xaa, 0x54, 0x09, 0x26, 0xa0, 0x30, 0x6b, 0x10, 0x3b, 0x90,
	0x1f, 0x0c, 0xe4, 0x0a, 0x49, 0x73, 0x71, 0x56, 0xe5, 0x64, 0x26, 0xc5, 0xa7, 0x24, 0x96, 0x24,
	0x4a, 0x30, 0x83, 0xb5, 0x70, 0x80, 0x04, 0x5c, 0x80, 0x7c, 0x90, 0x64, 0x4e, 0x55, 0x6e, 0x22,
	0x44, 0x92, 0x05, 0x22, 0x09, 0x12, 0x00, 0x4b, 0x1a, 0x71, 0x09, 0xfb, 0x3b, 0x05, 0xfb, 0xfb,
	0xb8, 0x86, 0xb8, 0xc6, 0x27, 0x55, 0x65, 0x16, 0x18, 0x41, 0x94, 0xb1, 0x82, 0x94, 0x39, 0x31,
	0x49, 0x30, 0x06, 0x09, 0xe6, 0x43, 0xa5, 0x9d, 0x40, 0xb2, 0x20, 0x3d, 0x4a, 0x51, 0x5c, 0x5c,
	0x20, 0x27, 0x7a, 0xa4, 0x26, 0xa6, 0xa4, 0x16, 0x09, 0x09, 0x71, 0xb1, 0x94, 0x54, 0x16, 0xa4,
	0x02, 0x5d, 0xca, 0xa4, 0xc1, 0x19, 0x04, 0x66, 0x0b, 0xc9, 0x70, 0x71, 0x66, 0xe6, 0xa5, 0xa4,
	0x56, 0x80, 0xcd, 0x62, 0x02, 0x5b, 0x89, 0x10, 0x10, 0x92, 0xe2, 0xe2, 0x00, 0xd1, 0x60, 0x8f,
	0x30, 0x03, 0x75, 0xb1, 0x06, 0xc1, 0xf9, 0x4e, 0x8a, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x00, 0xe2,
	0x07, 0x40, 0x3c, 0xe3, 0xb1, 0x1c, 0x03, 0x17, 0x6f, 0x72, 0x51, 0x7e, 0x71, 0x52, 0xa5, 0x5e,
	0x52, 0x66, 0x5e, 0x62, 0x51, 0xa5, 0x07, 0x33, 0x20, 0x00, 0x00, 0xff, 0xff, 0xac, 0xe2, 0x10,
	0xa6, 0x3d, 0x01, 0x00, 0x00,
}
