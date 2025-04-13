// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: bookamovie.proto

package bookamoviev2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Movie struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Genre         string                 `protobuf:"bytes,2,opt,name=genre,proto3" json:"genre,omitempty"`
	Country       string                 `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Premier       *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=premier,proto3" json:"premier,omitempty"`
	Duration      *durationpb.Duration   `protobuf:"bytes,5,opt,name=duration,proto3" json:"duration,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Movie) Reset() {
	*x = Movie{}
	mi := &file_bookamovie_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_bookamovie_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movie.ProtoReflect.Descriptor instead.
func (*Movie) Descriptor() ([]byte, []int) {
	return file_bookamovie_proto_rawDescGZIP(), []int{0}
}

func (x *Movie) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Movie) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *Movie) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Movie) GetPremier() *timestamppb.Timestamp {
	if x != nil {
		return x.Premier
	}
	return nil
}

func (x *Movie) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

type Cinema struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Location      string                 `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Cinema) Reset() {
	*x = Cinema{}
	mi := &file_bookamovie_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cinema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cinema) ProtoMessage() {}

func (x *Cinema) ProtoReflect() protoreflect.Message {
	mi := &file_bookamovie_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cinema.ProtoReflect.Descriptor instead.
func (*Cinema) Descriptor() ([]byte, []int) {
	return file_bookamovie_proto_rawDescGZIP(), []int{1}
}

func (x *Cinema) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cinema) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type Session struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Screen        int32                  `protobuf:"varint,1,opt,name=screen,proto3" json:"screen,omitempty"`
	Seat          int32                  `protobuf:"varint,2,opt,name=seat,proto3" json:"seat,omitempty"`
	Date          *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Session) Reset() {
	*x = Session{}
	mi := &file_bookamovie_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_bookamovie_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_bookamovie_proto_rawDescGZIP(), []int{2}
}

func (x *Session) GetScreen() int32 {
	if x != nil {
		return x.Screen
	}
	return 0
}

func (x *Session) GetSeat() int32 {
	if x != nil {
		return x.Seat
	}
	return 0
}

func (x *Session) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ticket        string                 `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_bookamovie_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_bookamovie_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_bookamovie_proto_rawDescGZIP(), []int{3}
}

func (x *Order) GetTicket() string {
	if x != nil {
		return x.Ticket
	}
	return ""
}

type BookRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cinema        *Cinema                `protobuf:"bytes,1,opt,name=cinema,proto3" json:"cinema,omitempty"`
	Movie         *Movie                 `protobuf:"bytes,2,opt,name=movie,proto3" json:"movie,omitempty"`
	Session       *Session               `protobuf:"bytes,3,opt,name=session,proto3" json:"session,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BookRequest) Reset() {
	*x = BookRequest{}
	mi := &file_bookamovie_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookRequest) ProtoMessage() {}

func (x *BookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookamovie_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookRequest.ProtoReflect.Descriptor instead.
func (*BookRequest) Descriptor() ([]byte, []int) {
	return file_bookamovie_proto_rawDescGZIP(), []int{4}
}

func (x *BookRequest) GetCinema() *Cinema {
	if x != nil {
		return x.Cinema
	}
	return nil
}

func (x *BookRequest) GetMovie() *Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

func (x *BookRequest) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

type BookResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Order         *Order                 `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	Cinema        *Cinema                `protobuf:"bytes,2,opt,name=cinema,proto3" json:"cinema,omitempty"`
	Movie         *Movie                 `protobuf:"bytes,3,opt,name=movie,proto3" json:"movie,omitempty"`
	Session       *Session               `protobuf:"bytes,4,opt,name=session,proto3" json:"session,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BookResponse) Reset() {
	*x = BookResponse{}
	mi := &file_bookamovie_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookResponse) ProtoMessage() {}

