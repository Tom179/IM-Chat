package xerr

import "github.com/zeromicro/x/errors"

func New(code int, msg string) error { //创建grpcError
	return errors.New(code, msg)
}

func NewDBError() error {
	return errors.New(DB_ERROR, ErrMsg(DB_ERROR))
}

func NewInternalErr() error {
	return errors.New(SERVER_COMMON_ERROR, ErrMsg(SERVER_COMMON_ERROR))
}
 