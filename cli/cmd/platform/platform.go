/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package platform

import (
	"fmt"

	"github.com/spf13/cobra"
)

// platformCmd represents the platform command
var PlatformCmd = &cobra.Command{
	Use:   "platform [command]",
	Short: "Manage bug bounty platforms within daneel",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("platform called")
	},
}

func init() {

}
