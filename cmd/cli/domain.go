package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "Manage domains within a program",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("domain called")
	},
}

func init() {
	domainCmd.PersistentFlags().StringVar(&programName, "program", "", "Bug bounty program")
}
