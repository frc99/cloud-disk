package logic

import (
	"cloud-disk/core/define"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, user_Identity string) (resp *types.UserFileListReply, err error) {
	// todo: add your logic here and delete this line
	uf := make([]*types.UserFile, 0)
	var cnt int64
	resp = new(types.UserFileListReply)
	size := req.Size
	if size == 0 {
		size = define.PageSize
	}
	page := req.Page
	//从第一页开始
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * size

	ur := new(models.UserRepository)
	//l.svcCtx.Engine,ShowSQL(true)
	err = l.svcCtx.Engine.Table("user_repository").Where("parent_id = ? AND user_identity = ? ", ur.Id, user_Identity).
		Select("user_repository.id, user_repository.identity, user_repository.repository_identity, user_repository.ext,"+
			"user_repository.name, repository_pool.path, repository_pool.size").
		Join("LEFT", "repository_pool", "user_repository.repository_identity = repository_pool.identity").
		Where("user_repository.deleted_at = ? OR user_repository.deleted_at IS NULL", time.Time{}.Format(define.Datetime)).
		Limit(size, offset).Find(&uf)
	if err != nil {
		return
	}
	resp.List = uf
	resp.Count = cnt
	// 查询用户文件总数
	cnt, err = l.svcCtx.Engine.Where("parent_id = ? AND user_identity = ? ", ur.Id, user_Identity).Count(new(models.UserRepository))
	if err != nil {
		return
	}
	resp.List = uf
	resp.Count = cnt
	return
}
