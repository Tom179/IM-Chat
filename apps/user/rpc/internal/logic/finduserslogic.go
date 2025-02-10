package logic

import (
	"context"
	"fmt"

	"Im-chat/Chat/apps/user/models"
	"Im-chat/Chat/apps/user/rpc/internal/svc"
	"Im-chat/Chat/apps/user/rpc/user"
	"Im-chat/Chat/pkg/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
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

	var (
		userEntitys []*models.Users
		err         error
	)
	if in.Phone != "" {
		userEntity, err := l.svcCtx.UsersModel.FindByPhone(l.ctx, in.Phone)
		if err == nil {
			userEntitys = append(userEntitys, userEntity)
		}
	} else if in.Name != "" {
		userEntitys, err = l.svcCtx.UsersModel.FindByName(l.ctx, in.Name)
	} else if len(in.Ids) > 0 {
		userEntitys, err = l.svcCtx.FindByIds(l.ctx, in.Ids)
		fmt.Println("根据ids找到userEntityes数量为:", len(userEntitys))
	}
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBError(), "查找用户失败 %v", err)
	}

	var resp []*user.UserEntity
	copier.Copy(&resp, &userEntitys)

	return &user.FindUserResp{User: resp}, nil
}
