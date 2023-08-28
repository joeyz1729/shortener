package sequence

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Redis struct {
	Addr       string
	RedisCache *redis.Redis
}

func NewRedis(redisAddr string) Sequence {
	rdb, err := redis.NewRedis(redis.RedisConf{
		Host: redisAddr,
	})
	if err != nil {
		logx.Errorw("new redis failed",
			logx.Field("err", err),
		)
		return nil
	}
	return &Redis{
		Addr:       redisAddr,
		RedisCache: rdb,
	}
}

func (r *Redis) Next() (seq uint64, err error) {
	return seq, nil
}
