package cmd

import (
	"github.com/spf13/cobra"
)

var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "List all trees",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	listCmd.AddCommand(treeCmd)
}
