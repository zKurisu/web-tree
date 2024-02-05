package utils

import (
	// "github.com/davecgh/go-spew/spew"
	"log"
	"net/url"
	"reflect"
)

// The Root tree name should be the same with
// file name
func CheckTreeName() {
	allTreeName := GetAllTreeName()
	for _, name := range allTreeName {
		t := GetTree(name)
		if t.Name != name {
			log.Fatal(`[` + name + `]` + " :The Root tree name should be the same with file name")
		}
	}
	log.Println("CheckTreeName OK")
}

func IsUrl(s string) error {
	_, err := url.ParseRequestURI(s)
	return err
}

func IsInList(slice []string, str string) bool {
	for _, s := range slice {
		if str == s {
			return true
		}
	}
	return false
}

func IsNodeExist(nodes []*Node, node *Node) bool {
	for _, v := range nodes {
		if node == v {
			return true
		}
	}
	return false
}

func IsTreeExist(name string) bool {
	if IsInList(GetAllTreeName(), name) {
		return true
	} else {
		return false
	}
}

func IsSubTree(tree *Tree) bool {
	for _, v := range RootTree.GetAllSubtree() {
		if v == tree {
			return false
		}
	}
	return true
}

func IsNameValid(name string) bool {
	if name == "root" || name == "" {
		return false
	}
	return true
}

func isNodeEqual(n1, n2 *Node) bool {
	return reflect.DeepEqual(n1, n2)
}

func IsTreeEqual(t1, t2 *Tree) bool {
	// log.Println("For Cur:")
	// spew.Dump(t1)
	// log.Println("For Ori:")
	// spew.Dump(t2)
	// log.Println("For tree: " + t1.Name + " and " + t2.Name)
	// log.Println("For tree: " + t1.Name)
	// log.Println(t1.Nodes)
	// log.Println("For ori:" + t2.Name)
	// log.Println(t2.Nodes)
	// log.Printf("t1: %p, t2: %p", t1, t2)
	var isEqual bool = true
	// log.Println("Before testing name and len")
	if t1.Name != t2.Name || len(t1.Nodes) != len(t2.Nodes) || len(t1.SubTrees) != len(t2.SubTrees) {
		isEqual = false
	} else {
		// log.Println("After testing name and len")
		// log.Println("Len is same for " + t1.Name + " and " + t2.Name)
		// log.Printf("%d %d", len(t1.Nodes), len(t2.Nodes))
		for i := 0; i < len(t1.Nodes); i++ {
			if !isNodeEqual(t1.Nodes[i], t2.Nodes[i]) {
				log.Println("Node does not equal")
				return false
			}
		}
		for i := 0; i < len(t1.SubTrees); i++ {
			isEqual = IsTreeEqual(t1.SubTrees[i], t2.SubTrees[i])
			if !isEqual {
				return isEqual
			}
		}
	}
	return isEqual
}

func (tree *Tree) IsUpdate() bool {
	// log.Println("Before get ori")
	ori := getTree(tree.Name)
	// if ori == nil {
	// 	log.Fatal("Fail to get tree " + tree.Name)
	// }
	// log.Println("After get ori")
	// log.Println("For ori")
	// spew.Dump(ori)
	// log.Println("For cur")
	// spew.Dump(tree)
	// log.Printf("Cur: %p, Ori: %p", tree, ori)
	// log.Println("For ori")
	// spew.Dump(ori.SubTrees[1].SubTrees[0].Nodes)
	// log.Println("For tree")
	// spew.Dump(tree.SubTrees[1].SubTrees[0].Nodes)
	// log.Println("For tree: " + tree.Name)
	// log.Println("The number of node for Ori is: ")
	// log.Println(len(ori.Nodes))
	// log.Println("The number of node for tree is: ")
	// log.Println(len(tree.Nodes))
	return !IsTreeEqual(tree, ori)
}
