package utils

import (
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

func IsTreeEqual(t1, t2 Tree) bool {
	if len(t1.Nodes) != len(t2.Nodes) || len(t1.SubTrees) != len(t2.SubTrees) {
		return false
	} else {
		for i := 0; i < len(t1.Nodes); i++ {
			if !isNodeEqual(t1.Nodes[i], t2.Nodes[i]) {
				return false
			}
		}
		for i := 0; i < len(t1.SubTrees); i++ {
			return IsTreeEqual(*t1.SubTrees[i], *t2.SubTrees[i])
		}
	}
	return true
}

func (tree Tree) IsUpdate() bool {
	ori := getTree(tree.Name)
	return !IsTreeEqual(tree, *ori)
}
