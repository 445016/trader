package market

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewFetchCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "fetch",
		Short: "fetch market",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("sss")
		},
	}

	return command
}
