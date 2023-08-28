package logic

import (
	"context"

	"github.com/YiZou89/shortener/internal/svc"
	"github.com/YiZou89/shortener/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	// not nil, invalid log url, check if already convert(query from database), input can not be a short url
	



	
	// 2. generate

	// 3. convert

	// 4. store

	return
}
