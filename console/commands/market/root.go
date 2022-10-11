package market

import (
	"github.com/spf13/cobra"
)

func NewMarketCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "market",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	//
	command.AddCommand(NewFetchCommand())
	//command.AddCommand(NewListCommand())
	//command.AddCommand(NewCreateCommand())
	//command.AddCommand(NewDeleteCommand())
	//command.AddCommand(NewLintCommand())
	//command.AddCommand(NewSuspendCommand())
	//command.AddCommand(NewResumeCommand())

	return command
}
