package program

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Command flags
var Platform string

var ProgramCmd = &cobra.Command{
	Use:   "program",
	Short: "Manage bug bounty programs within daneel",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("program called")
	},
}

func init() {
	ProgramCmd.Flags().StringVarP(&Platform, "platform", "p", "", "Bug bounty platform")
}
