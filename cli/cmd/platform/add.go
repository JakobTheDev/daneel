package platform

import (
	"fmt"
	"strings"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [platform]",
	Short: "Add a bug bounty platform to daneel",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		platformName := strings.ToLower(args[0])

		err := models.AddPlatform(models.Platform{DisplayName: platformName})
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	PlatformCmd.AddCommand(addCmd)
}
