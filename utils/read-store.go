package utils

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"web-tree/conf"
)

var STORE_DIR = conf.GetStoreDir()

type Node struct {
	Link  []string    `yaml:"link"`
	Alias []string    `yaml:"alias"`
	Desc  []string    `yaml:"description"`
	Icon  string      `yaml:"icon"`
	Label []string    `yaml:"label"`
	Style interface{} `yaml:"style"`
}

type Tree struct {
	Name     string  `yaml:"name"`
	SubTrees []*Tree `yaml:"tree"`
	Nodes    []Node  `yaml:"nodes"`
}

type Store struct {
	Trees []Tree
}

func isInList(str string, slice []string) bool {
	for _, s := range slice {
		if str == s {
			return true
		}
	}
	return false
}

func NewTree(name string) *Tree {
	return &Tree{
		Name:     name,
		SubTrees: []*Tree{},
		Nodes:    []Node{},
	}
}

func GetTrees() []string {
	fileList := []string{}
	err := filepath.Walk(STORE_DIR, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if !info.IsDir() {
			fileList = append(fileList, RemoveFileExtention(info.Name()))
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return fileList
}

func GetTree(name string) Tree {
	var t Tree
	treeList := GetTrees()

	if isInList(name, treeList) {
		treePath := filepath.Join(STORE_DIR, name)
		yamlContent, err := os.ReadFile(treePath)
		if err != nil {
			log.Fatal(err)
		}
		if err = yaml.Unmarshal(yamlContent, &t); err != nil {
			log.Fatal(err)
		}
	}

	return t
}

func GetSubTree(tree Tree, name string) []Tree {
	list := []Tree{}
	if len(tree.SubTrees) == 0 {
		return list
	}
	for _, subtree := range tree.SubTrees {
		if subtree.Name == name {
			list = append(list, *subtree)
		}
		if len(subtree.SubTrees) != 0 {
			for _, subsubtree := range GetSubTree(*subtree, name) {
				list = append(list, subsubtree)
			}
		}
	}
	return list
}

// hint means "link" or "alias"
func GetNode(tree Tree, hint string) []Node {
	list := []Node{}
	pattern := regexp.QuoteMeta(hint)
	reg := regexp.MustCompile(pattern)

	if len(tree.Nodes) != 0 {
		for _, node := range tree.Nodes {
			listLen := len(list)
			for _, link := range node.Link {
				if reg.MatchString(link) {
					list = append(list, node)
					break
				}
			}
			if listLen == len(list) { // Doesn't find through links, then from alias
				for _, alias := range node.Alias {
					if reg.MatchString(alias) {
						list = append(list, node)
						break
					}
				}
			}
		}
	}

	if len(tree.SubTrees) != 0 {
		for _, subtree := range tree.SubTrees {
			for _, node := range GetNode(*subtree, hint) {
				list = append(list, node)
			}
		}
	}
	return list
}
