package logic

import (
	"Im-chat/Chat/apps/user/models"
	"Im-chat/Chat/pkg/ctxdata"
	"Im-chat/Chat/pkg/encrypt"
	"Im-chat/Chat/pkg/wuid"
	"context"
	"database/sql"
	"errors"
	"time"

	"Im-chat/Chat/apps/user/rpc/internal/svc"
	"Im-chat/Chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var (
	ErrPhoneRegisted = errors.New("手机号码已注册")
)

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {

	userEntity, err := l.svcCtx.UsersModel.FindByPhone(l.ctx, in.Phone)
	if err != nil && err != models.ErrNotFound {
		return nil, err
	}

	if userEntity != nil { //用户已存在，注册失败
		return nil, ErrPhoneRegisted
	}

	userEntity = &models.Users{
		Id:       wuid.GenUid(l.svcCtx.Config.Mysql.Datasource),
		Avatar:   in.Avatar,
		Nickname: in.Nickname,
		Phone:    in.Phone,
		Gender: sql.NullInt64{
			Int64: int64(in.Gender),
			Valid: true,
		},
		//todo status、CreateAt和UpdateAt什么时候写如？
	}

	if len(in.Password) > 0 { //密码加密
		genPass, err := encrypt.GenPasswordHash([]byte(in.Password))
		if err != nil {
			return nil, err
		}
		userEntity.Password = sql.NullString{
			string(genPass),
			true,
		}
	}

	_, err = l.svcCtx.UsersModel.Insert(l.ctx, userEntity)
	if err != nil {
		return nil, err
	}
	now := time.Now().Unix()
	token, err := ctxdata.GetJwtToken(l.svcCtx.Config.Jwt.AccessKey, now, l.svcCtx.Config.Jwt.AccessExpire, userEntity.Id)
	if err != nil {
		return nil, err
	}

	return &user.RegisterResp{Token: token, Expire: now + l.svcCtx.Config.Jwt.AccessExpire}, nil
}
