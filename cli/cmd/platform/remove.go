package platform

import (
	"fmt"
	"strings"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [PLATFORM]",
	Short: "Remove a bug bounty platform from daneel (soft delete)",
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
	PlatformCmd.AddCommand(removeCmd)
}
