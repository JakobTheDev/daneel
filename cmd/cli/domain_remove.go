package cli

import (
	"log"
	"strings"

	"github.com/JakobTheDev/daneel/internal/domain"
	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/spf13/cobra"
)

var domainRemoveCmd = &cobra.Command{
	Use:   "remove [domain]",
	Short: "Remove a domain from a program",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		domainName := strings.ToLower(args[0])

		err := domain.RemoveDomain(models.Domain{Name: domainName})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	domainCmd.AddCommand(domainRemoveCmd)
}
