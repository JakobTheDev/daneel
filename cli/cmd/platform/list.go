package platform

import (
	"fmt"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/spf13/cobra"
)

// Flags
var showInactive bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists bug bounty platforms tracked by daneel",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		platforms, err := models.ListPlatforms(showInactive)
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

	listCmd.Flags().BoolVarP(&showInactive, "inactive", "i", false, "Show inactive platforms")
}
