package user

import (
	"context"

	"Im-chat/Chat/apps/user/api/internal/svc"
	"Im-chat/Chat/apps/user/api/internal/types"
	"Im-chat/Chat/apps/user/rpc/user"
	"Im-chat/Chat/pkg/ctxdata"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line
	uid := ctxdata.GetUid(l.ctx)
	userInfoResp, err := l.svcCtx.User.GetUserInfo(l.ctx, &user.GetUserInfoReq{
		Id: uid,
	})

	if err != nil {
		return nil, err
	}

	var res types.User
	copier.Copy(&res, userInfoResp.User)

	return &types.UserInfoResp{Info: res}, nil

}
