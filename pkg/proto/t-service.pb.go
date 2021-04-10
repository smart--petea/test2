// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: t-service.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fsyms []string `protobuf:"bytes,1,rep,name=fsyms,proto3" json:"fsyms,omitempty"`
	Tsyms []string `protobuf:"bytes,2,rep,name=tsyms,proto3" json:"tsyms,omitempty"`
}

func (x *TServiceRequest) Reset() {
	*x = TServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_t_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TServiceRequest) ProtoMessage() {}

func (x *TServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_t_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TServiceRequest.ProtoReflect.Descriptor instead.
func (*TServiceRequest) Descriptor() ([]byte, []int) {
	return file_t_service_proto_rawDescGZIP(), []int{0}
}

func (x *TServiceRequest) GetFsyms() []string {
	if x != nil {
		return x.Fsyms
	}
	return nil
}

func (x *TServiceRequest) GetTsyms() []string {
	if x != nil {
		return x.Tsyms
	}
	return nil
}

type TRawCurrency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price           float64 `protobuf:"fixed64,1,opt,name=price,proto3" json:"price,omitempty"`
	Lastupdate      float64 `protobuf:"fixed64,2,opt,name=lastupdate,proto3" json:"lastupdate,omitempty"`
	Volume24Hour    float64 `protobuf:"fixed64,3,opt,name=volume24hour,proto3" json:"volume24hour,omitempty"`
	Volume24Hourto  float64 `protobuf:"fixed64,4,opt,name=volume24hourto,proto3" json:"volume24hourto,omitempty"`
	Open24Hour      float64 `protobuf:"fixed64,5,opt,name=open24hour,proto3" json:"open24hour,omitempty"`
	High24Hour      float64 `protobuf:"fixed64,6,opt,name=high24hour,proto3" json:"high24hour,omitempty"`
	Low24Hour       float64 `protobuf:"fixed64,7,opt,name=low24hour,proto3" json:"low24hour,omitempty"`
	Change24Hour    float64 `protobuf:"fixed64,8,opt,name=change24hour,proto3" json:"change24hour,omitempty"`
	Changepct24Hour float64 `protobuf:"fixed64,9,opt,name=changepct24hour,proto3" json:"changepct24hour,omitempty"`
	Supply          float64 `protobuf:"fixed64,10,opt,name=supply,proto3" json:"supply,omitempty"`
	Mktcap          float64 `protobuf:"fixed64,11,opt,name=mktcap,proto3" json:"mktcap,omitempty"`
}

func (x *TRawCurrency) Reset() {
	*x = TRawCurrency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_t_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TRawCurrency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TRawCurrency) ProtoMessage() {}

func (x *TRawCurrency) ProtoReflect() protoreflect.Message {
	mi := &file_t_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TRawCurrency.ProtoReflect.Descriptor instead.
func (*TRawCurrency) Descriptor() ([]byte, []int) {
	return file_t_service_proto_rawDescGZIP(), []int{1}
}

func (x *TRawCurrency) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *TRawCurrency) GetLastupdate() float64 {
	if x != nil {
		return x.Lastupdate
	}
	return 0
}

func (x *TRawCurrency) GetVolume24Hour() float64 {
	if x != nil {
		return x.Volume24Hour
	}
	return 0
}

func (x *TRawCurrency) GetVolume24Hourto() float64 {
	if x != nil {
		return x.Volume24Hourto
	}
	return 0
}

func (x *TRawCurrency) GetOpen24Hour() float64 {
	if x != nil {
		return x.Open24Hour
	}
	return 0
}

func (x *TRawCurrency) GetHigh24Hour() float64 {
	if x != nil {
		return x.High24Hour
	}
	return 0
}

func (x *TRawCurrency) GetLow24Hour() float64 {
	if x != nil {
		return x.Low24Hour
	}
	return 0
}

func (x *TRawCurrency) GetChange24Hour() float64 {
	if x != nil {
		return x.Change24Hour
	}
	return 0
}

func (x *TRawCurrency) GetChangepct24Hour() float64 {
	if x != nil {
		return x.Changepct24Hour
	}
	return 0
}

func (x *TRawCurrency) GetSupply() float64 {
	if x != nil {
		return x.Supply
	}
	return 0
}

func (x *TRawCurrency) GetMktcap() float64 {
	if x != nil {
		return x.Mktcap
	}
	return 0
}

