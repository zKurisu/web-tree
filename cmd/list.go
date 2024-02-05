package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List meta info",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Please run with sub command")
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
}
