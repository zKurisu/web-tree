package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"log"
	"web-tree/utils"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add new tree/node",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		/*
		   Example command :
		   webtree add --tree=hello/world,another
		   webtree add --tree=hello/world --node \
		   --link="https://orkarin.com","https://kurisu.com" \
		   --alias="alias1","alias2" \
		   --desc="desc1","desc2" \
		   --label="label1","label2"
		*/
		nameList := utils.Split2List(treeName)
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
			treeLevels := utils.SplitTreeLevel(name)
			root.DeepAddNewSubTree(treeLevels)
			tree := root.DeepFindSubTree(treeLevels)

			if isNode {
				if alias == "" && link == "" {
					return errors.New(`If adding a new node, a link and a alias must be given`)
				}
				node, err := utils.NewNode(
					utils.Split2List(link),
					utils.Split2List(alias),
					utils.Split2List(desc),
					icon,
					utils.Split2List(label),
					style,
				)
				if err != nil {
					log.Fatal(err)
				}
				tree.AddNode(node)

			} else if alias != "" || link != "" || desc != "" || label != "" {
				return errors.New("Flag link, alias, desc, label should be used with --node")
			}
		}

		if err := utils.WriteAll(); err != nil {
			log.Fatal(err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	/*
	   --tree/-t
	   --node/-n
	   --alias/-a
	   --link/-l
	   --desc/-d
	   --label
	*/
	addCmd.Flags().StringVarP(&treeName, "tree", "t", "", "target web tree")
	addCmd.Flags().BoolVarP(&isNode, "node", "n", true, "target node")
	addCmd.Flags().StringVarP(&alias, "alias", "a", "", "Alias of target node")
	addCmd.Flags().StringVarP(&link, "link", "l", "", "Urls of target node")
	addCmd.Flags().StringVarP(&desc, "desc", "d", "", "Descriptions of target node")
	addCmd.Flags().StringVarP(&label, "label", "", "", "Labels of target node")
}
