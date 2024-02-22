package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strings"
	"web-tree/utils"
)

var (
	moveCmd = &cobra.Command{
		Use:   "move",
		Short: "Move tree/node to another tree",
		Long:  "",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
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

			if len(args) == 0 {
				log.Fatal("A target tree need to be specified as a argument")
			}

			root := utils.GetRootTree()
			targetName := args[0]
			targetTree := root.DeepFindSubTree(targetName)

			if targetTree == nil {
				log.Fatal("Could not find tree " + targetName)
			}
			// log.Println("Echo " + args[0])
			for _, name := range nameList {
				if targetTree.DeepFindSubTree(name) != nil {
					log.Println("Tree " + targetName + " already have a subtree named " + name)
					log.Println("Jump the moving of " + name)
				} else {
					tree := root.DeepFindSubTree(name)

					if tree == nil {
						log.Fatal("Could not find tree " + name)
					}
					if !isNode {
						log.Println("Does not open node flag...")
						if alias != "" || link != "" {
							log.Fatal("Flag: --alias and --link should be used with --node")
						}
						targetTree.AppendSubTree(tree)
						if !root.IsSubTree(tree) {
							treeLevels := utils.Split2List(name)
							preTreeName := strings.Join(treeLevels[:len(treeLevels)-1], "/")
							root.DeepFindSubTree(preTreeName).DelSubTree(tree.Name)
						} else {
							root.DelSubTree(tree.Name)
						}
					} else {
						log.Println("Open node flag...")

						hints := []string{}

						if alias == "" && link == "" {
							log.Fatal("Flag: --alias and --link should be used with --node")
						}
						if alias != "" {
							hints = utils.MergeList(hints, utils.Split2List(alias)).([]string)
						}
						if link != "" {
							hints = utils.MergeList(hints, utils.Split2List(link)).([]string)
						}
						node := tree.FindNode(hints)
						if node == nil {
							log.Fatal("[move] Could not find node")
						}
						targetTree.AppendNode(node)
						tree.DelNode(hints)
					}
				}
			}

			if err := utils.WriteAll(); err != nil {
				log.Fatal(err)
			}
			log.Println("Finish moving...")
		},
	}
)

func init() {
	rootCmd.AddCommand(moveCmd)

	/*
	   --tree/-t
	   --node/-n
	   --alias/-a
	   --link/-l
	*/
	moveCmd.Flags().StringVarP(&treeName, "tree", "t", "", "target web tree")
	moveCmd.Flags().BoolVarP(&isNode, "node", "n", false, "target node")
	moveCmd.Flags().StringVarP(&alias, "alias", "a", "", "Alias of target node")
	moveCmd.Flags().StringVarP(&link, "link", "l", "", "Urls of target node")
}
