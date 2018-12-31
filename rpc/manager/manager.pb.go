// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/manager/manager.proto

package manager // import "github.com/R-a-dio/valkyrie/rpc/manager"

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

type StatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusRequest) Reset()         { *m = StatusRequest{} }
func (m *StatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()    {}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_manager_d516f118c655c20c, []int{0}
}
func (m *StatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusRequest.Unmarshal(m, b)
}
func (m *StatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusRequest.Marshal(b, m, deterministic)
}
func (dst *StatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequest.Merge(dst, src)
}
func (m *StatusRequest) XXX_Size() int {
	return xxx_messageInfo_StatusRequest.Size(m)
}
func (m *StatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequest proto.InternalMessageInfo

type StatusResponse struct {
	User                 *User         `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Song                 *Song         `protobuf:"bytes,2,opt,name=song" json:"song,omitempty"`
	ListenerInfo         *ListenerInfo `protobuf:"bytes,3,opt,name=listener_info,json=listenerInfo" json:"listener_info,omitempty"`
	Thread               *Thread       `protobuf:"bytes,4,opt,name=thread" json:"thread,omitempty"`
	BotConfig            *BotConfig    `protobuf:"bytes,5,opt,name=bot_config,json=botConfig" json:"bot_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_manager_d516f118c655c20c, []int{1}
}
func (m *StatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResponse.Unmarshal(m, b)
}
func (m *StatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResponse.Marshal(b, m, deterministic)
}
func (dst *StatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse.Merge(dst, src)
}
func (m *StatusResponse) XXX_Size() int {
	return xxx_messageInfo_StatusResponse.Size(m)
}
func (m *StatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse proto.InternalMessageInfo

func (m *StatusResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *StatusResponse) GetSong() *Song {
	if m != nil {
		return m.Song
	}
	return nil
}

func (m *StatusResponse) GetListenerInfo() *ListenerInfo {
	if m != nil {
		return m.ListenerInfo
	}
	return nil
}

func (m *StatusResponse) GetThread() *Thread {
	if m != nil {
		return m.Thread
	}
	return nil
}

func (m *StatusResponse) GetBotConfig() *BotConfig {
	if m != nil {
		return m.BotConfig
	}
	return nil
}

type Song struct {
	// song identifier
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// short metadata
	Metadata string `protobuf:"bytes,2,opt,name=metadata" json:"metadata,omitempty"`
	// song start time in unix epoch
	StartTime uint64 `protobuf:"varint,3,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	// song end time in unix epoch, can be inaccurate
	EndTime uint64 `protobuf:"varint,4,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	// database track identifier
	TrackId              int32    `protobuf:"varint,5,opt,name=track_id,json=trackId" json:"track_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Song) Reset()         { *m = Song{} }
func (m *Song) String() string { return proto.CompactTextString(m) }
func (*Song) ProtoMessage()    {}
func (*Song) Descriptor() ([]byte, []int) {
	return fileDescriptor_manager_d516f118c655c20c, []int{2}
}
func (m *Song) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Song.Unmarshal(m, b)
}
func (m *Song) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Song.Marshal(b, m, deterministic)
}
func (dst *Song) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Song.Merge(dst, src)
}
func (m *Song) XXX_Size() int {
	return xxx_messageInfo_Song.Size(m)
}
func (m *Song) XXX_DiscardUnknown() {
	xxx_messageInfo_Song.DiscardUnknown(m)
}

var xxx_messageInfo_Song proto.InternalMessageInfo

func (m *Song) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Song) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

func (m *Song) GetStartTime() uint64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Song) GetEndTime() uint64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *Song) GetTrackId() int32 {
	if m != nil {
		return m.TrackId
	}
	return 0
}

type Thread struct {
	// thread string, most of the time an URL
	Thread               string   `protobuf:"bytes,1,opt,name=thread" json:"thread,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Thread) Reset()         { *m = Thread{} }
func (m *Thread) String() string { return proto.CompactTextString(m) }
func (*Thread) ProtoMessage()    {}
func (*Thread) Descriptor() ([]byte, []int) {
	return fileDescriptor_manager_d516f118c655c20c, []int{3}
}
func (m *Thread) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Thread.Unmarshal(m, b)
}
func (m *Thread) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Thread.Marshal(b, m, deterministic)
}
func (dst *Thread) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Thread.Merge(dst, src)
}
func (m *Thread) XXX_Size() int {
	return xxx_messageInfo_Thread.Size(m)
}
func (m *Thread) XXX_DiscardUnknown() {
	xxx_messageInfo_Thread.DiscardUnknown(m)
}

var xxx_messageInfo_Thread proto.InternalMessageInfo

func (m *Thread) GetThread() string {
	if m != nil {
		return m.Thread
	}
	return ""
}

