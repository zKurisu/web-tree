package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"web-tree/utils"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add new tree/node",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		if treeName == "" {
			return errors.New("A tree name must be specified")
		}
		if isNode {
			if alias == "" && link == "" {
				return errors.New(`If adding a new node, a link and a alias must be given`)
			}
		} else if alias != "" || link != "" || desc != "" || label != "" {
			return errors.New("Flag link, alias, desc, label should be used with --node")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		/*
		   Example command :
		   webtree add --tree=hello/world,another
		   webtree add --tree=hello/world --node \
		   --link="https://orkarin.com","https://kurisu.com" \
		   --alias="alias1","alias2" \
		   --desc="desc1","desc2" \
		   --label="label1","label2"
		*/
		treeLevel := utils.SplitTreeLevel(treeName)
		if !utils.IsTreeExist(treeLevel[0]) {

		}
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
