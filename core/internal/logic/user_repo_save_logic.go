package logic

import (
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"cloud-disk/core/utils"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepoSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepoSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepoSaveLogic {
	return &UserRepoSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepoSaveLogic) UserRepoSave(req *types.UserRepoSaveRequest, userIdentity string) (resp *types.UserRepoSaveReply, err error) {
	// todo: add your logic here and delete this line
	ur := &models.UserRepository{
		Identity:           utils.GetUUID(),
		UserIdentity:       userIdentity,
		ParentId:           int64(req.Parent_id),
		RepositoryIdentity: req.Repository_identity,
		Ext:                req.Ext,
		Name:               req.Name,
	}
	_, err = l.svcCtx.Engine.Insert(ur)
	if err != nil {
		return
	}
	return &types.UserRepoSaveReply{}, nil
	return
}