type TRaw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Currencies map[string]*TRawCurrency `protobuf:"bytes,1,rep,name=currencies,proto3" json:"currencies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TRaw) Reset() {
	*x = TRaw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_t_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TRaw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TRaw) ProtoMessage() {}

func (x *TRaw) ProtoReflect() protoreflect.Message {
	mi := &file_t_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TRaw.ProtoReflect.Descriptor instead.
func (*TRaw) Descriptor() ([]byte, []int) {
	return file_t_service_proto_rawDescGZIP(), []int{2}
}

func (x *TRaw) GetCurrencies() map[string]*TRawCurrency {
	if x != nil {
		return x.Currencies
	}
	return nil
}

type TDisplayCurrency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price           float64 `protobuf:"fixed64,1,opt,name=price,proto3" json:"price,omitempty"`
	Volume24Hour    float64 `protobuf:"fixed64,2,opt,name=volume24hour,proto3" json:"volume24hour,omitempty"`
	Volume24Hourto  float64 `protobuf:"fixed64,3,opt,name=volume24hourto,proto3" json:"volume24hourto,omitempty"`
	Open24Hour      float64 `protobuf:"fixed64,4,opt,name=open24hour,proto3" json:"open24hour,omitempty"`
	High24Hour      float64 `protobuf:"fixed64,5,opt,name=high24hour,proto3" json:"high24hour,omitempty"`
	Low24Hour       float64 `protobuf:"fixed64,6,opt,name=low24hour,proto3" json:"low24hour,omitempty"`
	Change24Hour    float64 `protobuf:"fixed64,7,opt,name=change24hour,proto3" json:"change24hour,omitempty"`
	Changepct24Hour float64 `protobuf:"fixed64,8,opt,name=changepct24hour,proto3" json:"changepct24hour,omitempty"`
}

func (x *TDisplayCurrency) Reset() {
	*x = TDisplayCurrency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_t_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TDisplayCurrency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TDisplayCurrency) ProtoMessage() {}

func (x *TDisplayCurrency) ProtoReflect() protoreflect.Message {
	mi := &file_t_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TDisplayCurrency.ProtoReflect.Descriptor instead.
func (*TDisplayCurrency) Descriptor() ([]byte, []int) {
	return file_t_service_proto_rawDescGZIP(), []int{3}
}

func (x *TDisplayCurrency) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *TDisplayCurrency) GetVolume24Hour() float64 {
	if x != nil {
		return x.Volume24Hour
	}
	return 0
}

func (x *TDisplayCurrency) GetVolume24Hourto() float64 {
	if x != nil {
		return x.Volume24Hourto
	}
	return 0
}

func (x *TDisplayCurrency) GetOpen24Hour() float64 {
	if x != nil {
		return x.Open24Hour
	}
	return 0
}

func (x *TDisplayCurrency) GetHigh24Hour() float64 {
	if x != nil {
		return x.High24Hour
	}
	return 0
}

func (x *TDisplayCurrency) GetLow24Hour() float64 {
	if x != nil {
		return x.Low24Hour
	}
	return 0
}

func (x *TDisplayCurrency) GetChange24Hour() float64 {
	if x != nil {
		return x.Change24Hour
	}
	return 0
}

func (x *TDisplayCurrency) GetChangepct24Hour() float64 {
	if x != nil {
		return x.Changepct24Hour
	}
	return 0
}

type TDisplay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Currencies map[string]*TDisplayCurrency `protobuf:"bytes,1,rep,name=currencies,proto3" json:"currencies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TDisplay) Reset() {
	*x = TDisplay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_t_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TDisplay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TDisplay) ProtoMessage() {}

func (x *TDisplay) ProtoReflect() protoreflect.Message {
	mi := &file_t_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TDisplay.ProtoReflect.Descriptor instead.
func (*TDisplay) Descriptor() ([]byte, []int) {
	return file_t_service_proto_rawDescGZIP(), []int{4}
}

func (x *TDisplay) GetCurrencies() map[string]*TDisplayCurrency {
	if x != nil {
		return x.Currencies
	}
	return nil
}

type TServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Raw     map[string]*TRaw     `protobuf:"bytes,1,rep,name=raw,proto3" json:"raw,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Display map[string]*TDisplay `protobuf:"bytes,2,rep,name=display,proto3" json:"display,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TServiceResponse) Reset() {
	*x = TServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_t_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TServiceResponse) ProtoMessage() {}

