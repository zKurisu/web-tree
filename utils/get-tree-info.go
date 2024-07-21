package utils

import (
	"strings"
)

func (tree Tree) GetTreeName() string {
	return tree.Name
}

func (tree Tree) GetTreeBaseName() string {
	levels := SplitTreeLevel(tree.Name)
	name := tree.GetTreeName()
	if len(levels) > 1 {
		name = levels[len(levels)-1]
	}
	return name
}

func (tree Tree) GetFatherName() string {
	levels := strings.Split(tree.GetTreeName(), "/")
	fatherName := strings.Join(levels[0:len(levels)-1], "/")
	if fatherName == "" {
		return "root"
	} else {
		return fatherName
	}
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

func (t Tree) DeepGetTreePosi(tName string, initY int) (int, int) {
	x := 0
	y := initY

	levels := SplitTreeLevel(tName)
	targetTreeName := t.subTreeNameWrap(levels[0])
	for i, subTree := range t.GetAllSubtree() {
		if subTree.GetTreeName() == targetTreeName {
			if len(levels) > 1 {
				y += 1
				return subTree.DeepGetTreePosi(strings.Join(levels[1:], "/"), y)
			} else {
				x = i
			}
		}
	}

	return x, y
}

func (t Tree) GetSubTreePosi(tName string) (int, int) {
	x := 0
	y := len(SplitTreeLevel(tName)) - 1
	for i, subTree := range t.GetAllSubtree() {
		if subTree.GetTreeName() == tName {
			x = i
		}
	}
	return x, y
}

func (t Tree) GetTreePosiList(tName string, posiList [][]int, treeList []*Tree) ([][]int, []*Tree) {
	levels := SplitTreeLevel(tName)
	targetTreeName := t.subTreeNameWrap(levels[0])

	for _, subTree := range t.GetAllSubtree() {
		if subTree.GetTreeName() == targetTreeName {
			x, y := t.GetSubTreePosi(targetTreeName)
			posiList = append(posiList, []int{x, y})
			treeList = append(treeList, subTree)
			if len(levels) > 1 {
				return subTree.GetTreePosiList(strings.Join(levels[1:], "/"), posiList, treeList)
			}
		}
	}
	return posiList, treeList
}

func (t Tree) GetTreeDepth() int {
	depth := 0

	for len(t.GetAllSubtree()) > 0 {
		t = *t.GetAllSubtree()[0]
		depth++
	}
	if len(t.GetNodes()) > 0 {
		depth++
	}
	return depth
}
