package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"web-tree/utils"
)

var (
	labelCmd = &cobra.Command{
		Use:   "label",
		Short: "List all labels",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			root := utils.GetRootTree()
			nodeList := root.FindAllNode("")
			var labelList []string

			for _, node := range nodeList {
				labelList = utils.MergeList(labelList, node.Label).([]string)
			}
			labelList = utils.RemoveDup(labelList)

			fmt.Printf("[Total: %d]\n", len(labelList))
			for _, label := range labelList {
				fmt.Println(label)
			}
		},
	}
)

func init() {
	listCmd.AddCommand(labelCmd)
}
