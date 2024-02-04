package cmd

import (
	"github.com/spf13/cobra"
)

var (
	nodeCmd = &cobra.Command{
		Use:   "node",
		Short: "List all nodes",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

func init() {
	listCmd.AddCommand(nodeCmd)
}
