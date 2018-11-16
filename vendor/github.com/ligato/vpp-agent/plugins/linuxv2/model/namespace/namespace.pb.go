// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: namespace.proto

package linux_namespace

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

type NetNamespace_ReferenceType int32

const (
	NetNamespace_NETNS_REF_UNDEFINED    NetNamespace_ReferenceType = 0
	NetNamespace_NETNS_REF_NSID         NetNamespace_ReferenceType = 1
	NetNamespace_NETNS_REF_PID          NetNamespace_ReferenceType = 2
	NetNamespace_NETNS_REF_FD           NetNamespace_ReferenceType = 3
	NetNamespace_NETNS_REF_MICROSERVICE NetNamespace_ReferenceType = 4
)

var NetNamespace_ReferenceType_name = map[int32]string{
	0: "NETNS_REF_UNDEFINED",
	1: "NETNS_REF_NSID",
	2: "NETNS_REF_PID",
	3: "NETNS_REF_FD",
	4: "NETNS_REF_MICROSERVICE",
}
var NetNamespace_ReferenceType_value = map[string]int32{
	"NETNS_REF_UNDEFINED":    0,
	"NETNS_REF_NSID":         1,
	"NETNS_REF_PID":          2,
	"NETNS_REF_FD":           3,
	"NETNS_REF_MICROSERVICE": 4,
}

func (x NetNamespace_ReferenceType) String() string {
	return proto.EnumName(NetNamespace_ReferenceType_name, int32(x))
}
func (NetNamespace_ReferenceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_namespace_42caec4d73627a40, []int{0, 0}
}

type NetNamespace struct {
	Type                 NetNamespace_ReferenceType `protobuf:"varint,1,opt,name=type,proto3,enum=linux.namespace.NetNamespace_ReferenceType" json:"type,omitempty"`
	Reference            string                     `protobuf:"bytes,2,opt,name=reference,proto3" json:"reference,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *NetNamespace) Reset()         { *m = NetNamespace{} }
func (m *NetNamespace) String() string { return proto.CompactTextString(m) }
func (*NetNamespace) ProtoMessage()    {}
func (*NetNamespace) Descriptor() ([]byte, []int) {
	return fileDescriptor_namespace_42caec4d73627a40, []int{0}
}
func (m *NetNamespace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetNamespace.Unmarshal(m, b)
}
func (m *NetNamespace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetNamespace.Marshal(b, m, deterministic)
}
func (dst *NetNamespace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetNamespace.Merge(dst, src)
}
func (m *NetNamespace) XXX_Size() int {
	return xxx_messageInfo_NetNamespace.Size(m)
}
func (m *NetNamespace) XXX_DiscardUnknown() {
	xxx_messageInfo_NetNamespace.DiscardUnknown(m)
}

var xxx_messageInfo_NetNamespace proto.InternalMessageInfo

func (m *NetNamespace) GetType() NetNamespace_ReferenceType {
	if m != nil {
		return m.Type
	}
	return NetNamespace_NETNS_REF_UNDEFINED
}

func (m *NetNamespace) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func init() {
	proto.RegisterType((*NetNamespace)(nil), "linux.namespace.NetNamespace")
	proto.RegisterEnum("linux.namespace.NetNamespace_ReferenceType", NetNamespace_ReferenceType_name, NetNamespace_ReferenceType_value)
}

func init() { proto.RegisterFile("namespace.proto", fileDescriptor_namespace_42caec4d73627a40) }

var fileDescriptor_namespace_42caec4d73627a40 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x4b, 0xcc, 0x4d,
	0x2d, 0x2e, 0x48, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcf, 0xc9, 0xcc,
	0x2b, 0xad, 0xd0, 0x83, 0x0b, 0x2b, 0xbd, 0x61, 0xe4, 0xe2, 0xf1, 0x4b, 0x2d, 0xf1, 0x83, 0x09,
	0x08, 0xd9, 0x73, 0xb1, 0x94, 0x54, 0x16, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x19, 0x69,
	0xeb, 0xa1, 0x69, 0xd0, 0x43, 0x56, 0xac, 0x17, 0x94, 0x9a, 0x96, 0x5a, 0x94, 0x9a, 0x97, 0x9c,
	0x1a, 0x52, 0x59, 0x90, 0x1a, 0x04, 0xd6, 0x28, 0x24, 0xc3, 0xc5, 0x59, 0x04, 0x13, 0x96, 0x60,
	0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x08, 0x28, 0xd5, 0x72, 0xf1, 0xa2, 0x68, 0x12, 0x12, 0xe7,
	0x12, 0xf6, 0x73, 0x0d, 0xf1, 0x0b, 0x8e, 0x0f, 0x72, 0x75, 0x8b, 0x0f, 0xf5, 0x73, 0x71, 0x75,
	0xf3, 0xf4, 0x73, 0x75, 0x11, 0x60, 0x10, 0x12, 0xe2, 0xe2, 0x43, 0x48, 0xf8, 0x05, 0x7b, 0xba,
	0x08, 0x30, 0x0a, 0x09, 0x72, 0xf1, 0x22, 0xc4, 0x02, 0x3c, 0x5d, 0x04, 0x98, 0x84, 0x04, 0xb8,
	0x78, 0x10, 0x42, 0x6e, 0x2e, 0x02, 0xcc, 0x42, 0x52, 0x5c, 0x62, 0x08, 0x11, 0x5f, 0x4f, 0xe7,
	0x20, 0xff, 0x60, 0xd7, 0xa0, 0x30, 0x4f, 0x67, 0x57, 0x01, 0x96, 0x24, 0x36, 0x70, 0x30, 0x18,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x16, 0x6a, 0x0d, 0x19, 0x01, 0x00, 0x00,
}
