package logic

import (
	"context"
	"database/sql"
	"time"

	"Im-chat/Chat/apps/social/rpc/internal/svc"
	"Im-chat/Chat/apps/social/rpc/social"
	"Im-chat/Chat/apps/social/socialmodels"
	"Im-chat/Chat/pkg/constants"
	"Im-chat/Chat/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendPutInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInLogic {
	return &FriendPutInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendPutInLogic) FriendPutIn(in *social.FriendPutInReq) (*social.FriendPutInResp, error) {
	//查询是否已经存在好友关系
	friends, err := l.svcCtx.FriendsModel.FindByUidAndFid(l.ctx, in.UserId, in.ReqUid)
	if err != nil && err != socialmodels.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewDBError(), "根据uid和fid查询Friend关系表出错 %v req %v", err, in)
	}
	if friends != nil {
		return &social.FriendPutInResp{}, err
	}

	//查询是否存在申请记录
	friendReqs, err := l.svcCtx.FriendRequestsModel.FindByReqUidAndUserId(l.ctx, in.ReqUid, in.UserId)
	if err != nil && err != socialmodels.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewDBError(), "根据uid和rid查询Friend_requests添加记录出错 %v req %v", err, in)
	}
	if friendReqs != nil {
		return &social.FriendPutInResp{}, err
	}

	_, err = l.svcCtx.FriendRequestsModel.Insert(l.ctx, &socialmodels.FriendRequests{
		UserId: in.UserId,
		ReqUid: in.ReqUid,
		ReqMsg: sql.NullString{
			Valid:  true,
			String: in.ReqMsg,
		},
		ReqTime: time.Unix(in.ReqTime, 0),
		HandleResult: sql.NullInt64{
			Int64: int64(constants.NoHandlerResult),
			Valid: true,
		},
	})

	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBError(), "新增好友申请错误%v %v", err, in)
	}

	return &social.FriendPutInResp{}, nil
}
