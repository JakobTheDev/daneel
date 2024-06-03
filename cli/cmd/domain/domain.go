package domain

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ProgramName string

var DomainCmd = &cobra.Command{
	Use:   "domain [command]",
	Short: "Manage domains within daneel",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("domain called")
	},
}

func init() {
	DomainCmd.PersistentFlags().StringVar(&ProgramName, "program", "", "Bug bounty program")
}
