package subdomain

import (
	"fmt"
	"strings"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [subdomain]",
	Short: "Remove a subdomain",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		subdomainName := strings.ToLower(args[0])

		err := models.RemoveSubdomain(models.Subdomain{SubdomainName: subdomainName})
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	SubdomainCmd.AddCommand(removeCmd)
}
