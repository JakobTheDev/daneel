package subdomain

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DomainName string
var ProgramName string

var SubdomainCmd = &cobra.Command{
	Use:   "subdomain",
	Short: "Manage subdomains",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("subdomain called")
	},
}

func init() {
}
