// Refactored code for subdomain

package subdomain

import (
	"fmt"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var showOutOfScope bool

var subdomainListCmd = &cobra.Command{
	Use:   "list [OPTIONS]",
	Short: "Lists subdomains",
	Run: func(command *cobra.Command, args []string) {
		subdomains, err := models.ListSubdomains(ProgramName, DomainName, showOutOfScope)
		if err != nil {
			fmt.Println(err)
		}

		if command.Flag("table").Value.String() == "true" {

			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()

			tbl := table.New("ID", "Domain Name", "Subdomain", "In Scope")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

			for _, subdomain := range subdomains {
				tbl.AddRow(subdomain.ID, subdomain.DomainName, subdomain.SubdomainName, subdomain.IsInScope)
			}

			tbl.Print()
		} else {
			for _, s := range subdomains {
				fmt.Println(s.SubdomainName)
			}
		}

	},
}

func init() {
	SubdomainCmd.AddCommand(subdomainListCmd)

	subdomainListCmd.Flags().StringVar(&ProgramName, "program", "", "Bug bounty program")
	subdomainListCmd.Flags().StringVar(&DomainName, "domain", "", "The subdomain's parent domain")
	subdomainListCmd.Flags().BoolVar(&showOutOfScope, "show-no-scope", false, "Show out-of-scope subdomains (default false)")
	subdomainListCmd.MarkFlagsOneRequired("program", "domain")
	subdomainListCmd.MarkFlagsMutuallyExclusive("program", "domain")
}
