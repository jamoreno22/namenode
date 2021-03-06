// Code generated by protoc-gen-go. DO NOT EDIT.
// source: NameNode.proto

package lab

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

type File struct {
	Log                  []byte   `protobuf:"bytes,1,opt,name=log,proto3" json:"log,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_e56228e413a39553, []int{0}
}

func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetLog() []byte {
	if m != nil {
		return m.Log
	}
	return nil
}

type Book struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Parts                int32    `protobuf:"varint,2,opt,name=parts,proto3" json:"parts,omitempty"`
	Chunks               []*Chunk `protobuf:"bytes,3,rep,name=chunks,proto3" json:"chunks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_e56228e413a39553, []int{1}
}

func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Book) GetParts() int32 {
	if m != nil {
		return m.Parts
	}
	return 0
}

func (m *Book) GetChunks() []*Chunk {
	if m != nil {
		return m.Chunks
	}
	return nil
}

type Chunk struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}
func (*Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_e56228e413a39553, []int{2}
}

func (m *Chunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chunk.Unmarshal(m, b)
}
func (m *Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chunk.Marshal(b, m, deterministic)
}
func (m *Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chunk.Merge(m, src)
}
func (m *Chunk) XXX_Size() int {
	return xxx_messageInfo_Chunk.Size(m)
}
func (m *Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_Chunk proto.InternalMessageInfo

func (m *Chunk) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Chunk) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Message struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_e56228e413a39553, []int{3}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Proposal struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Chunk                *Chunk   `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Proposal) Reset()         { *m = Proposal{} }
func (m *Proposal) String() string { return proto.CompactTextString(m) }
func (*Proposal) ProtoMessage()    {}
func (*Proposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_e56228e413a39553, []int{4}
}

func (m *Proposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proposal.Unmarshal(m, b)
}
func (m *Proposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proposal.Marshal(b, m, deterministic)
}
func (m *Proposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proposal.Merge(m, src)
}
func (m *Proposal) XXX_Size() int {
	return xxx_messageInfo_Proposal.Size(m)
}
func (m *Proposal) XXX_DiscardUnknown() {
	xxx_messageInfo_Proposal.DiscardUnknown(m)
}

var xxx_messageInfo_Proposal proto.InternalMessageInfo

func (m *Proposal) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Proposal) GetChunk() *Chunk {
	if m != nil {
		return m.Chunk
	}
	return nil
}

func init() {
	proto.RegisterType((*File)(nil), "lab.File")
	proto.RegisterType((*Book)(nil), "lab.Book")
	proto.RegisterType((*Chunk)(nil), "lab.Chunk")
	proto.RegisterType((*Message)(nil), "lab.Message")
	proto.RegisterType((*Proposal)(nil), "lab.Proposal")
}

func init() {
	proto.RegisterFile("NameNode.proto", fileDescriptor_e56228e413a39553)
}

var fileDescriptor_e56228e413a39553 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0xa5, 0x94, 0x22, 0x0c, 0x15, 0xc9, 0x84, 0x43, 0x43, 0x62, 0x42, 0xf6, 0xd4, 0x13, 0x18,
	0x38, 0x78, 0xf1, 0xe2, 0x47, 0x34, 0x1e, 0x24, 0xa6, 0xfa, 0x07, 0xb6, 0xb2, 0xe2, 0x86, 0xda,
	0x6d, 0xba, 0x8b, 0xf1, 0xa7, 0xf9, 0xf3, 0xdc, 0x1d, 0x8a, 0x91, 0x6a, 0xf4, 0xf6, 0xde, 0xbc,
	0xbe, 0xf7, 0x66, 0x9a, 0x85, 0xfe, 0x82, 0xbf, 0x8a, 0x85, 0x5a, 0x8a, 0x49, 0x51, 0x2a, 0xa3,
	0xd0, 0xcf, 0x78, 0xca, 0x22, 0x68, 0x5d, 0xcb, 0x4c, 0xe0, 0x00, 0xfc, 0x4c, 0xad, 0x22, 0x6f,
	0xec, 0xc5, 0x61, 0xe2, 0x20, 0x7b, 0x84, 0xd6, 0x85, 0x52, 0x6b, 0x44, 0x68, 0xe5, 0xd6, 0x48,
	0x52, 0x37, 0x21, 0x8c, 0x43, 0x08, 0x0a, 0x5e, 0x1a, 0x1d, 0x35, 0xed, 0x30, 0x48, 0xb6, 0x04,
	0x19, 0xb4, 0x9f, 0x5e, 0x36, 0xf9, 0x5a, 0x47, 0xfe, 0xd8, 0x8f, 0x7b, 0x33, 0x98, 0xd8, 0x86,
	0xc9, 0xa5, 0x1b, 0x25, 0x95, 0xc2, 0xa6, 0x10, 0xd0, 0xe0, 0xd7, 0x58, 0x3b, 0x5b, 0x72, 0xc3,
	0x29, 0x35, 0x4c, 0x08, 0xb3, 0x63, 0x38, 0xb8, 0x13, 0x5a, 0xf3, 0x15, 0xc9, 0x46, 0xbc, 0x9b,
	0x9d, 0xc5, 0x61, 0x76, 0x06, 0x9d, 0xfb, 0x52, 0x15, 0x4a, 0xf3, 0x0c, 0xfb, 0xd0, 0x94, 0x45,
	0xa5, 0x5a, 0x84, 0x63, 0x08, 0xa8, 0x95, 0xf2, 0xf6, 0xd7, 0xd9, 0x0a, 0xb3, 0x8f, 0x26, 0x74,
	0x76, 0x7f, 0x05, 0xa7, 0x70, 0x74, 0x23, 0xcc, 0xf9, 0x1b, 0x97, 0x69, 0x26, 0xdc, 0xe9, 0x1a,
	0x43, 0xb2, 0x54, 0xfd, 0xa3, 0x3d, 0xc6, 0x1a, 0x95, 0xe1, 0x4a, 0x6a, 0x53, 0xca, 0x74, 0x63,
	0xa4, 0xca, 0xff, 0x31, 0xc4, 0xd0, 0xb3, 0x06, 0x17, 0x7d, 0x9b, 0x3f, 0x2b, 0xec, 0x92, 0xec,
	0xe8, 0x8f, 0x2f, 0x67, 0x10, 0x3e, 0x88, 0x7c, 0xf9, 0x75, 0xda, 0x21, 0xe9, 0x3b, 0x3a, 0xda,
	0xa7, 0xac, 0x11, 0x7b, 0x27, 0x1e, 0x9e, 0xc2, 0xd0, 0xa6, 0xd3, 0x7d, 0x7f, 0xec, 0x54, 0xb7,
	0x5a, 0xe3, 0x1c, 0x06, 0xdf, 0xcb, 0x9c, 0xb9, 0x5e, 0x58, 0xdb, 0x2f, 0xf6, 0xd2, 0x36, 0x3d,
	0xa2, 0xf9, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xee, 0xec, 0xe6, 0x56, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NameNodeClient is the client API for NameNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NameNodeClient interface {
	GetAvaibleBooks(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	GetDistribution(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	GetBookInfo(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Message, error)
	//Centralizado
	SendProposal(ctx context.Context, opts ...grpc.CallOption) (NameNode_SendProposalClient, error)
	GetChunkDistribution(ctx context.Context, in *Message, opts ...grpc.CallOption) (NameNode_GetChunkDistributionClient, error)
	SendProposalDist(ctx context.Context, opts ...grpc.CallOption) (NameNode_SendProposalDistClient, error)
}

type nameNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewNameNodeClient(cc grpc.ClientConnInterface) NameNodeClient {
	return &nameNodeClient{cc}
}

func (c *nameNodeClient) GetAvaibleBooks(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/lab.NameNode/GetAvaibleBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameNodeClient) GetDistribution(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/lab.NameNode/GetDistribution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameNodeClient) GetBookInfo(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/lab.NameNode/GetBookInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameNodeClient) SendProposal(ctx context.Context, opts ...grpc.CallOption) (NameNode_SendProposalClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NameNode_serviceDesc.Streams[0], "/lab.NameNode/SendProposal", opts...)
	if err != nil {
		return nil, err
	}
	x := &nameNodeSendProposalClient{stream}
	return x, nil
}

type NameNode_SendProposalClient interface {
	Send(*Proposal) error
	Recv() (*Proposal, error)
	grpc.ClientStream
}

type nameNodeSendProposalClient struct {
	grpc.ClientStream
}

func (x *nameNodeSendProposalClient) Send(m *Proposal) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nameNodeSendProposalClient) Recv() (*Proposal, error) {
	m := new(Proposal)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nameNodeClient) GetChunkDistribution(ctx context.Context, in *Message, opts ...grpc.CallOption) (NameNode_GetChunkDistributionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NameNode_serviceDesc.Streams[1], "/lab.NameNode/GetChunkDistribution", opts...)
	if err != nil {
		return nil, err
	}
	x := &nameNodeGetChunkDistributionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NameNode_GetChunkDistributionClient interface {
	Recv() (*Proposal, error)
	grpc.ClientStream
}

type nameNodeGetChunkDistributionClient struct {
	grpc.ClientStream
}

func (x *nameNodeGetChunkDistributionClient) Recv() (*Proposal, error) {
	m := new(Proposal)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nameNodeClient) SendProposalDist(ctx context.Context, opts ...grpc.CallOption) (NameNode_SendProposalDistClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NameNode_serviceDesc.Streams[2], "/lab.NameNode/SendProposalDist", opts...)
	if err != nil {
		return nil, err
	}
	x := &nameNodeSendProposalDistClient{stream}
	return x, nil
}

type NameNode_SendProposalDistClient interface {
	Send(*Proposal) error
	CloseAndRecv() (*Message, error)
	grpc.ClientStream
}

type nameNodeSendProposalDistClient struct {
	grpc.ClientStream
}

func (x *nameNodeSendProposalDistClient) Send(m *Proposal) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nameNodeSendProposalDistClient) CloseAndRecv() (*Message, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NameNodeServer is the server API for NameNode service.
type NameNodeServer interface {
	GetAvaibleBooks(context.Context, *Message) (*Message, error)
	GetDistribution(context.Context, *Message) (*Message, error)
	GetBookInfo(context.Context, *Book) (*Message, error)
	//Centralizado
	SendProposal(NameNode_SendProposalServer) error
	GetChunkDistribution(*Message, NameNode_GetChunkDistributionServer) error
	SendProposalDist(NameNode_SendProposalDistServer) error
}

// UnimplementedNameNodeServer can be embedded to have forward compatible implementations.
type UnimplementedNameNodeServer struct {
}

func (*UnimplementedNameNodeServer) GetAvaibleBooks(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvaibleBooks not implemented")
}
func (*UnimplementedNameNodeServer) GetDistribution(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDistribution not implemented")
}
func (*UnimplementedNameNodeServer) GetBookInfo(ctx context.Context, req *Book) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookInfo not implemented")
}
func (*UnimplementedNameNodeServer) SendProposal(srv NameNode_SendProposalServer) error {
	return status.Errorf(codes.Unimplemented, "method SendProposal not implemented")
}
func (*UnimplementedNameNodeServer) GetChunkDistribution(req *Message, srv NameNode_GetChunkDistributionServer) error {
	return status.Errorf(codes.Unimplemented, "method GetChunkDistribution not implemented")
}
func (*UnimplementedNameNodeServer) SendProposalDist(srv NameNode_SendProposalDistServer) error {
	return status.Errorf(codes.Unimplemented, "method SendProposalDist not implemented")
}

func RegisterNameNodeServer(s *grpc.Server, srv NameNodeServer) {
	s.RegisterService(&_NameNode_serviceDesc, srv)
}

func _NameNode_GetAvaibleBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServer).GetAvaibleBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lab.NameNode/GetAvaibleBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServer).GetAvaibleBooks(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameNode_GetDistribution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServer).GetDistribution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lab.NameNode/GetDistribution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServer).GetDistribution(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameNode_GetBookInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServer).GetBookInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lab.NameNode/GetBookInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServer).GetBookInfo(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameNode_SendProposal_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NameNodeServer).SendProposal(&nameNodeSendProposalServer{stream})
}

type NameNode_SendProposalServer interface {
	Send(*Proposal) error
	Recv() (*Proposal, error)
	grpc.ServerStream
}

type nameNodeSendProposalServer struct {
	grpc.ServerStream
}

func (x *nameNodeSendProposalServer) Send(m *Proposal) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nameNodeSendProposalServer) Recv() (*Proposal, error) {
	m := new(Proposal)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _NameNode_GetChunkDistribution_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Message)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NameNodeServer).GetChunkDistribution(m, &nameNodeGetChunkDistributionServer{stream})
}

type NameNode_GetChunkDistributionServer interface {
	Send(*Proposal) error
	grpc.ServerStream
}

type nameNodeGetChunkDistributionServer struct {
	grpc.ServerStream
}

func (x *nameNodeGetChunkDistributionServer) Send(m *Proposal) error {
	return x.ServerStream.SendMsg(m)
}

func _NameNode_SendProposalDist_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NameNodeServer).SendProposalDist(&nameNodeSendProposalDistServer{stream})
}

type NameNode_SendProposalDistServer interface {
	SendAndClose(*Message) error
	Recv() (*Proposal, error)
	grpc.ServerStream
}

type nameNodeSendProposalDistServer struct {
	grpc.ServerStream
}

func (x *nameNodeSendProposalDistServer) SendAndClose(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nameNodeSendProposalDistServer) Recv() (*Proposal, error) {
	m := new(Proposal)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _NameNode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lab.NameNode",
	HandlerType: (*NameNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAvaibleBooks",
			Handler:    _NameNode_GetAvaibleBooks_Handler,
		},
		{
			MethodName: "GetDistribution",
			Handler:    _NameNode_GetDistribution_Handler,
		},
		{
			MethodName: "GetBookInfo",
			Handler:    _NameNode_GetBookInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendProposal",
			Handler:       _NameNode_SendProposal_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetChunkDistribution",
			Handler:       _NameNode_GetChunkDistribution_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendProposalDist",
			Handler:       _NameNode_SendProposalDist_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "NameNode.proto",
}
