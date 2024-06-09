package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var programCmd = &cobra.Command{
	Use:   "program",
	Short: "Manage bug bounty programs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("program called")
	},
}

func init() {
	programCmd.Flags().StringVarP(&platformName, "platform", "p", "", "Bug bounty platform")
}
