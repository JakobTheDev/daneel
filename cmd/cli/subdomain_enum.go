package cli

import (
	"fmt"
	"log"

	"github.com/JakobTheDev/daneel/internal/domain"
	"github.com/JakobTheDev/daneel/internal/subdomain"
	"github.com/spf13/cobra"
)

var subdomainEnumCmd = &cobra.Command{
	Use:   "enum [domain]",
	Short: "Enumerate subdomains for a domain",
	Run: func(cmd *cobra.Command, args []string) {

		if programName != "" {
			domains, err := domain.ListDomains(programName, false)
			if err != nil {
				log.Fatal(err)
			}
			for _, d := range domains {
				enumerateSubdomainsByDomain(d.Name)
			}
		} else if domainName != "" {
			enumerateSubdomainsByDomain(domainName)
		} else {
			fmt.Println("No domain or program specified")
		}
	},
}

func enumerateSubdomainsByDomain(domainName string) {
	_, newSubdomains, err := subdomain.EnumerateSubdomainsByDomain(domainName)
	if err != nil {
		log.Fatal(err)
	}

	if len(newSubdomains) > 0 {
		fmt.Printf("%d new subdomains found:\n", len(newSubdomains))
		for _, s := range newSubdomains {
			fmt.Println(s)
		}
	} else {
		fmt.Println("No new subdomains found")
	}

}

func init() {
	subdomainCmd.AddCommand(subdomainEnumCmd)

	subdomainEnumCmd.Flags().StringVar(&programName, "program", "", "A bug bounty program to enumerate subdomains for")
	subdomainEnumCmd.Flags().StringVar(&domainName, "domain", "", "The domain to enumerate (must be tracked by daneel)")
	subdomainEnumCmd.MarkFlagsOneRequired("program", "domain")
	subdomainEnumCmd.MarkFlagsMutuallyExclusive("program", "domain")
}
