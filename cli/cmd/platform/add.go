package platform

import (
	"fmt"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [PLATFORM]",
	Short: "Add a bug bounty platform to daneel",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := models.AddPlatform(models.Platform{DisplayName: args[0]})
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	PlatformCmd.AddCommand(addCmd)
}
