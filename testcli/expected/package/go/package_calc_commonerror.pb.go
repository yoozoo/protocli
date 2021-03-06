// Code generated by protoc-gen-go. DO NOT EDIT.
// source: package_calc_commonerror.proto

package calcsvrmain

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AddReq struct {
	X                    int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddReq) Reset()         { *m = AddReq{} }
func (m *AddReq) String() string { return proto.CompactTextString(m) }
func (*AddReq) ProtoMessage()    {}
func (*AddReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_19293775ba74a490, []int{0}
}

func (m *AddReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddReq.Unmarshal(m, b)
}
func (m *AddReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddReq.Marshal(b, m, deterministic)
}
func (m *AddReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddReq.Merge(m, src)
}
func (m *AddReq) XXX_Size() int {
	return xxx_messageInfo_AddReq.Size(m)
}
func (m *AddReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddReq proto.InternalMessageInfo

func (m *AddReq) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *AddReq) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

type AddResp struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResp) Reset()         { *m = AddResp{} }
func (m *AddResp) String() string { return proto.CompactTextString(m) }
func (*AddResp) ProtoMessage()    {}
func (*AddResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_19293775ba74a490, []int{1}
}

func (m *AddResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResp.Unmarshal(m, b)
}
func (m *AddResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResp.Marshal(b, m, deterministic)
}
func (m *AddResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResp.Merge(m, src)
}
func (m *AddResp) XXX_Size() int {
	return xxx_messageInfo_AddResp.Size(m)
}
func (m *AddResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResp.DiscardUnknown(m)
}

var xxx_messageInfo_AddResp proto.InternalMessageInfo

func (m *AddResp) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type AddError struct {
	Req                  *AddReq  `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddError) Reset()         { *m = AddError{} }
func (m *AddError) String() string { return proto.CompactTextString(m) }
func (*AddError) ProtoMessage()    {}
func (*AddError) Descriptor() ([]byte, []int) {
	return fileDescriptor_19293775ba74a490, []int{2}
}

func (m *AddError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddError.Unmarshal(m, b)
}
func (m *AddError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddError.Marshal(b, m, deterministic)
}
func (m *AddError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddError.Merge(m, src)
}
func (m *AddError) XXX_Size() int {
	return xxx_messageInfo_AddError.Size(m)
}
func (m *AddError) XXX_DiscardUnknown() {
	xxx_messageInfo_AddError.DiscardUnknown(m)
}

var xxx_messageInfo_AddError proto.InternalMessageInfo

func (m *AddError) GetReq() *AddReq {
	if m != nil {
		return m.Req
	}
	return nil
}

func (m *AddError) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*AddReq)(nil), "calcmain.AddReq")
	proto.RegisterType((*AddResp)(nil), "calcmain.AddResp")
	proto.RegisterType((*AddError)(nil), "calcmain.AddError")
}

func init() { proto.RegisterFile("package_calc_commonerror.proto", fileDescriptor_19293775ba74a490) }

var fileDescriptor_19293775ba74a490 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x48, 0x4c, 0xce,
	0x4e, 0x4c, 0x4f, 0x8d, 0x4f, 0x4e, 0xcc, 0x49, 0x8e, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x4b,
	0x2d, 0x2a, 0xca, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x89, 0xe7, 0x26,
	0x66, 0xe6, 0x49, 0xf1, 0x40, 0x24, 0x21, 0xe2, 0x52, 0x22, 0x30, 0x7d, 0x89, 0x29, 0x29, 0x41,
	0xa9, 0x85, 0x10, 0x51, 0x25, 0x15, 0x2e, 0x36, 0x47, 0x30, 0x5f, 0x88, 0x87, 0x8b, 0xb1, 0x42,
	0x82, 0x51, 0x81, 0x51, 0x83, 0x35, 0x88, 0xb1, 0x02, 0xc4, 0xab, 0x94, 0x60, 0x82, 0xf0, 0x2a,
	0x95, 0x14, 0xb9, 0xd8, 0xc1, 0xaa, 0x8a, 0x0b, 0x84, 0xc4, 0xb8, 0xd8, 0x8a, 0x52, 0x8b, 0x4b,
	0x73, 0x4a, 0xa0, 0x6a, 0xa1, 0x3c, 0x25, 0x17, 0x2e, 0x0e, 0xc7, 0x94, 0x14, 0x57, 0x90, 0x43,
	0x84, 0x94, 0xb8, 0x98, 0x8b, 0x52, 0x0b, 0xc1, 0x0a, 0xb8, 0x8d, 0x04, 0xf4, 0x60, 0x0e, 0xd2,
	0x83, 0xd8, 0x14, 0x04, 0x92, 0x14, 0x12, 0xe1, 0x62, 0x05, 0xbb, 0x1a, 0x6c, 0x09, 0x67, 0x10,
	0x84, 0x63, 0xd4, 0xcf, 0xc8, 0xc5, 0xed, 0x9c, 0x98, 0x93, 0x1c, 0x9c, 0x5a, 0x54, 0x96, 0x99,
	0x9c, 0x2a, 0x64, 0xc1, 0xc5, 0x9c, 0x98, 0x92, 0x22, 0x84, 0x61, 0x86, 0x94, 0x20, 0x9a, 0x48,
	0x71, 0x81, 0x12, 0xcf, 0xaf, 0xcf, 0x12, 0x08, 0x37, 0x98, 0x72, 0xb1, 0x24, 0xa6, 0xa4, 0x18,
	0x09, 0xf1, 0x80, 0x15, 0x12, 0xab, 0x4d, 0x8a, 0xbf, 0xe9, 0x8b, 0x04, 0xb7, 0x33, 0x38, 0xdc,
	0xc0, 0x02, 0x4e, 0xbc, 0x51, 0xdc, 0x20, 0x2d, 0xc5, 0x65, 0x45, 0x20, 0x5d, 0x49, 0x6c, 0xe0,
	0x60, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x22, 0x5d, 0x42, 0x86, 0x01, 0x00, 0x00,
}
