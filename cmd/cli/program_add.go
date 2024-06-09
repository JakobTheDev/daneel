package cli

import (
	"fmt"
	"strings"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/spf13/cobra"
)

var programAddCmd = &cobra.Command{
	Use:   "add [program]",
	Short: "Add a bug bounty program",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		program := models.Program{
			DisplayName:  strings.ToLower(args[0]),
			PlatformName: strings.ToLower(platformName),
			IsPrivate:    isPrivate,
		}

		err := models.AddProgram(program)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	programCmd.AddCommand(programAddCmd)

	programAddCmd.Flags().StringVarP(&platformName, "platform", "p", "", "The platform the program is on (required)")
	programAddCmd.Flags().BoolVar(&isPrivate, "private", false, "Mark the program as private (default false)")

	programAddCmd.MarkFlagRequired("platform")
}
