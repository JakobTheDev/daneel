package platform

import (
	"fmt"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists bug bounty platforms tracked by daneel",
	Run: func(cmd *cobra.Command, args []string) {
		platforms, err := models.ListPlatforms()
		if err != nil {
			fmt.Println(err)
		}

		for _, p := range platforms {
			fmt.Println(p.DisplayName)
		}
	},
}

func init() {
	PlatformCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
