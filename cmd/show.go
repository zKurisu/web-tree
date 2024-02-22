package cmd

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"log"
	"web-tree/utils"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show detail of a meta info",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		root := utils.GetRootTree()
		// If --label=<label> is specified
		if label != "" {
			if treeName != "" || isNode || link != "" || alias != "" {
				log.Fatal("Using --label alone, without --tree, --node, --link and --alias")
			}
			nodeList := root.FindAllNode("")
			for _, node := range nodeList {
				for _, nodelabel := range utils.RemoveDup(node.Label) {
					if label == nodelabel {
						fmt.Println(node.Link)
					}
				}
			}
		} else {
			nameList := utils.Split2List(treeName)
			if len(nameList) == 0 {
				log.Fatal("A tree name must be specified")
			} else {
				for _, name := range nameList {
					if !utils.IsNameValid(name) {
						log.Fatal("The tree name is not valid (edit)")
					}
				}
			}

			for _, name := range nameList {
				tree := root.DeepFindSubTree(name)
				if tree == nil {
					log.Fatal("[show] Could not find tree: " + name)
				}

				if !isNode {
					spew.Dump(tree)
				} else {
					hints := []string{}
					if link != "" {
						links := utils.Split2List(link)
						hints = utils.MergeList(hints, links).([]string)
					}
					if alias != "" {
						aliasList := utils.Split2List(alias)
						hints = utils.MergeList(hints, aliasList).([]string)
					}
					node := tree.FindNode(hints)
					if node == nil {
						log.Fatal("[edit] could not find node")
					}
					spew.Dump(node)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	/*
	   --tree/-t
	   --node/-n
	   --alias/-a
	   --link/-l
	   --label
	*/
	showCmd.Flags().StringVarP(&treeName, "tree", "t", "", "target web tree")
	showCmd.Flags().BoolVarP(&isNode, "node", "n", false, "target node")
	showCmd.Flags().StringVarP(&alias, "alias", "a", "", "Alias of target node")
	showCmd.Flags().StringVarP(&link, "link", "l", "", "Urls of target node")
	showCmd.Flags().StringVarP(&label, "label", "", "", "Labels of target node")
}
