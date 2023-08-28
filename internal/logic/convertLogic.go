package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/YiZou89/shortener/internal/svc"
	"github.com/YiZou89/shortener/internal/types"
	"github.com/YiZou89/shortener/model"
	"github.com/YiZou89/shortener/pkg/base62"
	"github.com/YiZou89/shortener/pkg/connect"
	"github.com/YiZou89/shortener/pkg/md5"
	"github.com/YiZou89/shortener/pkg/urltool"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ConvertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConvertLogic) Convert(req *types.ConvertRequest) (resp *types.ConvertResponse, err error) {
	// todo: add your logic here and delete this line
	// 1. validator
	// not nil, invalid long url,
	if ok := connect.Get(req.LongUrl); !ok {
		return nil, errors.New("invalid long url")
	}

	// check if already convert(query from database),
	md5Val := md5.Sum([]byte(req.LongUrl))
	u, err := l.svcCtx.ShortUrlMapModel.FindOneByMd5(l.ctx, sql.NullString{String: md5Val, Valid: true})
	if err != sqlx.ErrNotFound {
		if err == nil {
			if err := l.svcCtx.Filter.Add([]byte(u.Surl.String)); err != nil {
				logx.Errorw("add short url into bloom filter failed",
					logx.Field("err", err),
				)
				return nil, err
			}
			return nil, fmt.Errorf("already convert: %s", u.Surl.String)
		}
		logx.Errorw("find one by md5 failed",
			logx.Field("err", err),
		)
		return nil, err
	}

	// input can not be a short url
	basePath, err := urltool.GetBasePath(req.LongUrl)
	if err != nil {
		logx.Errorw("parse url failed",
			logx.Field("err", err),
			logx.Field("long url", req.LongUrl),
		)
		return nil, err
	}
	_, err = l.svcCtx.ShortUrlMapModel.FindOneBySurl(l.ctx, sql.NullString{
		String: basePath,
		Valid:  true,
	})
	if err != sqlx.ErrNotFound {
		if err == nil {
			return nil, errors.New("this url is already a short link")
		}
		logx.Errorw("find one by surl failed",
			logx.Field("err", err),
			logx.Field("surl", basePath),
		)
		return nil, err
	}

	var short string
	for {
		// 2. generate
		seq, err := l.svcCtx.Sequence.Next()
		if err != nil {
			logx.Errorw("sequence gen next id failed",
				logx.Field("err", err),
			)
			return nil, err
		}
		logx.Infow("gen next id success",
			logx.Field("seq", seq),
		)

		// 3. convert
		short = base62.Int2String(seq)
		if _, ok := l.svcCtx.ShortUrlBlackList[short]; !ok {
			break
		}

	}

	// 4. store
	if err := l.svcCtx.Filter.Add([]byte(short)); err != nil {
		logx.Errorw("add short url into bloom filter failed",
			logx.Field("err", err),
		)
		return nil, err
	}
	_, err = l.svcCtx.ShortUrlMapModel.Insert(l.ctx, &model.ShortUrlMap{
		Lurl: sql.NullString{String: req.LongUrl, Valid: true},
		Surl: sql.NullString{String: short, Valid: true},
		Md5:  sql.NullString{String: md5Val, Valid: true},
	})
	if err != nil {
		logx.Errorw("store short url failed",
			logx.Field("err", err),
		)
		return nil, err
	}
	

	resp = new(types.ConvertResponse)
	resp.ShortUrl = l.svcCtx.Config.ShortDomain + "/" + short
	return resp, nil
}
