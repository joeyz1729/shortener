package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	ShortUrlDB ShortUrlDB

	Sequence struct {
		DSN string
	}

	BaseString string

	ShortUrlBlackList []string

	ShortDomain string
}

type ShortUrlDB struct {
	DSN string
}
