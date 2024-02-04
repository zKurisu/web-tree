package cmd

import (
	"github.com/spf13/cobra"
)

var (
	delCmd = &cobra.Command{
		Use:   "del",
		Short: "Delete tree/node or link/alias/desc in node",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

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
