// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: l3.proto

package linux_l3

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type StaticRoute_Scope int32

const (
	StaticRoute_UNDEFINED StaticRoute_Scope = 0
	StaticRoute_GLOBAL    StaticRoute_Scope = 1
	StaticRoute_SITE      StaticRoute_Scope = 2
	StaticRoute_LINK      StaticRoute_Scope = 3
	StaticRoute_HOST      StaticRoute_Scope = 4
)

var StaticRoute_Scope_name = map[int32]string{
	0: "UNDEFINED",
	1: "GLOBAL",
	2: "SITE",
	3: "LINK",
	4: "HOST",
}
var StaticRoute_Scope_value = map[string]int32{
	"UNDEFINED": 0,
	"GLOBAL":    1,
	"SITE":      2,
	"LINK":      3,
	"HOST":      4,
}

func (x StaticRoute_Scope) String() string {
	return proto.EnumName(StaticRoute_Scope_name, int32(x))
}
func (StaticRoute_Scope) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_l3_0ba4bb4c78f0dd5f, []int{0, 0}
}

type StaticRoute struct {
	OutgoingInterface    string            `protobuf:"bytes,1,opt,name=outgoing_interface,json=outgoingInterface,proto3" json:"outgoing_interface,omitempty"`
	Scope                StaticRoute_Scope `protobuf:"varint,2,opt,name=scope,proto3,enum=linux.l3.StaticRoute_Scope" json:"scope,omitempty"`
	DstNetwork           string            `protobuf:"bytes,3,opt,name=dst_network,json=dstNetwork,proto3" json:"dst_network,omitempty"`
	GwAddr               string            `protobuf:"bytes,4,opt,name=gw_addr,json=gwAddr,proto3" json:"gw_addr,omitempty"`
	Metric               uint32            `protobuf:"varint,5,opt,name=metric,proto3" json:"metric,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StaticRoute) Reset()         { *m = StaticRoute{} }
func (m *StaticRoute) String() string { return proto.CompactTextString(m) }
func (*StaticRoute) ProtoMessage()    {}
func (*StaticRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_l3_0ba4bb4c78f0dd5f, []int{0}
}
func (m *StaticRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StaticRoute.Unmarshal(m, b)
}
func (m *StaticRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StaticRoute.Marshal(b, m, deterministic)
}
func (dst *StaticRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StaticRoute.Merge(dst, src)
}
func (m *StaticRoute) XXX_Size() int {
	return xxx_messageInfo_StaticRoute.Size(m)
}
func (m *StaticRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_StaticRoute.DiscardUnknown(m)
}

var xxx_messageInfo_StaticRoute proto.InternalMessageInfo

func (m *StaticRoute) GetOutgoingInterface() string {
	if m != nil {
		return m.OutgoingInterface
	}
	return ""
}

func (m *StaticRoute) GetScope() StaticRoute_Scope {
	if m != nil {
		return m.Scope
	}
	return StaticRoute_UNDEFINED
}

func (m *StaticRoute) GetDstNetwork() string {
	if m != nil {
		return m.DstNetwork
	}
	return ""
}

func (m *StaticRoute) GetGwAddr() string {
	if m != nil {
		return m.GwAddr
	}
	return ""
}

func (m *StaticRoute) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

type StaticARPEntry struct {
	Interface            string   `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
	IpAddress            string   `protobuf:"bytes,2,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	HwAddress            string   `protobuf:"bytes,3,opt,name=hw_address,json=hwAddress,proto3" json:"hw_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StaticARPEntry) Reset()         { *m = StaticARPEntry{} }
func (m *StaticARPEntry) String() string { return proto.CompactTextString(m) }
func (*StaticARPEntry) ProtoMessage()    {}
func (*StaticARPEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_l3_0ba4bb4c78f0dd5f, []int{1}
}
func (m *StaticARPEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StaticARPEntry.Unmarshal(m, b)
}
func (m *StaticARPEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StaticARPEntry.Marshal(b, m, deterministic)
}
func (dst *StaticARPEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StaticARPEntry.Merge(dst, src)
}
func (m *StaticARPEntry) XXX_Size() int {
	return xxx_messageInfo_StaticARPEntry.Size(m)
}
func (m *StaticARPEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_StaticARPEntry.DiscardUnknown(m)
}

var xxx_messageInfo_StaticARPEntry proto.InternalMessageInfo

func (m *StaticARPEntry) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *StaticARPEntry) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *StaticARPEntry) GetHwAddress() string {
	if m != nil {
		return m.HwAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*StaticRoute)(nil), "linux.l3.StaticRoute")
	proto.RegisterType((*StaticARPEntry)(nil), "linux.l3.StaticARPEntry")
	proto.RegisterEnum("linux.l3.StaticRoute_Scope", StaticRoute_Scope_name, StaticRoute_Scope_value)
}

