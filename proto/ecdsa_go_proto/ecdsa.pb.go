// Copyright 2017 Google Inc.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//      http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
////////////////////////////////////////////////////////////////////////////////

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/ecdsa.proto

package ecdsa_go_proto // import "github.com/google/tink/proto/ecdsa_go_proto"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common_go_proto "github.com/google/tink/proto/common_go_proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EcdsaSignatureEncoding int32

const (
	EcdsaSignatureEncoding_UNKNOWN_ENCODING EcdsaSignatureEncoding = 0
	// The signature's format is r || s, where r and s are zero-padded and have the same size in
	// bytes as the order of the curve. For example, for NIST P-256 curve, r and s are zero-padded to
	// 32 bytes.
	EcdsaSignatureEncoding_IEEE_P1363 EcdsaSignatureEncoding = 1
	// The signature is encoded using ASN.1
	// (https://tools.ietf.org/html/rfc5480#appendix-A):
	// ECDSA-Sig-Value :: = SEQUENCE {
	//  r INTEGER,
	//  s INTEGER
	// }
	EcdsaSignatureEncoding_DER EcdsaSignatureEncoding = 2
)

var EcdsaSignatureEncoding_name = map[int32]string{
	0: "UNKNOWN_ENCODING",
	1: "IEEE_P1363",
	2: "DER",
}

var EcdsaSignatureEncoding_value = map[string]int32{
	"UNKNOWN_ENCODING": 0,
	"IEEE_P1363":       1,
	"DER":              2,
}

func (x EcdsaSignatureEncoding) String() string {
	return proto.EnumName(EcdsaSignatureEncoding_name, int32(x))
}

func (EcdsaSignatureEncoding) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b92ace61c26b8fc2, []int{0}
}

// Protos for Ecdsa.
type EcdsaParams struct {
	// Required.
	HashType common_go_proto.HashType `protobuf:"varint,1,opt,name=hash_type,json=hashType,proto3,enum=google.crypto.tink.HashType" json:"hash_type,omitempty"`
	// Required.
	Curve common_go_proto.EllipticCurveType `protobuf:"varint,2,opt,name=curve,proto3,enum=google.crypto.tink.EllipticCurveType" json:"curve,omitempty"`
	// Required.
	Encoding             EcdsaSignatureEncoding `protobuf:"varint,3,opt,name=encoding,proto3,enum=google.crypto.tink.EcdsaSignatureEncoding" json:"encoding,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *EcdsaParams) Reset()         { *m = EcdsaParams{} }
func (m *EcdsaParams) String() string { return proto.CompactTextString(m) }
func (*EcdsaParams) ProtoMessage()    {}
func (*EcdsaParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_b92ace61c26b8fc2, []int{0}
}
func (m *EcdsaParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EcdsaParams.Unmarshal(m, b)
}
func (m *EcdsaParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EcdsaParams.Marshal(b, m, deterministic)
}
func (dst *EcdsaParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EcdsaParams.Merge(dst, src)
}
func (m *EcdsaParams) XXX_Size() int {
	return xxx_messageInfo_EcdsaParams.Size(m)
}
func (m *EcdsaParams) XXX_DiscardUnknown() {
	xxx_messageInfo_EcdsaParams.DiscardUnknown(m)
}

var xxx_messageInfo_EcdsaParams proto.InternalMessageInfo

func (m *EcdsaParams) GetHashType() common_go_proto.HashType {
	if m != nil {
		return m.HashType
	}
	return common_go_proto.HashType_UNKNOWN_HASH
}

func (m *EcdsaParams) GetCurve() common_go_proto.EllipticCurveType {
	if m != nil {
		return m.Curve
	}
	return common_go_proto.EllipticCurveType_UNKNOWN_CURVE
}

func (m *EcdsaParams) GetEncoding() EcdsaSignatureEncoding {
	if m != nil {
		return m.Encoding
	}
	return EcdsaSignatureEncoding_UNKNOWN_ENCODING
}

// key_type: type.googleapis.com/google.crypto.tink.EcdsaPublicKey
type EcdsaPublicKey struct {
	// Required.
	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Required.
	Params *EcdsaParams `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	// Affine coordinates of the public key in bigendian representation. The
	// public key is a point (x, y) on the curve defined by params.curve. For
	// ECDH, it is crucial to verify whether the public key point (x, y) is on the
	// private's key curve. For ECDSA, such verification is a defense in depth.
	// Required.
	X []byte `protobuf:"bytes,3,opt,name=x,proto3" json:"x,omitempty"`
	// Required.
	Y                    []byte   `protobuf:"bytes,4,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EcdsaPublicKey) Reset()         { *m = EcdsaPublicKey{} }
func (m *EcdsaPublicKey) String() string { return proto.CompactTextString(m) }
func (*EcdsaPublicKey) ProtoMessage()    {}
func (*EcdsaPublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_b92ace61c26b8fc2, []int{1}
}
func (m *EcdsaPublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EcdsaPublicKey.Unmarshal(m, b)
}
func (m *EcdsaPublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EcdsaPublicKey.Marshal(b, m, deterministic)
}
func (dst *EcdsaPublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EcdsaPublicKey.Merge(dst, src)
}
func (m *EcdsaPublicKey) XXX_Size() int {
	return xxx_messageInfo_EcdsaPublicKey.Size(m)
}
func (m *EcdsaPublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_EcdsaPublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_EcdsaPublicKey proto.InternalMessageInfo

func (m *EcdsaPublicKey) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *EcdsaPublicKey) GetParams() *EcdsaParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *EcdsaPublicKey) GetX() []byte {
	if m != nil {
		return m.X
	}
	return nil
}

func (m *EcdsaPublicKey) GetY() []byte {
	if m != nil {
		return m.Y
	}
	return nil
}

// key_type: type.googleapis.com/google.crypto.tink.EcdsaPrivateKey
type EcdsaPrivateKey struct {
	// Required.
	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Required.
	PublicKey *EcdsaPublicKey `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// Unsigned big integer in bigendian representation.
	// Required.
	KeyValue             []byte   `protobuf:"bytes,3,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EcdsaPrivateKey) Reset()         { *m = EcdsaPrivateKey{} }
func (m *EcdsaPrivateKey) String() string { return proto.CompactTextString(m) }
func (*EcdsaPrivateKey) ProtoMessage()    {}
func (*EcdsaPrivateKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_b92ace61c26b8fc2, []int{2}
}
func (m *EcdsaPrivateKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EcdsaPrivateKey.Unmarshal(m, b)
}
func (m *EcdsaPrivateKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EcdsaPrivateKey.Marshal(b, m, deterministic)
}
func (dst *EcdsaPrivateKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EcdsaPrivateKey.Merge(dst, src)
}
func (m *EcdsaPrivateKey) XXX_Size() int {
	return xxx_messageInfo_EcdsaPrivateKey.Size(m)
}
func (m *EcdsaPrivateKey) XXX_DiscardUnknown() {
	xxx_messageInfo_EcdsaPrivateKey.DiscardUnknown(m)
}

var xxx_messageInfo_EcdsaPrivateKey proto.InternalMessageInfo

func (m *EcdsaPrivateKey) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *EcdsaPrivateKey) GetPublicKey() *EcdsaPublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *EcdsaPrivateKey) GetKeyValue() []byte {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

type EcdsaKeyFormat struct {
	// Required.
	Params               *EcdsaParams `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *EcdsaKeyFormat) Reset()         { *m = EcdsaKeyFormat{} }