type BotConfig struct {
	RequestsEnabled      bool     `protobuf:"varint,1,opt,name=requests_enabled,json=requestsEnabled" json:"requests_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BotConfig) Reset()         { *m = BotConfig{} }
func (m *BotConfig) String() string { return proto.CompactTextString(m) }
func (*BotConfig) ProtoMessage()    {}
func (*BotConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_manager_d516f118c655c20c, []int{4}
}
func (m *BotConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BotConfig.Unmarshal(m, b)
}
func (m *BotConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BotConfig.Marshal(b, m, deterministic)
}
func (dst *BotConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BotConfig.Merge(dst, src)
}
func (m *BotConfig) XXX_Size() int {
	return xxx_messageInfo_BotConfig.Size(m)
}
func (m *BotConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_BotConfig.DiscardUnknown(m)
}

var xxx_messageInfo_BotConfig proto.InternalMessageInfo

func (m *BotConfig) GetRequestsEnabled() bool {
	if m != nil {
		return m.RequestsEnabled
	}
	return false
}

type User struct {
	// user identifier
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// user nickname, this is only a display-name
	Nickname string `protobuf:"bytes,2,opt,name=nickname" json:"nickname,omitempty"`
	// indicates if this user is a robot or not
	IsRobot              bool     `protobuf:"varint,3,opt,name=is_robot,json=isRobot" json:"is_robot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_manager_d516f118c655c20c, []int{5}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *User) GetIsRobot() bool {
	if m != nil {
		return m.IsRobot
	}
	return false
}

