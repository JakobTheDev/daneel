package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var subdomainCmd = &cobra.Command{
	Use:   "subdomain",
	Short: "Manage subdomains",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("subdomain called")
	},
}

func init() {
}