func init() { proto.RegisterFile("l3.proto", fileDescriptor_l3_0ba4bb4c78f0dd5f) }

var fileDescriptor_l3_0ba4bb4c78f0dd5f = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x4f, 0x4f, 0x83, 0x40,
	0x10, 0xc5, 0xa5, 0x7f, 0xb0, 0x4c, 0xd3, 0x06, 0xe7, 0xa0, 0x24, 0x6a, 0x6c, 0x7a, 0xea, 0x45,
	0x12, 0xed, 0x17, 0x10, 0x53, 0x54, 0x62, 0x43, 0x0d, 0xd4, 0x33, 0x41, 0x76, 0xa5, 0x1b, 0x91,
	0x25, 0xcb, 0x12, 0xf4, 0xec, 0x17, 0x37, 0x2c, 0x45, 0x8d, 0xb7, 0xd9, 0xf7, 0x9b, 0xbc, 0x7d,
	0x2f, 0x03, 0xa3, 0x6c, 0x69, 0x17, 0x82, 0x4b, 0x8e, 0xa3, 0x8c, 0xe5, 0xd5, 0x87, 0x9d, 0x2d,
	0xe7, 0x5f, 0x3d, 0x18, 0x87, 0x32, 0x96, 0x2c, 0x09, 0x78, 0x25, 0x29, 0x5e, 0x02, 0xf2, 0x4a,
	0xa6, 0x9c, 0xe5, 0x69, 0xc4, 0x72, 0x49, 0xc5, 0x6b, 0x9c, 0x50, 0x4b, 0x9b, 0x69, 0x0b, 0x23,
	0x38, 0xea, 0x88, 0xd7, 0x01, 0xbc, 0x82, 0x61, 0x99, 0xf0, 0x82, 0x5a, 0xbd, 0x99, 0xb6, 0x98,
	0x5e, 0x9f, 0xda, 0x9d, 0xb1, 0xfd, 0xc7, 0xd4, 0x0e, 0x9b, 0x95, 0xa0, 0xdd, 0xc4, 0x0b, 0x18,
	0x93, 0x52, 0x46, 0x39, 0x95, 0x35, 0x17, 0x6f, 0x56, 0x5f, 0x59, 0x03, 0x29, 0xa5, 0xdf, 0x2a,
	0x78, 0x02, 0x87, 0x69, 0x1d, 0xc5, 0x84, 0x08, 0x6b, 0xa0, 0xa0, 0x9e, 0xd6, 0x0e, 0x21, 0x02,
	0x8f, 0x41, 0x7f, 0xa7, 0x52, 0xb0, 0xc4, 0x1a, 0xce, 0xb4, 0xc5, 0x24, 0xd8, 0xbf, 0xe6, 0x37,
	0x30, 0x54, 0x3f, 0xe0, 0x04, 0x8c, 0x67, 0x7f, 0xe5, 0xde, 0x79, 0xbe, 0xbb, 0x32, 0x0f, 0x10,
	0x40, 0xbf, 0x5f, 0x6f, 0x6e, 0x9d, 0xb5, 0xa9, 0xe1, 0x08, 0x06, 0xa1, 0xb7, 0x75, 0xcd, 0x5e,
	0x33, 0xad, 0x3d, 0xff, 0xd1, 0xec, 0x37, 0xd3, 0xc3, 0x26, 0xdc, 0x9a, 0x83, 0x79, 0x06, 0xd3,
	0x36, 0xaf, 0x13, 0x3c, 0xb9, 0xb9, 0x14, 0x9f, 0x78, 0x06, 0xc6, 0xff, 0xfa, 0xbf, 0x02, 0x9e,
	0x03, 0xb0, 0x42, 0x45, 0xa4, 0x65, 0xa9, 0xba, 0x37, 0xb8, 0x70, 0x5a, 0xa1, 0xc1, 0xbb, 0xfa,
	0x07, 0xb7, 0x0d, 0x8d, 0x5d, 0xbd, 0xc7, 0x2f, 0xba, 0x3a, 0xc2, 0xf2, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0x13, 0x86, 0xfd, 0x10, 0x90, 0x01, 0x00, 0x00,
}
