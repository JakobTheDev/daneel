package cmd

import (
	"os"

	"github.com/JakobTheDev/daneel/cmd/domain"
	"github.com/JakobTheDev/daneel/cmd/platform"
	"github.com/JakobTheDev/daneel/cmd/program"
	"github.com/spf13/cobra"
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
	RootCmd.AddCommand(platform.PlatformCmd)
	RootCmd.AddCommand(program.ProgramCmd)
	RootCmd.AddCommand(domain.DomainCmd)

	RootCmd.PersistentFlags().BoolP("table", "t", false, "Print output in a table (default false)")
}
