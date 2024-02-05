package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"web-tree/utils"
)

var (
	nodeCmd = &cobra.Command{
		Use:   "node",
		Short: "List all nodes",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			root := utils.GetRootTree()
			nodeList := root.FindAllNode("")
			for _, node := range nodeList {
				linkString := utils.List2String(node.Link)
				aliasString := utils.List2String(node.Alias)
				descString := utils.List2String(node.Desc)
				iconString := icon
				labelString := utils.List2String(node.Label)
				fmt.Println("================================")
				fmt.Println("Links : " + linkString)
				fmt.Println("Alias : " + aliasString)
				fmt.Println("Descs : " + descString)
				fmt.Println("Icon  : " + iconString)
				fmt.Println("Labels: " + labelString)
				fmt.Println("================================")
			}
		},
	}
)

func init() {
	listCmd.AddCommand(nodeCmd)
}
