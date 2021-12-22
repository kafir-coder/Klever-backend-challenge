// Code generated by protoc-gen-go. DO NOT EDIT.
// source: token.proto

package token

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AddTokenRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price                float64  `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddTokenRequest) Reset()         { *m = AddTokenRequest{} }
func (m *AddTokenRequest) String() string { return proto.CompactTextString(m) }
func (*AddTokenRequest) ProtoMessage()    {}
func (*AddTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aff0bcd502840ab, []int{0}
}

func (m *AddTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddTokenRequest.Unmarshal(m, b)
}
func (m *AddTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddTokenRequest.Marshal(b, m, deterministic)
}
func (m *AddTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddTokenRequest.Merge(m, src)
}
func (m *AddTokenRequest) XXX_Size() int {
	return xxx_messageInfo_AddTokenRequest.Size(m)
}
func (m *AddTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddTokenRequest proto.InternalMessageInfo

func (m *AddTokenRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddTokenRequest) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

type NumberOfVotes struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	NVotes               uint32   `protobuf:"varint,2,opt,name=n_votes,json=nVotes,proto3" json:"n_votes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NumberOfVotes) Reset()         { *m = NumberOfVotes{} }
func (m *NumberOfVotes) String() string { return proto.CompactTextString(m) }
func (*NumberOfVotes) ProtoMessage()    {}
func (*NumberOfVotes) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aff0bcd502840ab, []int{1}
}

func (m *NumberOfVotes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NumberOfVotes.Unmarshal(m, b)
}
func (m *NumberOfVotes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NumberOfVotes.Marshal(b, m, deterministic)
}
func (m *NumberOfVotes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumberOfVotes.Merge(m, src)
}
func (m *NumberOfVotes) XXX_Size() int {
	return xxx_messageInfo_NumberOfVotes.Size(m)
}
func (m *NumberOfVotes) XXX_DiscardUnknown() {
	xxx_messageInfo_NumberOfVotes.DiscardUnknown(m)
}

var xxx_messageInfo_NumberOfVotes proto.InternalMessageInfo

func (m *NumberOfVotes) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *NumberOfVotes) GetNVotes() uint32 {
	if m != nil {
		return m.NVotes
	}
	return 0
}

type GetTokenRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokenRequest) Reset()         { *m = GetTokenRequest{} }
func (m *GetTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GetTokenRequest) ProtoMessage()    {}
func (*GetTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aff0bcd502840ab, []int{2}
}

func (m *GetTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenRequest.Unmarshal(m, b)
}
func (m *GetTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenRequest.Marshal(b, m, deterministic)
}
func (m *GetTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenRequest.Merge(m, src)
}
func (m *GetTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GetTokenRequest.Size(m)
}
func (m *GetTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenRequest proto.InternalMessageInfo

func (m *GetTokenRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Token struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price                float64  `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aff0bcd502840ab, []int{3}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Token) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Token) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

type ListTokenRequest struct {
	Limit                uint32   `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page                 uint32   `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTokenRequest) Reset()         { *m = ListTokenRequest{} }
func (m *ListTokenRequest) String() string { return proto.CompactTextString(m) }
func (*ListTokenRequest) ProtoMessage()    {}
func (*ListTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aff0bcd502840ab, []int{4}
}

func (m *ListTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTokenRequest.Unmarshal(m, b)
}
func (m *ListTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTokenRequest.Marshal(b, m, deterministic)
}
func (m *ListTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTokenRequest.Merge(m, src)
}
func (m *ListTokenRequest) XXX_Size() int {
	return xxx_messageInfo_ListTokenRequest.Size(m)
}
func (m *ListTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTokenRequest proto.InternalMessageInfo

func (m *ListTokenRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListTokenRequest) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type DeleteTokenRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTokenRequest) Reset()         { *m = DeleteTokenRequest{} }
func (m *DeleteTokenRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTokenRequest) ProtoMessage()    {}
func (*DeleteTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aff0bcd502840ab, []int{5}
}

func (m *DeleteTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTokenRequest.Unmarshal(m, b)
}
func (m *DeleteTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTokenRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTokenRequest.Merge(m, src)
}
func (m *DeleteTokenRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTokenRequest.Size(m)
}
func (m *DeleteTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTokenRequest proto.InternalMessageInfo

func (m *DeleteTokenRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UpdateTokenRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price                float64  `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTokenRequest) Reset()         { *m = UpdateTokenRequest{} }
func (m *UpdateTokenRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTokenRequest) ProtoMessage()    {}
func (*UpdateTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aff0bcd502840ab, []int{6}
}

func (m *UpdateTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTokenRequest.Unmarshal(m, b)
}
func (m *UpdateTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTokenRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTokenRequest.Merge(m, src)
}
func (m *UpdateTokenRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTokenRequest.Size(m)
}
func (m *UpdateTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTokenRequest proto.InternalMessageInfo

func (m *UpdateTokenRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateTokenRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateTokenRequest) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

type Votes struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Votes) Reset()         { *m = Votes{} }
func (m *Votes) String() string { return proto.CompactTextString(m) }
func (*Votes) ProtoMessage()    {}
func (*Votes) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aff0bcd502840ab, []int{7}
}

