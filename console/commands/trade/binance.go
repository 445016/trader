package trade

import (
	"github.com/spf13/cobra"
)

func NewBinanceCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "binance",
		Short: "binance trade",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	command.AddCommand(NewBinanceBuyCommand())
	command.AddCommand(NewBinanceKlineCommand())

	return command
}
