package program

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [OPTIONS]",
	Short: "Lists bug bounty programs tracked by daneel",
	Run: func(cmd *cobra.Command, args []string) {
		programs, err := models.ListPrograms()
		if err != nil {
			fmt.Println(err)
		}

		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		fmt.Fprintln(w, "Program Name", "\t", "Platform Name")
		for _, p := range programs {
			fmt.Fprintln(w, p.DisplayName, "\t", p.PlatformName)
		}
		w.Flush()
	},
}

func init() {
	ProgramCmd.AddCommand(listCmd)
}