func (m *EcdsaKeyFormat) String() string { return proto.CompactTextString(m) }
func (*EcdsaKeyFormat) ProtoMessage()    {}
func (*EcdsaKeyFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_b92ace61c26b8fc2, []int{3}
}
func (m *EcdsaKeyFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EcdsaKeyFormat.Unmarshal(m, b)
}
func (m *EcdsaKeyFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EcdsaKeyFormat.Marshal(b, m, deterministic)
}
func (dst *EcdsaKeyFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EcdsaKeyFormat.Merge(dst, src)
}
func (m *EcdsaKeyFormat) XXX_Size() int {
	return xxx_messageInfo_EcdsaKeyFormat.Size(m)
}
func (m *EcdsaKeyFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_EcdsaKeyFormat.DiscardUnknown(m)
}

var xxx_messageInfo_EcdsaKeyFormat proto.InternalMessageInfo

func (m *EcdsaKeyFormat) GetParams() *EcdsaParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterType((*EcdsaParams)(nil), "google.crypto.tink.EcdsaParams")
	proto.RegisterType((*EcdsaPublicKey)(nil), "google.crypto.tink.EcdsaPublicKey")
	proto.RegisterType((*EcdsaPrivateKey)(nil), "google.crypto.tink.EcdsaPrivateKey")
	proto.RegisterType((*EcdsaKeyFormat)(nil), "google.crypto.tink.EcdsaKeyFormat")
	proto.RegisterEnum("google.crypto.tink.EcdsaSignatureEncoding", EcdsaSignatureEncoding_name, EcdsaSignatureEncoding_value)
}

func init() { proto.RegisterFile("proto/ecdsa.proto", fileDescriptor_b92ace61c26b8fc2) }

