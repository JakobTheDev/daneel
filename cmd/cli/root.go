package cli

import (
	"os"

	"github.com/spf13/cobra"
)

// Flags
var (
	domainName     string
	isOutOfScope   bool
	isPrivate      bool
	platformName   string
	programName    string
	showInactive   bool
	showOutOfScope bool
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "daneel",
	Short: "A bug bounty robot",
	Long: `
               ______   _______  __    _  _______  _______  ___     
    /\__/\    |      | |   _   ||  |  | ||       ||       ||   |    
   / _   _\   |  _    ||  |_|  ||   |_| ||    ___||    ___||   |    
  | | | | |   | | |   ||       ||       ||   |___ |   |___ |   |    
  | |_| |_|   | |_|   ||       ||  _    ||    ___||    ___||   |___ 
   \_____/    |       ||   _   || | |   ||   |___ |   |___ |       |
              |______| |__| |__||_|  |__||_______||_______||_______|

              A bug bounty bot
	`,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(platformCmd)
	RootCmd.AddCommand(programCmd)
	RootCmd.AddCommand(domainCmd)
	RootCmd.AddCommand(subdomainCmd)

	RootCmd.PersistentFlags().BoolP("table", "t", false, "Print output in a table (default false)")
}
