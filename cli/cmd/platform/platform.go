package platform

import (
	"fmt"

	"github.com/spf13/cobra"
)

var PlatformCmd = &cobra.Command{
	Use:   "platform [command]",
	Short: "Manage bug bounty platforms within daneel",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("platform called")
	},
}

func init() {

}
