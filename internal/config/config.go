package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	ShortUrlDB ShortUrlDB

	Sequence struct {
		DSN string
	}

	CacheRedis cache.CacheConf
	

	BaseString string

	ShortUrlBlackList []string

	ShortDomain string

	
}

type ShortUrlDB struct {
	DSN string
}
