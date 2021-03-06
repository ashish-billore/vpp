// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sfc.proto

package sfc

/*
Package pod defines data model for Kubernetes Pod.
*/

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

// Pod is a collection of containers that can run on a host.
// This resource is created by clients and scheduled onto hosts.
type Sfc struct {
	// Name of the pod unique within the namespace.
	// Cannot be updated.
	Pod string `protobuf:"bytes,1,opt,name=pod,proto3" json:"pod,omitempty"`
	// Node the pod is deployed into into.
	// Can be updated
	Node                 string   `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sfc) Reset()         { *m = Sfc{} }
func (m *Sfc) String() string { return proto.CompactTextString(m) }
func (*Sfc) ProtoMessage()    {}
func (*Sfc) Descriptor() ([]byte, []int) {
	return fileDescriptor_sfc_1bc893331ea1c04f, []int{0}
}
func (m *Sfc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sfc.Unmarshal(m, b)
}
func (m *Sfc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sfc.Marshal(b, m, deterministic)
}
func (dst *Sfc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sfc.Merge(dst, src)
}
func (m *Sfc) XXX_Size() int {
	return xxx_messageInfo_Sfc.Size(m)
}
func (m *Sfc) XXX_DiscardUnknown() {
	xxx_messageInfo_Sfc.DiscardUnknown(m)
}

var xxx_messageInfo_Sfc proto.InternalMessageInfo

func (m *Sfc) GetPod() string {
	if m != nil {
		return m.Pod
	}
	return ""
}

func (m *Sfc) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func init() {
	proto.RegisterType((*Sfc)(nil), "sfc.Sfc")
}

func init() { proto.RegisterFile("sfc.proto", fileDescriptor_sfc_1bc893331ea1c04f) }

var fileDescriptor_sfc_1bc893331ea1c04f = []byte{
	// 81 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x4e, 0x4b, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x4e, 0x4b, 0x56, 0xd2, 0xe6, 0x62, 0x0e, 0x4e,
	0x4b, 0x16, 0x12, 0xe0, 0x62, 0x2e, 0xc8, 0x4f, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02,
	0x31, 0x85, 0x84, 0xb8, 0x58, 0xf2, 0xf2, 0x53, 0x52, 0x25, 0x98, 0xc0, 0x42, 0x60, 0x76, 0x12,
	0x1b, 0x58, 0xa3, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x75, 0x16, 0x53, 0x45, 0x00, 0x00,
	0x00,
}
