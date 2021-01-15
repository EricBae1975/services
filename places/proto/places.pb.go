// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/places.proto

package places

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Location struct {
	Id                   string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Metadata             map[string]string     `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Timestamp            *timestamp.Timestamp  `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Latitude             *wrappers.DoubleValue `protobuf:"bytes,5,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            *wrappers.DoubleValue `protobuf:"bytes,6,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b635ff9d2e2d652, []int{0}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Location) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Location) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Location) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Location) GetLatitude() *wrappers.DoubleValue {
	if m != nil {
		return m.Latitude
	}
	return nil
}

func (m *Location) GetLongitude() *wrappers.DoubleValue {
	if m != nil {
		return m.Longitude
	}
	return nil
}

type SaveRequest struct {
	Places               []*Location `protobuf:"bytes,1,rep,name=places,proto3" json:"places,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SaveRequest) Reset()         { *m = SaveRequest{} }
func (m *SaveRequest) String() string { return proto.CompactTextString(m) }
func (*SaveRequest) ProtoMessage()    {}
func (*SaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b635ff9d2e2d652, []int{1}
}

func (m *SaveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveRequest.Unmarshal(m, b)
}
func (m *SaveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveRequest.Marshal(b, m, deterministic)
}
func (m *SaveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveRequest.Merge(m, src)
}
func (m *SaveRequest) XXX_Size() int {
	return xxx_messageInfo_SaveRequest.Size(m)
}
func (m *SaveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveRequest proto.InternalMessageInfo

func (m *SaveRequest) GetPlaces() []*Location {
	if m != nil {
		return m.Places
	}
	return nil
}

type SaveResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveResponse) Reset()         { *m = SaveResponse{} }
func (m *SaveResponse) String() string { return proto.CompactTextString(m) }
func (*SaveResponse) ProtoMessage()    {}
func (*SaveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b635ff9d2e2d652, []int{2}
}

func (m *SaveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveResponse.Unmarshal(m, b)
}
func (m *SaveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveResponse.Marshal(b, m, deterministic)
}
func (m *SaveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveResponse.Merge(m, src)
}
func (m *SaveResponse) XXX_Size() int {
	return xxx_messageInfo_SaveResponse.Size(m)
}
func (m *SaveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SaveResponse proto.InternalMessageInfo

type LastRequest struct {
	Ids                  []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LastRequest) Reset()         { *m = LastRequest{} }
func (m *LastRequest) String() string { return proto.CompactTextString(m) }
func (*LastRequest) ProtoMessage()    {}
func (*LastRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b635ff9d2e2d652, []int{3}
}

func (m *LastRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LastRequest.Unmarshal(m, b)
}
func (m *LastRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LastRequest.Marshal(b, m, deterministic)
}
func (m *LastRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LastRequest.Merge(m, src)
}
func (m *LastRequest) XXX_Size() int {
	return xxx_messageInfo_LastRequest.Size(m)
}
func (m *LastRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LastRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LastRequest proto.InternalMessageInfo

func (m *LastRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type ListResponse struct {
	Places               []*Location `protobuf:"bytes,1,rep,name=places,proto3" json:"places,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b635ff9d2e2d652, []int{4}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetPlaces() []*Location {
	if m != nil {
		return m.Places
	}
	return nil
}

