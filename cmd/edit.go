package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"web-tree/utils"
)

// editUsage string = "webtree edit --tree=<name> --tname=<newName>"
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// log.Fatal("Test log fatal")
		// log.Println("Run..")
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
		// log.Println("Name is valid..")

		root := utils.GetRootTree()

		for _, name := range nameList {
			treeLevels := utils.SplitTreeLevel(name)
			tree := root.DeepFindSubTree(treeLevels)
			if tree == nil {
				log.Fatal("Could not find tree " + name)
			}
			// log.Println("Find tree " + tree.Name)

			if !isNode {
				// log.Println("Editing tree name..")
				if newName == "" {
					log.Fatal("When editing the tree name, --tname=<newName> should be provided")
				} else if newLink != "" || newAlias != "" || newDesc != "" || newLabel != "" {
					log.Fatal("Flag: --nlink, --nalias, --ndesc, --nlabel should be used with --node")
				} else {
					tree.ChangeFileName(newName)
					tree.Name = newName
				}
			} else {
				hints := []string{}
				if newName != "" {
					tree.ChangeFileName(newName)
					tree.Name = newName
				}
				if link != "" {
					links := utils.Split2List(link)
					hints = utils.MergeList(hints, links)
				}
				if alias != "" {
					aliasList := utils.Split2List(alias)
					hints = utils.MergeList(hints, aliasList)
				}
				node := tree.FindNode(hints)
				if node == nil {
					log.Fatal("[edit] could not find node")
				}
				if newLink != "" {
					node.Link = utils.Split2List(newLink)
				}
				if newAlias != "" {
					node.Alias = utils.Split2List(newAlias)
				}
				if newDesc != "" {
					node.Desc = utils.Split2List(newDesc)
				}
				if newLabel != "" {
					node.Label = utils.Split2List(newLabel)
				}
			}
		}

		// log.Println("Before Running WriteAll")
		if err := utils.WriteAll(); err != nil {
			log.Println("Running WriteAll")
			log.Fatal(err)
		}
		log.Println("Finish edit...")
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	/*
	   --tree/-t
	   --node/-n
	   --alias/-a
	   --link/-l
	   --tname
	   --nlink
	   --nalias
	   --ndesc
	   --nlabel
	*/
	editCmd.Flags().StringVarP(&treeName, "tree", "t", "", "target web tree")
	editCmd.Flags().BoolVarP(&isNode, "node", "n", false, "target node")
	editCmd.Flags().StringVarP(&alias, "alias", "a", "", "Alias of target node")
	editCmd.Flags().StringVarP(&link, "link", "l", "", "Urls of target node")
	editCmd.Flags().StringVarP(&newName, "tname", "", "", "New tree name")
	editCmd.Flags().StringVarP(&newLink, "nlink", "", "", "New links for target node")
	editCmd.Flags().StringVarP(&newAlias, "nalias", "", "", "New alias for target node")
	editCmd.Flags().StringVarP(&newDesc, "ndesc", "", "", "New descriptions for target node")
	editCmd.Flags().StringVarP(&newLabel, "nlabel", "", "", "New labels for target node")
}