func (m *Votes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Votes.Unmarshal(m, b)
}
func (m *Votes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Votes.Marshal(b, m, deterministic)
}
func (m *Votes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Votes.Merge(m, src)
}
func (m *Votes) XXX_Size() int {
	return xxx_messageInfo_Votes.Size(m)
}
func (m *Votes) XXX_DiscardUnknown() {
	xxx_messageInfo_Votes.DiscardUnknown(m)
}

var xxx_messageInfo_Votes proto.InternalMessageInfo

func (m *Votes) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Votes) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

type TokenList struct {
	Tokens               []*Token `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenList) Reset()         { *m = TokenList{} }
func (m *TokenList) String() string { return proto.CompactTextString(m) }
func (*TokenList) ProtoMessage()    {}
func (*TokenList) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aff0bcd502840ab, []int{8}
}

func (m *TokenList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenList.Unmarshal(m, b)
}
func (m *TokenList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenList.Marshal(b, m, deterministic)
}
func (m *TokenList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenList.Merge(m, src)
}
func (m *TokenList) XXX_Size() int {
	return xxx_messageInfo_TokenList.Size(m)
}
func (m *TokenList) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenList.DiscardUnknown(m)
}

var xxx_messageInfo_TokenList proto.InternalMessageInfo

func (m *TokenList) GetTokens() []*Token {
	if m != nil {
		return m.Tokens
	}
	return nil
}

type UpvoteVoteRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpvoteVoteRequest) Reset()         { *m = UpvoteVoteRequest{} }
func (m *UpvoteVoteRequest) String() string { return proto.CompactTextString(m) }
func (*UpvoteVoteRequest) ProtoMessage()    {}
func (*UpvoteVoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aff0bcd502840ab, []int{9}
}

func (m *UpvoteVoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpvoteVoteRequest.Unmarshal(m, b)
}
func (m *UpvoteVoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpvoteVoteRequest.Marshal(b, m, deterministic)
}
func (m *UpvoteVoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpvoteVoteRequest.Merge(m, src)
}
func (m *UpvoteVoteRequest) XXX_Size() int {
	return xxx_messageInfo_UpvoteVoteRequest.Size(m)
}
func (m *UpvoteVoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpvoteVoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpvoteVoteRequest proto.InternalMessageInfo

func (m *UpvoteVoteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DownVoteRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownVoteRequest) Reset()         { *m = DownVoteRequest{} }
func (m *DownVoteRequest) String() string { return proto.CompactTextString(m) }
func (*DownVoteRequest) ProtoMessage()    {}
func (*DownVoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aff0bcd502840ab, []int{10}
}

func (m *DownVoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownVoteRequest.Unmarshal(m, b)
}
func (m *DownVoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownVoteRequest.Marshal(b, m, deterministic)
}
func (m *DownVoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownVoteRequest.Merge(m, src)
}
func (m *DownVoteRequest) XXX_Size() int {
	return xxx_messageInfo_DownVoteRequest.Size(m)
}
func (m *DownVoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DownVoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DownVoteRequest proto.InternalMessageInfo

func (m *DownVoteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type WatchTokenRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WatchTokenRequest) Reset()         { *m = WatchTokenRequest{} }
func (m *WatchTokenRequest) String() string { return proto.CompactTextString(m) }
func (*WatchTokenRequest) ProtoMessage()    {}
func (*WatchTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aff0bcd502840ab, []int{11}
}

func (m *WatchTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchTokenRequest.Unmarshal(m, b)
}
func (m *WatchTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchTokenRequest.Marshal(b, m, deterministic)
}
func (m *WatchTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchTokenRequest.Merge(m, src)
}
func (m *WatchTokenRequest) XXX_Size() int {
	return xxx_messageInfo_WatchTokenRequest.Size(m)
}
func (m *WatchTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WatchTokenRequest proto.InternalMessageInfo

func (m *WatchTokenRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Ok struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	NoData               bool     `protobuf:"varint,2,opt,name=noData,proto3" json:"noData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ok) Reset()         { *m = Ok{} }
func (m *Ok) String() string { return proto.CompactTextString(m) }
func (*Ok) ProtoMessage()    {}
func (*Ok) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aff0bcd502840ab, []int{12}
}

