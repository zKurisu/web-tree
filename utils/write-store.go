package utils

import (
	"errors"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"web-tree/conf"
)

// Compare to root tree list
func (tree *Tree) WriteTree() error {
	if !RootTree.IsSubTree(tree) {
		return errors.New("It is not a sub tree of root")
	}

	Backup(tree.Name)

	data, err := yaml.Marshal(tree)
	if err != nil {
		log.Fatal(err)
	}

	filePath := getRootSubTreeFile(tree.Name)
	fd, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	if err = fd.Truncate(0); err != nil {
		log.Fatal(err)
	}

	_, err = fd.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	return nil
}

func WriteAll() error {
	// log.Println("Begin Write")
	for _, tree := range RootTree.GetAllSubtree() {
		if tree.IsUpdate() {
			// log.Println("Tree " + tree.Name + " is updated")
			if err := tree.WriteTree(); err != nil {
				return errors.New("Wrong when Write " + tree.GetTreeName())
			}
		}
		//       else {
		// 	log.Println("Tree " + tree.Name + " does not update")
		// }
	}
	return nil
}

// Create a file under the data dir
func addRootSubTree(name string) {
	if name == "" {
		log.Fatal("Can not use empty as name")
	} else if IsInList(RootTree.GetAllSubtreeName(), AddFileExtention(name)) {
		log.Fatal("Already exit " + name + ".yaml")
	}
	path := filepath.Join(conf.GetStoreDir(), AddFileExtention(name))

	pattern := regexp.QuoteMeta("%")
	re := regexp.MustCompile(pattern)

	treeTemp := conf.GetTreeTemp()
	content := re.ReplaceAllString(treeTemp, `"`+name+`"`)

	fd, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	if _, err = fd.WriteString(content); err != nil {
		log.Fatal(err)
	}
}

// Delete the file under the data dir
func delRootSubTree(name string) {
	Backup(name)
	if name == "" {
		log.Fatal("Can not use empty as name")
	} else if !IsInList(RootTree.GetAllSubtreeName(), RemoveFileExtention(name)) {
		log.Fatal("Does not exit " + name + ".yaml")
	}
	path := filepath.Join(conf.GetStoreDir(), AddFileExtention(name))
	os.Remove(path)
}

func (tree *Tree) AddNewSubTree(name string) error {
	if tree.GetTreeName() != "root" {
		name = tree.subTreeNameWrap(name)
	}
	// log.Println("Before check exist, tree is " + tree.GetTreeName())
	if IsInList(tree.GetAllSubtreeName(), name) {
		return errors.New("In function [AddNewSubTree], subtree " + name + " already exist for " + tree.GetTreeName())
	} else if name == "" {
		return errors.New("In function [AddNewSubTree], subtree name can not be empty")
	}
	// log.Println("After check exist")
	if tree.GetTreeName() == "root" {
		addRootSubTree(name)
	}
	newTree, err := NewTree(name)
	if err != nil {
		return err
	}
	tree.SubTrees = append(tree.SubTrees, newTree)
	return nil
}

func (tree *Tree) DeepAddNewSubTree(name string) error {
	levels := SplitTreeLevel(name)
	curLevel := levels[0]
	// log.Println("Before adding " + curLevel + " to tree " + tree.GetTreeName())
	tree.AddNewSubTree(curLevel)

	if len(levels) > 1 {
		remainLevels := levels[1:]
		name = strings.Join(remainLevels, "/")
		if tree.GetTreeName() == "root" {
			tree.FindSubTree(curLevel).DeepAddNewSubTree(name)
		} else {
			tree.FindSubTree(tree.subTreeNameWrap(curLevel)).DeepAddNewSubTree(name)
		}
	}
	return nil
}

func (tree *Tree) AppendSubTree(subtree *Tree) {
	tree.SubTrees = append(tree.SubTrees, subtree)
}

func (tree *Tree) DelSubTree(name string) {
	if tree.Name == "root" {
		delRootSubTree(name)
	}
	list := []*Tree{}
	for _, subtree := range tree.SubTrees {
		if subtree.Name != name {
			list = append(list, subtree)
		}
	}
	tree.SubTrees = list
}

func (tree *Tree) DeepDelSubTree(name string) error {
	levels := SplitTreeLevel(name)
	if len(levels) > 1 {
		targetTreeName := strings.Join(levels[:len(levels)-1], "/")
		targetSubTreeName := levels[len(levels)-1]
		targetTree := tree.DeepFindSubTree(targetTreeName)
		if targetTree != nil {
			targetTree.DelSubTree(targetSubTreeName)
		} else {
			return errors.New("Tree " + targetTreeName + " does not exist")
		}
	} else {
		tree.DelSubTree(name)
	}
	return nil
}

func (tree *Tree) AppendNode(node *Node) {
	tree.Nodes = append(tree.Nodes, node)
}

func (tree *Tree) DelNode(hints []string) {
	hints = RemoveDup(hints)
	delList := []*Node{}
	for _, node := range tree.Nodes {
		// log.Println("Node call")
		if node.MatchHint(hints) == len(hints) {
			delList = append(delList, node)
		}
	}

	if len(delList) == 0 {
		log.Println("Can not find the node")
	} else {
		list := []*Node{}
		for _, node := range tree.Nodes {
			if !IsNodeExist(delList, node) {
				list = append(list, node)
			}
		}
		tree.Nodes = list
	}
}

func (tree *Tree) subTreeNameWrap(name string) string {
	return tree.Name + "/" + name
}
