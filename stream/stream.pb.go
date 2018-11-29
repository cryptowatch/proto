// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stream/stream.proto

package ProtobufStream

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import markets "code.cryptowat.ch/proto/markets"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AuthenticationResult_Status int32

const (
	AuthenticationResult_UNKNOWN       AuthenticationResult_Status = 0
	AuthenticationResult_AUTHENTICATED AuthenticationResult_Status = 1
	AuthenticationResult_BAD_NONCE     AuthenticationResult_Status = 2
	AuthenticationResult_BAD_TOKEN     AuthenticationResult_Status = 3
	AuthenticationResult_TOKEN_EXPIRED AuthenticationResult_Status = 4
)

var AuthenticationResult_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "AUTHENTICATED",
	2: "BAD_NONCE",
	3: "BAD_TOKEN",
	4: "TOKEN_EXPIRED",
}
var AuthenticationResult_Status_value = map[string]int32{
	"UNKNOWN":       0,
	"AUTHENTICATED": 1,
	"BAD_NONCE":     2,
	"BAD_TOKEN":     3,
	"TOKEN_EXPIRED": 4,
}

func (x AuthenticationResult_Status) String() string {
	return proto.EnumName(AuthenticationResult_Status_name, int32(x))
}
func (AuthenticationResult_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_stream_aa04848dfef74b52, []int{1, 0}
}

