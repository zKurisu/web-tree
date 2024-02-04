package cmd

import (
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show detail of a meta info",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

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
