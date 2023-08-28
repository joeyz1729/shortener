package svc

import (
	"github.com/YiZou89/shortener/internal/config"
	"github.com/YiZou89/shortener/model"
	"github.com/YiZou89/shortener/sequence"
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	ShortUrlMapModel model.ShortUrlMapModel

	Sequence sequence.Sequence

	ShortUrlBlackList map[string]struct{}

	Filter *bloom.Filter // memory version bloom filter
}

func NewServiceContext(c config.Config) *ServiceContext {
	// mysql db initialize
	conn := sqlx.NewMysql(c.ShortUrlDB.DSN)

	// black list initialize
	m := make(map[string]struct{}, len(c.ShortUrlBlackList))
	for _, v := range c.ShortUrlBlackList {
		m[v] = struct{}{}
	}

	// bloom filter initialize
	store := redis.New(c.CacheRedis[0].Host, func(r *redis.Redis) {
		r.Type = redis.NodeType
	})
	filter := bloom.New(store, "bloom", 20*1<<20)

	return &ServiceContext{
		Config:            c,
		ShortUrlMapModel:  model.NewShortUrlMapModel(conn, c.CacheRedis),
		Sequence:          sequence.NewMySQL(c.Sequence.DSN),
		ShortUrlBlackList: m,
		Filter:            filter,
	}
}

func loadDataToBloom() {
	
}
