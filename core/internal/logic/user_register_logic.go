package logic

import (
	"cloud-disk/core/models"
	"cloud-disk/core/utils"
	"context"
	"errors"
	"github.com/gomodule/redigo/redis"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterReply, err error) {
	//判断code是否一致
	val, err := models.Rdb.Do("get", req.Email)
	code, err := redis.String(val, nil)
	if err != nil {
		return nil, errors.New("未获取邮箱验证码")
	}
	if code != req.Code {
		err = errors.New(req.Code + "验证码错误" + code)
		return
	}
	//判断用户名是否已存在
	cnt, err := l.svcCtx.Engine.Where("name = ?", req.Name).Count(new(models.UserBasic))
	if err != nil {
		return
	}
	if cnt > 0 {
		err = errors.New("用户已存在")
		return
	}
	//数据入库
	user := &models.UserBasic{
		Identity: utils.GetUUID(),
		Name:     req.Name,
		Password: utils.Md5(req.Password),
		Email:    req.Email,
	}
	//返回了插入的个数
	_, err = l.svcCtx.Engine.Insert(user)
	if err != nil {
		return nil, errors.New("没有插入成功")
	}

	return
}
