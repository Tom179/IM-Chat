package rpcserver

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	zerr "github.com/zeromicro/x/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// 错误日志拦截器
func LogInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	resp, err = handler(ctx, req)
	if err == nil {
		return resp, nil
	}

	logx.WithContext(ctx).Errorf("【RPC SRV ERR】 $v", err)

	//logic中实现的业务中error.Wrapf返回的是将grpc错误转换为了标准错误，需要我们解析回grpc错误
	causeErr := errors.Cause(err)
	if e, ok := causeErr.(*zerr.CodeMsg); ok { //如果是，说明是我们自定义的错误
		err = status.Error(codes.Code(e.Code), e.Msg)
	}

	return resp, err
}
