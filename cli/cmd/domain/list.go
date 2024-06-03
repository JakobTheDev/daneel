package domain

import (
	"fmt"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var showOutOfScope bool

var listCmd = &cobra.Command{
	Use:   "list [OPTIONS]",
	Short: "Lists domains tracked by daneel",
	Run: func(command *cobra.Command, args []string) {
		programs, err := models.ListDomain(ProgramName, showOutOfScope)
		if err != nil {
			fmt.Println(err)
		}

		if command.Flag("table").Value.String() == "true" {

			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()

			tbl := table.New("ID", "Program Name", "DomainName", "In Scope")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

			for _, program := range programs {
				tbl.AddRow(program.ID, program.ProgramName, program.DomainName, program.IsInScope)
			}

			tbl.Print()
		} else {
			for _, p := range programs {
				fmt.Println(p.DomainName)
			}
		}

	},
}

func init() {
	DomainCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&showOutOfScope, "show-no-scope", false, "Show out-of-scope domains (default false)")
}
