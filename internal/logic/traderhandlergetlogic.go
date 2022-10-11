package logic

import (
	"context"

	"trader/internal/svc"
	"trader/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TraderHandlerGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTraderHandlerGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TraderHandlerGetLogic {
	return &TraderHandlerGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TraderHandlerGetLogic) TraderHandlerGet(req *types.TradingViewHookRequest) (resp *types.TradingViewHookResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
