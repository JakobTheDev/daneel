package cli

import (
	"fmt"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var programListCmd = &cobra.Command{
	Use:   "list",
	Short: "List bug bounty programs",
	Run: func(cmd *cobra.Command, args []string) {
		programs, err := models.ListPrograms()
		if err != nil {
			fmt.Println(err)
		}

		if cmd.Flag("table").Value.String() == "true" {

			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()

			tbl := table.New("ID", "Program Name", "Platform Name", "Active", "Private")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

			for _, program := range programs {
				tbl.AddRow(program.Id, program.DisplayName, program.PlatformName, program.IsActive, program.IsPrivate)
			}

			tbl.Print()
		} else {
			for _, p := range programs {
				fmt.Println(p.DisplayName)
			}
		}
	},
}

func init() {
	programCmd.AddCommand(programListCmd)
}
