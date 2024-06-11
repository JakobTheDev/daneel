package cli

import (
	"fmt"
	"log"

	"github.com/JakobTheDev/daneel/internal/platform"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var platformListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists bug bounty platforms",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		platforms, err := platform.ListPlatforms(showInactive)
		if err != nil {
			log.Fatal(err)
		}

		if cmd.Flag("table").Value.String() == "true" {
			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()

			tbl := table.New("ID", "Platform Name", "Active")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

			for _, platform := range platforms {
				tbl.AddRow(platform.ID, platform.Name, platform.IsActive)
			}

			tbl.Print()
		} else {
			for _, p := range platforms {
				fmt.Println(p.Name)
			}
		}
	},
}

func init() {
	platformCmd.AddCommand(platformListCmd)

	platformListCmd.Flags().BoolVarP(&showInactive, "inactive", "i", false, "Show inactive platforms")
}
