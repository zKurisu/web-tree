package main

import (
	"fmt"
	// "gopkg.in/yaml.v3"
	// "web-tree/conf"
	// "web-tree/cmd"
	"web-tree/ui"
	"web-tree/utils"
)

func main() {
	// cmd.Execute()
	utils.CheckTreeName()
	root := utils.GetRootTree()
	allTrees := root.DeepGetAllSubtreePath("")
	list := ui.SuggestionInit()
	fmt.Println(allTrees)
	fmt.Println(utils.SplitTreeLevel(allTrees[1]))
	for _, elem := range list {
		fmt.Println(elem)
	}
	// t := root.FindSubTree("balabala")
	// n := t.FindNode([]string{"test1"})
	// if utils.IsNode(n) {
	// 	fmt.Println("It's node")
	// } else {
	// 	fmt.Println("It's not a tree")
	// }
	// b := root.FindSubTree("baba")
	//
	// b.ChangeFileName("haha")

	// node1, _ := utils.NewNode(
	// 	[]string{"https://test1.com", "https://test2.com"},
	// 	[]string{"test4", "test5"},
	// 	[]string{"desc1", "desc2"},
	// 	"icon",
	// 	[]string{"label1", "label2"},
	// 	"style",
	// )
	// h.AddNode(node1)
	// h.AddNewSubTree("babalala")
	// utils.WriteAll()
}

// root := utils.GetRootTree()
// h := root.FindSubTree("hello")
// t := root.FindSubTree("test")
//
// node1 := utils.NewNode(
// 	[]string{"https://test1.com", "https://test2.com"},
// 	[]string{"test1", "test2"},
// 	[]string{"desc1", "desc2"},
// 	"icon",
// 	[]string{"label1", "label2"},
// 	"style",
// )
// node2 := utils.NewNode(
// 	[]string{"https://test1.com", "https://test2.com"},
// 	[]string{"test1", "test2"},
// 	[]string{"desc1", "desc2"},
// 	"icon",
// 	[]string{"label1", "label2"},
// 	"style",
// )
//
// tree1 := utils.NewTree("tree1")
// tree2 := utils.NewTree("tree2")
//
// tree1.AddNode(node1)
// tree1.AddSubTree(h)
//
// tree2.AddNode(node2)
// tree2.AddSubTree(t)
//
// if utils.IsTreeEqual(*tree1, *tree2) {
// 	log.Println("Equal")
// } else {
// 	log.Println("Not equal")
// }
