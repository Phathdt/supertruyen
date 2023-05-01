// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/chapter.proto

package supertruyen_proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChapterServiceClient is the client API for ChapterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChapterServiceClient interface {
	GetTotalChapter(ctx context.Context, in *GetTotalChapterRequest, opts ...grpc.CallOption) (*GetTotalChapterResponse, error)
}

type chapterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChapterServiceClient(cc grpc.ClientConnInterface) ChapterServiceClient {
	return &chapterServiceClient{cc}
}

func (c *chapterServiceClient) GetTotalChapter(ctx context.Context, in *GetTotalChapterRequest, opts ...grpc.CallOption) (*GetTotalChapterResponse, error) {
	out := new(GetTotalChapterResponse)
	err := c.cc.Invoke(ctx, "/demo.ChapterService/GetTotalChapter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChapterServiceServer is the server API for ChapterService service.
// All implementations should embed UnimplementedChapterServiceServer
// for forward compatibility
type ChapterServiceServer interface {
	GetTotalChapter(context.Context, *GetTotalChapterRequest) (*GetTotalChapterResponse, error)
}

// UnimplementedChapterServiceServer should be embedded to have forward compatible implementations.
type UnimplementedChapterServiceServer struct {
}

func (UnimplementedChapterServiceServer) GetTotalChapter(context.Context, *GetTotalChapterRequest) (*GetTotalChapterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTotalChapter not implemented")
}

// UnsafeChapterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChapterServiceServer will
// result in compilation errors.
type UnsafeChapterServiceServer interface {
	mustEmbedUnimplementedChapterServiceServer()
}

func RegisterChapterServiceServer(s grpc.ServiceRegistrar, srv ChapterServiceServer) {
	s.RegisterService(&ChapterService_ServiceDesc, srv)
}

func _ChapterService_GetTotalChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTotalChapterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChapterServiceServer).GetTotalChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.ChapterService/GetTotalChapter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChapterServiceServer).GetTotalChapter(ctx, req.(*GetTotalChapterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChapterService_ServiceDesc is the grpc.ServiceDesc for ChapterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChapterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "demo.ChapterService",
	HandlerType: (*ChapterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTotalChapter",
			Handler:    _ChapterService_GetTotalChapter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/chapter.proto",
}
