package cli

import (
	"fmt"
	"strings"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/spf13/cobra"
)

var platformRemoveCmd = &cobra.Command{
	Use:   "remove [platform]",
	Short: "Remove a bug bounty platform",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		platformName := strings.ToLower(args[0])

		err := models.RemovePlatform(models.Platform{DisplayName: platformName})
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	platformCmd.AddCommand(platformRemoveCmd)
}
