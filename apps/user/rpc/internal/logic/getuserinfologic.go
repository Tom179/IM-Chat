package logic

import (
	"context"
	"errors"

	"Im-chat/Chat/apps/user/models"
	"Im-chat/Chat/apps/user/rpc/internal/svc"
	"Im-chat/Chat/apps/user/rpc/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUserNotFound = errors.New("用户不存在")

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line
	one, err := l.svcCtx.UsersModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == models.ErrNotFound { //用户不存在
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}

	// fmt.Println("查询到用户为:", one)

	var respUser user.UserEntity
	copier.Copy(&respUser, one)

	return &user.GetUserInfoResp{User: &respUser}, nil
}