func (m *Ok) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ok.Unmarshal(m, b)
}
func (m *Ok) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ok.Marshal(b, m, deterministic)
}
func (m *Ok) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ok.Merge(m, src)
}
func (m *Ok) XXX_Size() int {
	return xxx_messageInfo_Ok.Size(m)
}
func (m *Ok) XXX_DiscardUnknown() {
	xxx_messageInfo_Ok.DiscardUnknown(m)
}

var xxx_messageInfo_Ok proto.InternalMessageInfo

func (m *Ok) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *Ok) GetNoData() bool {
	if m != nil {
		return m.NoData
	}
	return false
}

func init() {
	proto.RegisterType((*AddTokenRequest)(nil), "token.add_token_request")
	proto.RegisterType((*NumberOfVotes)(nil), "token.number_of_votes")
	proto.RegisterType((*GetTokenRequest)(nil), "token.get_token_request")
	proto.RegisterType((*Token)(nil), "token.token")
	proto.RegisterType((*ListTokenRequest)(nil), "token.list_token_request")
	proto.RegisterType((*DeleteTokenRequest)(nil), "token.delete_token_request")
	proto.RegisterType((*UpdateTokenRequest)(nil), "token.update_token_request")
	proto.RegisterType((*Votes)(nil), "token.votes")
	proto.RegisterType((*TokenList)(nil), "token.token_list")
	proto.RegisterType((*UpvoteVoteRequest)(nil), "token.upvote_vote_request")
	proto.RegisterType((*DownVoteRequest)(nil), "token.down_vote_request")
	proto.RegisterType((*WatchTokenRequest)(nil), "token.watch_token_request")
	proto.RegisterType((*Ok)(nil), "token.ok")
}

func init() { proto.RegisterFile("token.proto", fileDescriptor_3aff0bcd502840ab) }

