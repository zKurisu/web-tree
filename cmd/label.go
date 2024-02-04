package cmd

import (
	"github.com/spf13/cobra"
)

var (
	labelCmd = &cobra.Command{
		Use:   "label",
		Short: "List all labels",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

func init() {
	listCmd.AddCommand(labelCmd)
}
