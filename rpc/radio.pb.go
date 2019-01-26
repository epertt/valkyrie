// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/radio.proto

package rpc // import "github.com/R-a-dio/valkyrie/rpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/empty"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Song struct {
	// song identifier (esong.id)
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// short metadata
	Metadata             string   `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Song) Reset()         { *m = Song{} }
func (m *Song) String() string { return proto.CompactTextString(m) }
func (*Song) ProtoMessage()    {}
func (*Song) Descriptor() ([]byte, []int) {
	return fileDescriptor_radio_69e3ec7604407055, []int{0}
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

type StatusResponse struct {
	// the current user that is streaming
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// information about the current song
	Info *StreamInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	// information about the current listeners
	ListenerInfo *ListenerInfo `protobuf:"bytes,3,opt,name=listener_info,json=listenerInfo,proto3" json:"listener_info,omitempty"`
	// the current thread to be shown on the website or elsewhere
	Thread string `protobuf:"bytes,4,opt,name=thread,proto3" json:"thread,omitempty"`
	// the current configuration of the streamer
	StreamerConfig       *StreamerConfig `protobuf:"bytes,5,opt,name=streamer_config,json=streamerConfig,proto3" json:"streamer_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_radio_69e3ec7604407055, []int{1}
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

func (m *StatusResponse) GetInfo() *StreamInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *StatusResponse) GetListenerInfo() *ListenerInfo {
	if m != nil {
		return m.ListenerInfo
	}
	return nil
}

func (m *StatusResponse) GetThread() string {
	if m != nil {
		return m.Thread
	}
	return ""
}

func (m *StatusResponse) GetStreamerConfig() *StreamerConfig {
	if m != nil {
		return m.StreamerConfig
	}
	return nil
}

type StreamInfo struct {
	// the current song being played
	Song *Song `protobuf:"bytes,1,opt,name=song,proto3" json:"song,omitempty"`
	// the time this song started playing
	StartTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// the time this song will end playing
	EndTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// last time this song was played
	LastPlayed           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=last_played,json=lastPlayed,proto3" json:"last_played,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StreamInfo) Reset()         { *m = StreamInfo{} }
func (m *StreamInfo) String() string { return proto.CompactTextString(m) }
func (*StreamInfo) ProtoMessage()    {}
func (*StreamInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_radio_69e3ec7604407055, []int{2}
}
func (m *StreamInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamInfo.Unmarshal(m, b)
}
func (m *StreamInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamInfo.Marshal(b, m, deterministic)
}
func (dst *StreamInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamInfo.Merge(dst, src)
}
func (m *StreamInfo) XXX_Size() int {
	return xxx_messageInfo_StreamInfo.Size(m)
}
func (m *StreamInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamInfo.DiscardUnknown(m)
}

var xxx_messageInfo_StreamInfo proto.InternalMessageInfo

func (m *StreamInfo) GetSong() *Song {
	if m != nil {
		return m.Song
	}
	return nil
}

func (m *StreamInfo) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *StreamInfo) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *StreamInfo) GetLastPlayed() *timestamp.Timestamp {
	if m != nil {
		return m.LastPlayed
	}
	return nil
}

