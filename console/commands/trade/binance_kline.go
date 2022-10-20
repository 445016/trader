package trade

import (
	"flag"
	"fmt"
	"github.com/adshao/go-binance/v2"
	logger "github.com/dirname/binance/logging"
	"github.com/fatih/color"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
	"strconv"
	fconfig "trader/internal/config"
)


func NewBinanceKlineCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "Kline",
		Short: "Kline Data",
		Run: func(cmd *cobra.Command, args []string) {
			flag.Parse()

			var c fconfig.Config
			conf.MustLoad(*configFile, &c)
			logger.Enable(false)
			var LastVolume,LastBuyVolume float64
			var LastSellVolume decimal.Decimal
			wsKlineHandler := func(event *binance.WsKlineEvent) {
				//fmt.Println(event.Kline)
				fmt.Println(event.Kline)

				FVolume, _ := strconv.ParseFloat(event.Kline.Volume, 64)
				fActiveBuyVolume, _ := strconv.ParseFloat(event.Kline.ActiveBuyVolume, 64)
				fmt.Printf("交易对：%s\r\n",event.Kline.Symbol)
				fmt.Printf("K线间隔：%s\r\n",event.Kline.Interval)
				fmt.Printf("K线起始时间：%d\r\n",event.Kline.StartTime / 1e3)
				fmt.Printf("K线结束时间：%d\r\n",event.Kline.EndTime / 1e3)
				fmt.Printf("K线成交笔数：%d\r\n",event.Kline.TradeNum)
				fmt.Printf("K线成交量：%s\r\n",event.Kline.Volume)
				fmt.Printf("单笔成交量：%f\r\n",FVolume - LastVolume)
				SellVolume := decimal.NewFromFloat(FVolume).Sub(decimal.NewFromFloat(fActiveBuyVolume))
				FActiveBuyVolume, _ := strconv.ParseFloat(event.Kline.ActiveBuyVolume, 64)
				ColorPrintF := color.New()
				ColorPrintF.Add(color.Bold)
				ColorPrintF.Add(color.FgGreen)
				fmt.Printf("购买量：%f\r\n",FActiveBuyVolume)
				ColorPrintF.Printf("单笔购买量：%f\r\n",FActiveBuyVolume - LastBuyVolume)
				fmt.Printf("卖出量：%s\r\n",SellVolume)
				fmt.Printf("单笔卖出量：%s\r\n",SellVolume.Sub(LastSellVolume))
				fClose,_ := strconv.ParseFloat(event.Kline.Close, 64)
				fmt.Printf("总金额：%f\r\n",FVolume * fClose)
				fmt.Printf("主动购买金额：%s\r\n",event.Kline.ActiveBuyQuoteVolume)
				if event.Kline.IsFinal == true {
					LastVolume  = 0
					LastBuyVolume = 0
					LastSellVolume = decimal.NewFromFloat(0)
				} else {
					LastVolume = FVolume
					LastBuyVolume = FActiveBuyVolume
					LastSellVolume = SellVolume
				}
				//fmt.Printf("累计主动购买量：%f\r\n",ActiveBuyVolumeTotal)
				//SellVolume, _ := strconv.ParseInt(event.Kline.ActiveBuyVolume, 10,64)
				//LastTradeNum = event.Kline.TradeNum
			}
			errHandler := func(err error) {
				fmt.Println(err)
			}

			doneC, _, err := binance.WsKlineServe("BTCUSDT", "1m", wsKlineHandler, errHandler)
			if err != nil {
				fmt.Println(err)
				return
			}
			<-doneC
		},
	}

	return command
}


