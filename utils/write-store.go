package utils

import (
	"errors"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"regexp"
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

	filePath := tree.GetFile()
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
	for _, tree := range RootTree.SubTrees {
		if tree.IsUpdate() {
			// log.Println("Tree " + tree.Name + " is updated")
			if err := tree.WriteTree(); err != nil {
				return errors.New("Wrong when Write " + tree.Name)
			}
		}
		//       else {
		// 	log.Println("Tree " + tree.Name + " does not update")
		// }
	}
	return nil
}

// Create a file under the data dir
func addTree(name string) {
	if name == "" {
		log.Fatal("Can not use empty as name")
	} else if IsInList(GetAllTreeName(), AddFileExtention(name)) {
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
func delTree(name string) {
	Backup(name)
	if name == "" {
		log.Fatal("Can not use empty as name")
	} else if !IsInList(GetAllTreeName(), RemoveFileExtention(name)) {
		log.Fatal("Does not exit " + name + ".yaml")
	}
	path := filepath.Join(conf.GetStoreDir(), AddFileExtention(name))
	os.Remove(path)
}

func (tree *Tree) AddNewSubTree(name string) {
	if tree.Name == "root" {
		addTree(name)
	}
	newTree, err := NewTree(name)
	if err != nil {
		log.Fatal(err)
	}
	tree.SubTrees = append(tree.SubTrees, newTree)
}

func (tree *Tree) DeepAddNewSubTree(names []string) error {
	if len(names) == 0 {
		return nil
	}
	if !IsInList(tree.GetSubtreesName(), names[0]) {
		tree.AddNewSubTree(names[0])
	}
	return tree.FindSubTree(names[0]).DeepAddNewSubTree(names[1:])
}

func (tree *Tree) AddSubTree(subtree *Tree) {
	tree.SubTrees = append(tree.SubTrees, subtree)
}

func (tree *Tree) DelSubTree(name string) {
	if tree.Name == "root" {
		delTree(name)
	}
	list := []*Tree{}
	for _, subtree := range tree.SubTrees {
		if subtree.Name != name {
			list = append(list, subtree)
		}
	}
	tree.SubTrees = list
}

func (tree *Tree) AddNode(node *Node) {
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
