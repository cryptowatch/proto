// Code generated by protoc-gen-go. DO NOT EDIT.
// source: broker/private.proto

package ProtobufBroker

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FundingType int32

const (
	FundingType_Spot   FundingType = 0
	FundingType_Margin FundingType = 1
)

var FundingType_name = map[int32]string{
	0: "Spot",
	1: "Margin",
}

var FundingType_value = map[string]int32{
	"Spot":   0,
	"Margin": 1,
}

func (x FundingType) String() string {
	return proto.EnumName(FundingType_name, int32(x))
}

func (FundingType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2643a2aaad9ebcde, []int{0}
}

type PrivateOrder_Type int32

const (
	PrivateOrder_Market                  PrivateOrder_Type = 0
	PrivateOrder_Limit                   PrivateOrder_Type = 1
	PrivateOrder_StopLoss                PrivateOrder_Type = 2
	PrivateOrder_StopLossLimit           PrivateOrder_Type = 3
	PrivateOrder_TakeProfit              PrivateOrder_Type = 4
	PrivateOrder_TakeProfitLimit         PrivateOrder_Type = 5
	PrivateOrder_StopLossTakeProfit      PrivateOrder_Type = 6
	PrivateOrder_StopLossTakeProfitLimit PrivateOrder_Type = 7
	PrivateOrder_TrailingStopLoss        PrivateOrder_Type = 8
	PrivateOrder_TrailingStopLossLimit   PrivateOrder_Type = 9
	PrivateOrder_StopLossAndLimit        PrivateOrder_Type = 10
	PrivateOrder_FillOrKill              PrivateOrder_Type = 11
	PrivateOrder_SettlePosition          PrivateOrder_Type = 12
)

var PrivateOrder_Type_name = map[int32]string{
	0:  "Market",
	1:  "Limit",
	2:  "StopLoss",
	3:  "StopLossLimit",
	4:  "TakeProfit",
	5:  "TakeProfitLimit",
	6:  "StopLossTakeProfit",
	7:  "StopLossTakeProfitLimit",
	8:  "TrailingStopLoss",
	9:  "TrailingStopLossLimit",
	10: "StopLossAndLimit",
	11: "FillOrKill",
	12: "SettlePosition",
}

var PrivateOrder_Type_value = map[string]int32{
	"Market":                  0,
	"Limit":                   1,
	"StopLoss":                2,
	"StopLossLimit":           3,
	"TakeProfit":              4,
	"TakeProfitLimit":         5,
	"StopLossTakeProfit":      6,
	"StopLossTakeProfitLimit": 7,
	"TrailingStopLoss":        8,
	"TrailingStopLossLimit":   9,
	"StopLossAndLimit":        10,
	"FillOrKill":              11,
	"SettlePosition":          12,
}

func (x PrivateOrder_Type) String() string {
	return proto.EnumName(PrivateOrder_Type_name, int32(x))
}

func (PrivateOrder_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2643a2aaad9ebcde, []int{0, 0}
}

type PrivateOrder_PriceParamType int32

const (
	PrivateOrder_AbsoluteValue         PrivateOrder_PriceParamType = 0
	PrivateOrder_OffsetValue           PrivateOrder_PriceParamType = 1
	PrivateOrder_PrecentageOffsetValue PrivateOrder_PriceParamType = 2
)

var PrivateOrder_PriceParamType_name = map[int32]string{
	0: "AbsoluteValue",
	1: "OffsetValue",
	2: "PrecentageOffsetValue",
}

var PrivateOrder_PriceParamType_value = map[string]int32{
	"AbsoluteValue":         0,
	"OffsetValue":           1,
	"PrecentageOffsetValue": 2,
}

func (x PrivateOrder_PriceParamType) String() string {
	return proto.EnumName(PrivateOrder_PriceParamType_name, int32(x))
}

func (PrivateOrder_PriceParamType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2643a2aaad9ebcde, []int{0, 1}
}

