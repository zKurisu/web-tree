package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"web-tree/utils"
)

var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "List all trees",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		treeList := utils.GetAllTreeName()
		for _, name := range treeList {
			fmt.Println(name)
		}
	},
}

func init() {
	listCmd.AddCommand(treeCmd)
}
