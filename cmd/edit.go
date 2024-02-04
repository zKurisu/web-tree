package cmd

import (
	"github.com/spf13/cobra"
)

// editUsage string = "webtree edit --tree=<name> --tname=<newName>"
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

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
