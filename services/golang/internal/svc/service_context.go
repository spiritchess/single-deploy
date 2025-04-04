package svc

import (
	"golang/internal/config"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config   config.Config
	Database sqlx.SqlConn
	Redis    redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	database := sqlx.NewMysql(c.DatabaseDsn)
	rds := redis.MustNewRedis(c.RedisConf)
	return &ServiceContext{
		Config:   c,
		Database: database,
		Redis:    *rds,
	}
}