type PrivateOrder struct {
	Id                 string                     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Time               int64                      `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Price              float32                    `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Amount             float32                    `protobuf:"fixed32,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Side               int32                      `protobuf:"varint,5,opt,name=side,proto3" json:"side,omitempty"`
	Type               PrivateOrder_Type          `protobuf:"varint,6,opt,name=type,proto3,enum=ProtobufBroker.PrivateOrder_Type" json:"type,omitempty"`
	FundingType        FundingType                `protobuf:"varint,7,opt,name=fundingType,proto3,enum=ProtobufBroker.FundingType" json:"fundingType,omitempty"`
	PriceParams        []*PrivateOrder_PriceParam `protobuf:"bytes,8,rep,name=priceParams,proto3" json:"priceParams,omitempty"`
	AmountParam        float32                    `protobuf:"fixed32,9,opt,name=amountParam,proto3" json:"amountParam,omitempty"`
	AmountParamString  string                     `protobuf:"bytes,16,opt,name=amountParamString,proto3" json:"amountParamString,omitempty"`
	AmountFilledString string                     `protobuf:"bytes,17,opt,name=amountFilledString,proto3" json:"amountFilledString,omitempty"`
	Leverage           string                     `protobuf:"bytes,10,opt,name=leverage,proto3" json:"leverage,omitempty"`
	// For trailing orders
	// DEPRECATED; use strings instead
	CurrentStop             float32                    `protobuf:"fixed32,11,opt,name=currentStop,proto3" json:"currentStop,omitempty"`
	InitialStop             float32                    `protobuf:"fixed32,12,opt,name=initialStop,proto3" json:"initialStop,omitempty"`
	CurrentStopString       string                     `protobuf:"bytes,18,opt,name=currentStopString,proto3" json:"currentStopString,omitempty"`
	InitialStopString       string                     `protobuf:"bytes,19,opt,name=initialStopString,proto3" json:"initialStopString,omitempty"`
	StartTime               int64                      `protobuf:"varint,13,opt,name=startTime,proto3" json:"startTime,omitempty"`
	ExpireTime              int64                      `protobuf:"varint,14,opt,name=expireTime,proto3" json:"expireTime,omitempty"`
	Rate                    float32                    `protobuf:"fixed32,15,opt,name=rate,proto3" json:"rate,omitempty"`
	HasClosingOrder         bool                       `protobuf:"varint,20,opt,name=hasClosingOrder,proto3" json:"hasClosingOrder,omitempty"`
	ClosingOrderType        PrivateOrder_Type          `protobuf:"varint,21,opt,name=closingOrderType,proto3,enum=ProtobufBroker.PrivateOrder_Type" json:"closingOrderType,omitempty"`
	ClosingOrderPriceParams []*PrivateOrder_PriceParam `protobuf:"bytes,22,rep,name=closingOrderPriceParams,proto3" json:"closingOrderPriceParams,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                   `json:"-"`
	XXX_unrecognized        []byte                     `json:"-"`
	XXX_sizecache           int32                      `json:"-"`
}

func (m *PrivateOrder) Reset()         { *m = PrivateOrder{} }
func (m *PrivateOrder) String() string { return proto.CompactTextString(m) }
func (*PrivateOrder) ProtoMessage()    {}
func (*PrivateOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_2643a2aaad9ebcde, []int{0}
}

func (m *PrivateOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrivateOrder.Unmarshal(m, b)
}
func (m *PrivateOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrivateOrder.Marshal(b, m, deterministic)
}
func (m *PrivateOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivateOrder.Merge(m, src)
}
func (m *PrivateOrder) XXX_Size() int {
	return xxx_messageInfo_PrivateOrder.Size(m)
}
func (m *PrivateOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivateOrder.DiscardUnknown(m)
}

var xxx_messageInfo_PrivateOrder proto.InternalMessageInfo

func (m *PrivateOrder) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PrivateOrder) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *PrivateOrder) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *PrivateOrder) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *PrivateOrder) GetSide() int32 {
	if m != nil {
		return m.Side
	}
	return 0
}

func (m *PrivateOrder) GetType() PrivateOrder_Type {
	if m != nil {
		return m.Type
	}
	return PrivateOrder_Market
}

func (m *PrivateOrder) GetFundingType() FundingType {
	if m != nil {
		return m.FundingType
	}
	return FundingType_Spot
}

func (m *PrivateOrder) GetPriceParams() []*PrivateOrder_PriceParam {
	if m != nil {
		return m.PriceParams
	}
	return nil
}

func (m *PrivateOrder) GetAmountParam() float32 {
	if m != nil {
		return m.AmountParam
	}
	return 0
}

func (m *PrivateOrder) GetAmountParamString() string {
	if m != nil {
		return m.AmountParamString
	}
	return ""
}

