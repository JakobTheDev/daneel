package cli

import (
	"fmt"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/spf13/cobra"
)

var subdomainAddCmd = &cobra.Command{
	Use:   "add [subdomain]",
	Short: "Add a subdomain to daneel",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err, _ := models.AddSubdomain(models.Subdomain{
			SubdomainName: args[0],
			IsInScope:     !isOutOfScope,
			DomainName:    domainName})
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	subdomainCmd.AddCommand(subdomainAddCmd)

	subdomainAddCmd.Flags().StringVar(&domainName, "domain", "", "The subdomain's parent domain")
	subdomainAddCmd.Flags().BoolVar(&isOutOfScope, "no-scope", false, "Mark the subdomain as out of scope (default false)")
	subdomainAddCmd.MarkFlagRequired("domain")
}
