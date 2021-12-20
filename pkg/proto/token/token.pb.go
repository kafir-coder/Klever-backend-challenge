// Code generated by protoc-gen-go. DO NOT EDIT.
// source: token.proto

package token

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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
	return fileDescriptor_3aff0bcd502840ab, []int{1}
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
	return fileDescriptor_3aff0bcd502840ab, []int{2}
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
	return fileDescriptor_3aff0bcd502840ab, []int{3}
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
	return fileDescriptor_3aff0bcd502840ab, []int{4}
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
	return fileDescriptor_3aff0bcd502840ab, []int{5}
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
	return fileDescriptor_3aff0bcd502840ab, []int{6}
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
	return fileDescriptor_3aff0bcd502840ab, []int{7}
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
	return fileDescriptor_3aff0bcd502840ab, []int{8}
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
	return fileDescriptor_3aff0bcd502840ab, []int{9}
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
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ok) Reset()         { *m = Ok{} }
func (m *Ok) String() string { return proto.CompactTextString(m) }
func (*Ok) ProtoMessage()    {}
func (*Ok) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aff0bcd502840ab, []int{10}
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

func init() {
	proto.RegisterType((*AddTokenRequest)(nil), "token.add_token_request")
	proto.RegisterType((*GetTokenRequest)(nil), "token.get_token_request")
	proto.RegisterType((*Token)(nil), "token.token")
	proto.RegisterType((*ListTokenRequest)(nil), "token.list_token_request")
	proto.RegisterType((*DeleteTokenRequest)(nil), "token.delete_token_request")
	proto.RegisterType((*Votes)(nil), "token.votes")
	proto.RegisterType((*TokenList)(nil), "token.token_list")
	proto.RegisterType((*UpvoteVoteRequest)(nil), "token.upvote_vote_request")
	proto.RegisterType((*DownVoteRequest)(nil), "token.down_vote_request")
	proto.RegisterType((*WatchTokenRequest)(nil), "token.watch_token_request")
	proto.RegisterType((*Ok)(nil), "token.ok")
}

func init() { proto.RegisterFile("token.proto", fileDescriptor_3aff0bcd502840ab) }

var fileDescriptor_3aff0bcd502840ab = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x95, 0x5d, 0xb2, 0x34, 0x53, 0xba, 0x4b, 0xbc, 0x91, 0x08, 0x81, 0x43, 0x95, 0x65, 0xa1,
	0xa7, 0x82, 0x82, 0x40, 0x42, 0x08, 0x24, 0x7a, 0xe1, 0x1e, 0x16, 0x0e, 0x5c, 0xaa, 0x50, 0x5b,
	0xc5, 0x4a, 0xb6, 0x0e, 0x89, 0xdb, 0xfe, 0x03, 0xff, 0xc0, 0x8f, 0x72, 0x42, 0x1e, 0xbb, 0x28,
	0x4d, 0x0a, 0xbd, 0x44, 0xf6, 0x8c, 0xdf, 0x9b, 0x37, 0xf3, 0x26, 0x30, 0xd2, 0xaa, 0x10, 0xeb,
	0x59, 0x55, 0x2b, 0xad, 0x98, 0x87, 0x97, 0xf8, 0xc1, 0x36, 0x2f, 0x25, 0xcf, 0xb5, 0x78, 0xbe,
	0x3f, 0xd8, 0x7c, 0xf2, 0x11, 0x82, 0x9c, 0xf3, 0x05, 0xbe, 0x5a, 0xd4, 0xe2, 0xc7, 0x46, 0x34,
	0x9a, 0x3d, 0x86, 0x3b, 0xeb, 0xfc, 0x56, 0x44, 0x64, 0x42, 0xa6, 0xfe, 0x7c, 0xf8, 0x7b, 0xee,
	0xfd, 0x24, 0xf4, 0x3e, 0xc9, 0x30, 0xca, 0x42, 0xf0, 0xaa, 0x5a, 0x2e, 0x45, 0x44, 0x27, 0x64,
	0x4a, 0x32, 0x7b, 0x49, 0xae, 0x20, 0x58, 0x09, 0xdd, 0x21, 0x3a, 0x07, 0x2a, 0xb9, 0xa5, 0xc9,
	0xa8, 0xe4, 0xc9, 0x07, 0xb0, 0x7a, 0xba, 0x09, 0xc6, 0x5c, 0x45, 0x8a, 0x91, 0x4e, 0x9d, 0x41,
	0xbb, 0xce, 0x7b, 0x60, 0xa5, 0x6c, 0xba, 0x85, 0x42, 0xf0, 0x4a, 0x79, 0x2b, 0x35, 0x52, 0x8e,
	0x33, 0x7b, 0x31, 0xac, 0x55, 0xbe, 0xb2, 0xac, 0xe3, 0x0c, 0xcf, 0xc9, 0x53, 0x08, 0xb9, 0x28,
	0x85, 0x16, 0x27, 0xa4, 0x3e, 0x03, 0x6f, 0xab, 0xb4, 0x68, 0x7a, 0x52, 0xcf, 0x81, 0x6a, 0xe5,
	0x84, 0x52, 0xad, 0x92, 0x14, 0xc0, 0x32, 0x19, 0x59, 0xec, 0x09, 0x9c, 0xe1, 0xad, 0x89, 0xc8,
	0x64, 0x30, 0x1d, 0xa5, 0xf7, 0x66, 0xd6, 0x0d, 0xfc, 0x66, 0x2e, 0x97, 0x5c, 0xc3, 0xe5, 0xa6,
	0x32, 0xf4, 0x0b, 0xfc, 0xfc, 0x4b, 0xc3, 0x15, 0x04, 0x5c, 0xed, 0xd6, 0xff, 0x7f, 0x74, 0x0d,
	0x97, 0xbb, 0x5c, 0x2f, 0xbf, 0x9f, 0xe8, 0x27, 0x04, 0xaa, 0x0a, 0x13, 0x55, 0x05, 0x46, 0x87,
	0x19, 0x55, 0x45, 0xfa, 0x6b, 0x00, 0xec, 0xc6, 0xe0, 0x3e, 0xa3, 0x9c, 0x4f, 0xa2, 0xde, 0xca,
	0xa5, 0x60, 0x29, 0x0c, 0x73, 0xce, 0x31, 0xc1, 0x22, 0xd7, 0x41, 0x6f, 0x4d, 0xe2, 0x83, 0xde,
	0x0c, 0x66, 0x25, 0xf4, 0x21, 0xa6, 0xb7, 0x11, 0x1d, 0xcc, 0x1b, 0xf0, 0xcd, 0xd4, 0x2c, 0xe8,
	0xa1, 0x4b, 0xf5, 0xed, 0x8d, 0x83, 0x36, 0xca, 0x0e, 0xfa, 0x15, 0x8c, 0xac, 0x8f, 0x16, 0xfc,
	0xc8, 0xbd, 0x38, 0xe6, 0x6d, 0xec, 0xbb, 0xa4, 0x2a, 0xd8, 0x6b, 0x18, 0xdb, 0xc9, 0xdf, 0x28,
	0x0b, 0x8c, 0x5d, 0xee, 0x88, 0x1f, 0x87, 0xb8, 0x0b, 0x63, 0x45, 0x1b, 0xb9, 0x6f, 0xb2, 0x67,
	0x51, 0x1b, 0xf7, 0x0e, 0x2e, 0xd0, 0x1d, 0x84, 0x7c, 0xc1, 0x85, 0xda, 0x57, 0x3c, 0xe2, 0xda,
	0xdf, 0xf1, 0xe0, 0xea, 0xbd, 0x20, 0x73, 0xff, 0xeb, 0xdd, 0xd9, 0x5b, 0x0c, 0x7d, 0x3b, 0xc3,
	0x1f, 0xf6, 0xe5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xd9, 0x01, 0x2d, 0xdf, 0x03, 0x00,
	0x00,
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
	ListToken(ctx context.Context, in *ListTokenRequest, opts ...grpc.CallOption) (*TokenList, error)
	DeleteToken(ctx context.Context, in *DeleteTokenRequest, opts ...grpc.CallOption) (*Ok, error)
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

