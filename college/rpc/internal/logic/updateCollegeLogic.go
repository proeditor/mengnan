package logic

import (
	"context"

	"github.com/proeditor/mengnan/college/rpc/college-rpc"
	"github.com/proeditor/mengnan/college/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCollegeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCollegeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCollegeLogic {
	return &UpdateCollegeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCollegeLogic) UpdateCollege(in *college_rpc.UpdateCollegeRequest) (*college_rpc.UpdateCollegeResponse, error) {
	// todo: add your logic here and delete this line

	return &college_rpc.UpdateCollegeResponse{}, nil
}
