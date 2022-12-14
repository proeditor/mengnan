// Code generated by goctl. DO NOT EDIT!
// Source: college.proto

package server

import (
	"context"

	"github.com/proeditor/mengnan/college/rpc/college-rpc"
	"github.com/proeditor/mengnan/college/rpc/internal/logic"
	"github.com/proeditor/mengnan/college/rpc/internal/svc"
)

type CollegeServer struct {
	svcCtx *svc.ServiceContext
	college_rpc.UnimplementedCollegeServer
}

func NewCollegeServer(svcCtx *svc.ServiceContext) *CollegeServer {
	return &CollegeServer{
		svcCtx: svcCtx,
	}
}

func (s *CollegeServer) AddCollege(ctx context.Context, in *college_rpc.AddCollegeRequest) (*college_rpc.AddCollegeResponse, error) {
	l := logic.NewAddCollegeLogic(ctx, s.svcCtx)
	return l.AddCollege(in)
}

func (s *CollegeServer) UpdateCollege(ctx context.Context, in *college_rpc.UpdateCollegeRequest) (*college_rpc.UpdateCollegeResponse, error) {
	l := logic.NewUpdateCollegeLogic(ctx, s.svcCtx)
	return l.UpdateCollege(in)
}
