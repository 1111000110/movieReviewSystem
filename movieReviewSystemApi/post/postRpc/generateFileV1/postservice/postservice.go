// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.5
// Source: post.proto

package postservice

import (
	"context"

	"movieReviewSystem/movieReviewSystemApi/post/postRpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	PostCreateReq         = __.PostCreateReq
	PostCreateResp        = __.PostCreateResp
	PostDeleteByUserIdReq = __.PostDeleteByUserIdReq
	PostDeleteResp        = __.PostDeleteResp
	PostFindByPostIdReq   = __.PostFindByPostIdReq
	PostFindResp          = __.PostFindResp

	PostService interface {
		PostCreate(ctx context.Context, in *PostCreateReq, opts ...grpc.CallOption) (*PostCreateResp, error)
		PostFindByPostId(ctx context.Context, in *PostFindByPostIdReq, opts ...grpc.CallOption) (*PostFindResp, error)
		PostDeleteByUserId(ctx context.Context, in *PostDeleteByUserIdReq, opts ...grpc.CallOption) (*PostDeleteResp, error)
	}

	defaultPostService struct {
		cli zrpc.Client
	}
)

func NewPostService(cli zrpc.Client) PostService {
	return &defaultPostService{
		cli: cli,
	}
}

func (m *defaultPostService) PostCreate(ctx context.Context, in *PostCreateReq, opts ...grpc.CallOption) (*PostCreateResp, error) {
	client := __.NewPostServiceClient(m.cli.Conn())
	return client.PostCreate(ctx, in, opts...)
}

func (m *defaultPostService) PostFindByPostId(ctx context.Context, in *PostFindByPostIdReq, opts ...grpc.CallOption) (*PostFindResp, error) {
	client := __.NewPostServiceClient(m.cli.Conn())
	return client.PostFindByPostId(ctx, in, opts...)
}

func (m *defaultPostService) PostDeleteByUserId(ctx context.Context, in *PostDeleteByUserIdReq, opts ...grpc.CallOption) (*PostDeleteResp, error) {
	client := __.NewPostServiceClient(m.cli.Conn())
	return client.PostDeleteByUserId(ctx, in, opts...)
}
