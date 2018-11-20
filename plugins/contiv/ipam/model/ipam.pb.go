// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipam.proto

/*
Package model is a generated protocol buffer package.

It is generated from these files:
	ipam.proto

It has these top-level messages:
	AllocatedIP
*/
package model

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

// AllocatedIP represent an IP address allocated for a pod
type AllocatedIP struct {
	// id represents the assigned IP address
	ID uint32 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	// pod is an identifier tied to assigned IP
	Pod string `protobuf:"bytes,2,opt,name=pod" json:"pod,omitempty"`
}

func (m *AllocatedIP) Reset()                    { *m = AllocatedIP{} }
func (m *AllocatedIP) String() string            { return proto.CompactTextString(m) }
func (*AllocatedIP) ProtoMessage()               {}
func (*AllocatedIP) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AllocatedIP) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *AllocatedIP) GetPod() string {
	if m != nil {
		return m.Pod
	}
	return ""
}

func init() {
	proto.RegisterType((*AllocatedIP)(nil), "model.AllocatedIP")
}

func init() { proto.RegisterFile("ipam.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 95 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x2c, 0x48, 0xcc,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0xcd, 0x4f, 0x49, 0xcd, 0x51, 0xd2, 0xe7,
	0xe2, 0x76, 0xcc, 0xc9, 0xc9, 0x4f, 0x4e, 0x2c, 0x49, 0x4d, 0xf1, 0x0c, 0x10, 0xe2, 0xe3, 0x62,
	0xf2, 0x74, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0d, 0x62, 0xf2, 0x74, 0x11, 0x12, 0xe0, 0x62,
	0x2e, 0xc8, 0x4f, 0x91, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x93, 0xd8, 0xc0, 0xda,
	0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x43, 0x71, 0x9e, 0xa6, 0x4c, 0x00, 0x00, 0x00,
}