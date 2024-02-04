package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "webtree",
	Short: "A url storage with many features for maintain",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// Open the TUI
	},
}

func Execute() {
	rootCmd.Execute()
}

func init() {
	var version = "0.01"
	// --version/-v
	rootCmd.SetVersionTemplate(version)

	/*
	   --tree/-t
	   --node/-n
	   --alias/-a
	   --link/-l
	   --browser/-b
	*/
	rootCmd.Flags().StringVarP(&treeName, "tree", "t", "root", "target web tree")
	rootCmd.Flags().BoolVarP(&isNode, "node", "n", false, "target node")
	rootCmd.Flags().StringVarP(&alias, "alias", "a", "", "Alias of target node")
	rootCmd.Flags().StringVarP(&link, "link", "l", "", "Urls of target node")
	rootCmd.Flags().StringVarP(&browser, "browser", "b", "", "Open url with target browser")
}
