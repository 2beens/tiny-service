// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: proto/tiny_stock_exchange.proto

package internal

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListStocksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListStocksRequest) Reset() {
	*x = ListStocksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tiny_stock_exchange_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStocksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStocksRequest) ProtoMessage() {}

func (x *ListStocksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tiny_stock_exchange_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStocksRequest.ProtoReflect.Descriptor instead.
func (*ListStocksRequest) Descriptor() ([]byte, []int) {
	return file_proto_tiny_stock_exchange_proto_rawDescGZIP(), []int{0}
}

type ListStockValueDeltasRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListStockValueDeltasRequest) Reset() {
	*x = ListStockValueDeltasRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tiny_stock_exchange_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStockValueDeltasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStockValueDeltasRequest) ProtoMessage() {}

func (x *ListStockValueDeltasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tiny_stock_exchange_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStockValueDeltasRequest.ProtoReflect.Descriptor instead.
func (*ListStockValueDeltasRequest) Descriptor() ([]byte, []int) {
	return file_proto_tiny_stock_exchange_proto_rawDescGZIP(), []int{1}
}

type Stock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker string `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Stock) Reset() {
	*x = Stock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tiny_stock_exchange_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stock) ProtoMessage() {}

func (x *Stock) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tiny_stock_exchange_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stock.ProtoReflect.Descriptor instead.
func (*Stock) Descriptor() ([]byte, []int) {
	return file_proto_tiny_stock_exchange_proto_rawDescGZIP(), []int{2}
}

func (x *Stock) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *Stock) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// StockValueDelta represents a chagne of stock price at the specified timestamp
// e.g. 153153153123, +0.11
// e.g. 153153153532, -0.24
type StockValueDelta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int32 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Delta     int32 `protobuf:"varint,2,opt,name=delta,proto3" json:"delta,omitempty"`
}

func (x *StockValueDelta) Reset() {
	*x = StockValueDelta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tiny_stock_exchange_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockValueDelta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockValueDelta) ProtoMessage() {}

func (x *StockValueDelta) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tiny_stock_exchange_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockValueDelta.ProtoReflect.Descriptor instead.
func (*StockValueDelta) Descriptor() ([]byte, []int) {
	return file_proto_tiny_stock_exchange_proto_rawDescGZIP(), []int{3}
}

func (x *StockValueDelta) GetTimestamp() int32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *StockValueDelta) GetDelta() int32 {
	if x != nil {
		return x.Delta
	}
	return 0
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tiny_stock_exchange_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tiny_stock_exchange_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_proto_tiny_stock_exchange_proto_rawDescGZIP(), []int{4}
}

func (x *Result) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Result) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_tiny_stock_exchange_proto protoreflect.FileDescriptor

var file_proto_tiny_stock_exchange_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x69, 0x6e, 0x79, 0x5f, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x74, 0x69, 0x6e, 0x79, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x65, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1d, 0x0a, 0x1b, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x44, 0x65, 0x6c,
	0x74, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x33, 0x0a, 0x05, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x45, 0x0a, 0x0f, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x44, 0x65, 0x6c,
	0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x22, 0x3c, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0x8e, 0x04, 0x0a, 0x11, 0x54, 0x69, 0x6e, 0x79, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x45, 0x0a, 0x08, 0x4e, 0x65,
	0x77, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1a, 0x2e, 0x74, 0x69, 0x6e, 0x79, 0x5f, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x1a, 0x1b, 0x2e, 0x74, 0x69, 0x6e, 0x79, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x12, 0x48, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x12, 0x1a, 0x2e, 0x74, 0x69, 0x6e, 0x79, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x65, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x1a, 0x1b, 0x2e, 0x74,
	0x69, 0x6e, 0x79, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1a, 0x2e, 0x74, 0x69, 0x6e,
	0x79, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x1a, 0x1b, 0x2e, 0x74, 0x69, 0x6e, 0x79, 0x5f, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x73, 0x12, 0x26, 0x2e, 0x74, 0x69, 0x6e, 0x79, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x69,
	0x6e, 0x79, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x22, 0x00, 0x30, 0x01, 0x12, 0x54, 0x0a, 0x0d, 0x4e,
	0x65, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x24, 0x2e, 0x74,
	0x69, 0x6e, 0x79, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x44, 0x65, 0x6c,
	0x74, 0x61, 0x1a, 0x1b, 0x2e, 0x74, 0x69, 0x6e, 0x79, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x12, 0x72, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x73, 0x12, 0x30, 0x2e, 0x74, 0x69, 0x6e, 0x79,
	0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x44, 0x65,
	0x6c, 0x74, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x74, 0x69,
	0x6e, 0x79, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x44, 0x65, 0x6c, 0x74,
	0x61, 0x22, 0x00, 0x30, 0x01, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x32, 0x62, 0x65, 0x65, 0x6e, 0x73, 0x2f, 0x74, 0x69, 0x6e, 0x79, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_tiny_stock_exchange_proto_rawDescOnce sync.Once
	file_proto_tiny_stock_exchange_proto_rawDescData = file_proto_tiny_stock_exchange_proto_rawDesc
)

func file_proto_tiny_stock_exchange_proto_rawDescGZIP() []byte {
	file_proto_tiny_stock_exchange_proto_rawDescOnce.Do(func() {
		file_proto_tiny_stock_exchange_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_tiny_stock_exchange_proto_rawDescData)
	})
	return file_proto_tiny_stock_exchange_proto_rawDescData
}

var file_proto_tiny_stock_exchange_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_tiny_stock_exchange_proto_goTypes = []interface{}{
	(*ListStocksRequest)(nil),           // 0: tiny_stock_exchange.ListStocksRequest
	(*ListStockValueDeltasRequest)(nil), // 1: tiny_stock_exchange.ListStockValueDeltasRequest
	(*Stock)(nil),                       // 2: tiny_stock_exchange.Stock
	(*StockValueDelta)(nil),             // 3: tiny_stock_exchange.StockValueDelta
	(*Result)(nil),                      // 4: tiny_stock_exchange.Result
}
var file_proto_tiny_stock_exchange_proto_depIdxs = []int32{
	2, // 0: tiny_stock_exchange.TinyStockExchange.NewStock:input_type -> tiny_stock_exchange.Stock
	2, // 1: tiny_stock_exchange.TinyStockExchange.RemoveStock:input_type -> tiny_stock_exchange.Stock
	2, // 2: tiny_stock_exchange.TinyStockExchange.UpdateStock:input_type -> tiny_stock_exchange.Stock
	0, // 3: tiny_stock_exchange.TinyStockExchange.ListStocks:input_type -> tiny_stock_exchange.ListStocksRequest
	3, // 4: tiny_stock_exchange.TinyStockExchange.NewValueDelta:input_type -> tiny_stock_exchange.StockValueDelta
	1, // 5: tiny_stock_exchange.TinyStockExchange.ListStockValueDeltas:input_type -> tiny_stock_exchange.ListStockValueDeltasRequest
	4, // 6: tiny_stock_exchange.TinyStockExchange.NewStock:output_type -> tiny_stock_exchange.Result
	4, // 7: tiny_stock_exchange.TinyStockExchange.RemoveStock:output_type -> tiny_stock_exchange.Result
	4, // 8: tiny_stock_exchange.TinyStockExchange.UpdateStock:output_type -> tiny_stock_exchange.Result
	2, // 9: tiny_stock_exchange.TinyStockExchange.ListStocks:output_type -> tiny_stock_exchange.Stock
	4, // 10: tiny_stock_exchange.TinyStockExchange.NewValueDelta:output_type -> tiny_stock_exchange.Result
	3, // 11: tiny_stock_exchange.TinyStockExchange.ListStockValueDeltas:output_type -> tiny_stock_exchange.StockValueDelta
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_tiny_stock_exchange_proto_init() }
func file_proto_tiny_stock_exchange_proto_init() {
	if File_proto_tiny_stock_exchange_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_tiny_stock_exchange_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStocksRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_tiny_stock_exchange_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStockValueDeltasRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_tiny_stock_exchange_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stock); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_tiny_stock_exchange_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockValueDelta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_tiny_stock_exchange_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_tiny_stock_exchange_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_tiny_stock_exchange_proto_goTypes,
		DependencyIndexes: file_proto_tiny_stock_exchange_proto_depIdxs,
		MessageInfos:      file_proto_tiny_stock_exchange_proto_msgTypes,
	}.Build()
	File_proto_tiny_stock_exchange_proto = out.File
	file_proto_tiny_stock_exchange_proto_rawDesc = nil
	file_proto_tiny_stock_exchange_proto_goTypes = nil
	file_proto_tiny_stock_exchange_proto_depIdxs = nil
}
