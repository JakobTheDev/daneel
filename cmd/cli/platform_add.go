package cli

import (
	"log"
	"strings"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/JakobTheDev/daneel/internal/platform"
	"github.com/spf13/cobra"
)

var platformAddCmd = &cobra.Command{
	Use:   "add [platform]",
	Short: "Add a bug bounty platform",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		platformName := strings.ToLower(args[0])

		err := platform.AddPlatform(models.Platform{Name: platformName})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	platformCmd.AddCommand(platformAddCmd)
}