var fileDescriptor_b92ace61c26b8fc2 = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xd1, 0x6a, 0xd4, 0x40,
	0x14, 0x86, 0x9d, 0xad, 0x6e, 0x77, 0x4f, 0xeb, 0xba, 0x0e, 0x22, 0x41, 0x0b, 0x4a, 0x40, 0x90,
	0x0a, 0x59, 0xec, 0x82, 0x22, 0x5e, 0xd9, 0x76, 0x5a, 0x97, 0x85, 0x74, 0x89, 0xb5, 0x82, 0x37,
	0x61, 0x76, 0x3a, 0x24, 0xc3, 0x26, 0x99, 0x61, 0x32, 0x59, 0x3a, 0x57, 0x3e, 0x80, 0x6f, 0xe1,
	0xfb, 0xf8, 0x4e, 0x92, 0x99, 0xb4, 0x08, 0x6e, 0xbd, 0xe8, 0xdd, 0xfc, 0xe4, 0xff, 0x72, 0xfe,
	0xff, 0x70, 0xe0, 0xb1, 0xd2, 0xd2, 0xc8, 0x09, 0x67, 0x97, 0x35, 0x8d, 0xdc, 0x1b, 0xe3, 0x4c,
	0xca, 0xac, 0xe0, 0x11, 0xd3, 0x56, 0x19, 0x19, 0x19, 0x51, 0xad, 0x9e, 0x61, 0x6f, 0x63, 0xb2,
	0x2c, 0x65, 0xe5, 0x7d, 0xe1, 0x6f, 0x04, 0x3b, 0xa4, 0xe5, 0x16, 0x54, 0xd3, 0xb2, 0xc6, 0x1f,
	0x60, 0x98, 0xd3, 0x3a, 0x4f, 0x8d, 0x55, 0x3c, 0x40, 0x2f, 0xd1, 0xeb, 0xd1, 0xc1, 0x5e, 0xf4,
	0xef, 0xbf, 0xa2, 0xcf, 0xb4, 0xce, 0xcf, 0xad, 0xe2, 0xc9, 0x20, 0xef, 0x5e, 0xf8, 0x23, 0x3c,
	0x60, 0x8d, 0x5e, 0xf3, 0xa0, 0xe7, 0xb0, 0x57, 0x9b, 0x30, 0x52, 0x14, 0x42, 0x19, 0xc1, 0x8e,
	0x5a, 0xa3, 0xe3, 0x3d, 0x83, 0x4f, 0x60, 0xc0, 0x2b, 0x26, 0x2f, 0x45, 0x95, 0x05, 0x5b, 0x8e,
	0xdf, 0xdf, 0xc8, 0xb7, 0x51, 0xbf, 0x88, 0xac, 0xa2, 0xa6, 0xd1, 0x9c, 0x74, 0x44, 0x72, 0xc3,
	0x86, 0x3f, 0x60, 0xe4, 0xeb, 0x34, 0xcb, 0x42, 0xb0, 0x39, 0xb7, 0x38, 0x80, 0xed, 0x35, 0xd7,
	0xb5, 0x90, 0x95, 0xeb, 0xf3, 0x30, 0xb9, 0x96, 0xf8, 0x3d, 0xf4, 0x95, 0x6b, 0xed, 0x12, 0xef,
	0x1c, 0xbc, 0xb8, 0x75, 0xa2, 0x5f, 0x4e, 0xd2, 0xd9, 0xf1, 0x2e, 0xa0, 0x2b, 0x97, 0x72, 0x37,
	0x41, 0x57, 0xad, 0xb2, 0xc1, 0x7d, 0xaf, 0x6c, 0xf8, 0x13, 0xc1, 0x23, 0xcf, 0x68, 0xb1, 0xa6,
	0x86, 0xff, 0x3f, 0xc2, 0x27, 0x00, 0xe5, 0x92, 0xa6, 0x2b, 0x6e, 0xbb, 0x18, 0xe1, 0xed, 0x31,
	0xae, 0x4b, 0x25, 0x43, 0x75, 0xd3, 0xef, 0x39, 0x0c, 0x57, 0xdc, 0xa6, 0x6b, 0x5a, 0x34, 0xbc,
	0x0b, 0x35, 0x58, 0x71, 0x7b, 0xd1, 0xea, 0x70, 0xd6, 0xad, 0x63, 0xce, 0xed, 0x89, 0xd4, 0x25,
	0x35, 0x77, 0x2e, 0xbd, 0x7f, 0x0a, 0x4f, 0x37, 0x6f, 0x1f, 0x3f, 0x81, 0xf1, 0xd7, 0x78, 0x1e,
	0x9f, 0x7d, 0x8b, 0x53, 0x12, 0x1f, 0x9d, 0x1d, 0xcf, 0xe2, 0xd3, 0xf1, 0x3d, 0x3c, 0x02, 0x98,
	0x11, 0x42, 0xd2, 0xc5, 0xdb, 0xe9, 0xbb, 0xe9, 0x18, 0xe1, 0x6d, 0xd8, 0x3a, 0x26, 0xc9, 0xb8,
	0x77, 0x78, 0x01, 0x7b, 0x4c, 0x96, 0x9b, 0xc6, 0xba, 0x93, 0x5c, 0xa0, 0xef, 0x6f, 0x32, 0x61,
	0xf2, 0x66, 0x19, 0x31, 0x59, 0x4e, 0xbc, 0x6d, 0xd2, 0x7e, 0x9f, 0xfc, 0x75, 0xe6, 0x69, 0x26,
	0x53, 0x27, 0x7f, 0xf5, 0xfa, 0xe7, 0xb3, 0x78, 0xbe, 0x38, 0x5c, 0xf6, 0x9d, 0x9e, 0xfe, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x7d, 0x47, 0x15, 0xf0, 0x0e, 0x03, 0x00, 0x00,
}
