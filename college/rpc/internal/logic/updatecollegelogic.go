package logic

import (
	"context"

	"rpc/college/rpc"
	"rpc/internal/svc"

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

func (l *UpdateCollegeLogic) UpdateCollege(in *rpc.CollegeUpdateRequest) (*rpc.CollegeUpdateResponse, error) {
	// todo: add your logic here and delete this line

	return &rpc.CollegeUpdateResponse{}, nil
}
