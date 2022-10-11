package logic

import (
	"context"
	"fmt"

	"trader/internal/svc"
	"trader/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TradingviewHookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTradingviewHookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TradingviewHookLogic {
	return &TradingviewHookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TradingviewHookLogic) TradingviewHook(req *types.TradingViewHookRequest) (resp *types.TradingViewHookResponse, err error) {
	fmt.Println("hook")

	return
}
