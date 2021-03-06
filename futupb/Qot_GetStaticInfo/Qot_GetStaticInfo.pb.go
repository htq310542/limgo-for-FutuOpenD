// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetStaticInfo/Qot_GetStaticInfo.proto

package Qot_GetStaticInfo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "limgo/futupb/Common"
import Qot_Common "limgo/futupb/Qot_Common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type C2S struct {
	Market               *int32   `protobuf:"varint,1,opt,name=market" json:"market,omitempty"`
	SecType              *int32   `protobuf:"varint,2,opt,name=secType" json:"secType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetStaticInfo_dcc2e39209a9ae41, []int{0}
}
func (m *C2S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2S.Unmarshal(m, b)
}
func (m *C2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2S.Marshal(b, m, deterministic)
}
func (dst *C2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S.Merge(dst, src)
}
func (m *C2S) XXX_Size() int {
	return xxx_messageInfo_C2S.Size(m)
}
func (m *C2S) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S.DiscardUnknown(m)
}

var xxx_messageInfo_C2S proto.InternalMessageInfo

func (m *C2S) GetMarket() int32 {
	if m != nil && m.Market != nil {
		return *m.Market
	}
	return 0
}

func (m *C2S) GetSecType() int32 {
	if m != nil && m.SecType != nil {
		return *m.SecType
	}
	return 0
}

type S2C struct {
	StaticInfoList       []*Qot_Common.SecurityStaticInfo `protobuf:"bytes,1,rep,name=staticInfoList" json:"staticInfoList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetStaticInfo_dcc2e39209a9ae41, []int{1}
}
func (m *S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C.Unmarshal(m, b)
}
func (m *S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C.Marshal(b, m, deterministic)
}
func (dst *S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C.Merge(dst, src)
}
func (m *S2C) XXX_Size() int {
	return xxx_messageInfo_S2C.Size(m)
}
func (m *S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_S2C proto.InternalMessageInfo

func (m *S2C) GetStaticInfoList() []*Qot_Common.SecurityStaticInfo {
	if m != nil {
		return m.StaticInfoList
	}
	return nil
}

type Request struct {
	C2S                  *C2S     `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetStaticInfo_dcc2e39209a9ae41, []int{2}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetC2S() *C2S {
	if m != nil {
		return m.C2S
	}
	return nil
}

type Response struct {
	RetType              *int32   `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"`
	RetMsg               *string  `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode              *int32   `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C                  *S2C     `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetStaticInfo_dcc2e39209a9ae41, []int{3}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

const Default_Response_RetType int32 = -400

func (m *Response) GetRetType() int32 {
	if m != nil && m.RetType != nil {
		return *m.RetType
	}
	return Default_Response_RetType
}

func (m *Response) GetRetMsg() string {
	if m != nil && m.RetMsg != nil {
		return *m.RetMsg
	}
	return ""
}

func (m *Response) GetErrCode() int32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *Response) GetS2C() *S2C {
	if m != nil {
		return m.S2C
	}
	return nil
}

func init() {
	proto.RegisterType((*C2S)(nil), "Qot_GetStaticInfo.C2S")
	proto.RegisterType((*S2C)(nil), "Qot_GetStaticInfo.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetStaticInfo.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetStaticInfo.Response")
}

func init() {
	proto.RegisterFile("Qot_GetStaticInfo/Qot_GetStaticInfo.proto", fileDescriptor_Qot_GetStaticInfo_dcc2e39209a9ae41)
}

var fileDescriptor_Qot_GetStaticInfo_dcc2e39209a9ae41 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0x49, 0xd3, 0xfe, 0xf3, 0x77, 0x0a, 0x82, 0x2b, 0x94, 0xa5, 0x42, 0x09, 0x39, 0xc5,
	0x83, 0x69, 0x59, 0x05, 0xc1, 0xeb, 0x82, 0x22, 0xd8, 0x83, 0xbb, 0xde, 0xa5, 0xac, 0xa3, 0x04,
	0x49, 0x36, 0xee, 0x4e, 0x0f, 0xfd, 0x02, 0x7e, 0x6e, 0xd9, 0x24, 0xb5, 0x62, 0xf0, 0x94, 0xbc,
	0x79, 0x6f, 0x76, 0xde, 0x0f, 0xce, 0x1f, 0x2d, 0x3d, 0xdf, 0x21, 0x69, 0xda, 0x50, 0x69, 0xee,
	0xeb, 0x57, 0xbb, 0x1c, 0x4c, 0x8a, 0xc6, 0x59, 0xb2, 0xec, 0x64, 0x60, 0xcc, 0x4f, 0xa5, 0xad,
	0x2a, 0x5b, 0x2f, 0xbb, 0x4f, 0x97, 0x9b, 0x9f, 0x85, 0x5c, 0x6f, 0x1c, 0x7e, 0x3b, 0x33, 0xbb,
	0x86, 0x58, 0x0a, 0xcd, 0x66, 0xf0, 0xaf, 0xda, 0xb8, 0x77, 0x24, 0x1e, 0xa5, 0x51, 0x3e, 0x51,
	0xbd, 0x62, 0x1c, 0x12, 0x8f, 0xe6, 0x69, 0xd7, 0x20, 0x1f, 0xb5, 0xc6, 0x5e, 0x66, 0x6b, 0x88,
	0xb5, 0x90, 0xec, 0x16, 0x8e, 0xfd, 0xf7, 0xfd, 0x87, 0xd2, 0x87, 0x07, 0xe2, 0x7c, 0x2a, 0x16,
	0xc5, 0x8f, 0x53, 0x1a, 0xcd, 0xd6, 0x95, 0xb4, 0x3b, 0x34, 0x55, 0xbf, 0xb6, 0xb2, 0x4b, 0x48,
	0x14, 0x7e, 0x6c, 0xd1, 0x13, 0xcb, 0x21, 0x36, 0xc2, 0xf3, 0x28, 0x1d, 0xe5, 0x53, 0x31, 0x2b,
	0x86, 0xf8, 0x52, 0x68, 0x15, 0x22, 0xd9, 0x67, 0x04, 0xff, 0x15, 0xfa, 0xc6, 0xd6, 0x1e, 0xd9,
	0x02, 0x12, 0x87, 0xd4, 0x56, 0x0d, 0xab, 0x93, 0x9b, 0xf1, 0xc5, 0xd5, 0x6a, 0xa5, 0xf6, 0xc3,
	0x80, 0xe8, 0x90, 0xd6, 0xfe, 0xad, 0x25, 0x39, 0x52, 0xbd, 0x0a, 0x88, 0xe8, 0x9c, 0xb4, 0x2f,
	0xc8, 0xe3, 0x0e, 0xb1, 0x97, 0xa1, 0x88, 0x17, 0x86, 0x8f, 0xd3, 0xe8, 0x8f, 0x22, 0x5a, 0x48,
	0x15, 0x22, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x82, 0x5c, 0x0c, 0x14, 0xb6, 0x01, 0x00, 0x00,
}
