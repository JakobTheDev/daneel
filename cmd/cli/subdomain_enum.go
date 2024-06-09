package cli

import (
	"fmt"
	"log"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/JakobTheDev/daneel/internal/tools"
	"github.com/spf13/cobra"
)

var subdomainEnumCmd = &cobra.Command{
	Use:   "enum [domain]",
	Short: "Enumerate subdomains for a domain",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			err        error
			isInserted bool
		)

		if programName != "" {
			// Check the program exists
			// Get domains for the program
			// for each, enumerate subdomains
			// Insert into database
			// Handle results
		} else if domainName != "" {
			// Check the domain exists
			var domain models.Domain
			domain, err = models.GetDomain(domainName)
			if err != nil {
				log.Fatalf("Failed to get domain. Make sure it's been added to a program.")
				log.Fatal(err)
			}
			// Enumerate subdomains for the domain
			var subdomains []string
			subdomains, err = tools.RunSubfinder(domain.DomainName)
			if err != nil {
				log.Fatalf("Error running subfinder: %v", err)
			}
			if len(subdomains) == 0 {
				log.Println("No subdomains found")
				return
			}
			log.Printf("Found %d subdomains\n", len(subdomains))
			// Insert into database
			var newSubdomains []string
			for _, subdomain := range subdomains {
				err, isInserted = models.AddSubdomain(models.Subdomain{DomainName: domain.DomainName, SubdomainName: subdomain})
				if err != nil {
					log.Fatalf("Error adding subdomain to database: %v", err)
				}
				if isInserted {
					newSubdomains = append(newSubdomains, subdomain)
				}
			}
			// Handle results
		} else {
			fmt.Println("No domain or program specified")
		}
	},
}

func init() {
	subdomainCmd.AddCommand(subdomainEnumCmd)

	subdomainEnumCmd.Flags().StringVar(&programName, "program", "", "A bug bounty program to enumerate subdomains for")
	subdomainEnumCmd.Flags().StringVar(&domainName, "domain", "", "The domain to enumerate (must be tracked by daneel)")
	subdomainEnumCmd.MarkFlagsOneRequired("program", "domain")
	subdomainEnumCmd.MarkFlagsMutuallyExclusive("program", "domain")
}
