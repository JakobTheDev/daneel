package cli

import (
	"log"
	"strings"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/JakobTheDev/daneel/internal/platform"
	"github.com/spf13/cobra"
)

var platformRemoveCmd = &cobra.Command{
	Use:   "remove [platform]",
	Short: "Remove a bug bounty platform",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		platformName := strings.ToLower(args[0])

		err := platform.RemovePlatform(models.Platform{Name: platformName})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	platformCmd.AddCommand(platformRemoveCmd)
}
