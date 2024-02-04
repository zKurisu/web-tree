package cmd

import (
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List meta info",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
}
