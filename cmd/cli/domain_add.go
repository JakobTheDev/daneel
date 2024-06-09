package cli

import (
	"fmt"

	"github.com/JakobTheDev/daneel/internal/domain"
	"github.com/spf13/cobra"
)

var addDomainCmd = &cobra.Command{
	Use:   "add [domain]",
	Short: "Add a domain to a program",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := domain.AddDomain(domain.Domain{
			DomainName:  args[0],
			IsInScope:   !isOutOfScope,
			ProgramName: programName})
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	domainCmd.AddCommand(addDomainCmd)

	addDomainCmd.Flags().BoolVar(&isOutOfScope, "no-scope", false, "Mark the domain as out of scope (default false)")
	addDomainCmd.MarkPersistentFlagRequired("program")
}