func (m *PrivateOrder) GetAmountFilledString() string {
	if m != nil {
		return m.AmountFilledString
	}
	return ""
}

func (m *PrivateOrder) GetLeverage() string {
	if m != nil {
		return m.Leverage
	}
	return ""
}

func (m *PrivateOrder) GetCurrentStop() float32 {
	if m != nil {
		return m.CurrentStop
	}
	return 0
}

func (m *PrivateOrder) GetInitialStop() float32 {
	if m != nil {
		return m.InitialStop
	}
	return 0
}

func (m *PrivateOrder) GetCurrentStopString() string {
	if m != nil {
		return m.CurrentStopString
	}
	return ""
}

func (m *PrivateOrder) GetInitialStopString() string {
	if m != nil {
		return m.InitialStopString
	}
	return ""
}

func (m *PrivateOrder) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *PrivateOrder) GetExpireTime() int64 {
	if m != nil {
		return m.ExpireTime
	}
	return 0
}

func (m *PrivateOrder) GetRate() float32 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func (m *PrivateOrder) GetHasClosingOrder() bool {
	if m != nil {
		return m.HasClosingOrder
	}
	return false
}

func (m *PrivateOrder) GetClosingOrderType() PrivateOrder_Type {
	if m != nil {
		return m.ClosingOrderType
	}
	return PrivateOrder_Market
}

func (m *PrivateOrder) GetClosingOrderPriceParams() []*PrivateOrder_PriceParam {
	if m != nil {
		return m.ClosingOrderPriceParams
	}
	return nil
}

type PrivateOrder_PriceParam struct {
	Value                float32                     `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	ValueString          string                      `protobuf:"bytes,3,opt,name=valueString,proto3" json:"valueString,omitempty"`
	Type                 PrivateOrder_PriceParamType `protobuf:"varint,2,opt,name=type,proto3,enum=ProtobufBroker.PrivateOrder_PriceParamType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *PrivateOrder_PriceParam) Reset()         { *m = PrivateOrder_PriceParam{} }
func (m *PrivateOrder_PriceParam) String() string { return proto.CompactTextString(m) }
func (*PrivateOrder_PriceParam) ProtoMessage()    {}
func (*PrivateOrder_PriceParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_2643a2aaad9ebcde, []int{0, 0}
}

func (m *PrivateOrder_PriceParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrivateOrder_PriceParam.Unmarshal(m, b)
}
func (m *PrivateOrder_PriceParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrivateOrder_PriceParam.Marshal(b, m, deterministic)
}
func (m *PrivateOrder_PriceParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivateOrder_PriceParam.Merge(m, src)
}
func (m *PrivateOrder_PriceParam) XXX_Size() int {
	return xxx_messageInfo_PrivateOrder_PriceParam.Size(m)
}
func (m *PrivateOrder_PriceParam) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivateOrder_PriceParam.DiscardUnknown(m)
}

var xxx_messageInfo_PrivateOrder_PriceParam proto.InternalMessageInfo

func (m *PrivateOrder_PriceParam) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *PrivateOrder_PriceParam) GetValueString() string {
	if m != nil {
		return m.ValueString
	}
	return ""
}

func (m *PrivateOrder_PriceParam) GetType() PrivateOrder_PriceParamType {
	if m != nil {
		return m.Type
	}
	return PrivateOrder_AbsoluteValue
}

type PrivateTrade struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ExternalId           string   `protobuf:"bytes,7,opt,name=externalId,proto3" json:"externalId,omitempty"`
	OrderId              string   `protobuf:"bytes,8,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Time                 int64    `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	TimeMillis           int64    `protobuf:"varint,6,opt,name=timeMillis,proto3" json:"timeMillis,omitempty"`
	Price                float32  `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	PriceString          string   `protobuf:"bytes,9,opt,name=priceString,proto3" json:"priceString,omitempty"`
	Amount               float32  `protobuf:"fixed32,4,opt,name=amount,proto3" json:"amount,omitempty"`
	AmountString         string   `protobuf:"bytes,10,opt,name=amountString,proto3" json:"amountString,omitempty"`
	Side                 int32    `protobuf:"varint,5,opt,name=side,proto3" json:"side,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrivateTrade) Reset()         { *m = PrivateTrade{} }
