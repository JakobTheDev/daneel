package cli

import (
	"fmt"
	"log"

	"github.com/JakobTheDev/daneel/internal/domain"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var domainListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists domains within a program",
	Run: func(command *cobra.Command, args []string) {
		programs, err := domain.ListDomains(programName, showOutOfScope)
		if err != nil {
			log.Fatal(err)
		}

		if command.Flag("table").Value.String() == "true" {

			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()

			tbl := table.New("ID", "Program Name", "Domain Name", "In Scope")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

			for _, program := range programs {
				tbl.AddRow(program.ID, program.ProgramName, program.Name, program.IsInScope)
			}

			tbl.Print()
		} else {
			for _, p := range programs {
				fmt.Println(p.Name)
			}
		}

	},
}

func init() {
	domainCmd.AddCommand(domainListCmd)

	domainListCmd.Flags().BoolVar(&showOutOfScope, "show-no-scope", false, "Show out-of-scope domains (default false)")
}