var fileDescriptor_3aff0bcd502840ab = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcf, 0x8f, 0xd2, 0x40,
	0x14, 0xc7, 0xd3, 0x22, 0x2c, 0x3c, 0x64, 0x09, 0xb3, 0x64, 0x45, 0xbc, 0x90, 0xae, 0xab, 0x1c,
	0x0c, 0xd9, 0x60, 0xdc, 0x8b, 0xd1, 0xa8, 0x31, 0xf1, 0x6a, 0xea, 0xea, 0xc1, 0x0b, 0x99, 0x65,
	0x9e, 0xd8, 0x14, 0x66, 0x6a, 0x3b, 0x5d, 0xfe, 0x58, 0xff, 0x19, 0x33, 0x6f, 0x46, 0xd2, 0x5f,
	0x4a, 0xf6, 0x42, 0xfa, 0xed, 0x9b, 0xcf, 0x9b, 0xef, 0xfb, 0x41, 0xa1, 0xaf, 0x55, 0x8c, 0x72,
	0x91, 0xa4, 0x4a, 0x2b, 0xd6, 0x26, 0x11, 0xbc, 0x81, 0x11, 0x17, 0x62, 0x45, 0x62, 0x95, 0xe2,
	0xaf, 0x1c, 0x33, 0xcd, 0x18, 0x3c, 0x90, 0x7c, 0x87, 0x13, 0x6f, 0xe6, 0xcd, 0x7b, 0x21, 0x3d,
	0xb3, 0x31, 0xb4, 0x93, 0x34, 0x5a, 0xe3, 0xc4, 0x9f, 0x79, 0x73, 0x2f, 0xb4, 0x22, 0x78, 0x07,
	0x43, 0x99, 0xef, 0x6e, 0x31, 0x5d, 0xa9, 0x1f, 0xab, 0x3b, 0xa5, 0x31, 0x33, 0x07, 0xd5, 0x5e,
	0x62, 0xea, 0x68, 0x2b, 0xd8, 0x23, 0x38, 0x91, 0xf6, 0x00, 0x25, 0x18, 0x84, 0x1d, 0xf9, 0xcd,
	0xa8, 0xe0, 0x02, 0x46, 0x1b, 0xd4, 0x15, 0x03, 0xa7, 0xe0, 0x47, 0xc2, 0x25, 0xf0, 0x23, 0x11,
	0xbc, 0x07, 0x6b, 0xb7, 0x1a, 0x38, 0x38, 0xf5, 0x9b, 0x9c, 0xb6, 0x8a, 0x4e, 0xdf, 0x02, 0xdb,
	0x46, 0x59, 0xf5, 0xa2, 0x31, 0xb4, 0xb7, 0xd1, 0x2e, 0xd2, 0x94, 0x72, 0x10, 0x5a, 0x61, 0xb2,
	0x26, 0x7c, 0x83, 0xce, 0x29, 0x3d, 0x07, 0xcf, 0x60, 0x2c, 0x70, 0x8b, 0x1a, 0x8f, 0x58, 0xfd,
	0x0c, 0xe3, 0x3c, 0x11, 0xfc, 0xd8, 0xb9, 0x7b, 0x38, 0x7f, 0x0e, 0x6d, 0xdb, 0xd9, 0x6a, 0x8a,
	0x53, 0xf0, 0xb5, 0x72, 0x09, 0x7c, 0xad, 0x82, 0x25, 0x80, 0xbd, 0xd3, 0x14, 0xca, 0x9e, 0x42,
	0x87, 0x54, 0x36, 0xf1, 0x66, 0xad, 0x79, 0x7f, 0xf9, 0x70, 0x61, 0xc7, 0x4f, 0xbf, 0xa1, 0x8b,
	0x05, 0x97, 0x70, 0x96, 0x27, 0x26, 0x3d, 0x0d, 0xe7, 0x9f, 0x55, 0x5d, 0xc0, 0x48, 0xa8, 0xbd,
	0xfc, 0xff, 0xa1, 0x4b, 0x38, 0xdb, 0x73, 0xbd, 0xfe, 0x79, 0xa4, 0x43, 0x2f, 0xc0, 0x57, 0xb1,
	0x79, 0xab, 0x62, 0x7a, 0xdb, 0x0d, 0x8d, 0x3e, 0x87, 0x8e, 0x54, 0x1f, 0xb9, 0xe6, 0x54, 0x50,
	0x37, 0x74, 0x6a, 0xf9, 0xbb, 0x05, 0xec, 0xc6, 0xe4, 0xfb, 0x4a, 0x36, 0xbf, 0x60, 0x7a, 0x17,
	0xad, 0x91, 0x2d, 0xa1, 0xcb, 0x85, 0xa0, 0x00, 0x9b, 0xb8, 0xca, 0x6a, 0x8b, 0x3c, 0x2d, 0xd5,
	0x6c, 0x98, 0x0d, 0xea, 0x32, 0x53, 0xdb, 0xbd, 0x0a, 0x73, 0x0d, 0x3d, 0xd3, 0x4d, 0x0b, 0x3d,
	0x76, 0xa1, 0xfa, 0x22, 0x95, 0xa9, 0x2b, 0x8f, 0xbd, 0x82, 0xbe, 0x5d, 0x17, 0x4b, 0x3e, 0x71,
	0xe1, 0xa6, 0x15, 0x9a, 0xf6, 0x5c, 0x50, 0xc5, 0x06, 0xb3, 0xdb, 0x53, 0xc6, 0x9a, 0x36, 0xaa,
	0x88, 0x5d, 0xc3, 0xc0, 0x4e, 0xf1, 0x46, 0x59, 0x70, 0x7a, 0x00, 0x6b, 0xb3, 0x2d, 0x73, 0x43,
	0x33, 0xd6, 0x22, 0xf9, 0xb7, 0x31, 0xb5, 0x71, 0x17, 0xb9, 0x4f, 0x30, 0xa4, 0x49, 0x13, 0x42,
	0xff, 0xe3, 0xc3, 0x8d, 0x0d, 0x1b, 0x30, 0x3d, 0x77, 0xb1, 0xca, 0xa7, 0xe2, 0xca, 0xfb, 0xd0,
	0xfb, 0x7e, 0xb2, 0x78, 0x4d, 0xc1, 0xdb, 0x0e, 0x7d, 0x97, 0x5e, 0xfe, 0x09, 0x00, 0x00, 0xff,
	0xff, 0xaa, 0x1c, 0xfb, 0x1c, 0xa6, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TokenUpvoteServiceClient is the client API for TokenUpvoteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TokenUpvoteServiceClient interface {
	AddToken(ctx context.Context, in *AddTokenRequest, opts ...grpc.CallOption) (*Token, error)
	GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*Token, error)
	ListToken(ctx context.Context, in *ListTokenRequest, opts ...grpc.CallOption) (TokenUpvoteService_ListTokenClient, error)
	DeleteToken(ctx context.Context, in *DeleteTokenRequest, opts ...grpc.CallOption) (*Ok, error)
	UpdateToken(ctx context.Context, in *UpdateTokenRequest, opts ...grpc.CallOption) (*Ok, error)
	UpvoteToToken(ctx context.Context, in *UpvoteVoteRequest, opts ...grpc.CallOption) (*Ok, error)
	DownvoteToToken(ctx context.Context, in *DownVoteRequest, opts ...grpc.CallOption) (*Ok, error)
	WatchTokenVotes(ctx context.Context, in *WatchTokenRequest, opts ...grpc.CallOption) (TokenUpvoteService_WatchTokenVotesClient, error)
}