type NearRequest struct {
	Latitude  *wrappers.DoubleValue `protobuf:"bytes,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude *wrappers.DoubleValue `protobuf:"bytes,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	// radius to search within, units km
	Radius               *wrappers.DoubleValue `protobuf:"bytes,3,opt,name=radius,proto3" json:"radius,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *NearRequest) Reset()         { *m = NearRequest{} }
func (m *NearRequest) String() string { return proto.CompactTextString(m) }
func (*NearRequest) ProtoMessage()    {}
func (*NearRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b635ff9d2e2d652, []int{5}
}

func (m *NearRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NearRequest.Unmarshal(m, b)
}
func (m *NearRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NearRequest.Marshal(b, m, deterministic)
}
func (m *NearRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NearRequest.Merge(m, src)
}
func (m *NearRequest) XXX_Size() int {
	return xxx_messageInfo_NearRequest.Size(m)
}
func (m *NearRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NearRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NearRequest proto.InternalMessageInfo

func (m *NearRequest) GetLatitude() *wrappers.DoubleValue {
	if m != nil {
		return m.Latitude
	}
	return nil
}

func (m *NearRequest) GetLongitude() *wrappers.DoubleValue {
	if m != nil {
		return m.Longitude
	}
	return nil
}

func (m *NearRequest) GetRadius() *wrappers.DoubleValue {
	if m != nil {
		return m.Radius
	}
	return nil
}

type ReadRequest struct {
	Ids                  []string             `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	After                *timestamp.Timestamp `protobuf:"bytes,2,opt,name=after,proto3" json:"after,omitempty"`
	Before               *timestamp.Timestamp `protobuf:"bytes,3,opt,name=before,proto3" json:"before,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b635ff9d2e2d652, []int{6}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *ReadRequest) GetAfter() *timestamp.Timestamp {
	if m != nil {
		return m.After
	}
	return nil
}

func (m *ReadRequest) GetBefore() *timestamp.Timestamp {
	if m != nil {
		return m.Before
	}
	return nil
}

func init() {
	proto.RegisterType((*Location)(nil), "places.Location")
	proto.RegisterMapType((map[string]string)(nil), "places.Location.MetadataEntry")
	proto.RegisterType((*SaveRequest)(nil), "places.SaveRequest")
	proto.RegisterType((*SaveResponse)(nil), "places.SaveResponse")
	proto.RegisterType((*LastRequest)(nil), "places.LastRequest")
	proto.RegisterType((*ListResponse)(nil), "places.ListResponse")
	proto.RegisterType((*NearRequest)(nil), "places.NearRequest")
	proto.RegisterType((*ReadRequest)(nil), "places.ReadRequest")
}

func init() { proto.RegisterFile("proto/places.proto", fileDescriptor_3b635ff9d2e2d652) }

var fileDescriptor_3b635ff9d2e2d652 = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0x19, 0x3b, 0xb1, 0x92, 0xeb, 0x52, 0x45, 0x97, 0x2e, 0x2c, 0x0b, 0xb5, 0x91, 0x57,
	0x59, 0x39, 0x28, 0x45, 0x22, 0x0a, 0x5b, 0xd8, 0x05, 0x84, 0x0c, 0x62, 0x3f, 0xa9, 0x6f, 0x22,
	0x0b, 0xc7, 0x63, 0xec, 0x71, 0x51, 0x1f, 0x80, 0x87, 0xe2, 0x49, 0xd8, 0xf0, 0x30, 0x68, 0x7e,
	0x1c, 0xbb, 0x45, 0x8d, 0x1a, 0x76, 0x33, 0xe3, 0xf3, 0x5d, 0x9f, 0x39, 0x67, 0x00, 0xcb, 0x4a,
	0x48, 0x31, 0x2f, 0x73, 0x7e, 0x43, 0x75, 0xac, 0x37, 0xe8, 0x99, 0x5d, 0x78, 0xb5, 0x13, 0x62,
	0x97, 0xd3, 0x5c, 0x9f, 0x6e, 0x9a, 0xed, 0x5c, 0x66, 0x7b, 0xaa, 0x25, 0xdf, 0x97, 0x46, 0x18,
	0x5e, 0x3e, 0x14, 0xfc, 0xa8, 0x78, 0x59, 0x52, 0x65, 0x07, 0x45, 0xbf, 0x1d, 0x18, 0xad, 0xc5,
	0x0d, 0x97, 0x99, 0x28, 0xf0, 0x1c, 0x9c, 0x2c, 0x0d, 0xd8, 0x94, 0xcd, 0xc6, 0x89, 0x93, 0xa5,
	0x88, 0x30, 0x28, 0xf8, 0x9e, 0x02, 0x47, 0x9f, 0xe8, 0x35, 0xae, 0x60, 0xb4, 0x27, 0xc9, 0x53,
	0x2e, 0x79, 0xe0, 0x4e, 0xdd, 0x99, 0xbf, 0xb8, 0x8c, 0xad, 0xb5, 0x76, 0x4e, 0xfc, 0xc1, 0x0a,
	0xde, 0x17, 0xb2, 0xba, 0x4b, 0x0e, 0x7a, 0x5c, 0xc2, 0xf8, 0xe0, 0x2f, 0x18, 0x4c, 0xd9, 0xcc,
	0x5f, 0x84, 0xb1, 0x31, 0x18, 0xb7, 0x06, 0xe3, 0x2f, 0xad, 0x22, 0xe9, 0xc4, 0xb8, 0x84, 0x51,
	0xce, 0x65, 0x26, 0x9b, 0x94, 0x82, 0xa1, 0x06, 0x5f, 0xfe, 0x03, 0xbe, 0x13, 0xcd, 0x26, 0xa7,
	0xaf, 0x3c, 0x6f, 0x28, 0x39, 0xa8, 0x71, 0x05, 0xe3, 0x5c, 0x14, 0x3b, 0x83, 0x7a, 0x4f, 0x40,
	0x3b, 0x79, 0xf8, 0x16, 0x9e, 0xdf, 0xbb, 0x0a, 0x4e, 0xc0, 0xfd, 0x46, 0x77, 0x36, 0x21, 0xb5,
	0xc4, 0x0b, 0x18, 0xde, 0x2a, 0xcc, 0x66, 0x64, 0x36, 0x2b, 0x67, 0xc9, 0xa2, 0x37, 0xe0, 0x7f,
	0xe6, 0xb7, 0x94, 0xd0, 0xf7, 0x86, 0x6a, 0x89, 0x33, 0xb0, 0x9d, 0x05, 0x4c, 0xa7, 0x36, 0x79,
	0x98, 0x5a, 0x62, 0xbf, 0x47, 0xe7, 0x70, 0x66, 0xc0, 0xba, 0x14, 0x45, 0x4d, 0xd1, 0x15, 0xf8,
	0x6b, 0x5e, 0xcb, 0x76, 0xd0, 0x04, 0xdc, 0x2c, 0x35, 0x53, 0xc6, 0x89, 0x5a, 0x46, 0x4b, 0x38,
	0x5b, 0x67, 0x4a, 0x60, 0x80, 0x13, 0x7e, 0xf5, 0x8b, 0x81, 0xff, 0x91, 0x78, 0xd5, 0xce, 0xee,
	0xc7, 0xcc, 0xfe, 0x3f, 0x66, 0xe7, 0xa4, 0x98, 0xf1, 0x35, 0x78, 0x15, 0x4f, 0xb3, 0xa6, 0x0e,
	0xdc, 0x27, 0x80, 0x56, 0x1b, 0xfd, 0x64, 0xe0, 0x27, 0xc4, 0xd3, 0x47, 0x73, 0xc1, 0x57, 0x30,
	0xe4, 0x5b, 0x49, 0x95, 0xf5, 0x73, 0xec, 0xa9, 0x19, 0x21, 0x2e, 0xc0, 0xdb, 0xd0, 0x56, 0x54,
	0x64, 0x9d, 0x1c, 0x43, 0xac, 0x72, 0xf1, 0x87, 0x81, 0xf7, 0x49, 0xc7, 0x89, 0xd7, 0x30, 0x50,
	0xcd, 0xe1, 0x8b, 0x36, 0xf0, 0xde, 0x03, 0x08, 0x2f, 0xee, 0x1f, 0xda, 0x72, 0x9f, 0x29, 0x48,
	0xd5, 0xdb, 0x41, 0xbd, 0xb2, 0x3b, 0xa8, 0x5f, 0xb0, 0x81, 0x54, 0x6f, 0x1d, 0xd4, 0x6b, 0xf1,
	0x18, 0xa4, 0x02, 0xeb, 0xa0, 0x5e, 0x7c, 0x8f, 0x41, 0x1b, 0x4f, 0x5f, 0xfd, 0xfa, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x7a, 0x9f, 0xaf, 0x48, 0x86, 0x04, 0x00, 0x00,
}