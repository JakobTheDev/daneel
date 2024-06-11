package cli

import (
	"log"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/JakobTheDev/daneel/internal/subdomain"
	"github.com/spf13/cobra"
)

var subdomainAddCmd = &cobra.Command{
	Use:   "add [subdomain]",
	Short: "Add a subdomain to daneel",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := subdomain.AddSubdomain(models.Subdomain{
			Name:       args[0],
			IsInScope:  !isOutOfScope,
			DomainName: domainName})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	subdomainCmd.AddCommand(subdomainAddCmd)

	subdomainAddCmd.Flags().StringVar(&domainName, "domain", "", "The subdomain's parent domain")
	subdomainAddCmd.Flags().BoolVar(&isOutOfScope, "no-scope", false, "Mark the subdomain as out of scope (default false)")
	subdomainAddCmd.MarkFlagRequired("domain")
}