type tokenUpvoteServiceClient struct {
	cc *grpc.ClientConn
}

func NewTokenUpvoteServiceClient(cc *grpc.ClientConn) TokenUpvoteServiceClient {
	return &tokenUpvoteServiceClient{cc}
}

func (c *tokenUpvoteServiceClient) AddToken(ctx context.Context, in *AddTokenRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/token.TokenUpvoteService/addToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenUpvoteServiceClient) GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/token.TokenUpvoteService/getToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenUpvoteServiceClient) ListToken(ctx context.Context, in *ListTokenRequest, opts ...grpc.CallOption) (TokenUpvoteService_ListTokenClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TokenUpvoteService_serviceDesc.Streams[0], "/token.TokenUpvoteService/listToken", opts...)
	if err != nil {
		return nil, err
	}
	x := &tokenUpvoteServiceListTokenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TokenUpvoteService_ListTokenClient interface {
	Recv() (*Token, error)
	grpc.ClientStream
}

type tokenUpvoteServiceListTokenClient struct {
	grpc.ClientStream
}

func (x *tokenUpvoteServiceListTokenClient) Recv() (*Token, error) {
	m := new(Token)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tokenUpvoteServiceClient) DeleteToken(ctx context.Context, in *DeleteTokenRequest, opts ...grpc.CallOption) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/token.TokenUpvoteService/deleteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenUpvoteServiceClient) UpdateToken(ctx context.Context, in *UpdateTokenRequest, opts ...grpc.CallOption) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/token.TokenUpvoteService/updateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenUpvoteServiceClient) UpvoteToToken(ctx context.Context, in *UpvoteVoteRequest, opts ...grpc.CallOption) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/token.TokenUpvoteService/upvoteToToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenUpvoteServiceClient) DownvoteToToken(ctx context.Context, in *DownVoteRequest, opts ...grpc.CallOption) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/token.TokenUpvoteService/downvoteToToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenUpvoteServiceClient) WatchTokenVotes(ctx context.Context, in *WatchTokenRequest, opts ...grpc.CallOption) (TokenUpvoteService_WatchTokenVotesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TokenUpvoteService_serviceDesc.Streams[1], "/token.TokenUpvoteService/watchTokenVotes", opts...)
	if err != nil {
		return nil, err
	}
	x := &tokenUpvoteServiceWatchTokenVotesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TokenUpvoteService_WatchTokenVotesClient interface {
	Recv() (*NumberOfVotes, error)
	grpc.ClientStream
}

