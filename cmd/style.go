package cmd

import (
	"github.com/spf13/cobra"
)

var styleCmd = &cobra.Command{
	Use:   "style",
	Short: "List all styles",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	listCmd.AddCommand(styleCmd)
}
