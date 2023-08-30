// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: api/grpc/fetch-server.proto

package grpcapi

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

type Null struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Null) Reset() {
	*x = Null{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_fetch_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Null) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Null) ProtoMessage() {}

func (x *Null) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_fetch_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Null.ProtoReflect.Descriptor instead.
func (*Null) Descriptor() ([]byte, []int) {
	return file_api_grpc_fetch_server_proto_rawDescGZIP(), []int{0}
}

type StockId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockId string `protobuf:"bytes,1,opt,name=stock_id,json=stockId,proto3" json:"stock_id,omitempty"`
}

func (x *StockId) Reset() {
	*x = StockId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_fetch_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockId) ProtoMessage() {}

func (x *StockId) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_fetch_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockId.ProtoReflect.Descriptor instead.
func (*StockId) Descriptor() ([]byte, []int) {
	return file_api_grpc_fetch_server_proto_rawDescGZIP(), []int{1}
}

func (x *StockId) GetStockId() string {
	if x != nil {
		return x.StockId
	}
	return ""
}

// meaning of option_status:
//  0 - off
//  1 - on
//  -1 - none
//  other - error
type OptionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OptionStatus int32 `protobuf:"varint,1,opt,name=option_status,json=optionStatus,proto3" json:"option_status,omitempty"`
}

func (x *OptionStatus) Reset() {
	*x = OptionStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_fetch_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionStatus) ProtoMessage() {}

func (x *OptionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_fetch_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionStatus.ProtoReflect.Descriptor instead.
func (*OptionStatus) Descriptor() ([]byte, []int) {
	return file_api_grpc_fetch_server_proto_rawDescGZIP(), []int{2}
}

func (x *OptionStatus) GetOptionStatus() int32 {
	if x != nil {
		return x.OptionStatus
	}
	return 0
}

type StockConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockId       string        `protobuf:"bytes,1,opt,name=stock_id,json=stockId,proto3" json:"stock_id,omitempty"`
	Relayable     *OptionStatus `protobuf:"bytes,2,opt,name=relayable,proto3" json:"relayable,omitempty"`
	Transmittable *OptionStatus `protobuf:"bytes,3,opt,name=transmittable,proto3" json:"transmittable,omitempty"`
	Storeable     *OptionStatus `protobuf:"bytes,4,opt,name=storeable,proto3" json:"storeable,omitempty"`
}

func (x *StockConfig) Reset() {
	*x = StockConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_fetch_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockConfig) ProtoMessage() {}

func (x *StockConfig) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_fetch_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockConfig.ProtoReflect.Descriptor instead.
func (*StockConfig) Descriptor() ([]byte, []int) {
	return file_api_grpc_fetch_server_proto_rawDescGZIP(), []int{3}
}

func (x *StockConfig) GetStockId() string {
	if x != nil {
		return x.StockId
	}
	return ""
}

func (x *StockConfig) GetRelayable() *OptionStatus {
	if x != nil {
		return x.Relayable
	}
	return nil
}

func (x *StockConfig) GetTransmittable() *OptionStatus {
	if x != nil {
		return x.Transmittable
	}
	return nil
}

func (x *StockConfig) GetStoreable() *OptionStatus {
	if x != nil {
		return x.Storeable
	}
	return nil
}

type StockConfigList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockConfig []*StockConfig `protobuf:"bytes,1,rep,name=stock_config,json=stockConfig,proto3" json:"stock_config,omitempty"`
}

func (x *StockConfigList) Reset() {
	*x = StockConfigList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_fetch_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockConfigList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockConfigList) ProtoMessage() {}

func (x *StockConfigList) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_fetch_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockConfigList.ProtoReflect.Descriptor instead.
func (*StockConfigList) Descriptor() ([]byte, []int) {
	return file_api_grpc_fetch_server_proto_rawDescGZIP(), []int{4}
}

func (x *StockConfigList) GetStockConfig() []*StockConfig {
	if x != nil {
		return x.StockConfig
	}
	return nil
}

type ReturnMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ReturnMessage) Reset() {
	*x = ReturnMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_fetch_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReturnMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReturnMessage) ProtoMessage() {}

func (x *ReturnMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_fetch_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReturnMessage.ProtoReflect.Descriptor instead.
func (*ReturnMessage) Descriptor() ([]byte, []int) {
	return file_api_grpc_fetch_server_proto_rawDescGZIP(), []int{5}
}

func (x *ReturnMessage) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *ReturnMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ReturnMessageList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReturnMessage []*ReturnMessage `protobuf:"bytes,1,rep,name=return_message,json=returnMessage,proto3" json:"return_message,omitempty"`
}

func (x *ReturnMessageList) Reset() {
	*x = ReturnMessageList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_fetch_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReturnMessageList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReturnMessageList) ProtoMessage() {}

func (x *ReturnMessageList) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_fetch_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReturnMessageList.ProtoReflect.Descriptor instead.
func (*ReturnMessageList) Descriptor() ([]byte, []int) {
	return file_api_grpc_fetch_server_proto_rawDescGZIP(), []int{6}
}

func (x *ReturnMessageList) GetReturnMessage() []*ReturnMessage {
	if x != nil {
		return x.ReturnMessage
	}
	return nil
}

