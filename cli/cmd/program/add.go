package program

import (
	"fmt"
	"strings"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/spf13/cobra"
)

var platformName string
var isPrivate bool

var addCmd = &cobra.Command{
	Use:   "add [program]",
	Short: "Adds a bug bounty program to daneel",
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
	ProgramCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&platformName, "platform", "p", "", "The platform the program is on (required)")
	addCmd.Flags().BoolVar(&isPrivate, "private", false, "Mark the program as private (default false)")

	addCmd.MarkFlagRequired("platform")
}
