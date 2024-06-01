package cmd

import (
	"os"

	"github.com/JakobTheDev/daneel/cmd/platform"
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

              A bug bounty robot
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(platform.PlatformCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.daneel.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}