func (m *PrivateTrade) String() string { return proto.CompactTextString(m) }
func (*PrivateTrade) ProtoMessage()    {}
func (*PrivateTrade) Descriptor() ([]byte, []int) {
	return fileDescriptor_2643a2aaad9ebcde, []int{1}
}

func (m *PrivateTrade) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrivateTrade.Unmarshal(m, b)
}
func (m *PrivateTrade) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrivateTrade.Marshal(b, m, deterministic)
}
func (m *PrivateTrade) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivateTrade.Merge(m, src)
}
func (m *PrivateTrade) XXX_Size() int {
	return xxx_messageInfo_PrivateTrade.Size(m)
}
func (m *PrivateTrade) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivateTrade.DiscardUnknown(m)
}

var xxx_messageInfo_PrivateTrade proto.InternalMessageInfo

func (m *PrivateTrade) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PrivateTrade) GetExternalId() string {
	if m != nil {
		return m.ExternalId
	}
	return ""
}

func (m *PrivateTrade) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *PrivateTrade) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *PrivateTrade) GetTimeMillis() int64 {
	if m != nil {
		return m.TimeMillis
	}
	return 0
}

func (m *PrivateTrade) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *PrivateTrade) GetPriceString() string {
	if m != nil {
		return m.PriceString
	}
	return ""
}

func (m *PrivateTrade) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *PrivateTrade) GetAmountString() string {
	if m != nil {
		return m.AmountString
	}
	return ""
}

func (m *PrivateTrade) GetSide() int32 {
	if m != nil {
		return m.Side
	}
	return 0
}

type PrivatePosition struct {
	Id             string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Time           int64   `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Side           int32   `protobuf:"varint,3,opt,name=side,proto3" json:"side,omitempty"`
	AvgPrice       float32 `protobuf:"fixed32,4,opt,name=avgPrice,proto3" json:"avgPrice,omitempty"`
	AvgPriceString string  `protobuf:"bytes,9,opt,name=avgPriceString,proto3" json:"avgPriceString,omitempty"`
	// DEPRECATED; use strings instead
	AmountOpen         float32  `protobuf:"fixed32,5,opt,name=amountOpen,proto3" json:"amountOpen,omitempty"`
	AmountClosed       float32  `protobuf:"fixed32,6,opt,name=amountClosed,proto3" json:"amountClosed,omitempty"`
	AmountOpenString   string   `protobuf:"bytes,10,opt,name=amountOpenString,proto3" json:"amountOpenString,omitempty"`
	AmountClosedString string   `protobuf:"bytes,11,opt,name=amountClosedString,proto3" json:"amountClosedString,omitempty"`
	OrderIds           []string `protobuf:"bytes,7,rep,name=orderIds,proto3" json:"orderIds,omitempty"`
	// NOTE:
	// Trade ids are sent as strings here, while historically they
	// have been represented as integers. We need to transition to using
	// strings everywhere because of the rectangle/square rule; using
	// ints is simply not compatible with some exchanges which use UUIDs (like Kraken lol)
	TradeIds             []string `protobuf:"bytes,8,rep,name=tradeIds,proto3" json:"tradeIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrivatePosition) Reset()         { *m = PrivatePosition{} }
func (m *PrivatePosition) String() string { return proto.CompactTextString(m) }
func (*PrivatePosition) ProtoMessage()    {}
func (*PrivatePosition) Descriptor() ([]byte, []int) {
	return fileDescriptor_2643a2aaad9ebcde, []int{2}
}

func (m *PrivatePosition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrivatePosition.Unmarshal(m, b)
}
func (m *PrivatePosition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrivatePosition.Marshal(b, m, deterministic)
}
func (m *PrivatePosition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivatePosition.Merge(m, src)
}
func (m *PrivatePosition) XXX_Size() int {
	return xxx_messageInfo_PrivatePosition.Size(m)
}
func (m *PrivatePosition) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivatePosition.DiscardUnknown(m)
}

var xxx_messageInfo_PrivatePosition proto.InternalMessageInfo

func (m *PrivatePosition) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PrivatePosition) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *PrivatePosition) GetSide() int32 {
	if m != nil {
		return m.Side
	}
	return 0
}

func (m *PrivatePosition) GetAvgPrice() float32 {
	if m != nil {
		return m.AvgPrice
	}
	return 0
}

func (m *PrivatePosition) GetAvgPriceString() string {
	if m != nil {
		return m.AvgPriceString
	}
	return ""
}

