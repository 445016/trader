package commands

import (
	"github.com/spf13/cobra"
	"trader/console/commands/market"
	"trader/console/commands/trade"

)

const (
	// CLIName is the name of the CLI
	CLIName = "trader"
)

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   CLIName,
		Short: "",
		Long: `trader命令行脚本`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	var logLevel string
	var glogLevel int
	var verbose bool
	command.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if verbose {
			logLevel = "debug"
			glogLevel = 6
		}
	}
	command.AddCommand(market.NewMarketCommand())
	command.AddCommand(trade.NewTradeCommand())
	command.PersistentFlags().StringVar(&logLevel, "loglevel", "info", "Set the logging level. One of: debug|info|warn|error")
	command.PersistentFlags().IntVar(&glogLevel, "gloglevel", 0, "Set the glog logging level")
	command.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enabled verbose logging, i.e. --loglevel debug")

	return command
}
