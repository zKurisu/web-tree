package utils

func GetTreeName(tree Tree) string {
	return tree.Name
}

func GetSubtrees(tree Tree) []*Tree {
	return tree.SubTrees
}

func GetSubtreesName(tree Tree) []string {
	list := []string{}
	for _, subtree := range GetSubtrees(tree) {
		if subtree.Name != "" {
			list = append(list, subtree.Name)
		}
	}
	return list
}

func GetNodes(tree Tree) []Node {
	return tree.Nodes
}

func IsTreeExist(name string) bool {
	if isInList(name, GetTrees()) {
		return true
	} else {
		return false
	}
}