var File_api_grpc_fetch_server_proto protoreflect.FileDescriptor

var file_api_grpc_fetch_server_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x65, 0x74, 0x63, 0x68,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67,
	0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x22, 0x06, 0x0a, 0x04, 0x4e, 0x75, 0x6c, 0x6c, 0x22, 0x24,
	0x0a, 0x07, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x0c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xcf, 0x01, 0x0a, 0x0b, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x09, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70,
	0x69, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09,
	0x72, 0x65, 0x6c, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x3b, 0x0a, 0x0d, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69,
	0x74, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x61, 0x70, 0x69, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x4a, 0x0a, 0x0f, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x37,
	0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x41, 0x0a, 0x0d, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x52, 0x0a, 0x11, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x3d, 0x0a, 0x0e, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70,
	0x69, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x0d, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xab,
	0x02, 0x0a, 0x11, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x46, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x6e, 0x65, 0x12, 0x14, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x15,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x4d, 0x61, 0x6e, 0x79, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x3d, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f,
	0x6e, 0x65, 0x12, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x49, 0x64, 0x1a, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x41, 0x6c,
	0x6c, 0x12, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x75, 0x6c, 0x6c,
	0x1a, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x6f, 0x62, 0x6f, 0x6f,
	0x6c, 0x65, 0x61, 0x6e, 0x2f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x67, 0x72,
	0x70, 0x63, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_grpc_fetch_server_proto_rawDescOnce sync.Once
	file_api_grpc_fetch_server_proto_rawDescData = file_api_grpc_fetch_server_proto_rawDesc
)

func file_api_grpc_fetch_server_proto_rawDescGZIP() []byte {
	file_api_grpc_fetch_server_proto_rawDescOnce.Do(func() {
		file_api_grpc_fetch_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_fetch_server_proto_rawDescData)
	})
	return file_api_grpc_fetch_server_proto_rawDescData
}

var file_api_grpc_fetch_server_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_grpc_fetch_server_proto_goTypes = []interface{}{
	(*Null)(nil),              // 0: grpcapi.Null
	(*StockId)(nil),           // 1: grpcapi.StockId
	(*OptionStatus)(nil),      // 2: grpcapi.OptionStatus
	(*StockConfig)(nil),       // 3: grpcapi.StockConfig
	(*StockConfigList)(nil),   // 4: grpcapi.StockConfigList
	(*ReturnMessage)(nil),     // 5: grpcapi.ReturnMessage
	(*ReturnMessageList)(nil), // 6: grpcapi.ReturnMessageList
}
var file_api_grpc_fetch_server_proto_depIdxs = []int32{
	2, // 0: grpcapi.StockConfig.relayable:type_name -> grpcapi.OptionStatus
	2, // 1: grpcapi.StockConfig.transmittable:type_name -> grpcapi.OptionStatus
	2, // 2: grpcapi.StockConfig.storeable:type_name -> grpcapi.OptionStatus
	3, // 3: grpcapi.StockConfigList.stock_config:type_name -> grpcapi.StockConfig
	5, // 4: grpcapi.ReturnMessageList.return_message:type_name -> grpcapi.ReturnMessage
	3, // 5: grpcapi.StockConfigurator.UpdateStockConfigOne:input_type -> grpcapi.StockConfig
	4, // 6: grpcapi.StockConfigurator.UpdateStockConfigMany:input_type -> grpcapi.StockConfigList
	1, // 7: grpcapi.StockConfigurator.GetStockConfigOne:input_type -> grpcapi.StockId
	0, // 8: grpcapi.StockConfigurator.GetStockConfigAll:input_type -> grpcapi.Null
	5, // 9: grpcapi.StockConfigurator.UpdateStockConfigOne:output_type -> grpcapi.ReturnMessage
	6, // 10: grpcapi.StockConfigurator.UpdateStockConfigMany:output_type -> grpcapi.ReturnMessageList
	3, // 11: grpcapi.StockConfigurator.GetStockConfigOne:output_type -> grpcapi.StockConfig
	4, // 12: grpcapi.StockConfigurator.GetStockConfigAll:output_type -> grpcapi.StockConfigList
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_grpc_fetch_server_proto_init() }
func file_api_grpc_fetch_server_proto_init() {
	if File_api_grpc_fetch_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_fetch_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Null); i {
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
		file_api_grpc_fetch_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockId); i {
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
		file_api_grpc_fetch_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionStatus); i {
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
		file_api_grpc_fetch_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockConfig); i {
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
		file_api_grpc_fetch_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockConfigList); i {
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
		file_api_grpc_fetch_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReturnMessage); i {
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
		file_api_grpc_fetch_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReturnMessageList); i {
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
			RawDescriptor: file_api_grpc_fetch_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_fetch_server_proto_goTypes,
		DependencyIndexes: file_api_grpc_fetch_server_proto_depIdxs,
		MessageInfos:      file_api_grpc_fetch_server_proto_msgTypes,
	}.Build()
	File_api_grpc_fetch_server_proto = out.File
	file_api_grpc_fetch_server_proto_rawDesc = nil
	file_api_grpc_fetch_server_proto_goTypes = nil
	file_api_grpc_fetch_server_proto_depIdxs = nil
}
