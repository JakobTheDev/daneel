package domain

import (
	"fmt"
	"strings"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [domain]",
	Short: "Remove a domain from a program",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		domainName := strings.ToLower(args[0])

		err := models.RemoveDomain(models.Domain{DomainName: domainName})
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	DomainCmd.AddCommand(removeCmd)
}