type tokenUpvoteServiceWatchTokenVotesClient struct {
	grpc.ClientStream
}

func (x *tokenUpvoteServiceWatchTokenVotesClient) Recv() (*NumberOfVotes, error) {
	m := new(NumberOfVotes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TokenUpvoteServiceServer is the server API for TokenUpvoteService service.
type TokenUpvoteServiceServer interface {
	AddToken(context.Context, *AddTokenRequest) (*Token, error)
	GetToken(context.Context, *GetTokenRequest) (*Token, error)
	ListToken(*ListTokenRequest, TokenUpvoteService_ListTokenServer) error
	DeleteToken(context.Context, *DeleteTokenRequest) (*Ok, error)
	UpdateToken(context.Context, *UpdateTokenRequest) (*Ok, error)
	UpvoteToToken(context.Context, *UpvoteVoteRequest) (*Ok, error)
	DownvoteToToken(context.Context, *DownVoteRequest) (*Ok, error)
	WatchTokenVotes(*WatchTokenRequest, TokenUpvoteService_WatchTokenVotesServer) error
}

// UnimplementedTokenUpvoteServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTokenUpvoteServiceServer struct {
}

func (*UnimplementedTokenUpvoteServiceServer) AddToken(ctx context.Context, req *AddTokenRequest) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToken not implemented")
}
func (*UnimplementedTokenUpvoteServiceServer) GetToken(ctx context.Context, req *GetTokenRequest) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (*UnimplementedTokenUpvoteServiceServer) ListToken(req *ListTokenRequest, srv TokenUpvoteService_ListTokenServer) error {
	return status.Errorf(codes.Unimplemented, "method ListToken not implemented")
}
func (*UnimplementedTokenUpvoteServiceServer) DeleteToken(ctx context.Context, req *DeleteTokenRequest) (*Ok, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteToken not implemented")
}
func (*UnimplementedTokenUpvoteServiceServer) UpdateToken(ctx context.Context, req *UpdateTokenRequest) (*Ok, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateToken not implemented")
}
func (*UnimplementedTokenUpvoteServiceServer) UpvoteToToken(ctx context.Context, req *UpvoteVoteRequest) (*Ok, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpvoteToToken not implemented")
}
func (*UnimplementedTokenUpvoteServiceServer) DownvoteToToken(ctx context.Context, req *DownVoteRequest) (*Ok, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownvoteToToken not implemented")
}
func (*UnimplementedTokenUpvoteServiceServer) WatchTokenVotes(req *WatchTokenRequest, srv TokenUpvoteService_WatchTokenVotesServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchTokenVotes not implemented")
}

