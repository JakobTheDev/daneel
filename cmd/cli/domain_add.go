package cli

import (
	"log"

	"github.com/JakobTheDev/daneel/internal/domain"
	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/spf13/cobra"
)

var addDomainCmd = &cobra.Command{
	Use:   "add [domain]",
	Short: "Add a domain to a program",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := domain.AddDomain(models.Domain{
			Name:        args[0],
			IsInScope:   !isOutOfScope,
			ProgramName: programName})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	domainCmd.AddCommand(addDomainCmd)

	addDomainCmd.Flags().BoolVar(&isOutOfScope, "no-scope", false, "Mark the domain as out of scope (default false)")
	addDomainCmd.MarkPersistentFlagRequired("program")
}
