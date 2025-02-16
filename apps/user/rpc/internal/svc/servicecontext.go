package svc

import (
	"Im-chat/Chat/apps/user/models"
	"Im-chat/Chat/apps/user/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

//微服务配置初始化模块

type ServiceContext struct {
	Config config.Config  
	models.UsersModel 
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)

	return &ServiceContext{
		Config: c,
		//todo 这个mysql的缓存到底是什么？
		UsersModel: models.NewUsersModel(sqlConn, c.Cache),
	}
}
