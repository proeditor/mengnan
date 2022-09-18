package logic

import (
	"context"

	"collegerpc/collegerpc"
	"collegerpc/internal/svc"

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

func (l *AddCollegeLogic) AddCollege(in *collegerpc.CollegeAddRequest) (*collegerpc.CollegeAddResponse, error) {
	// todo: add your logic here and delete this line

	return &collegerpc.CollegeAddResponse{}, nil
}
