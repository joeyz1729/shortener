package svc

import (
	"github.com/YiZou89/shortener/internal/config"
	"github.com/YiZou89/shortener/model"
	"github.com/YiZou89/shortener/sequence"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	ShortUrlMapModel model.ShortUrlMapModel

	Sequence sequence.Sequence

	ShortUrlBlackList map[string]struct{}
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.ShortUrlDB.DSN)
	m := make(map[string]struct{}, len(c.ShortUrlBlackList))
	for _, v := range c.ShortUrlBlackList {
		m[v] = struct{}{}
	}
	return &ServiceContext{
		Config:           c,
		ShortUrlMapModel: model.NewShortUrlMapModel(conn),
		Sequence:         sequence.NewMySQL(c.Sequence.DSN),
		ShortUrlBlackList: m,
	}
}