type StreamerConfig struct {
	// can users request songs to be played right now
	RequestsEnabled bool `protobuf:"varint,1,opt,name=requests_enabled,json=requestsEnabled,proto3" json:"requests_enabled,omitempty"`
	// the queue implementation to use for the streamer
	QueueUsed            string   `protobuf:"bytes,2,opt,name=queue_used,json=queueUsed,proto3" json:"queue_used,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamerConfig) Reset()         { *m = StreamerConfig{} }
func (m *StreamerConfig) String() string { return proto.CompactTextString(m) }
func (*StreamerConfig) ProtoMessage()    {}
func (*StreamerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_radio_69e3ec7604407055, []int{3}
}
func (m *StreamerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamerConfig.Unmarshal(m, b)
}
func (m *StreamerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamerConfig.Marshal(b, m, deterministic)
}
func (dst *StreamerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamerConfig.Merge(dst, src)
}
func (m *StreamerConfig) XXX_Size() int {
	return xxx_messageInfo_StreamerConfig.Size(m)
}
func (m *StreamerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_StreamerConfig proto.InternalMessageInfo

func (m *StreamerConfig) GetRequestsEnabled() bool {
	if m != nil {
		return m.RequestsEnabled
	}
	return false
}

func (m *StreamerConfig) GetQueueUsed() string {
	if m != nil {
		return m.QueueUsed
	}
	return ""
}

type User struct {
	// user identifier
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// user nickname, this is only a display-name
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	// indicates if this user is a robot or not
	IsRobot              bool     `protobuf:"varint,3,opt,name=is_robot,json=isRobot,proto3" json:"is_robot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_radio_69e3ec7604407055, []int{4}
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
	// the amount of listeners to the stream
	Listeners            int64    `protobuf:"varint,1,opt,name=listeners,proto3" json:"listeners,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListenerInfo) Reset()         { *m = ListenerInfo{} }
func (m *ListenerInfo) String() string { return proto.CompactTextString(m) }
func (*ListenerInfo) ProtoMessage()    {}
func (*ListenerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_radio_69e3ec7604407055, []int{5}
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

type SongAnnouncement struct {
	Info                 *StreamInfo   `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Listeners            *ListenerInfo `protobuf:"bytes,2,opt,name=listeners,proto3" json:"listeners,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SongAnnouncement) Reset()         { *m = SongAnnouncement{} }
func (m *SongAnnouncement) String() string { return proto.CompactTextString(m) }
func (*SongAnnouncement) ProtoMessage()    {}
func (*SongAnnouncement) Descriptor() ([]byte, []int) {
	return fileDescriptor_radio_69e3ec7604407055, []int{6}
}
func (m *SongAnnouncement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SongAnnouncement.Unmarshal(m, b)
}
func (m *SongAnnouncement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SongAnnouncement.Marshal(b, m, deterministic)
}
func (dst *SongAnnouncement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SongAnnouncement.Merge(dst, src)
}
func (m *SongAnnouncement) XXX_Size() int {
	return xxx_messageInfo_SongAnnouncement.Size(m)
}
func (m *SongAnnouncement) XXX_DiscardUnknown() {
	xxx_messageInfo_SongAnnouncement.DiscardUnknown(m)
}

var xxx_messageInfo_SongAnnouncement proto.InternalMessageInfo

func (m *SongAnnouncement) GetInfo() *StreamInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *SongAnnouncement) GetListeners() *ListenerInfo {
	if m != nil {
		return m.Listeners
	}
	return nil
}

type SongRequestAnnouncement struct {
	Song                 *Song    `protobuf:"bytes,1,opt,name=song,proto3" json:"song,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SongRequestAnnouncement) Reset()         { *m = SongRequestAnnouncement{} }
func (m *SongRequestAnnouncement) String() string { return proto.CompactTextString(m) }
func (*SongRequestAnnouncement) ProtoMessage()    {}
func (*SongRequestAnnouncement) Descriptor() ([]byte, []int) {
	return fileDescriptor_radio_69e3ec7604407055, []int{7}
}
func (m *SongRequestAnnouncement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SongRequestAnnouncement.Unmarshal(m, b)
}
func (m *SongRequestAnnouncement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SongRequestAnnouncement.Marshal(b, m, deterministic)
}
func (dst *SongRequestAnnouncement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SongRequestAnnouncement.Merge(dst, src)
}
func (m *SongRequestAnnouncement) XXX_Size() int {
	return xxx_messageInfo_SongRequestAnnouncement.Size(m)
}
func (m *SongRequestAnnouncement) XXX_DiscardUnknown() {
	xxx_messageInfo_SongRequestAnnouncement.DiscardUnknown(m)
}

var xxx_messageInfo_SongRequestAnnouncement proto.InternalMessageInfo

func (m *SongRequestAnnouncement) GetSong() *Song {
	if m != nil {
		return m.Song
	}
	return nil
}

type QueueEntry struct {
	Song *Song `protobuf:"bytes,1,opt,name=song,proto3" json:"song,omitempty"`
	// is_user_request indicates if this was a request made by a human
	IsUserRequest bool `protobuf:"varint,2,opt,name=is_user_request,json=isUserRequest,proto3" json:"is_user_request,omitempty"`
	// user_identifier is the way we identify the user that added this to the
	// queue; This can be anything that uniquely identifies a user
	UserIdentifier string `protobuf:"bytes,3,opt,name=user_identifier,json=userIdentifier,proto3" json:"user_identifier,omitempty"`
	// expected_start_time is the expected time this song will start playing
	ExpectedStartTime    *timestamp.Timestamp `protobuf:"bytes,4,opt,name=expected_start_time,json=expectedStartTime,proto3" json:"expected_start_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *QueueEntry) Reset()         { *m = QueueEntry{} }
func (m *QueueEntry) String() string { return proto.CompactTextString(m) }
func (*QueueEntry) ProtoMessage()    {}
func (*QueueEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_radio_69e3ec7604407055, []int{8}
}
func (m *QueueEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueEntry.Unmarshal(m, b)
}
func (m *QueueEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueEntry.Marshal(b, m, deterministic)
}
func (dst *QueueEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueEntry.Merge(dst, src)
}
func (m *QueueEntry) XXX_Size() int {
	return xxx_messageInfo_QueueEntry.Size(m)
}
func (m *QueueEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueEntry.DiscardUnknown(m)
}

var xxx_messageInfo_QueueEntry proto.InternalMessageInfo

func (m *QueueEntry) GetSong() *Song {
	if m != nil {
		return m.Song
	}
	return nil
}

func (m *QueueEntry) GetIsUserRequest() bool {
	if m != nil {
		return m.IsUserRequest
	}
	return false
}

func (m *QueueEntry) GetUserIdentifier() string {
	if m != nil {
		return m.UserIdentifier
	}
	return ""
}

func (m *QueueEntry) GetExpectedStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.ExpectedStartTime
	}
	return nil
}

type QueueContent struct {
	// the name of the queue implementation
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// the entries in the queue
	Entries              []*QueueEntry `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *QueueContent) Reset()         { *m = QueueContent{} }
func (m *QueueContent) String() string { return proto.CompactTextString(m) }
func (*QueueContent) ProtoMessage()    {}
func (*QueueContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_radio_69e3ec7604407055, []int{9}
}
func (m *QueueContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueContent.Unmarshal(m, b)
}
func (m *QueueContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueContent.Marshal(b, m, deterministic)
}
func (dst *QueueContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueContent.Merge(dst, src)
}
func (m *QueueContent) XXX_Size() int {
	return xxx_messageInfo_QueueContent.Size(m)
}
func (m *QueueContent) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueContent.DiscardUnknown(m)
}

