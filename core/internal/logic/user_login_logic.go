package logic

import (
	"cloud-disk/core/define"
	"cloud-disk/core/models"
	"cloud-disk/core/utils"
	"context"
	"errors"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line
	//1、从数据库里查询当前用户
	user := new(models.UserBasic)
	ok, err := l.svcCtx.Engine.Where("name = ? AND password = ?", req.Name, utils.Md5(req.Password)).Get(user)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("用户名或密码错误")
	}
	//2、生成token
	token, err := utils.GenerateToken(user.Id, user.Identity, user.Name, define.TokenExpire)
	if err != nil {
		return nil, err
	}
	//3、生成用于刷新token的refreshToken
	refreshToken, err := utils.GenerateToken(user.Id, user.Identity, user.Name, define.RefreshTokenExpire)
	resp = new(types.LoginReply)
	resp.Token = token
	resp.RefreshToken = refreshToken
	return
}