func (x *TServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_t_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TServiceResponse.ProtoReflect.Descriptor instead.
func (*TServiceResponse) Descriptor() ([]byte, []int) {
	return file_t_service_proto_rawDescGZIP(), []int{5}
}

func (x *TServiceResponse) GetRaw() map[string]*TRaw {
	if x != nil {
		return x.Raw
	}
	return nil
}

func (x *TServiceResponse) GetDisplay() map[string]*TDisplay {
	if x != nil {
		return x.Display
	}
	return nil
}

var File_t_service_proto protoreflect.FileDescriptor

var file_t_service_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x0f, 0x54, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66,
	0x73, 0x79, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x66, 0x73, 0x79, 0x6d,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x73, 0x79, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x73, 0x79, 0x6d, 0x73, 0x22, 0xec, 0x02, 0x0a, 0x0c, 0x54, 0x52, 0x61, 0x77,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x32, 0x34, 0x68, 0x6f, 0x75, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x32, 0x34, 0x68, 0x6f,
	0x75, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x32, 0x34, 0x68, 0x6f,
	0x75, 0x72, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x32, 0x34, 0x68, 0x6f, 0x75, 0x72, 0x74, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x70,
	0x65, 0x6e, 0x32, 0x34, 0x68, 0x6f, 0x75, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a,
	0x6f, 0x70, 0x65, 0x6e, 0x32, 0x34, 0x68, 0x6f, 0x75, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x69,
	0x67, 0x68, 0x32, 0x34, 0x68, 0x6f, 0x75, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a,
	0x68, 0x69, 0x67, 0x68, 0x32, 0x34, 0x68, 0x6f, 0x75, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f,
	0x77, 0x32, 0x34, 0x68, 0x6f, 0x75, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c,
	0x6f, 0x77, 0x32, 0x34, 0x68, 0x6f, 0x75, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x32, 0x34, 0x68, 0x6f, 0x75, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x32, 0x34, 0x68, 0x6f, 0x75, 0x72, 0x12, 0x28, 0x0a, 0x0f,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x70, 0x63, 0x74, 0x32, 0x34, 0x68, 0x6f, 0x75, 0x72, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x70, 0x63, 0x74,
	0x32, 0x34, 0x68, 0x6f, 0x75, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x6b, 0x74, 0x63, 0x61, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06,
	0x6d, 0x6b, 0x74, 0x63, 0x61, 0x70, 0x22, 0x97, 0x01, 0x0a, 0x04, 0x54, 0x52, 0x61, 0x77, 0x12,
	0x3b, 0x0a, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x52, 0x61, 0x77,
	0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x1a, 0x52, 0x0a, 0x0f,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x52, 0x61, 0x77, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xa0, 0x02, 0x0a, 0x10, 0x54, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x32, 0x34, 0x68, 0x6f, 0x75, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0c, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x32, 0x34, 0x68, 0x6f, 0x75, 0x72, 0x12,
	0x26, 0x0a, 0x0e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x32, 0x34, 0x68, 0x6f, 0x75, 0x72, 0x74,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x32,
	0x34, 0x68, 0x6f, 0x75, 0x72, 0x74, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x6e, 0x32,
	0x34, 0x68, 0x6f, 0x75, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x6f, 0x70, 0x65,
	0x6e, 0x32, 0x34, 0x68, 0x6f, 0x75, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x32,
	0x34, 0x68, 0x6f, 0x75, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x68, 0x69, 0x67,
	0x68, 0x32, 0x34, 0x68, 0x6f, 0x75, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x77, 0x32, 0x34,
	0x68, 0x6f, 0x75, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x77, 0x32,
	0x34, 0x68, 0x6f, 0x75, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x32,
	0x34, 0x68, 0x6f, 0x75, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x32, 0x34, 0x68, 0x6f, 0x75, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x70, 0x63, 0x74, 0x32, 0x34, 0x68, 0x6f, 0x75, 0x72, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x70, 0x63, 0x74, 0x32, 0x34, 0x68,
	0x6f, 0x75, 0x72, 0x22, 0xa3, 0x01, 0x0a, 0x08, 0x54, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x12, 0x3f, 0x0a, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x44, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65,
	0x73, 0x1a, 0x56, 0x0a, 0x0f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x44,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x98, 0x02, 0x0a, 0x10, 0x54, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32,
	0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x61, 0x77, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x72,
	0x61, 0x77, 0x12, 0x3e, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x1a, 0x43, 0x0a, 0x08, 0x52, 0x61, 0x77, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x52, 0x61, 0x77, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4b, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x32, 0x42, 0x0a, 0x08, 0x54, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x36, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_t_service_proto_rawDescOnce sync.Once
	file_t_service_proto_rawDescData = file_t_service_proto_rawDesc
)

