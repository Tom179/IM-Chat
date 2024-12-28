package logic

import (
	"context"

	"Im-chat/Chat/apps/user/rpc/internal/svc"
	"Im-chat/Chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUsersLogic {
	return &FindUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUsersLogic) FindUsers(in *user.FindUserReq) (*user.FindUserResp, error) {
	// todo: add your logic here and delete this line

	return &user.FindUserResp{}, nil
}
