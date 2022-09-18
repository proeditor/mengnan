package logic

import (
	"context"

	"college/api/internal/svc"
	"college/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollegeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCollegeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollegeLogic {
	return &CollegeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CollegeLogic) College(req *types.CollegeRequest) (resp *types.CollegeResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