func file_t_service_proto_rawDescGZIP() []byte {
	file_t_service_proto_rawDescOnce.Do(func() {
		file_t_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_t_service_proto_rawDescData)
	})
	return file_t_service_proto_rawDescData
}

var file_t_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_t_service_proto_goTypes = []interface{}{
	(*TServiceRequest)(nil),  // 0: proto.TServiceRequest
	(*TRawCurrency)(nil),     // 1: proto.TRawCurrency
	(*TRaw)(nil),             // 2: proto.TRaw
	(*TDisplayCurrency)(nil), // 3: proto.TDisplayCurrency
	(*TDisplay)(nil),         // 4: proto.TDisplay
	(*TServiceResponse)(nil), // 5: proto.TServiceResponse
	nil,                      // 6: proto.TRaw.CurrenciesEntry
	nil,                      // 7: proto.TDisplay.CurrenciesEntry
	nil,                      // 8: proto.TServiceResponse.RawEntry
	nil,                      // 9: proto.TServiceResponse.DisplayEntry
}
var file_t_service_proto_depIdxs = []int32{
	6, // 0: proto.TRaw.currencies:type_name -> proto.TRaw.CurrenciesEntry
	7, // 1: proto.TDisplay.currencies:type_name -> proto.TDisplay.CurrenciesEntry
	8, // 2: proto.TServiceResponse.raw:type_name -> proto.TServiceResponse.RawEntry
	9, // 3: proto.TServiceResponse.display:type_name -> proto.TServiceResponse.DisplayEntry
	1, // 4: proto.TRaw.CurrenciesEntry.value:type_name -> proto.TRawCurrency
	3, // 5: proto.TDisplay.CurrenciesEntry.value:type_name -> proto.TDisplayCurrency
	2, // 6: proto.TServiceResponse.RawEntry.value:type_name -> proto.TRaw
	4, // 7: proto.TServiceResponse.DisplayEntry.value:type_name -> proto.TDisplay
	0, // 8: proto.TService.Get:input_type -> proto.TServiceRequest
	5, // 9: proto.TService.Get:output_type -> proto.TServiceResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_t_service_proto_init() }
func file_t_service_proto_init() {
	if File_t_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_t_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TServiceRequest); i {
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
		file_t_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TRawCurrency); i {
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
		file_t_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TRaw); i {
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
		file_t_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TDisplayCurrency); i {
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
		file_t_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TDisplay); i {
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
		file_t_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TServiceResponse); i {
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
			RawDescriptor: file_t_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_t_service_proto_goTypes,
		DependencyIndexes: file_t_service_proto_depIdxs,
		MessageInfos:      file_t_service_proto_msgTypes,
	}.Build()
	File_t_service_proto = out.File
	file_t_service_proto_rawDesc = nil
	file_t_service_proto_goTypes = nil
	file_t_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TServiceClient is the client API for TService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TServiceClient interface {
	Get(ctx context.Context, in *TServiceRequest, opts ...grpc.CallOption) (*TServiceResponse, error)
}

type tServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTServiceClient(cc grpc.ClientConnInterface) TServiceClient {
	return &tServiceClient{cc}
}

func (c *tServiceClient) Get(ctx context.Context, in *TServiceRequest, opts ...grpc.CallOption) (*TServiceResponse, error) {
	out := new(TServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.TService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TServiceServer is the server API for TService service.
type TServiceServer interface {
	Get(context.Context, *TServiceRequest) (*TServiceResponse, error)
}

// UnimplementedTServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTServiceServer struct {
}

func (*UnimplementedTServiceServer) Get(context.Context, *TServiceRequest) (*TServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterTServiceServer(s *grpc.Server, srv TServiceServer) {
	s.RegisterService(&_TService_serviceDesc, srv)
}

func _TService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TServiceServer).Get(ctx, req.(*TServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TService",
	HandlerType: (*TServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _TService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "t-service.proto",
}