func (m *PrivatePosition) GetAmountOpen() float32 {
	if m != nil {
		return m.AmountOpen
	}
	return 0
}

func (m *PrivatePosition) GetAmountClosed() float32 {
	if m != nil {
		return m.AmountClosed
	}
	return 0
}

func (m *PrivatePosition) GetAmountOpenString() string {
	if m != nil {
		return m.AmountOpenString
	}
	return ""
}

func (m *PrivatePosition) GetAmountClosedString() string {
	if m != nil {
		return m.AmountClosedString
	}
	return ""
}

func (m *PrivatePosition) GetOrderIds() []string {
	if m != nil {
		return m.OrderIds
	}
	return nil
}

func (m *PrivatePosition) GetTradeIds() []string {
	if m != nil {
		return m.TradeIds
	}
	return nil
}

type Balance struct {
	Currency             string   `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount               float32  `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
	AmountString         string   `protobuf:"bytes,3,opt,name=amountString,proto3" json:"amountString,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Balance) Reset()         { *m = Balance{} }
func (m *Balance) String() string { return proto.CompactTextString(m) }
func (*Balance) ProtoMessage()    {}
func (*Balance) Descriptor() ([]byte, []int) {
	return fileDescriptor_2643a2aaad9ebcde, []int{3}
}

func (m *Balance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Balance.Unmarshal(m, b)
}
func (m *Balance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Balance.Marshal(b, m, deterministic)
}
func (m *Balance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Balance.Merge(m, src)
}
func (m *Balance) XXX_Size() int {
	return xxx_messageInfo_Balance.Size(m)
}
func (m *Balance) XXX_DiscardUnknown() {
	xxx_messageInfo_Balance.DiscardUnknown(m)
}

var xxx_messageInfo_Balance proto.InternalMessageInfo

func (m *Balance) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Balance) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Balance) GetAmountString() string {
	if m != nil {
		return m.AmountString
	}
	return ""
}

type Balances struct {
	FundingType          FundingType `protobuf:"varint,1,opt,name=fundingType,proto3,enum=ProtobufBroker.FundingType" json:"fundingType,omitempty"`
	Balances             []*Balance  `protobuf:"bytes,2,rep,name=balances,proto3" json:"balances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Balances) Reset()         { *m = Balances{} }
func (m *Balances) String() string { return proto.CompactTextString(m) }
func (*Balances) ProtoMessage()    {}
func (*Balances) Descriptor() ([]byte, []int) {
	return fileDescriptor_2643a2aaad9ebcde, []int{4}
}

func (m *Balances) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Balances.Unmarshal(m, b)
}
func (m *Balances) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Balances.Marshal(b, m, deterministic)
}
func (m *Balances) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Balances.Merge(m, src)
}
func (m *Balances) XXX_Size() int {
	return xxx_messageInfo_Balances.Size(m)
}
func (m *Balances) XXX_DiscardUnknown() {
	xxx_messageInfo_Balances.DiscardUnknown(m)
}

var xxx_messageInfo_Balances proto.InternalMessageInfo

func (m *Balances) GetFundingType() FundingType {
	if m != nil {
		return m.FundingType
	}
	return FundingType_Spot
}

func (m *Balances) GetBalances() []*Balance {
	if m != nil {
		return m.Balances
	}
	return nil
}

func init() {
	proto.RegisterEnum("ProtobufBroker.FundingType", FundingType_name, FundingType_value)
	proto.RegisterEnum("ProtobufBroker.PrivateOrder_Type", PrivateOrder_Type_name, PrivateOrder_Type_value)
	proto.RegisterEnum("ProtobufBroker.PrivateOrder_PriceParamType", PrivateOrder_PriceParamType_name, PrivateOrder_PriceParamType_value)
	proto.RegisterType((*PrivateOrder)(nil), "ProtobufBroker.PrivateOrder")
	proto.RegisterType((*PrivateOrder_PriceParam)(nil), "ProtobufBroker.PrivateOrder.PriceParam")
	proto.RegisterType((*PrivateTrade)(nil), "ProtobufBroker.PrivateTrade")
	proto.RegisterType((*PrivatePosition)(nil), "ProtobufBroker.PrivatePosition")
	proto.RegisterType((*Balance)(nil), "ProtobufBroker.Balance")
	proto.RegisterType((*Balances)(nil), "ProtobufBroker.Balances")
}

