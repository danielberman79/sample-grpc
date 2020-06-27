// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/comment.proto

package comment

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type CreateRequest struct {
	Comment              string   `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f6ed0492636b486, []int{0}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Comment              string               `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	Name                 string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f6ed0492636b486, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Response) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *Response) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Response) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "comment.CreateRequest")
	proto.RegisterType((*Response)(nil), "comment.Response")
}

func init() {
	proto.RegisterFile("api/comment.proto", fileDescriptor_4f6ed0492636b486)
}

var fileDescriptor_4f6ed0492636b486 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0x4d, 0x2c, 0xd5, 0x8e, 0x58, 0xe8, 0x20, 0x25, 0xe4, 0x62, 0xc9, 0xa9, 0xa7, 0x0d,
	0x54, 0x3c, 0x78, 0xf0, 0xa0, 0x3d, 0x78, 0x8f, 0x9e, 0xbc, 0xc8, 0x36, 0x19, 0xcb, 0x42, 0xf7,
	0xc3, 0xec, 0xc4, 0x3f, 0xe0, 0x1f, 0x17, 0x76, 0xb3, 0x7e, 0x40, 0x6f, 0x3b, 0x2f, 0x33, 0xcf,
	0x3e, 0x33, 0xb0, 0x90, 0x4e, 0xd5, 0xad, 0xd5, 0x9a, 0x0c, 0x0b, 0xd7, 0x5b, 0xb6, 0x78, 0x36,
	0x96, 0xe5, 0xf5, 0xde, 0xda, 0xfd, 0x81, 0xea, 0x10, 0xef, 0x86, 0xf7, 0x9a, 0x95, 0x26, 0xcf,
	0x52, 0xbb, 0xd8, 0x59, 0xdd, 0xc3, 0xe5, 0xb6, 0x27, 0xc9, 0xd4, 0xd0, 0xc7, 0x40, 0x9e, 0xb1,
	0x80, 0x34, 0x5c, 0x64, 0xab, 0x6c, 0x3d, 0x6b, 0x52, 0x89, 0x08, 0x13, 0x23, 0x35, 0x15, 0x79,
	0x88, 0xc3, 0xbb, 0xfa, 0xca, 0xe0, 0xbc, 0x21, 0xef, 0xac, 0xf1, 0x84, 0x73, 0xc8, 0x55, 0x37,
	0x4e, 0xe5, 0xaa, 0xfb, 0x8b, 0xca, 0x8f, 0xa3, 0x4e, 0x7f, 0x51, 0x78, 0x07, 0xd0, 0x06, 0x93,
	0xee, 0x4d, 0x72, 0x31, 0x59, 0x65, 0xeb, 0x8b, 0x4d, 0x29, 0xa2, 0xbf, 0x48, 0xfe, 0xe2, 0x25,
	0xf9, 0x37, 0xb3, 0xb1, 0xfb, 0x81, 0x37, 0x4f, 0x30, 0xdf, 0x46, 0xf2, 0x33, 0xf5, 0x9f, 0xaa,
	0x25, 0xbc, 0x85, 0x69, 0x5c, 0x0b, 0x97, 0x22, 0x9d, 0xe6, 0xdf, 0x9e, 0xe5, 0xe2, 0x27, 0x4f,
	0xfe, 0xd5, 0xc9, 0xe3, 0xf2, 0xf5, 0x4a, 0x19, 0xa6, 0xde, 0xc8, 0x43, 0x2d, 0x9d, 0x4b, 0x57,
	0xdd, 0x4d, 0xc3, 0xff, 0x37, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x86, 0xb6, 0xf9, 0x41, 0x6b,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CommentServiceClient is the client API for CommentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommentServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Response, error)
}

type commentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentServiceClient(cc grpc.ClientConnInterface) CommentServiceClient {
	return &commentServiceClient{cc}
}

func (c *commentServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/comment.CommentService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServiceServer is the server API for CommentService service.
type CommentServiceServer interface {
	Create(context.Context, *CreateRequest) (*Response, error)
}

// UnimplementedCommentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCommentServiceServer struct {
}

func (*UnimplementedCommentServiceServer) Create(ctx context.Context, req *CreateRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterCommentServiceServer(s *grpc.Server, srv CommentServiceServer) {
	s.RegisterService(&_CommentService_serviceDesc, srv)
}

func _CommentService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.CommentService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "comment.CommentService",
	HandlerType: (*CommentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CommentService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/comment.proto",
}