func (x *BookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bookamovie_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookResponse.ProtoReflect.Descriptor instead.
func (*BookResponse) Descriptor() ([]byte, []int) {
	return file_bookamovie_proto_rawDescGZIP(), []int{5}
}

func (x *BookResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

func (x *BookResponse) GetCinema() *Cinema {
	if x != nil {
		return x.Cinema
	}
	return nil
}

func (x *BookResponse) GetMovie() *Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

func (x *BookResponse) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

var File_bookamovie_proto protoreflect.FileDescriptor

var file_bookamovie_proto_rawDesc = string([]byte{
	0x0a, 0x10, 0x62, 0x6f, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xba, 0x01, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x34, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x70, 0x72,
	0x65, 0x6d, 0x69, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x38, 0x0a, 0x06,
	0x43, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x65, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x61,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x65, 0x61, 0x74, 0x12, 0x2e, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x1f, 0x0a,
	0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x91,
	0x01, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a,
	0x0a, 0x06, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x43, 0x69, 0x6e, 0x65,
	0x6d, 0x61, 0x52, 0x06, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x12, 0x27, 0x0a, 0x05, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x61, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x05, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0xbb, 0x01, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x06,
	0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x43, 0x69, 0x6e, 0x65, 0x6d, 0x61,
	0x52, 0x06, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x12, 0x27, 0x0a, 0x05, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x61, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x05, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x12, 0x2d, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x32, 0x47, 0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6b, 0x61, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x39,
	0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x61, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x62, 0x6f, 0x6f,
	0x6b, 0x61, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2f, 0x76, 0x32, 0x3b, 0x62, 0x6f, 0x6f, 0x6b, 0x61,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_bookamovie_proto_rawDescOnce sync.Once
	file_bookamovie_proto_rawDescData []byte
)

func file_bookamovie_proto_rawDescGZIP() []byte {
	file_bookamovie_proto_rawDescOnce.Do(func() {
		file_bookamovie_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_bookamovie_proto_rawDesc), len(file_bookamovie_proto_rawDesc)))
	})
	return file_bookamovie_proto_rawDescData
}

var file_bookamovie_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_bookamovie_proto_goTypes = []any{
	(*Movie)(nil),                 // 0: bookamovie.Movie
	(*Cinema)(nil),                // 1: bookamovie.Cinema
	(*Session)(nil),               // 2: bookamovie.Session
	(*Order)(nil),                 // 3: bookamovie.Order
	(*BookRequest)(nil),           // 4: bookamovie.BookRequest
	(*BookResponse)(nil),          // 5: bookamovie.BookResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 7: google.protobuf.Duration
}
var file_bookamovie_proto_depIdxs = []int32{
	6,  // 0: bookamovie.Movie.premier:type_name -> google.protobuf.Timestamp
	7,  // 1: bookamovie.Movie.duration:type_name -> google.protobuf.Duration
	6,  // 2: bookamovie.Session.date:type_name -> google.protobuf.Timestamp
	1,  // 3: bookamovie.BookRequest.cinema:type_name -> bookamovie.Cinema
	0,  // 4: bookamovie.BookRequest.movie:type_name -> bookamovie.Movie
	2,  // 5: bookamovie.BookRequest.session:type_name -> bookamovie.Session
	3,  // 6: bookamovie.BookResponse.order:type_name -> bookamovie.Order
	1,  // 7: bookamovie.BookResponse.cinema:type_name -> bookamovie.Cinema
	0,  // 8: bookamovie.BookResponse.movie:type_name -> bookamovie.Movie
	2,  // 9: bookamovie.BookResponse.session:type_name -> bookamovie.Session
	4,  // 10: bookamovie.BookaMovie.Book:input_type -> bookamovie.BookRequest
	5,  // 11: bookamovie.BookaMovie.Book:output_type -> bookamovie.BookResponse
	11, // [11:12] is the sub-list for method output_type
	10, // [10:11] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_bookamovie_proto_init() }
func file_bookamovie_proto_init() {
	if File_bookamovie_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_bookamovie_proto_rawDesc), len(file_bookamovie_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bookamovie_proto_goTypes,
		DependencyIndexes: file_bookamovie_proto_depIdxs,
		MessageInfos:      file_bookamovie_proto_msgTypes,
	}.Build()
	File_bookamovie_proto = out.File
	file_bookamovie_proto_goTypes = nil
	file_bookamovie_proto_depIdxs = nil
}
