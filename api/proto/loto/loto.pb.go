// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.0
// source: proto/loto.proto

package loto

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type LotoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LotoRequest) Reset() {
	*x = LotoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_loto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LotoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LotoRequest) ProtoMessage() {}

func (x *LotoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_loto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LotoRequest.ProtoReflect.Descriptor instead.
func (*LotoRequest) Descriptor() ([]byte, []int) {
	return file_proto_loto_proto_rawDescGZIP(), []int{0}
}

type LotoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool            `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string          `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	List   []*Participants `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *LotoResponse) Reset() {
	*x = LotoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_loto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LotoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LotoResponse) ProtoMessage() {}

func (x *LotoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_loto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LotoResponse.ProtoReflect.Descriptor instead.
func (*LotoResponse) Descriptor() ([]byte, []int) {
	return file_proto_loto_proto_rawDescGZIP(), []int{1}
}

func (x *LotoResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *LotoResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *LotoResponse) GetList() []*Participants {
	if x != nil {
		return x.List
	}
	return nil
}

type Participants struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketNumber int64  `protobuf:"varint,1,opt,name=ticketNumber,proto3" json:"ticketNumber,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Surname      string `protobuf:"bytes,3,opt,name=surname,proto3" json:"surname,omitempty"`
	MiddleName   string `protobuf:"bytes,4,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	Phone        string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	Spi          string `protobuf:"bytes,6,opt,name=spi,proto3" json:"spi,omitempty"`
	Prize        string `protobuf:"bytes,7,opt,name=prize,proto3" json:"prize,omitempty"`
}

func (x *Participants) Reset() {
	*x = Participants{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_loto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Participants) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Participants) ProtoMessage() {}

func (x *Participants) ProtoReflect() protoreflect.Message {
	mi := &file_proto_loto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Participants.ProtoReflect.Descriptor instead.
func (*Participants) Descriptor() ([]byte, []int) {
	return file_proto_loto_proto_rawDescGZIP(), []int{2}
}

func (x *Participants) GetTicketNumber() int64 {
	if x != nil {
		return x.TicketNumber
	}
	return 0
}

func (x *Participants) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Participants) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *Participants) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *Participants) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Participants) GetSpi() string {
	if x != nil {
		return x.Spi
	}
	return ""
}

func (x *Participants) GetPrize() string {
	if x != nil {
		return x.Prize
	}
	return ""
}

type UploadDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UploadDataRequest) Reset() {
	*x = UploadDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_loto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadDataRequest) ProtoMessage() {}

func (x *UploadDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_loto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadDataRequest.ProtoReflect.Descriptor instead.
func (*UploadDataRequest) Descriptor() ([]byte, []int) {
	return file_proto_loto_proto_rawDescGZIP(), []int{3}
}

type UploadDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UploadDataResponse) Reset() {
	*x = UploadDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_loto_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadDataResponse) ProtoMessage() {}

func (x *UploadDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_loto_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadDataResponse.ProtoReflect.Descriptor instead.
func (*UploadDataResponse) Descriptor() ([]byte, []int) {
	return file_proto_loto_proto_rawDescGZIP(), []int{4}
}

func (x *UploadDataResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UploadDataResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type FillPrizesTableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FillPrizesTableRequest) Reset() {
	*x = FillPrizesTableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_loto_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FillPrizesTableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FillPrizesTableRequest) ProtoMessage() {}

func (x *FillPrizesTableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_loto_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FillPrizesTableRequest.ProtoReflect.Descriptor instead.
func (*FillPrizesTableRequest) Descriptor() ([]byte, []int) {
	return file_proto_loto_proto_rawDescGZIP(), []int{5}
}

type FillPrizesTableResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *FillPrizesTableResponse) Reset() {
	*x = FillPrizesTableResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_loto_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FillPrizesTableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FillPrizesTableResponse) ProtoMessage() {}

func (x *FillPrizesTableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_loto_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FillPrizesTableResponse.ProtoReflect.Descriptor instead.
func (*FillPrizesTableResponse) Descriptor() ([]byte, []int) {
	return file_proto_loto_proto_rawDescGZIP(), []int{6}
}

func (x *FillPrizesTableResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *FillPrizesTableResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_loto_proto protoreflect.FileDescriptor

var file_proto_loto_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x6c, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x4c, 0x6f, 0x74, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x64, 0x0a, 0x0c, 0x4c, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x26, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0xbf, 0x01, 0x0a, 0x0c,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0c,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x70, 0x69, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x73, 0x70, 0x69, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x7a, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x7a, 0x65, 0x22, 0x13, 0x0a,
	0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x48, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x18, 0x0a, 0x16,
	0x46, 0x69, 0x6c, 0x6c, 0x50, 0x72, 0x69, 0x7a, 0x65, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x17, 0x46, 0x69, 0x6c, 0x6c, 0x50, 0x72,
	0x69, 0x7a, 0x65, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x8f, 0x02, 0x0a, 0x04, 0x4c, 0x6f, 0x74, 0x6f, 0x12, 0x49,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x11, 0x2e, 0x6c,
	0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x6c, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x65, 0x74, 0x44, 0x72, 0x61, 0x77, 0x12, 0x58, 0x0a, 0x0a, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x2e, 0x6c, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x6c, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x11, 0x12, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x62, 0x0a, 0x0f, 0x46, 0x69, 0x6c, 0x6c, 0x50, 0x72, 0x69, 0x7a, 0x65,
	0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x6c, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x6c, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x16, 0x12, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x6c, 0x6c, 0x50, 0x72, 0x69, 0x7a,
	0x65, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x6c, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_loto_proto_rawDescOnce sync.Once
	file_proto_loto_proto_rawDescData = file_proto_loto_proto_rawDesc
)

func file_proto_loto_proto_rawDescGZIP() []byte {
	file_proto_loto_proto_rawDescOnce.Do(func() {
		file_proto_loto_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_loto_proto_rawDescData)
	})
	return file_proto_loto_proto_rawDescData
}

var file_proto_loto_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_loto_proto_goTypes = []interface{}{
	(*LotoRequest)(nil),             // 0: loto.LotoRequest
	(*LotoResponse)(nil),            // 1: loto.LotoResponse
	(*Participants)(nil),            // 2: loto.Participants
	(*UploadDataRequest)(nil),       // 3: loto.UploadDataRequest
	(*UploadDataResponse)(nil),      // 4: loto.UploadDataResponse
	(*FillPrizesTableRequest)(nil),  // 5: loto.FillPrizesTableRequest
	(*FillPrizesTableResponse)(nil), // 6: loto.FillPrizesTableResponse
}
var file_proto_loto_proto_depIdxs = []int32{
	2, // 0: loto.LotoResponse.list:type_name -> loto.Participants
	0, // 1: loto.Loto.GetWinners:input_type -> loto.LotoRequest
	3, // 2: loto.Loto.UploadData:input_type -> loto.UploadDataRequest
	3, // 3: loto.Loto.FillPrizesTable:input_type -> loto.UploadDataRequest
	1, // 4: loto.Loto.GetWinners:output_type -> loto.LotoResponse
	4, // 5: loto.Loto.UploadData:output_type -> loto.UploadDataResponse
	4, // 6: loto.Loto.FillPrizesTable:output_type -> loto.UploadDataResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_loto_proto_init() }
func file_proto_loto_proto_init() {
	if File_proto_loto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_loto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LotoRequest); i {
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
		file_proto_loto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LotoResponse); i {
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
		file_proto_loto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Participants); i {
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
		file_proto_loto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadDataRequest); i {
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
		file_proto_loto_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadDataResponse); i {
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
		file_proto_loto_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FillPrizesTableRequest); i {
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
		file_proto_loto_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FillPrizesTableResponse); i {
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
			RawDescriptor: file_proto_loto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_loto_proto_goTypes,
		DependencyIndexes: file_proto_loto_proto_depIdxs,
		MessageInfos:      file_proto_loto_proto_msgTypes,
	}.Build()
	File_proto_loto_proto = out.File
	file_proto_loto_proto_rawDesc = nil
	file_proto_loto_proto_goTypes = nil
	file_proto_loto_proto_depIdxs = nil
}