type StreamMessage struct {
	// Types that are valid to be assigned to Body:
	//	*StreamMessage_AuthenticationResult
	//	*StreamMessage_MarketUpdate
	//	*StreamMessage_PairUpdate
	//	*StreamMessage_AssetUpdate
	Body                 isStreamMessage_Body `protobuf_oneof:"body"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StreamMessage) Reset()         { *m = StreamMessage{} }
func (m *StreamMessage) String() string { return proto.CompactTextString(m) }
func (*StreamMessage) ProtoMessage()    {}
func (*StreamMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_stream_aa04848dfef74b52, []int{0}
}
func (m *StreamMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamMessage.Unmarshal(m, b)
}
func (m *StreamMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamMessage.Marshal(b, m, deterministic)
}
func (dst *StreamMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMessage.Merge(dst, src)
}
func (m *StreamMessage) XXX_Size() int {
	return xxx_messageInfo_StreamMessage.Size(m)
}
func (m *StreamMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMessage proto.InternalMessageInfo

type isStreamMessage_Body interface {
	isStreamMessage_Body()
}

type StreamMessage_AuthenticationResult struct {
	AuthenticationResult *AuthenticationResult `protobuf:"bytes,1,opt,name=authenticationResult,oneof"`
}
type StreamMessage_MarketUpdate struct {
	MarketUpdate *markets.MarketUpdateMessage `protobuf:"bytes,2,opt,name=marketUpdate,oneof"`
}
type StreamMessage_PairUpdate struct {
	PairUpdate *markets.PairUpdateMessage `protobuf:"bytes,3,opt,name=pairUpdate,oneof"`
}
type StreamMessage_AssetUpdate struct {
	AssetUpdate *markets.AssetUpdateMessage `protobuf:"bytes,4,opt,name=assetUpdate,oneof"`
}

func (*StreamMessage_AuthenticationResult) isStreamMessage_Body() {}
func (*StreamMessage_MarketUpdate) isStreamMessage_Body()         {}
func (*StreamMessage_PairUpdate) isStreamMessage_Body()           {}
func (*StreamMessage_AssetUpdate) isStreamMessage_Body()          {}

func (m *StreamMessage) GetBody() isStreamMessage_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *StreamMessage) GetAuthenticationResult() *AuthenticationResult {
	if x, ok := m.GetBody().(*StreamMessage_AuthenticationResult); ok {
		return x.AuthenticationResult
	}
	return nil
}

func (m *StreamMessage) GetMarketUpdate() *markets.MarketUpdateMessage {
	if x, ok := m.GetBody().(*StreamMessage_MarketUpdate); ok {
		return x.MarketUpdate
	}
	return nil
}

func (m *StreamMessage) GetPairUpdate() *markets.PairUpdateMessage {
	if x, ok := m.GetBody().(*StreamMessage_PairUpdate); ok {
		return x.PairUpdate
	}
	return nil
}

func (m *StreamMessage) GetAssetUpdate() *markets.AssetUpdateMessage {
	if x, ok := m.GetBody().(*StreamMessage_AssetUpdate); ok {
		return x.AssetUpdate
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StreamMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StreamMessage_OneofMarshaler, _StreamMessage_OneofUnmarshaler, _StreamMessage_OneofSizer, []interface{}{
		(*StreamMessage_AuthenticationResult)(nil),
		(*StreamMessage_MarketUpdate)(nil),
		(*StreamMessage_PairUpdate)(nil),
		(*StreamMessage_AssetUpdate)(nil),
	}
}

func _StreamMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StreamMessage)
	// body
	switch x := m.Body.(type) {
	case *StreamMessage_AuthenticationResult:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AuthenticationResult); err != nil {
			return err
		}
	case *StreamMessage_MarketUpdate:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MarketUpdate); err != nil {
			return err
		}
	case *StreamMessage_PairUpdate:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PairUpdate); err != nil {
			return err
		}
	case *StreamMessage_AssetUpdate:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AssetUpdate); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StreamMessage.Body has unexpected type %T", x)
	}
	return nil
}

func _StreamMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StreamMessage)
	switch tag {
	case 1: // body.authenticationResult
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AuthenticationResult)
		err := b.DecodeMessage(msg)
		m.Body = &StreamMessage_AuthenticationResult{msg}
		return true, err
	case 2: // body.marketUpdate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(markets.MarketUpdateMessage)
		err := b.DecodeMessage(msg)
		m.Body = &StreamMessage_MarketUpdate{msg}
		return true, err
	case 3: // body.pairUpdate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(markets.PairUpdateMessage)
		err := b.DecodeMessage(msg)
		m.Body = &StreamMessage_PairUpdate{msg}
		return true, err
	case 4: // body.assetUpdate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(markets.AssetUpdateMessage)
		err := b.DecodeMessage(msg)
		m.Body = &StreamMessage_AssetUpdate{msg}
		return true, err
	default:
		return false, nil
	}
}

func _StreamMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StreamMessage)
	// body
	switch x := m.Body.(type) {
	case *StreamMessage_AuthenticationResult:
		s := proto.Size(x.AuthenticationResult)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamMessage_MarketUpdate:
		s := proto.Size(x.MarketUpdate)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamMessage_PairUpdate:
		s := proto.Size(x.PairUpdate)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamMessage_AssetUpdate:
		s := proto.Size(x.AssetUpdate)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type AuthenticationResult struct {
	Status               AuthenticationResult_Status `protobuf:"varint,1,opt,name=status,enum=ProtobufStream.AuthenticationResult_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *AuthenticationResult) Reset()         { *m = AuthenticationResult{} }
func (m *AuthenticationResult) String() string { return proto.CompactTextString(m) }
func (*AuthenticationResult) ProtoMessage()    {}
func (*AuthenticationResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_stream_aa04848dfef74b52, []int{1}
}
func (m *AuthenticationResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationResult.Unmarshal(m, b)
}
func (m *AuthenticationResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationResult.Marshal(b, m, deterministic)
}
func (dst *AuthenticationResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationResult.Merge(dst, src)
}
func (m *AuthenticationResult) XXX_Size() int {
	return xxx_messageInfo_AuthenticationResult.Size(m)
}
func (m *AuthenticationResult) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationResult.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationResult proto.InternalMessageInfo

func (m *AuthenticationResult) GetStatus() AuthenticationResult_Status {
	if m != nil {
		return m.Status
	}
	return AuthenticationResult_UNKNOWN
}

func init() {
	proto.RegisterType((*StreamMessage)(nil), "ProtobufStream.StreamMessage")
	proto.RegisterType((*AuthenticationResult)(nil), "ProtobufStream.AuthenticationResult")
	proto.RegisterEnum("ProtobufStream.AuthenticationResult_Status", AuthenticationResult_Status_name, AuthenticationResult_Status_value)
}

func init() { proto.RegisterFile("stream/stream.proto", fileDescriptor_stream_aa04848dfef74b52) }

var fileDescriptor_stream_aa04848dfef74b52 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4e, 0xf2, 0x40,
	0x14, 0xc5, 0x29, 0x90, 0x7e, 0xf9, 0x2e, 0x42, 0xea, 0xc0, 0x82, 0xb0, 0x32, 0xd5, 0x85, 0x89,
	0x49, 0x49, 0xf0, 0x09, 0x0a, 0x9d, 0x08, 0x12, 0x06, 0x52, 0x4a, 0xfc, 0xb3, 0x21, 0x83, 0x8c,
	0xda, 0x28, 0xb4, 0xe9, 0x4c, 0x17, 0xbe, 0x98, 0x8f, 0xe1, 0x33, 0x99, 0xce, 0xb4, 0xb6, 0x48,
	0x17, 0xae, 0x6e, 0x7b, 0x72, 0xce, 0x6f, 0xee, 0x9c, 0x0c, 0xb4, 0xb9, 0x88, 0x18, 0xdd, 0xf5,
	0xd5, 0xb0, 0xc2, 0x28, 0x10, 0x01, 0x6a, 0x2d, 0x92, 0xb1, 0x89, 0x9f, 0x97, 0x52, 0xed, 0x75,
	0x76, 0x34, 0x7a, 0x63, 0x82, 0xf7, 0xd5, 0x54, 0xae, 0x1e, 0xca, 0xd4, 0x90, 0xfa, 0x51, 0xaa,
	0xb5, 0x33, 0x8d, 0x72, 0x9e, 0x19, 0xcd, 0xaf, 0x2a, 0x34, 0x15, 0x69, 0xc6, 0x38, 0xa7, 0x2f,
	0x0c, 0x3d, 0x42, 0x87, 0xc6, 0xe2, 0x95, 0xed, 0x85, 0xff, 0x44, 0x85, 0x1f, 0xec, 0x5d, 0xc6,
	0xe3, 0x77, 0xd1, 0xd5, 0xce, 0xb4, 0xcb, 0xc6, 0xe0, 0xc2, 0x3a, 0x3c, 0xdf, 0xb2, 0x4b, 0xbc,
	0xe3, 0x8a, 0x5b, 0xca, 0x40, 0xb7, 0x70, 0xa2, 0x96, 0x58, 0x85, 0x5b, 0x2a, 0x58, 0xb7, 0xfa,
	0x8b, 0x39, 0x53, 0x1b, 0x5a, 0xb3, 0x82, 0x29, 0xdd, 0x6b, 0x5c, 0x71, 0x0f, 0xb2, 0xc8, 0x01,
	0x48, 0x2e, 0x97, 0x92, 0x6a, 0x92, 0x64, 0x1e, 0x91, 0x16, 0x3f, 0x96, 0x9c, 0x53, 0xc8, 0xa1,
	0x1b, 0x68, 0xc8, 0x3a, 0x52, 0x4c, 0x5d, 0x62, 0xce, 0x8f, 0x30, 0x76, 0xee, 0xc9, 0x39, 0xc5,
	0xe4, 0x50, 0x87, 0xfa, 0x26, 0xd8, 0x7e, 0x98, 0x9f, 0x1a, 0x74, 0xca, 0x3a, 0x41, 0x23, 0xd0,
	0xb9, 0xa0, 0x22, 0xe6, 0xb2, 0xc9, 0xd6, 0xe0, 0xea, 0x2f, 0x4d, 0x5a, 0x4b, 0x19, 0x71, 0xd3,
	0xa8, 0xf9, 0x00, 0xba, 0x52, 0x50, 0x03, 0xfe, 0xad, 0xc8, 0x94, 0xcc, 0xef, 0x88, 0x51, 0x41,
	0xa7, 0xd0, 0xb4, 0x57, 0xde, 0x18, 0x13, 0x6f, 0x32, 0xb2, 0x3d, 0xec, 0x18, 0x1a, 0x6a, 0xc2,
	0xff, 0xa1, 0xed, 0xac, 0xc9, 0x9c, 0x8c, 0xb0, 0x51, 0xcd, 0x7e, 0xbd, 0xf9, 0x14, 0x13, 0xa3,
	0x96, 0x04, 0xe4, 0xe7, 0x1a, 0xdf, 0x2f, 0x26, 0x2e, 0x76, 0x8c, 0xfa, 0x46, 0x97, 0x0f, 0xe2,
	0xfa, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x4a, 0xff, 0x23, 0x76, 0x02, 0x00, 0x00,
}
