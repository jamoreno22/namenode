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
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x51, 0x4f, 0x4b, 0xc3, 0x30,
	0x14, 0x5f, 0xd6, 0x75, 0x6e, 0x6f, 0x75, 0xc8, 0x63, 0x87, 0x32, 0x10, 0x4a, 0x4e, 0x3d, 0xd5,
	0x51, 0x0f, 0x5e, 0x3c, 0xa9, 0x28, 0x1e, 0x1c, 0x52, 0xfd, 0x02, 0xa9, 0x8d, 0xb3, 0xac, 0x36,
	0xa5, 0xc9, 0xc0, 0xaf, 0xe3, 0x37, 0x35, 0x79, 0x6b, 0x85, 0xa2, 0xb7, 0xdf, 0x9f, 0xbc, 0xdf,
	0xef, 0x25, 0x81, 0xe5, 0x56, 0x7c, 0xca, 0xad, 0x2a, 0x64, 0xd2, 0xb4, 0xca, 0x28, 0xf4, 0x2a,
	0x91, 0xf3, 0x10, 0x26, 0xf7, 0x65, 0x25, 0xf1, 0x0c, 0xbc, 0x4a, 0xed, 0x42, 0x16, 0xb1, 0x38,
	0xc8, 0x1c, 0xe4, 0xaf, 0x30, 0xb9, 0x51, 0x6a, 0x8f, 0x08, 0x93, 0xda, 0x0e, 0x92, 0x35, 0xcf,
	0x08, 0xe3, 0x0a, 0xfc, 0x46, 0xb4, 0x46, 0x87, 0x63, 0x2b, 0xfa, 0xd9, 0x91, 0x20, 0x87, 0xe9,
	0xdb, 0xc7, 0xa1, 0xde, 0xeb, 0xd0, 0x8b, 0xbc, 0x78, 0x91, 0x42, 0x62, 0x1b, 0x92, 0x5b, 0x27,
	0x65, 0x9d, 0xc3, 0x2f, 0xc0, 0x27, 0xe1, 0xdf, 0x58, 0xab, 0x15, 0xc2, 0x08, 0x4a, 0x0d, 0x32,
	0xc2, 0xfc, 0x1c, 0x4e, 0x9e, 0xa4, 0xd6, 0x62, 0x47, 0xb6, 0x91, 0x5f, 0xa6, 0x1f, 0x71, 0x98,
	0x5f, 0xc3, 0xec, 0xb9, 0x55, 0x8d, 0xd2, 0xa2, 0xc2, 0x25, 0x8c, 0xcb, 0xa6, 0x73, 0x2d, 0xc2,
	0x08, 0x7c, 0x6a, 0xa5, 0xbc, 0xe1, 0x3a, 0x47, 0x23, 0xfd, 0x66, 0x30, 0xeb, 0x5f, 0x05, 0x63,
	0x58, 0x3c, 0x48, 0xe3, 0xee, 0xfc, 0x58, 0xbf, 0x2b, 0x9c, 0xd3, 0x71, 0x47, 0xd7, 0x01, 0xc1,
	0x6e, 0x0d, 0x3e, 0xc2, 0x14, 0x82, 0x17, 0x59, 0x17, 0xbf, 0xc5, 0xa7, 0xe4, 0xf7, 0x74, 0x3d,
	0xa4, 0x7c, 0x14, 0xb3, 0x0d, 0xc3, 0x2b, 0x58, 0xd9, 0x74, 0x6a, 0xbf, 0x2b, 0xb5, 0x69, 0xcb,
	0xfc, 0x60, 0x4a, 0x55, 0xe3, 0x20, 0xfb, 0xcf, 0xe8, 0x86, 0xe5, 0x53, 0xfa, 0xad, 0xcb, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x7e, 0xe9, 0x3a, 0xbf, 0x01, 0x00, 0x00,
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
	GetBookInfo(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Message, error)
	//Centralizado
	SendProposal(ctx context.Context, opts ...grpc.CallOption) (NameNode_SendProposalClient, error)
	GetChunkDistribution(ctx context.Context, in *Message, opts ...grpc.CallOption) (NameNode_GetChunkDistributionClient, error)
}

type nameNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewNameNodeClient(cc grpc.ClientConnInterface) NameNodeClient {
	return &nameNodeClient{cc}
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

// NameNodeServer is the server API for NameNode service.
type NameNodeServer interface {
	GetBookInfo(context.Context, *Book) (*Message, error)
	//Centralizado
	SendProposal(NameNode_SendProposalServer) error
	GetChunkDistribution(*Message, NameNode_GetChunkDistributionServer) error
}

// UnimplementedNameNodeServer can be embedded to have forward compatible implementations.
type UnimplementedNameNodeServer struct {
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

func RegisterNameNodeServer(s *grpc.Server, srv NameNodeServer) {
	s.RegisterService(&_NameNode_serviceDesc, srv)
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

var _NameNode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lab.NameNode",
	HandlerType: (*NameNodeServer)(nil),
	Methods: []grpc.MethodDesc{
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
	},
	Metadata: "NameNode.proto",
}
