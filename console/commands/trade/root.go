package trade

import (
"github.com/spf13/cobra"
)

func NewTradeCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "trade",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	//
	command.AddCommand(NewBinanceCommand())

	return command
}
