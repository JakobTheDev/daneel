package cli

import (
	"fmt"
	"log"

	"github.com/JakobTheDev/daneel/internal/subdomain"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var subdomainListCmd = &cobra.Command{
	Use:   "list [OPTIONS]",
	Short: "Lists subdomains",
	Run: func(command *cobra.Command, args []string) {
		subdomains, err := subdomain.ListSubdomains(programName, domainName, showOutOfScope)
		if err != nil {
			log.Fatal(err)
		}

		if command.Flag("table").Value.String() == "true" {

			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()

			tbl := table.New("ID", "Domain Name", "Subdomain", "In Scope")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

			for _, subdomain := range subdomains {
				tbl.AddRow(subdomain.ID, subdomain.DomainName, subdomain.Name, subdomain.IsInScope)
			}

			tbl.Print()
		} else {
			for _, s := range subdomains {
				fmt.Println(s.Name)
			}
		}

	},
}

func init() {
	subdomainCmd.AddCommand(subdomainListCmd)

	subdomainListCmd.Flags().StringVar(&programName, "program", "", "Bug bounty program")
	subdomainListCmd.Flags().StringVar(&domainName, "domain", "", "The subdomain's parent domain")
	subdomainListCmd.Flags().BoolVar(&showOutOfScope, "show-no-scope", false, "Show out-of-scope subdomains (default false)")
	subdomainListCmd.MarkFlagsOneRequired("program", "domain")
	subdomainListCmd.MarkFlagsMutuallyExclusive("program", "domain")
}