func init() { proto.RegisterFile("broker/private.proto", fileDescriptor_2643a2aaad9ebcde) }

var fileDescriptor_2643a2aaad9ebcde = []byte{
	// 921 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xdf, 0x6f, 0x1a, 0x47,
	0x10, 0xce, 0xfd, 0xc0, 0x1c, 0x03, 0x81, 0xf3, 0xc4, 0xb1, 0xb7, 0x4e, 0x15, 0x5d, 0xa9, 0xd4,
	0xa2, 0xb4, 0xa2, 0x52, 0xa2, 0x3e, 0x56, 0x55, 0x52, 0x29, 0x92, 0xd3, 0x58, 0xa0, 0x33, 0xea,
	0xfb, 0xc2, 0x2d, 0x74, 0xe5, 0xf3, 0x1d, 0xda, 0x5b, 0xac, 0xfa, 0xb9, 0xea, 0x4b, 0x9f, 0xfb,
	0x2f, 0xf6, 0xdf, 0xa8, 0xaa, 0x9d, 0x3d, 0x60, 0x01, 0x3b, 0x72, 0x9e, 0xd8, 0xf9, 0xe6, 0x9b,
	0xbd, 0x6f, 0x67, 0x66, 0x67, 0x81, 0x93, 0xa9, 0x2a, 0xaf, 0x85, 0xfa, 0x61, 0xa9, 0xe4, 0x2d,
	0xd7, 0x62, 0xb8, 0x54, 0xa5, 0x2e, 0xb1, 0x3b, 0x36, 0x3f, 0xd3, 0xd5, 0xfc, 0x1d, 0x79, 0xfb,
	0xff, 0x01, 0x74, 0xc6, 0x96, 0x31, 0x52, 0x99, 0x50, 0xd8, 0x05, 0x5f, 0x66, 0xcc, 0x4b, 0xbc,
	0x41, 0x2b, 0xf5, 0x65, 0x86, 0x08, 0xa1, 0x96, 0x37, 0x82, 0xf9, 0x89, 0x37, 0x08, 0x52, 0x5a,
	0xe3, 0x09, 0x34, 0x96, 0x4a, 0xce, 0x04, 0x0b, 0x12, 0x6f, 0xe0, 0xa7, 0xd6, 0xc0, 0x53, 0x38,
	0xe2, 0x37, 0xe5, 0xaa, 0xd0, 0x2c, 0x24, 0xb8, 0xb6, 0xcc, 0x0e, 0x95, 0xcc, 0x04, 0x6b, 0x24,
	0xde, 0xa0, 0x91, 0xd2, 0x1a, 0x7f, 0x84, 0x50, 0xdf, 0x2d, 0x05, 0x3b, 0x4a, 0xbc, 0x41, 0xf7,
	0xf5, 0x57, 0xc3, 0x5d, 0x55, 0x43, 0x57, 0xd1, 0x70, 0x72, 0xb7, 0x14, 0x29, 0xd1, 0xf1, 0x27,
	0x68, 0xcf, 0x57, 0x45, 0x26, 0x8b, 0x85, 0x01, 0x59, 0x93, 0xa2, 0x5f, 0xec, 0x47, 0xbf, 0xdf,
	0x52, 0x52, 0x97, 0x8f, 0x17, 0xd0, 0x26, 0xa9, 0x63, 0xae, 0xf8, 0x4d, 0xc5, 0xa2, 0x24, 0x18,
	0xb4, 0x5f, 0x7f, 0xfb, 0xc9, 0x8f, 0x8f, 0x37, 0xfc, 0xd4, 0x8d, 0xc5, 0x04, 0xda, 0xf6, 0x78,
	0x64, 0xb3, 0x16, 0x9d, 0xd8, 0x85, 0xf0, 0x7b, 0x38, 0x76, 0xcc, 0x2b, 0xad, 0x64, 0xb1, 0x60,
	0x31, 0xe5, 0xf5, 0xd0, 0x81, 0x43, 0x40, 0x0b, 0xbe, 0x97, 0x79, 0x2e, 0xb2, 0x9a, 0x7e, 0x4c,
	0xf4, 0x7b, 0x3c, 0x78, 0x0e, 0x51, 0x2e, 0x6e, 0x85, 0xe2, 0x0b, 0xc1, 0x80, 0x58, 0x1b, 0xdb,
	0x68, 0x9b, 0xad, 0x94, 0x12, 0x85, 0xbe, 0xd2, 0xe5, 0x92, 0xb5, 0xad, 0x36, 0x07, 0x32, 0x0c,
	0x59, 0x48, 0x2d, 0x79, 0x4e, 0x8c, 0x8e, 0x65, 0x38, 0x90, 0x51, 0xef, 0x04, 0xd4, 0x72, 0xd0,
	0xaa, 0x3f, 0x70, 0x18, 0xb6, 0x13, 0x5c, 0xb3, 0x9f, 0x59, 0xf6, 0x81, 0x03, 0xbf, 0x84, 0x56,
	0xa5, 0xb9, 0xd2, 0x13, 0xd3, 0x57, 0x4f, 0xa9, 0xaf, 0xb6, 0x00, 0xbe, 0x04, 0x10, 0x7f, 0x2c,
	0xa5, 0x12, 0xe4, 0xee, 0x92, 0xdb, 0x41, 0x4c, 0x3b, 0x29, 0xae, 0x05, 0xeb, 0x91, 0x68, 0x5a,
	0xe3, 0x00, 0x7a, 0xbf, 0xf3, 0xea, 0x97, 0xbc, 0xac, 0x64, 0xb1, 0xa0, 0xc2, 0xb1, 0x93, 0xc4,
	0x1b, 0x44, 0xe9, 0x3e, 0x8c, 0x97, 0x10, 0xcf, 0x1c, 0x9b, 0xda, 0xe8, 0xf9, 0x63, 0x9b, 0xf0,
	0x20, 0x14, 0x39, 0x9c, 0xb9, 0xd8, 0xd8, 0xe9, 0xae, 0xd3, 0xcf, 0xeb, 0xae, 0x87, 0xf6, 0x39,
	0xff, 0xcb, 0x03, 0xd8, 0xda, 0xe6, 0xee, 0xdd, 0xf2, 0x7c, 0x25, 0xe8, 0x8a, 0xfa, 0xa9, 0x35,
	0x4c, 0x41, 0x69, 0x51, 0xa7, 0x3e, 0xa0, 0xd4, 0xbb, 0x10, 0xfe, 0x5c, 0xdf, 0x38, 0x9f, 0x0e,
	0xfb, 0xdd, 0x23, 0x65, 0x6d, 0xef, 0x5e, 0xff, 0x4f, 0x1f, 0x42, 0x3a, 0x33, 0xc0, 0xd1, 0x25,
	0x57, 0xd7, 0x42, 0xc7, 0x4f, 0xb0, 0x05, 0x8d, 0x8f, 0xf2, 0x46, 0xea, 0xd8, 0xc3, 0x0e, 0x44,
	0xa6, 0xc6, 0x1f, 0xcb, 0xaa, 0x8a, 0x7d, 0x3c, 0x86, 0xa7, 0x6b, 0xcb, 0x12, 0x02, 0xec, 0x02,
	0x4c, 0xf8, 0xb5, 0x18, 0xab, 0x72, 0x2e, 0x75, 0x1c, 0xe2, 0x33, 0xe8, 0x6d, 0x6d, 0x4b, 0x6a,
	0xe0, 0x29, 0xe0, 0x3a, 0xce, 0x21, 0x1f, 0xe1, 0x0b, 0x38, 0x3b, 0xc4, 0x6d, 0x50, 0x13, 0x4f,
	0x20, 0x9e, 0x28, 0x2e, 0x73, 0x59, 0x2c, 0x36, 0x12, 0x22, 0xfc, 0x02, 0x9e, 0xef, 0xa3, 0x36,
	0xa0, 0x65, 0x02, 0xd6, 0xd0, 0xdb, 0x22, 0xb3, 0x28, 0x18, 0x81, 0xe6, 0x8e, 0x8d, 0xd4, 0xaf,
	0x32, 0xcf, 0xe3, 0x36, 0x22, 0x74, 0xaf, 0x84, 0xd6, 0xb9, 0x18, 0x97, 0x95, 0xd4, 0xb2, 0x2c,
	0xe2, 0x4e, 0x7f, 0x04, 0xdd, 0xdd, 0xec, 0x98, 0x93, 0xbe, 0x9d, 0x56, 0x65, 0xbe, 0xd2, 0xe2,
	0x37, 0x93, 0xef, 0xf8, 0x09, 0xf6, 0xa0, 0x3d, 0x9a, 0xcf, 0x2b, 0xa1, 0x2d, 0xe0, 0x19, 0x29,
	0x63, 0x25, 0x66, 0xa2, 0xd0, 0x7c, 0x21, 0x5c, 0x97, 0xdf, 0xff, 0xc7, 0xdf, 0x0c, 0xe0, 0x89,
	0xe2, 0x99, 0x70, 0x06, 0x70, 0x40, 0x03, 0x98, 0xee, 0x83, 0x16, 0xaa, 0xe0, 0xf9, 0x45, 0x46,
	0x23, 0xaf, 0x95, 0x3a, 0x08, 0x32, 0x68, 0x96, 0xa6, 0x6a, 0x17, 0x19, 0x8b, 0xc8, 0xb9, 0x36,
	0xef, 0x1d, 0xdd, 0x2f, 0x01, 0xcc, 0xef, 0xa5, 0xcc, 0x73, 0x59, 0xd1, 0xf8, 0x0d, 0x52, 0x07,
	0x79, 0x60, 0xb4, 0x27, 0xf5, 0xe0, 0xac, 0xdb, 0xab, 0x65, 0xdb, 0xcb, 0x81, 0x1e, 0x1c, 0xfe,
	0x7d, 0xe8, 0xd8, 0x55, 0x1d, 0x6a, 0x67, 0xd5, 0x0e, 0x76, 0xdf, 0x03, 0xd1, 0xff, 0xd7, 0x87,
	0x5e, 0x9d, 0x96, 0x75, 0xf6, 0x1f, 0xf5, 0x34, 0xad, 0xf7, 0x0a, 0x9c, 0xc7, 0xe6, 0x1c, 0x22,
	0x7e, 0xbb, 0xa0, 0xb2, 0xd5, 0xea, 0x36, 0x36, 0x7e, 0x03, 0xdd, 0xf5, 0x7a, 0xe7, 0x70, 0x7b,
	0xa8, 0xc9, 0x9b, 0xd5, 0x3c, 0x5a, 0x8a, 0x82, 0x94, 0xfa, 0xa9, 0x83, 0x6c, 0xcf, 0x69, 0xa6,
	0x8d, 0xc8, 0x28, 0xb3, 0x7e, 0xba, 0x83, 0xe1, 0x2b, 0x88, 0xb7, 0x11, 0x3b, 0xf9, 0x38, 0xc0,
	0xb7, 0xef, 0x81, 0x8d, 0xad, 0xd9, 0x6d, 0xf7, 0x3d, 0x70, 0x3d, 0xe6, 0x8c, 0x75, 0xd9, 0x2b,
	0xd6, 0x4c, 0x02, 0xf3, 0x1e, 0xac, 0x6d, 0xe3, 0xd3, 0xa6, 0xb5, 0x8c, 0x2f, 0xb2, 0xbe, 0xb5,
	0xdd, 0xe7, 0xd0, 0x7c, 0xc7, 0x73, 0x5e, 0xcc, 0x28, 0x4d, 0x76, 0xb2, 0xcf, 0xee, 0xea, 0x24,
	0x6f, 0x6c, 0xa7, 0xbc, 0xfe, 0x27, 0xcb, 0x1b, 0x1c, 0x96, 0xb7, 0xff, 0xb7, 0x07, 0x51, 0xfd,
	0x8d, 0x6a, 0xff, 0x05, 0xf7, 0x3e, 0xf3, 0x05, 0x7f, 0x03, 0xd1, 0xb4, 0xde, 0x8a, 0xf9, 0x34,
	0x60, 0xcf, 0xf6, 0x63, 0xeb, 0x4f, 0xa5, 0x1b, 0xe2, 0x87, 0x30, 0x0a, 0xe2, 0xf0, 0x43, 0x18,
	0x85, 0x71, 0xe3, 0xd5, 0xd7, 0xd0, 0x76, 0x36, 0xc7, 0x08, 0xc2, 0xab, 0x65, 0x69, 0x26, 0x99,
	0x9d, 0x6a, 0x0b, 0x59, 0xc4, 0xde, 0xf4, 0x88, 0xfe, 0x2b, 0xbd, 0xf9, 0x3f, 0x00, 0x00, 0xff,
	0xff, 0x36, 0xad, 0x00, 0xb0, 0x43, 0x09, 0x00, 0x00,
}
