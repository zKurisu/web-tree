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
	if !IsRootTree(tree) {
		return errors.New("Does not root tree")
	}

	Backup(tree.Name)

	data, err := yaml.Marshal(tree)
	if err != nil {
		log.Fatal(err)
	}

	filePath := GetTreeFile(tree)
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

// Create a file under the data dir
func AddTree(name string) {
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
func DelTree(name string) {
	Backup(name)
	if name == "" {
		log.Fatal("Can not use empty as name")
	} else if !IsInList(GetAllTreeName(), AddFileExtention(name)) {
		log.Fatal("Does not exit " + name + ".yaml")
	}
	path := filepath.Join(conf.GetStoreDir(), AddFileExtention(name))
	os.Remove(path)
}

func (tree *Tree) AddNewSubTree(name string) {
	tree.SubTrees = append(tree.SubTrees, NewTree(name))
}

func (tree *Tree) AddSubTree(subtree *Tree) {
	tree.SubTrees = append(tree.SubTrees, subtree)
}

func (tree *Tree) DelSubTree(name string) {
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

func (tree *Tree) DelNode(name string) {
	list := []*Node{}
	for _, node := range tree.Nodes {
		for _, link := range node.Link {
			if link != name {
				list = append(list, node)
			}
		}
		for _, alias := range node.Alias {
			if alias != name && !IsNodeExist(list, node) {
				list = append(list, node)
			}
		}
	}
	tree.Nodes = list
}
