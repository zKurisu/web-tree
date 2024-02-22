package utils

func (tree Tree) GetTreeName() string {
	return tree.Name
}

func (tree Tree) GetAllSubtree() []*Tree {
	return tree.SubTrees
}

func (tree Tree) DeepGetAllSubtree() []*Tree {
	list := []*Tree{}
	if len(tree.GetAllSubtree()) != 0 {
		for _, sub := range tree.GetAllSubtree() {
			list = append(list, sub)
			list = MergeList(list, sub.DeepGetAllSubtree()).([]*Tree)
		}
	}

	return list
}

// func (tree Tree) DeepGetAllSubtreePath(curPath string) []string {
// 	list := []string{}
// 	if len(tree.SubTrees) == 0 {
// 		return list
// 	}
// 	for _, sub := range tree.SubTrees {
// 		curPath := curPath + "/" + sub.Name
// 		if curPath[0] == '/' {
// 			curPath = curPath[1:len(curPath)]
// 		}
// 		list = append(list, curPath)
// 		list = MergeList(list, sub.DeepGetAllSubtreePath(curPath)).([]string)
// 	}
//
// 	return RemoveEmp(list)
// }

func (tree Tree) GetAllSubtreeName() []string {
	list := []string{}
	for _, subtree := range tree.GetAllSubtree() {
		if subtree.Name != "" {
			list = append(list, subtree.Name)
		}
	}
	return list
}

func (tree Tree) DeepGetAllSubtreeName() []string {
	allSubtree := tree.DeepGetAllSubtree()
	list := []string{}
	for _, sub := range allSubtree {
		list = append(list, sub.GetTreeName())
	}
	return list
}

func (tree Tree) GetNodes() []*Node {
	return tree.Nodes
}

// Begin with tree path
func (tree Tree) DeepGetAllNodeWithPath() ([]*Node, []string) {
	nodeList := []*Node{}
	pathList := []string{}

	if len(tree.GetNodes()) != 0 {
		for _, node := range tree.GetNodes() {
			treePath := tree.GetTreeName()

			nodeList = append(nodeList, node)
			pathList = append(pathList, treePath)
		}
	}
	if len(tree.GetAllSubtree()) == 0 {
		return nodeList, pathList
	} else {
		for _, sub := range tree.GetAllSubtree() {
			newNodeList, newPathList := sub.DeepGetAllNodeWithPath()
			nodeList = MergeList(nodeList, newNodeList).([]*Node)
			pathList = MergeList(pathList, newPathList).([]string)
		}
	}

	return nodeList, pathList
}
