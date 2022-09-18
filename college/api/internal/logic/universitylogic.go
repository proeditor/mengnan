package logic

import (
	"context"

	"college/api/internal/svc"
	"college/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UniversityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUniversityLogic(ctx context.Context, svcCtx *svc.ServiceContext) UniversityLogic {
	return UniversityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UniversityLogic) University(req types.CollegeRequest) (*types.CollegeResponse, error) {
	// todo: add your logic here and delete this line

	return &types.CollegeResponse{}, nil
}
