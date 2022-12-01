package logic

import (
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"cloud-disk/core/utils"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendRegisterLogic {
	return &MailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req *types.MailCodeRequest) (resp *types.MailCodeReply, err error) {
	//邮箱未被注册
	cnt, err := l.svcCtx.Engine.Where("email = ?", req.Email).Count(new(models.UserBasic))
	if err != nil {
		return
	}
	if cnt > 0 {
		err = errors.New("该邮箱已被注册")
		return
	}
	//获取验证码
	code := utils.RandCode()
	//存储验证码
	//_, err = models.Rdb.Do("set", req.Email, code, "EX", time.Second*time.Duration(define.CodeExpired))
	_, err = models.Rdb.Do("set", req.Email, code, "EX", 200)
	if err != nil {
		errors.New("redis存储失败")
	}
	//发送验证码
	err = utils.MailSendCode(req.Email, code)
	if err != nil {
		return nil, err
	}
	return
}
