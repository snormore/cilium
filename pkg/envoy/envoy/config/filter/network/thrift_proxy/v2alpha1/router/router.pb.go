// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/thrift_proxy/v2alpha1/router/router.proto

package router

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

// [#protodoc-title: Thrift Router]
// Thrift Router configuration.
type Router struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Router) Reset()         { *m = Router{} }
func (m *Router) String() string { return proto.CompactTextString(m) }
func (*Router) ProtoMessage()    {}
func (*Router) Descriptor() ([]byte, []int) {
	return fileDescriptor_router_67d77d14d4a36955, []int{0}
}
func (m *Router) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Router.Unmarshal(m, b)
}
func (m *Router) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Router.Marshal(b, m, deterministic)
}
func (dst *Router) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Router.Merge(dst, src)
}
func (m *Router) XXX_Size() int {
	return xxx_messageInfo_Router.Size(m)
}
func (m *Router) XXX_DiscardUnknown() {
	xxx_messageInfo_Router.DiscardUnknown(m)
}

var xxx_messageInfo_Router proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Router)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.router.Router")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/thrift_proxy/v2alpha1/router/router.proto", fileDescriptor_router_67d77d14d4a36955)
}

var fileDescriptor_router_67d77d14d4a36955 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4d, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x2f, 0xc9, 0x28, 0xca, 0x4c, 0x2b, 0x89, 0x2f,
	0x28, 0xca, 0xaf, 0xa8, 0xd4, 0x2f, 0x33, 0x4a, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0xd4, 0x2f, 0xca,
	0x2f, 0x05, 0x29, 0x82, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x16, 0x60, 0x63, 0xf4,
	0x20, 0xc6, 0xe8, 0x41, 0x8c, 0xd1, 0x83, 0x1a, 0xa3, 0x87, 0x6c, 0x8c, 0x1e, 0xcc, 0x18, 0x3d,
	0x88, 0x7e, 0x25, 0x0e, 0x2e, 0xb6, 0x20, 0x30, 0xcb, 0x89, 0x23, 0x8a, 0x0d, 0x22, 0x96, 0xc4,
	0x06, 0x36, 0xd4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x77, 0x03, 0xff, 0x08, 0x9d, 0x00, 0x00,
	0x00,
}