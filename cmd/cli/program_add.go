package cli

import (
	"log"
	"strings"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/JakobTheDev/daneel/internal/program"
	"github.com/spf13/cobra"
)

var programAddCmd = &cobra.Command{
	Use:   "add [program]",
	Short: "Add a bug bounty program",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		p := models.Program{
			Name:         strings.ToLower(args[0]),
			PlatformName: strings.ToLower(platformName),
			IsPrivate:    isPrivate,
		}

		err := program.AddProgram(p)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	programCmd.AddCommand(programAddCmd)

	programAddCmd.Flags().StringVarP(&platformName, "platform", "p", "", "The platform the program is on (required)")
	programAddCmd.Flags().BoolVar(&isPrivate, "private", false, "Mark the program as private (default false)")

	programAddCmd.MarkFlagRequired("platform")
}
