package logic

import (
	"context"

	"github.com/proeditor/mengnan/college/rpc/college-rpc"
	"github.com/proeditor/mengnan/college/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCollegeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCollegeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCollegeLogic {
	return &AddCollegeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCollegeLogic) AddCollege(in *college_rpc.AddCollegeRequest) (*college_rpc.AddCollegeResponse, error) {
	// todo: add your logic here and delete this line

	return &college_rpc.AddCollegeResponse{}, nil
}