var xxx_messageInfo_QueueContent proto.InternalMessageInfo

func (m *QueueContent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *QueueContent) GetEntries() []*QueueEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type TrackRequest struct {
	UserIdentifier       string   `protobuf:"bytes,1,opt,name=user_identifier,json=userIdentifier,proto3" json:"user_identifier,omitempty"`
	TrackId              int64    `protobuf:"varint,2,opt,name=track_id,json=trackId,proto3" json:"track_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrackRequest) Reset()         { *m = TrackRequest{} }
func (m *TrackRequest) String() string { return proto.CompactTextString(m) }
func (*TrackRequest) ProtoMessage()    {}
func (*TrackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_radio_69e3ec7604407055, []int{10}
}
func (m *TrackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackRequest.Unmarshal(m, b)
}
func (m *TrackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackRequest.Marshal(b, m, deterministic)
}
func (dst *TrackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackRequest.Merge(dst, src)
}
func (m *TrackRequest) XXX_Size() int {
	return xxx_messageInfo_TrackRequest.Size(m)
}
func (m *TrackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TrackRequest proto.InternalMessageInfo

func (m *TrackRequest) GetUserIdentifier() string {
	if m != nil {
		return m.UserIdentifier
	}
	return ""
}

func (m *TrackRequest) GetTrackId() int64 {
	if m != nil {
		return m.TrackId
	}
	return 0
}

type RequestResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestResponse) Reset()         { *m = RequestResponse{} }
func (m *RequestResponse) String() string { return proto.CompactTextString(m) }
func (*RequestResponse) ProtoMessage()    {}
func (*RequestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_radio_69e3ec7604407055, []int{11}
}
func (m *RequestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestResponse.Unmarshal(m, b)
}
func (m *RequestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestResponse.Marshal(b, m, deterministic)
}
func (dst *RequestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestResponse.Merge(dst, src)
}
func (m *RequestResponse) XXX_Size() int {
	return xxx_messageInfo_RequestResponse.Size(m)
}
func (m *RequestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RequestResponse proto.InternalMessageInfo

func (m *RequestResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *RequestResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Song)(nil), "radio.Song")
	proto.RegisterType((*StatusResponse)(nil), "radio.StatusResponse")
	proto.RegisterType((*StreamInfo)(nil), "radio.StreamInfo")
	proto.RegisterType((*StreamerConfig)(nil), "radio.StreamerConfig")
	proto.RegisterType((*User)(nil), "radio.User")
	proto.RegisterType((*ListenerInfo)(nil), "radio.ListenerInfo")
	proto.RegisterType((*SongAnnouncement)(nil), "radio.SongAnnouncement")
	proto.RegisterType((*SongRequestAnnouncement)(nil), "radio.SongRequestAnnouncement")
	proto.RegisterType((*QueueEntry)(nil), "radio.QueueEntry")
	proto.RegisterType((*QueueContent)(nil), "radio.QueueContent")
	proto.RegisterType((*TrackRequest)(nil), "radio.TrackRequest")
	proto.RegisterType((*RequestResponse)(nil), "radio.RequestResponse")
}

func init() { proto.RegisterFile("rpc/radio.proto", fileDescriptor_radio_69e3ec7604407055) }

var fileDescriptor_radio_69e3ec7604407055 = []byte{
	// 977 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdd, 0x6e, 0xdc, 0x44,
	0x14, 0x96, 0x37, 0x9b, 0xee, 0xee, 0xc9, 0x26, 0x9b, 0x4c, 0x44, 0x9a, 0x9a, 0xd2, 0x14, 0x23,
	0xa0, 0x08, 0xba, 0xab, 0x2e, 0x94, 0x96, 0x44, 0x80, 0xd2, 0x2a, 0x42, 0x41, 0x94, 0x9f, 0x71,
	0xca, 0x45, 0x6f, 0xac, 0x89, 0x7d, 0xb2, 0x1d, 0xc5, 0x9e, 0x71, 0x67, 0xc6, 0xa1, 0x79, 0x12,
	0xde, 0x82, 0x07, 0xe1, 0x19, 0xb8, 0x45, 0x48, 0x3c, 0x05, 0xf2, 0xd8, 0xde, 0x78, 0xb3, 0xeb,
	0x8d, 0xb8, 0xf3, 0xf9, 0x9b, 0x99, 0xf3, 0x9d, 0xef, 0x3b, 0x86, 0x81, 0x4a, 0xc3, 0x91, 0x62,
	0x11, 0x97, 0xc3, 0x54, 0x49, 0x23, 0xc9, 0xaa, 0x35, 0xdc, 0x77, 0x27, 0x52, 0x4e, 0x62, 0x1c,
	0x59, 0xe7, 0x69, 0x76, 0x36, 0xc2, 0x24, 0x35, 0x97, 0x45, 0x8e, 0x7b, 0xef, 0x7a, 0xf0, 0x37,
	0xc5, 0xd2, 0x14, 0x95, 0x2e, 0xe3, 0x7b, 0xd7, 0xe3, 0x86, 0x27, 0xa8, 0x0d, 0x4b, 0xd2, 0x22,
	0xc1, 0x1b, 0x43, 0xdb, 0x97, 0x62, 0x42, 0x36, 0xa0, 0xc5, 0xa3, 0x5d, 0xe7, 0xbe, 0xf3, 0x60,
	0x95, 0xb6, 0x78, 0x44, 0x5c, 0xe8, 0x26, 0x68, 0x58, 0xc4, 0x0c, 0xdb, 0x6d, 0xdd, 0x77, 0x1e,
	0xf4, 0xe8, 0xd4, 0xf6, 0xfe, 0x75, 0x60, 0xc3, 0x37, 0xcc, 0x64, 0x9a, 0xa2, 0x4e, 0xa5, 0xd0,
	0x48, 0xf6, 0xa0, 0x9d, 0x69, 0x54, 0xf6, 0x80, 0xb5, 0xf1, 0xda, 0xb0, 0xe8, 0xe3, 0xa5, 0x46,
	0x45, 0x6d, 0x80, 0x7c, 0x08, 0x6d, 0x2e, 0xce, 0xa4, 0x3d, 0x6b, 0x6d, 0xbc, 0x55, 0x26, 0xf8,
	0x46, 0x21, 0x4b, 0x8e, 0xc5, 0x99, 0xa4, 0x36, 0x4c, 0x9e, 0xc2, 0x7a, 0xcc, 0xb5, 0x41, 0x81,
	0x2a, 0xb0, 0xf9, 0x2b, 0x36, 0x7f, 0xbb, 0xcc, 0xff, 0xa1, 0x8c, 0xd9, 0x8a, 0x7e, 0x5c, 0xb3,
	0xc8, 0x0e, 0xdc, 0x32, 0xaf, 0x15, 0xb2, 0x68, 0xb7, 0x6d, 0x9f, 0x5b, 0x5a, 0xe4, 0x1b, 0x18,
	0x68, 0x7b, 0x0b, 0xaa, 0x20, 0x94, 0xe2, 0x8c, 0x4f, 0x76, 0x57, 0xed, 0x99, 0xef, 0xcc, 0xbc,
	0x01, 0xd5, 0x73, 0x1b, 0xa4, 0x1b, 0x7a, 0xc6, 0xf6, 0xfe, 0x72, 0x00, 0xae, 0x9e, 0x99, 0x37,
	0xaa, 0xa5, 0x98, 0x5c, 0x6b, 0x34, 0x87, 0x90, 0xda, 0x00, 0xf9, 0x0a, 0x40, 0x1b, 0xa6, 0x4c,
	0x90, 0x23, 0x5d, 0x3e, 0xdf, 0x1d, 0x16, 0x63, 0x18, 0x56, 0x63, 0x18, 0x9e, 0x54, 0x63, 0xa0,
	0x3d, 0x9b, 0x9d, 0xdb, 0xe4, 0x31, 0x74, 0x51, 0x44, 0x45, 0x61, 0xfb, 0xc6, 0xc2, 0x0e, 0x8a,
	0xc8, 0x96, 0x1d, 0xc0, 0x5a, 0xcc, 0xb4, 0x09, 0xd2, 0x98, 0x5d, 0x62, 0x54, 0x76, 0xb7, 0xac,
	0x12, 0xf2, 0xf4, 0x9f, 0x6d, 0xb6, 0xf7, 0x2a, 0x1f, 0x65, 0xbd, 0x61, 0xf2, 0x09, 0x6c, 0x2a,
	0x7c, 0x93, 0xa1, 0x36, 0x3a, 0x40, 0xc1, 0x4e, 0x63, 0x2c, 0x78, 0xd1, 0xa5, 0x83, 0xca, 0x7f,
	0x54, 0xb8, 0xc9, 0x7b, 0x00, 0x6f, 0x32, 0xcc, 0x30, 0xc8, 0x34, 0x46, 0x25, 0x4d, 0x7a, 0xd6,
	0xf3, 0x52, 0x63, 0xe4, 0xbd, 0x80, 0x76, 0xce, 0x80, 0x45, 0xdc, 0x12, 0x3c, 0x3c, 0x17, 0x2c,
	0xc1, 0x8a, 0x5b, 0x95, 0x4d, 0xee, 0x40, 0x97, 0xeb, 0x40, 0xc9, 0x53, 0x69, 0x2c, 0x78, 0x5d,
	0xda, 0xe1, 0x9a, 0xe6, 0xa6, 0xf7, 0x19, 0xf4, 0xeb, 0xf3, 0x27, 0x77, 0xa1, 0x57, 0x31, 0x40,
	0xdb, 0xd3, 0x57, 0xe8, 0x95, 0xc3, 0x8b, 0x61, 0x33, 0x9f, 0xca, 0xa1, 0x10, 0x32, 0x13, 0x21,
	0x26, 0x28, 0xcc, 0x94, 0x84, 0xce, 0x72, 0x12, 0x3e, 0xaa, 0x1f, 0xdc, 0x6a, 0x26, 0x60, 0xed,
	0xb6, 0x7d, 0xb8, 0x6d, 0x39, 0x50, 0x00, 0x34, 0x73, 0xe9, 0x4d, 0x8c, 0xf1, 0xfe, 0x74, 0x00,
	0x7e, 0xc9, 0x41, 0x3b, 0x12, 0x46, 0x5d, 0xde, 0xcc, 0xb0, 0x8f, 0x60, 0xc0, 0x75, 0x0e, 0xb9,
	0x0a, 0xca, 0x81, 0xd8, 0x47, 0x76, 0xe9, 0x3a, 0xd7, 0x56, 0x71, 0x85, 0x93, 0x7c, 0x0c, 0x03,
	0x9b, 0xc4, 0x23, 0x14, 0x86, 0x9f, 0x71, 0x54, 0x16, 0xd1, 0x1e, 0xdd, 0xc8, 0xdd, 0xc7, 0x53,
	0x2f, 0xf9, 0x1e, 0xb6, 0xf1, 0x6d, 0x8a, 0xa1, 0xc1, 0x28, 0xa8, 0x71, 0xf7, 0x66, 0x0a, 0x6e,
	0x55, 0x65, 0x7e, 0xc5, 0x61, 0xef, 0x27, 0xe8, 0xdb, 0x5e, 0x9e, 0x4b, 0x61, 0xf2, 0xee, 0x09,
	0xb4, 0xed, 0x9c, 0x1d, 0x7b, 0xb3, 0xfd, 0x26, 0x9f, 0x42, 0x07, 0x85, 0x51, 0x1c, 0x73, 0x74,
	0x57, 0x6a, 0x93, 0xb8, 0x42, 0x81, 0x56, 0x19, 0x1e, 0x85, 0xfe, 0x89, 0x62, 0xe1, 0xf9, 0x92,
	0xae, 0x9c, 0x85, 0x5d, 0xdd, 0x81, 0xae, 0xc9, 0x0b, 0x03, 0x5e, 0x50, 0x73, 0x85, 0x76, 0xac,
	0x7d, 0x1c, 0x79, 0x5f, 0xc3, 0xa0, 0x3c, 0x6e, 0xba, 0xc0, 0x76, 0xa1, 0xa3, 0xb3, 0x30, 0x44,
	0xad, 0x4b, 0xb2, 0x57, 0x26, 0xd9, 0x84, 0x95, 0x44, 0x4f, 0x4a, 0xa2, 0xe6, 0x9f, 0xe3, 0x7f,
	0x5a, 0xd0, 0x79, 0xc1, 0x04, 0x9b, 0xa0, 0x22, 0x4f, 0xe0, 0x56, 0xb1, 0x0a, 0xc9, 0xce, 0x1c,
	0x50, 0x47, 0xf9, 0xa2, 0x76, 0xaf, 0xf6, 0xcc, 0xcc, 0xc6, 0xfc, 0x00, 0x3a, 0x3e, 0x1a, 0xab,
	0x8f, 0xfa, 0xba, 0x74, 0xeb, 0x06, 0x79, 0x0c, 0xeb, 0x3e, 0x9a, 0xda, 0xfa, 0x99, 0xe7, 0xac,
	0x3b, 0xef, 0x22, 0x87, 0xb0, 0x35, 0x2d, 0x9b, 0xea, 0x7a, 0xf1, 0xbe, 0x73, 0x17, 0xbb, 0xc9,
	0x77, 0xd0, 0xf3, 0xd1, 0x9c, 0x14, 0x3b, 0xf4, 0xee, 0x5c, 0x6b, 0xbe, 0x51, 0x5c, 0x4c, 0x7e,
	0x65, 0x71, 0x86, 0xee, 0xd2, 0x28, 0x39, 0x80, 0x81, 0x8f, 0x66, 0x46, 0xb8, 0x8b, 0xc4, 0xe4,
	0x2e, 0x72, 0x8e, 0x7f, 0x77, 0xa0, 0x57, 0x89, 0x49, 0x91, 0x6f, 0xa1, 0x5f, 0x19, 0xf6, 0x9f,
	0x75, 0xbb, 0xa6, 0x8d, 0xba, 0xe4, 0xdc, 0x86, 0x51, 0x90, 0x63, 0x18, 0x54, 0x79, 0x15, 0x9d,
	0xee, 0xd5, 0xf5, 0x35, 0xaf, 0xde, 0xa6, 0xa3, 0xc6, 0x7f, 0xb4, 0xa0, 0x5b, 0x41, 0x46, 0x9e,
	0xc0, 0xaa, 0x55, 0x40, 0x23, 0x07, 0x9a, 0x1e, 0xb4, 0x0f, 0x6d, 0xdf, 0xc8, 0x94, 0xcc, 0x8b,
	0xec, 0x99, 0x94, 0x71, 0x01, 0x6f, 0x53, 0xed, 0x01, 0xf4, 0xcb, 0x07, 0x5b, 0x7d, 0x4c, 0x51,
	0xad, 0xab, 0xc5, 0xdd, 0x29, 0x9d, 0xd7, 0xe9, 0xbe, 0x6f, 0xc7, 0xbb, 0x9c, 0x19, 0x4d, 0x17,
	0x7f, 0x01, 0xab, 0x56, 0xa8, 0x8d, 0xdd, 0x6e, 0xd7, 0xe5, 0x5c, 0x2e, 0x82, 0xf1, 0xdf, 0x4e,
	0x55, 0xf6, 0x08, 0xe0, 0x30, 0x8a, 0xaa, 0x01, 0xd4, 0x17, 0x5c, 0xe3, 0x95, 0x4f, 0x61, 0x8d,
	0xa2, 0x46, 0x75, 0x81, 0x3f, 0xe2, 0xdb, 0x66, 0x98, 0xe7, 0xf7, 0x48, 0xae, 0x4f, 0x8a, 0x89,
	0xbc, 0x40, 0x32, 0x1f, 0x74, 0x97, 0xc0, 0x4e, 0xbe, 0x84, 0xce, 0x51, 0xb1, 0x82, 0xfe, 0x57,
	0x9f, 0xcf, 0xde, 0x7f, 0xb5, 0x37, 0xe1, 0xe6, 0x75, 0x76, 0x3a, 0x0c, 0x65, 0x32, 0xa2, 0x0f,
	0xd9, 0xc3, 0x88, 0xcb, 0xd1, 0x05, 0x8b, 0xcf, 0x2f, 0x15, 0xc7, 0x91, 0x4a, 0xc3, 0xd3, 0x5b,
	0xf6, 0x9c, 0xcf, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x3f, 0xa7, 0xeb, 0x0b, 0xf2, 0x09, 0x00,
	0x00,
}
