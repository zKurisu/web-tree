package utils

func (tree Tree) GetTreeName() string {
	return tree.Name
}

func (tree Tree) GetAllSubtree() []*Tree {
	return tree.SubTrees
}

func (tree Tree) GetSubtreesName() []string {
	list := []string{}
	for _, subtree := range tree.GetAllSubtree() {
		if subtree.Name != "" {
			list = append(list, subtree.Name)
		}
	}
	return list
}

func (tree Tree) GetNodes() []*Node {
	return tree.Nodes
}