type ListenerInfo struct {
	Listeners            int64    `protobuf:"varint,1,opt,name=listeners" json:"listeners,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListenerInfo) Reset()         { *m = ListenerInfo{} }
func (m *ListenerInfo) String() string { return proto.CompactTextString(m) }
func (*ListenerInfo) ProtoMessage()    {}
func (*ListenerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_manager_d516f118c655c20c, []int{6}
}
func (m *ListenerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerInfo.Unmarshal(m, b)
}
func (m *ListenerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerInfo.Marshal(b, m, deterministic)
}
func (dst *ListenerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerInfo.Merge(dst, src)
}
func (m *ListenerInfo) XXX_Size() int {
	return xxx_messageInfo_ListenerInfo.Size(m)
}
func (m *ListenerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerInfo proto.InternalMessageInfo

func (m *ListenerInfo) GetListeners() int64 {
	if m != nil {
		return m.Listeners
	}
	return 0
}

func init() {
	proto.RegisterType((*StatusRequest)(nil), "radio.internal.manager.StatusRequest")
	proto.RegisterType((*StatusResponse)(nil), "radio.internal.manager.StatusResponse")
	proto.RegisterType((*Song)(nil), "radio.internal.manager.Song")
	proto.RegisterType((*Thread)(nil), "radio.internal.manager.Thread")
	proto.RegisterType((*BotConfig)(nil), "radio.internal.manager.BotConfig")
	proto.RegisterType((*User)(nil), "radio.internal.manager.User")
	proto.RegisterType((*ListenerInfo)(nil), "radio.internal.manager.ListenerInfo")
}

func init() { proto.RegisterFile("rpc/manager/manager.proto", fileDescriptor_manager_d516f118c655c20c) }

var fileDescriptor_manager_d516f118c655c20c = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdb, 0x6a, 0xdb, 0x40,
	0x10, 0xc5, 0xb6, 0x7c, 0xd1, 0xd4, 0x89, 0xcb, 0x3e, 0x04, 0x47, 0xa4, 0x25, 0x15, 0xbd, 0x05,
	0x1a, 0xb9, 0xa4, 0x90, 0xe7, 0x92, 0x92, 0x07, 0xd3, 0xe6, 0x65, 0xe5, 0x52, 0x28, 0x14, 0xb1,
	0xb2, 0xc6, 0xce, 0x62, 0x6b, 0xd7, 0xdd, 0x1d, 0x17, 0xfa, 0x05, 0xfd, 0xa6, 0xfe, 0x56, 0xbf,
	0xa0, 0x78, 0x25, 0x39, 0x0e, 0xc4, 0x76, 0x9e, 0xec, 0x99, 0x73, 0x66, 0xa4, 0xb3, 0xe7, 0xac,
	0xe0, 0xd8, 0x2c, 0xc6, 0x83, 0x5c, 0x28, 0x31, 0x45, 0x53, 0xfd, 0x46, 0x0b, 0xa3, 0x49, 0xb3,
	0x23, 0x23, 0x32, 0xa9, 0x23, 0xa9, 0x08, 0x8d, 0x12, 0xf3, 0xa8, 0x44, 0xc3, 0x1e, 0x1c, 0xc4,
	0x24, 0x68, 0x69, 0x39, 0xfe, 0x5c, 0xa2, 0xa5, 0xf0, 0x6f, 0x1d, 0x0e, 0xab, 0x8e, 0x5d, 0x68,
	0x65, 0x91, 0xbd, 0x07, 0x6f, 0x69, 0xd1, 0xf4, 0x6b, 0xa7, 0xb5, 0xb7, 0x4f, 0x2e, 0x4e, 0xa2,
	0x87, 0x57, 0x45, 0x5f, 0x2d, 0x1a, 0xee, 0x98, 0xab, 0x09, 0xab, 0xd5, 0xb4, 0x5f, 0xdf, 0x3d,
	0x11, 0x6b, 0x35, 0xe5, 0x8e, 0xc9, 0x86, 0x70, 0x30, 0x97, 0x96, 0x50, 0xa1, 0x49, 0xa4, 0x9a,
	0xe8, 0x7e, 0xc3, 0x8d, 0xbe, 0xdc, 0x36, 0xfa, 0xa5, 0x24, 0x0f, 0xd5, 0x44, 0xf3, 0xee, 0x7c,
	0xa3, 0x62, 0x97, 0xd0, 0xa2, 0x5b, 0x83, 0x22, 0xeb, 0x7b, 0x6e, 0xc7, 0xf3, 0x6d, 0x3b, 0x46,
	0x8e, 0xc5, 0x4b, 0x36, 0xfb, 0x08, 0x90, 0x6a, 0x4a, 0xc6, 0x5a, 0x4d, 0xe4, 0xb4, 0xdf, 0x74,
	0xb3, 0x2f, 0xb6, 0xcd, 0x5e, 0x69, 0xfa, 0xe4, 0x88, 0xdc, 0x4f, 0xab, 0xbf, 0xe1, 0x9f, 0x1a,
	0x78, 0x2b, 0x4d, 0xec, 0x10, 0xea, 0x32, 0x73, 0xe7, 0xd5, 0xe4, 0x75, 0x99, 0xb1, 0x00, 0x3a,
	0x39, 0x92, 0xc8, 0x04, 0x09, 0x77, 0x26, 0x3e, 0x5f, 0xd7, 0xec, 0x19, 0x80, 0x25, 0x61, 0x28,
	0x21, 0x99, 0xa3, 0x93, 0xed, 0x71, 0xdf, 0x75, 0x46, 0x32, 0x47, 0x76, 0x0c, 0x1d, 0x54, 0x59,
	0x01, 0x7a, 0x0e, 0x6c, 0xa3, 0xca, 0x2a, 0x88, 0x8c, 0x18, 0xcf, 0x12, 0x99, 0xb9, 0xd7, 0x6d,
	0xf2, 0xb6, 0xab, 0x87, 0x59, 0x78, 0x0a, 0xad, 0x42, 0x1d, 0x3b, 0x5a, 0x9f, 0x46, 0xcd, 0x3d,
	0xb8, 0xac, 0xc2, 0x4b, 0xf0, 0xd7, 0x1a, 0xd8, 0x19, 0x3c, 0x35, 0x85, 0xff, 0x36, 0x41, 0x25,
	0xd2, 0x39, 0x16, 0xf4, 0x0e, 0xef, 0x55, 0xfd, 0xeb, 0xa2, 0x1d, 0xde, 0x80, 0xb7, 0x32, 0xfa,
	0x21, 0x89, 0x4a, 0x8e, 0x67, 0x4a, 0xe4, 0x58, 0x49, 0xac, 0xea, 0xd5, 0x8b, 0x4a, 0x9b, 0x18,
	0x9d, 0x6a, 0x72, 0x02, 0x3b, 0xbc, 0x2d, 0x2d, 0x5f, 0x95, 0xe1, 0x3b, 0xe8, 0x6e, 0x5a, 0xc9,
	0x4e, 0xc0, 0xaf, 0xcc, 0xb4, 0x6e, 0x7b, 0x83, 0xdf, 0x35, 0x2e, 0xfe, 0x35, 0xa0, 0x7d, 0x53,
	0x38, 0xc0, 0xbe, 0x41, 0xab, 0xc8, 0x29, 0x7b, 0xb5, 0x35, 0x5f, 0x9b, 0xc9, 0x0e, 0x5e, 0xef,
	0xa3, 0x95, 0x71, 0xbf, 0x86, 0x76, 0x8c, 0xe4, 0x44, 0xee, 0xcc, 0x7a, 0xb0, 0x13, 0x2d, 0xd7,
	0xb8, 0x38, 0xec, 0xbc, 0x00, 0xc1, 0x4e, 0x94, 0x8d, 0xa0, 0x1b, 0x23, 0xdd, 0x59, 0xb5, 0x3f,
	0x91, 0xc1, 0x7e, 0x0a, 0xfb, 0x0c, 0x7e, 0x8c, 0x54, 0x46, 0x64, 0xcf, 0x05, 0x09, 0xf6, 0xe0,
	0xec, 0x07, 0xf4, 0x62, 0xa4, 0x7b, 0x36, 0x3e, 0xea, 0xde, 0x06, 0x8f, 0x62, 0x5d, 0x9d, 0x7d,
	0x7f, 0x33, 0x95, 0x74, 0xbb, 0x4c, 0xa3, 0xb1, 0xce, 0x07, 0xfc, 0x5c, 0x9c, 0x67, 0x52, 0x0f,
	0x7e, 0x89, 0xf9, 0xec, 0xb7, 0x91, 0x38, 0xd8, 0xf8, 0xe6, 0xa5, 0x2d, 0xf7, 0xb1, 0xfb, 0xf0,
	0x3f, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xc3, 0xb5, 0x2e, 0x09, 0x05, 0x00, 0x00,
}
