// Code generated by goctl. DO NOT EDIT!
// Source: college.proto

package college

import (
	"context"

	"collegerpc/collegerpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CollegeAddRequest     = collegerpc.CollegeAddRequest
	CollegeAddResponse    = collegerpc.CollegeAddResponse
	CollegeUpdateRequest  = collegerpc.CollegeUpdateRequest
	CollegeUpdateResponse = collegerpc.CollegeUpdateResponse

	College interface {
		AddCollege(ctx context.Context, in *CollegeAddRequest, opts ...grpc.CallOption) (*CollegeAddResponse, error)
		UpdateCollege(ctx context.Context, in *CollegeUpdateRequest, opts ...grpc.CallOption) (*CollegeUpdateResponse, error)
	}

	defaultCollege struct {
		cli zrpc.Client
	}
)

func NewCollege(cli zrpc.Client) College {
	return &defaultCollege{
		cli: cli,
	}
}

func (m *defaultCollege) AddCollege(ctx context.Context, in *CollegeAddRequest, opts ...grpc.CallOption) (*CollegeAddResponse, error) {
	client := collegerpc.NewCollegeClient(m.cli.Conn())
	return client.AddCollege(ctx, in, opts...)
}

func (m *defaultCollege) UpdateCollege(ctx context.Context, in *CollegeUpdateRequest, opts ...grpc.CallOption) (*CollegeUpdateResponse, error) {
	client := collegerpc.NewCollegeClient(m.cli.Conn())
	return client.UpdateCollege(ctx, in, opts...)
}
