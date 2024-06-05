package platform

import (
	"fmt"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

// Flags
var showInactive bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists bug bounty platforms",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		platforms, err := models.ListPlatforms(showInactive)
		if err != nil {
			fmt.Println(err)
		}

		if cmd.Flag("table").Value.String() == "true" {

			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()

			tbl := table.New("ID", "Platform Name", "Active")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

			for _, platform := range platforms {
				tbl.AddRow(platform.ID, platform.DisplayName, platform.IsActive)
			}

			tbl.Print()
		} else {
			for _, p := range platforms {
				fmt.Println(p.DisplayName)
			}
		}
	},
}

func init() {
	PlatformCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&showInactive, "inactive", "i", false, "Show inactive platforms")
}
