package domain

import (
	"fmt"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/spf13/cobra"
)

var isOutOfScope bool

var addCmd = &cobra.Command{
	Use:   "add [domain]",
	Short: "Add a domain to daneel",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := models.AddDomain(models.Domain{
			DomainName:  args[0],
			IsInScope:   !isOutOfScope,
			ProgramName: ProgramName})
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	DomainCmd.AddCommand(addCmd)

	addCmd.Flags().BoolVar(&isOutOfScope, "no-scope", false, "Mark the program as out of scope (default false)")
	addCmd.MarkFlagRequired("program")
}
