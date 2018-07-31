// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Common.proto

package Common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 返回结果
type RetType int32

const (
	RetType_RetType_Succeed RetType = 0
	RetType_RetType_Failed  RetType = -1
	RetType_RetType_TimeOut RetType = -100
	RetType_RetType_Unknown RetType = -400
)

var RetType_name = map[int32]string{
	0:    "RetType_Succeed",
	-1:   "RetType_Failed",
	-100: "RetType_TimeOut",
	-400: "RetType_Unknown",
}
var RetType_value = map[string]int32{
	"RetType_Succeed": 0,
	"RetType_Failed":  -1,
	"RetType_TimeOut": -100,
	"RetType_Unknown": -400,
}

func (x RetType) Enum() *RetType {
	p := new(RetType)
	*p = x
	return p
}
func (x RetType) String() string {
	return proto.EnumName(RetType_name, int32(x))
}
func (x *RetType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RetType_value, data, "RetType")
	if err != nil {
		return err
	}
	*x = RetType(value)
	return nil
}
func (RetType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_Common_d9ca241298b4d843, []int{0}
}

// 包的唯一标识，用于回放攻击的识别和保护
type PacketID struct {
	ConnID               *uint64  `protobuf:"varint,1,req,name=connID" json:"connID,omitempty"`
	SerialNo             *uint32  `protobuf:"varint,2,req,name=serialNo" json:"serialNo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PacketID) Reset()         { *m = PacketID{} }
func (m *PacketID) String() string { return proto.CompactTextString(m) }
func (*PacketID) ProtoMessage()    {}
func (*PacketID) Descriptor() ([]byte, []int) {
	return fileDescriptor_Common_d9ca241298b4d843, []int{0}
}
func (m *PacketID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PacketID.Unmarshal(m, b)
}
func (m *PacketID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PacketID.Marshal(b, m, deterministic)
}
func (dst *PacketID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketID.Merge(dst, src)
}
func (m *PacketID) XXX_Size() int {
	return xxx_messageInfo_PacketID.Size(m)
}
func (m *PacketID) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketID.DiscardUnknown(m)
}

var xxx_messageInfo_PacketID proto.InternalMessageInfo

func (m *PacketID) GetConnID() uint64 {
	if m != nil && m.ConnID != nil {
		return *m.ConnID
	}
	return 0
}

func (m *PacketID) GetSerialNo() uint32 {
	if m != nil && m.SerialNo != nil {
		return *m.SerialNo
	}
	return 0
}

func init() {
	proto.RegisterType((*PacketID)(nil), "Common.PacketID")
	proto.RegisterEnum("Common.RetType", RetType_name, RetType_value)
}

func init() { proto.RegisterFile("Common.proto", fileDescriptor_Common_d9ca241298b4d843) }

var fileDescriptor_Common_d9ca241298b4d843 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x71, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0xec, 0xb8, 0x38,
	0x02, 0x12, 0x93, 0xb3, 0x53, 0x4b, 0x3c, 0x5d, 0x84, 0xc4, 0xb8, 0xd8, 0x92, 0xf3, 0xf3, 0xf2,
	0x3c, 0x5d, 0x24, 0x18, 0x15, 0x98, 0x34, 0x58, 0x82, 0xa0, 0x3c, 0x21, 0x29, 0x2e, 0x8e, 0xe2,
	0xd4, 0xa2, 0xcc, 0xc4, 0x1c, 0xbf, 0x7c, 0x09, 0x26, 0x05, 0x26, 0x0d, 0xde, 0x20, 0x38, 0x5f,
	0xab, 0x9c, 0x8b, 0x3d, 0x28, 0xb5, 0x24, 0xa4, 0xb2, 0x20, 0x55, 0x48, 0x98, 0x8b, 0x1f, 0xca,
	0x8c, 0x0f, 0x2e, 0x4d, 0x4e, 0x4e, 0x4d, 0x4d, 0x11, 0x60, 0x10, 0x92, 0xe6, 0xe2, 0x83, 0x09,
	0xba, 0x25, 0x66, 0xe6, 0xa4, 0xa6, 0x08, 0xfc, 0x87, 0x01, 0x46, 0x21, 0x19, 0x84, 0x8e, 0x90,
	0xcc, 0xdc, 0x54, 0xff, 0xd2, 0x12, 0x81, 0x39, 0x58, 0x65, 0x43, 0xf3, 0xb2, 0xf3, 0xf2, 0xcb,
	0xf3, 0x04, 0x3e, 0xfc, 0x81, 0xc9, 0x02, 0x02, 0x00, 0x00, 0xff, 0xff, 0x70, 0xaa, 0x1b, 0x81,
	0xcf, 0x00, 0x00, 0x00,
}