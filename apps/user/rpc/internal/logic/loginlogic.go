package logic

import (
	"Im-chat/Chat/apps/user/models"
	"Im-chat/Chat/apps/user/rpc/internal/svc"
	"Im-chat/Chat/apps/user/rpc/user"
	"Im-chat/Chat/pkg/ctxdata"
	"Im-chat/Chat/pkg/encrypt"
	"Im-chat/Chat/pkg/xerr"
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var ( //业务型错误常量
	ErrPhoneNotRegister = xerr.New(xerr.SERVER_COMMON_ERROR, "手机号未注册") //定义go-Zero错误
	ErrPwdError         = xerr.New(xerr.SERVER_COMMON_ERROR, "密码不正确")
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
			return nil, errors.WithStack(ErrPhoneNotRegister) //业务型错误，本来就有描述信息
		}
		return nil, errors.Wrapf(xerr.NewDBError(), "根据手机查询用户错误 %v, 请求 %v", err, in.Phone) //意外型错误需要记录好原因
	}

	if !encrypt.ValidatePasswordHash(in.Password, userEntity.Password.String) {
		return nil, errors.WithStack(ErrPwdError)
	}

	iat := time.Now().Unix()
	jwt, err := ctxdata.GetJwtToken(l.svcCtx.Config.Jwt.AccessKey, iat, l.svcCtx.Config.Jwt.AccessExpire, userEntity.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewInternalErr(), "生成jwt错误 %v", err)
	}

	//fmt.Println("expireAt:   ", iat+l.svcCtx.Config.Jwt.AccessExpire)
	return &user.LoginResp{Token: jwt, Expire: iat + l.svcCtx.Config.Jwt.AccessExpire}, nil
}
