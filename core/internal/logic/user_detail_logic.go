package logic

import (
	"cloud-disk/core/models"
	"context"
	"errors"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (resp *types.UserDetailReply, err error) {
	resp = &types.UserDetailReply{}
	userbasic := new(models.UserBasic)
	ok, err := models.Engine.Where("identity = ?", req.Identity).Get(userbasic)

	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("用户未找到")
	}
	resp.Name = userbasic.Name
	resp.Email = userbasic.Email
	return
}
