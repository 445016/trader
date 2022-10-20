package trade

import (
	"context"
	"flag"
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
	logger "github.com/dirname/binance/logging"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
	"time"
	fconfig "trader/internal/config"
)

var configFile = flag.String("f", "etc/trader-api.yaml", "the config file")


func NewBinanceBuyCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "buy",
		Short: "trade buy",
		Run: func(cmd *cobra.Command, args []string) {
			flag.Parse()

			var c fconfig.Config
			conf.MustLoad(*configFile, &c)
			logger.Enable(false)
			for {

			futuresClient := binance.NewFuturesClient(c.AppKey, c.AppSecret)

			starttime := time.Now().UnixNano() / 1e6
			order,err := futuresClient.NewCreateOrderService().Symbol("BTCUSDT").Side(futures.SideTypeSell).Type(futures.OrderTypeMarket).Quantity("0.001").PositionSide(futures.PositionSideTypeShort).Do(context.Background())
			//tradeClient = spotclient.NewTradeClient(c.SpotRestHost, c.AppKey, c.AppSecret)
			endtime := time.Now().UnixNano() / 1e6
			fmt.Printf("订单号:%v, 下单响应时长:%d ms,错误:%v", order.OrderID, endtime-starttime, err)

			}

		},
	}

	return command
}


