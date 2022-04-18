package logic

import (
	"context"

	"mall/service/user/api/internal/svc"
	"mall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserResetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserResetLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserResetLogic {
	return UserResetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserResetLogic) UserReset() (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
