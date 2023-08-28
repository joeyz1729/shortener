package logic

import (
	"context"
	"database/sql"
	"errors"

	"github.com/YiZou89/shortener/internal/svc"
	"github.com/YiZou89/shortener/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrFilterNotFound = errors.New("[bloom filter] short url does not exist")
	ErrNotFound = errors.New("[redis mysql] short url does not exist")
)


type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.ShowRequest) (resp *types.ShowResponse, err error) {
	// todo: add your logic here and delete this line
	ok, err := l.svcCtx.Filter.Exists([]byte(req.ShortUrl))
	if err != nil {
		logx.Errorw("bloom filter failed",
			logx.Field("err", err),
		)
		return nil, err
	}
	if !ok {
		return nil, ErrFilterNotFound
	}
	u, err := l.svcCtx.ShortUrlMapModel.FindOneBySurl(l.ctx, sql.NullString{String: req.ShortUrl, Valid: true})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		logx.Errorw("find one by surl failed",
			logx.Field("err", err),
		)
		return nil, err
	}

	return &types.ShowResponse{LongUrl: u.Lurl.String}, nil
}
