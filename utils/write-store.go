package utils

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"web-tree/conf"
)

func WriteTree() {

}

// Create a file under the data dir
func AddTree(name string) {
	if name == "" {
		log.Fatal("Can not use empty as name")
	} else if isInList(AddFileExtention(name), GetTrees()) {
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
	} else if !isInList(AddFileExtention(name), GetTrees()) {
		log.Fatal("Does not exit " + name + ".yaml")
	}
	path := filepath.Join(conf.GetStoreDir(), AddFileExtention(name))
	os.Remove(path)
}

func NewSubTree(tree *Tree, name string) {
	tree.SubTrees = append(tree.SubTrees, NewTree(name))
}

func AddSubTree(tree *Tree, subtree *Tree) {
	tree.SubTrees = append(tree.SubTrees, subtree)
}

func DelSubTree(tree *Tree, name string) {
	list := []*Tree{}
	for _, subtree := range tree.SubTrees {
		if subtree.Name != name {
			list = append(list, subtree)
		}
	}
	tree.SubTrees = list
}

func AddNode(tree *Tree, node Node) {

}

func DelNode(tree *Tree, node Node) {

}
