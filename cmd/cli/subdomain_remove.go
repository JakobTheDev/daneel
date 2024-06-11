package cli

import (
	"log"
	"strings"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/JakobTheDev/daneel/internal/subdomain"
	"github.com/spf13/cobra"
)

var subdomainRemoveCmd = &cobra.Command{
	Use:   "remove [subdomain]",
	Short: "Remove a subdomain",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		subdomainName := strings.ToLower(args[0])

		err := subdomain.RemoveSubdomain(models.Subdomain{Name: subdomainName})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	subdomainCmd.AddCommand(subdomainRemoveCmd)
}
