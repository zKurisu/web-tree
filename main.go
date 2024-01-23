package main

import (
	"fmt"
	// "gopkg.in/yaml.v3"
	// "web-tree/conf"
	"web-tree/utils"
)

func main() {
	t := utils.GetTree("linux")
	fmt.Printf("%#v", t)
	if utils.IsRootTree(t) {
		fmt.Println("It's Root Tree")
	}
	fmt.Printf("%p", t)
	fmt.Println()
	t.AddNewSubTree("Test")
	testTree := t.FindSubTree("Test")
	testNode := t.FindAllNode("kurisu")

	testNode[1].Link[1] = "haha"

	fmt.Printf("%#v", testNode[1])
	fmt.Println()

	testTree.AddNode(utils.NewNode(
		[]string{`www.test1.com`, `www.test2.com`},
		[]string{`test1`, `test2`},
		[]string{`desc1`, `desc2`},
		`icon`,
		[]string{`linux`, `perl`},
		[]string{`None`},
	))

	list := utils.AllRootTree
	fmt.Println()
	fmt.Printf("%#v", list)
	t.WriteTree()
}
