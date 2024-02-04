package cmd

import (
	"github.com/spf13/cobra"
)

var (
	moveCmd = &cobra.Command{
		Use:   "move",
		Short: "Move tree/node to another tree",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

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
