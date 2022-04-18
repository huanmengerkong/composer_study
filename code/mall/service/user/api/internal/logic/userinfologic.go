package logic

import (
	"context"
	"fmt"
	"mall/common"
	"mall/service/user/rpc/userclient"

	"mall/service/user/api/internal/svc"
	"mall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserInfoLogic {
	return UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	uid := l.ctx.Value("uid")
	logx.Info("解密token获取参数",uid)
	user,err :=l.svcCtx.UserRpc.UserInfo(l.ctx,&userclient.UserInfoRequest{Id: int64(common.GetInt64(uid))})
	fmt.Println(user)
	if err !=nil{
		return nil,err
	}
	var resps types.UserInfoResponse
	if user != nil {
		resps.Name = user.Name
		resps.Id = user.Id
		resps.Gender = user.Gender
		resps.Mobile = user.Mobile
	}
	return &resps,err
}
