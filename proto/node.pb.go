// Code generated by protoc-gen-go.
// source: node.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// node/ping
type Ping struct {
}

func (m *Ping) Reset()                    { *m = Ping{} }
func (m *Ping) String() string            { return proto1.CompactTextString(m) }
func (*Ping) ProtoMessage()               {}
func (*Ping) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type Pong struct {
}

func (m *Pong) Reset()                    { *m = Pong{} }
func (m *Pong) String() string            { return proto1.CompactTextString(m) }
func (*Pong) ProtoMessage()               {}
func (*Pong) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func init() {
	proto1.RegisterType((*Ping)(nil), "proto.Ping")
	proto1.RegisterType((*Pong)(nil), "proto.Pong")
}

func init() { proto1.RegisterFile("node.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 57 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0xcb, 0x4f, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x6c, 0x5c, 0x2c, 0x01, 0x99,
	0x79, 0xe9, 0x60, 0x3a, 0x3f, 0x2f, 0x3d, 0x89, 0x0d, 0x2c, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0xfe, 0xe2, 0xf3, 0xe8, 0x2b, 0x00, 0x00, 0x00,
}