func RegisterTokenUpvoteServiceServer(s *grpc.Server, srv TokenUpvoteServiceServer) {
	s.RegisterService(&_TokenUpvoteService_serviceDesc, srv)
}

func _TokenUpvoteService_AddToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenUpvoteServiceServer).AddToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/token.TokenUpvoteService/AddToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenUpvoteServiceServer).AddToken(ctx, req.(*AddTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenUpvoteService_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenUpvoteServiceServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/token.TokenUpvoteService/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenUpvoteServiceServer).GetToken(ctx, req.(*GetTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenUpvoteService_ListToken_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListTokenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TokenUpvoteServiceServer).ListToken(m, &tokenUpvoteServiceListTokenServer{stream})
}

type TokenUpvoteService_ListTokenServer interface {
	Send(*Token) error
	grpc.ServerStream
}

type tokenUpvoteServiceListTokenServer struct {
	grpc.ServerStream
}

func (x *tokenUpvoteServiceListTokenServer) Send(m *Token) error {
	return x.ServerStream.SendMsg(m)
}

func _TokenUpvoteService_DeleteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenUpvoteServiceServer).DeleteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/token.TokenUpvoteService/DeleteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenUpvoteServiceServer).DeleteToken(ctx, req.(*DeleteTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenUpvoteService_UpdateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenUpvoteServiceServer).UpdateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/token.TokenUpvoteService/UpdateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenUpvoteServiceServer).UpdateToken(ctx, req.(*UpdateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenUpvoteService_UpvoteToToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpvoteVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenUpvoteServiceServer).UpvoteToToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/token.TokenUpvoteService/UpvoteToToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenUpvoteServiceServer).UpvoteToToken(ctx, req.(*UpvoteVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenUpvoteService_DownvoteToToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenUpvoteServiceServer).DownvoteToToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/token.TokenUpvoteService/DownvoteToToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenUpvoteServiceServer).DownvoteToToken(ctx, req.(*DownVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenUpvoteService_WatchTokenVotes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchTokenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TokenUpvoteServiceServer).WatchTokenVotes(m, &tokenUpvoteServiceWatchTokenVotesServer{stream})
}

type TokenUpvoteService_WatchTokenVotesServer interface {
	Send(*NumberOfVotes) error
	grpc.ServerStream
}

type tokenUpvoteServiceWatchTokenVotesServer struct {
	grpc.ServerStream
}

func (x *tokenUpvoteServiceWatchTokenVotesServer) Send(m *NumberOfVotes) error {
	return x.ServerStream.SendMsg(m)
}

var _TokenUpvoteService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "token.TokenUpvoteService",
	HandlerType: (*TokenUpvoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addToken",
			Handler:    _TokenUpvoteService_AddToken_Handler,
		},
		{
			MethodName: "getToken",
			Handler:    _TokenUpvoteService_GetToken_Handler,
		},
		{
			MethodName: "deleteToken",
			Handler:    _TokenUpvoteService_DeleteToken_Handler,
		},
		{
			MethodName: "updateToken",
			Handler:    _TokenUpvoteService_UpdateToken_Handler,
		},
		{
			MethodName: "upvoteToToken",
			Handler:    _TokenUpvoteService_UpvoteToToken_Handler,
		},
		{
			MethodName: "downvoteToToken",
			Handler:    _TokenUpvoteService_DownvoteToToken_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "listToken",
			Handler:       _TokenUpvoteService_ListToken_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "watchTokenVotes",
			Handler:       _TokenUpvoteService_WatchTokenVotes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "token.proto",
}
