package cmd

import (
	"errors"
	// "github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"log"
	"web-tree/utils"
)

var (
	delCmd = &cobra.Command{
		Use:   "del",
		Short: "Delete tree/node or link/alias/desc in node",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			nameList := utils.Split2List(treeName)
			aliasList := utils.Split2List(alias)
			linkList := utils.Split2List(link)

			if len(nameList) == 0 {
				return errors.New("A tree name must be specified")
			} else {
				for _, name := range nameList {
					if !utils.IsNameValid(name) {
						return errors.New("The tree name is not valid (add)")
					}
				}
			}
			root := utils.GetRootTree()

			for _, name := range nameList {
				if root.DeepFindSubTree(name) == nil {
					return errors.New(name + " does not exist")
				} else {
					if !isNode {
						root.DeepDelSubTree(name)
						log.Println("Delete " + name + "....")
					} else if alias != "" {
						// log.Println(root.DeepFindSubTree(treeLevels).Nodes)
						// spew.Dump(root.DeepFindSubTree(treeLevels))
						root.DeepFindSubTree(name).DelNode(aliasList)
						// spew.Dump(root.DeepFindSubTree(treeLevels))
						// log.Println(root.DeepFindSubTree(treeLevels).Nodes)
						log.Println("Delete node with alias [" + alias + "] under " + name + "....")
					} else if link != "" {
						root.DeepFindSubTree(name).DelNode(linkList)
						log.Println("Delete node with link [" + link + "] under " + name + "....")
					} else {
						return errors.New("An alias or a link is needed to delete a node")
					}
				}
			}

			if err := utils.WriteAll(); err != nil {
				log.Fatal(err)
			}
			if !isNode {
				log.Println("Finish deleting" + treeName)
			} else {
				log.Println("Finish deleting node under " + treeName)
			}
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(delCmd)

	/*
	   --tree/-t
	   --node/-n
	   --alias/-a
	   --link/-l
	*/
	delCmd.Flags().StringVarP(&treeName, "tree", "t", "", "target web tree")
	delCmd.Flags().BoolVarP(&isNode, "node", "n", false, "target node")
	delCmd.Flags().StringVarP(&alias, "alias", "a", "", "Alias of target node")
	delCmd.Flags().StringVarP(&link, "link", "l", "", "Urls of target node")
}
