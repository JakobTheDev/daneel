package cli

import (
	"fmt"
	"strings"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/spf13/cobra"
)

var programRemoveCmd = &cobra.Command{
	Use:   "remove [program]",
	Short: "Remove a bug bounty program",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		err := models.RemoveProgram(models.Program{DisplayName: strings.ToLower(args[0])})
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	programCmd.AddCommand(programRemoveCmd)
}
