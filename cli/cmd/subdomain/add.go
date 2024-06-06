package subdomain

import (
	"fmt"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/spf13/cobra"
)

var isOutOfScope bool

var addCmd = &cobra.Command{
	Use:   "add [subdomain]",
	Short: "Add a subdomain to daneel",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err, _ := models.AddSubdomain(models.Subdomain{
			SubdomainName: args[0],
			IsInScope:     !isOutOfScope,
			DomainName:    DomainName})
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	SubdomainCmd.AddCommand(addCmd)

	addCmd.Flags().StringVar(&DomainName, "domain", "", "The subdomain's parent domain")
	addCmd.Flags().BoolVar(&isOutOfScope, "no-scope", false, "Mark the subdomain as out of scope (default false)")
	addCmd.MarkFlagRequired("domain")
}
