package utils

import (
	"log"
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

func IsRootTree(tree *Tree) bool {
	for _, v := range AllRootTree {
		if v == tree {
			return true
		}
	}
	return false
}
