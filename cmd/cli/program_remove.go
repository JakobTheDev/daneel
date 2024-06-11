package cli

import (
	"log"
	"strings"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/JakobTheDev/daneel/internal/program"
	"github.com/spf13/cobra"
)

var programRemoveCmd = &cobra.Command{
	Use:   "remove [program]",
	Short: "Remove a bug bounty program",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		err := program.RemoveProgram(models.Program{Name: strings.ToLower(args[0])})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	programCmd.AddCommand(programRemoveCmd)
}
