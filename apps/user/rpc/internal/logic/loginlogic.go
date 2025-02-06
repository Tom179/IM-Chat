package logic

import (
	"Im-chat/Chat/apps/user/models"
	"Im-chat/Chat/apps/user/rpc/internal/svc"
	"Im-chat/Chat/apps/user/rpc/user"
	"Im-chat/Chat/pkg/ctxdata"
	"Im-chat/Chat/pkg/encrypt"
	"context"
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrPhoneNotRegister = errors.New("手机号未注册")
	ErrPwdError         = errors.New("密码不正确")
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	userEntity, err := l.svcCtx.UsersModel.FindByPhone(l.ctx, in.Phone)
	if err != nil {
		if err == models.ErrNotFound { //如果没有找到，会返回models.Err类型，所以不用再去下面判断userEntity是否为空
			return nil, ErrPhoneNotRegister
		}
		return nil, err
	}

	if !encrypt.ValidatePasswordHash(in.Password, userEntity.Password.String) {
		return nil, ErrPwdError
	}

	iat := time.Now().Unix()
	jwt, err := ctxdata.GetJwtToken(l.svcCtx.Config.Jwt.AccessKey, iat, l.svcCtx.Config.Jwt.AccessExpire, userEntity.Id)
	if err != nil {
		return nil, err
	}

	//fmt.Println("expireAt:   ", iat+l.svcCtx.Config.Jwt.AccessExpire)
	return &user.LoginResp{Token: jwt, Expire: iat + l.svcCtx.Config.Jwt.AccessExpire}, nil
}