func (c *tokenUpvoteServiceClient) ListToken(ctx context.Context, in *ListTokenRequest, opts ...grpc.CallOption) (*TokenList, error) {
	out := new(TokenList)
	err := c.cc.Invoke(ctx, "/token.TokenUpvoteService/listToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenUpvoteServiceClient) DeleteToken(ctx context.Context, in *DeleteTokenRequest, opts ...grpc.CallOption) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/token.TokenUpvoteService/deleteToken", in, out, opts...)
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
	stream, err := c.cc.NewStream(ctx, &_TokenUpvoteService_serviceDesc.Streams[0], "/token.TokenUpvoteService/watchTokenVotes", opts...)
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
	Recv() (*Votes, error)
	grpc.ClientStream
}

type tokenUpvoteServiceWatchTokenVotesClient struct {
	grpc.ClientStream
}

func (x *tokenUpvoteServiceWatchTokenVotesClient) Recv() (*Votes, error) {
	m := new(Votes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TokenUpvoteServiceServer is the server API for TokenUpvoteService service.
type TokenUpvoteServiceServer interface {
	AddToken(context.Context, *AddTokenRequest) (*Token, error)
	GetToken(context.Context, *GetTokenRequest) (*Token, error)
	ListToken(context.Context, *ListTokenRequest) (*TokenList, error)
	DeleteToken(context.Context, *DeleteTokenRequest) (*Ok, error)
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
func (*UnimplementedTokenUpvoteServiceServer) ListToken(ctx context.Context, req *ListTokenRequest) (*TokenList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListToken not implemented")
}
func (*UnimplementedTokenUpvoteServiceServer) DeleteToken(ctx context.Context, req *DeleteTokenRequest) (*Ok, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteToken not implemented")
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

func _TokenUpvoteService_ListToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenUpvoteServiceServer).ListToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/token.TokenUpvoteService/ListToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenUpvoteServiceServer).ListToken(ctx, req.(*ListTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
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
	Send(*Votes) error
	grpc.ServerStream
}

type tokenUpvoteServiceWatchTokenVotesServer struct {
	grpc.ServerStream
}

func (x *tokenUpvoteServiceWatchTokenVotesServer) Send(m *Votes) error {
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
			MethodName: "listToken",
			Handler:    _TokenUpvoteService_ListToken_Handler,
		},
		{
			MethodName: "deleteToken",
			Handler:    _TokenUpvoteService_DeleteToken_Handler,
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
			StreamName:    "watchTokenVotes",
			Handler:       _TokenUpvoteService_WatchTokenVotes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "token.proto",
}
