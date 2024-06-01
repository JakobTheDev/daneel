package platform

import (
	"fmt"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add PLATFORM_NAME",
	Short: "Add a bug bounty platform to daneel",
	Run: func(cmd *cobra.Command, args []string) {
		err := models.AddPlatform(models.Platform{DisplayName: args[0]})
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	PlatformCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
