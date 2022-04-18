package logic

import (
	"context"

	"mall/service/user/api/internal/svc"
	"mall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) ImageLogic {
	return ImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImageLogic) Image(req types.UploadRequest) (resp *types.UploadResponse, err error) {
	// todo: add your logic here and delete this line
	var maps = make(map[string]map[string]string)
	maps["xxx"]["xxx1"]= "zhen"
	return
}
