package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var platformCmd = &cobra.Command{
	Use:   "platform [command]",
	Short: "Manage bug bounty platforms",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("platform called")
	},
}

func init() {

}
