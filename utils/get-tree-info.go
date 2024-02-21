package utils

func (tree Tree) GetTreeName() string {
	return tree.Name
}

func (tree Tree) GetAllSubtree() []*Tree {
	return tree.SubTrees
}

func (tree Tree) DeepGetAllSubtreePath(curPath string) []string {
	list := []string{}
	if len(tree.SubTrees) == 0 {
		return list
	}
	for _, sub := range tree.SubTrees {
		curPath := curPath + "/" + sub.Name
		if curPath[0] == '/' {
			curPath = curPath[1:len(curPath)]
		}
		list = append(list, curPath)
		list = MergeList(list, sub.DeepGetAllSubtreePath(curPath))
	}

	return RemoveEmp(list)
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
