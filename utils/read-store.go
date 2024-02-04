package utils

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"web-tree/conf"
)

var RootTree = Tree{
	Name:     "root",
	SubTrees: getAllRootTree(),
	Nodes:    []*Node{},
}
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
	Nodes    []*Node `yaml:"nodes"`
}
type Store struct {
	Trees []Tree
}

func NewTree(name string) *Tree {
	return &Tree{
		Name:     name,
		SubTrees: []*Tree{},
		Nodes:    []*Node{},
	}
}

func NewNode(links []string, alias []string, desc []string, icon string, labels []string, style interface{}) *Node {
	return &Node{
		Link:  links,
		Alias: alias,
		Desc:  desc,
		Icon:  icon,
		Label: labels,
		Style: style,
	}
}

func GetAllTreeName() []string {
	list := []string{}
	pattern := regexp.QuoteMeta(`.yaml`)
	re := regexp.MustCompile(pattern + `$`)

	err := filepath.Walk(STORE_DIR, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if !info.IsDir() {
			if re.MatchString(info.Name()) {
				list = append(list, RemoveFileExtention(info.Name()))
			}
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return list
}

func getTree(name string) *Tree {
	var t *Tree
	treeList := GetAllTreeName()

	if IsInList(treeList, name) {
		treePath := filepath.Join(STORE_DIR, AddFileExtention(name))
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

func getAllRootTree() []*Tree {
	list := []*Tree{}
	for _, name := range GetAllTreeName() {
		list = append(list, getTree(name))
	}
	return list
}

func GetTree(name string) *Tree {
	if IsTreeExist(name) {
		for _, tree := range RootTree.GetAllSubtree() {
			if tree.Name == name {
				return tree
			}
		}
	}
	return nil
}

func (tree Tree) FindSubTree(name string) *Tree {
	for _, subtree := range tree.SubTrees {
		if subtree.Name == name {
			return subtree
		}
	}
	return nil
}

func (tree Tree) FindAllSubTree(name string) []*Tree {
	list := []*Tree{}
	if len(tree.SubTrees) == 0 {
		return list
	}
	for _, subtree := range tree.SubTrees {
		if subtree.Name == name {
			list = append(list, subtree)
		}
		if len(subtree.SubTrees) != 0 {
			for _, subsubtree := range subtree.FindAllSubTree(name) {
				list = append(list, subsubtree)
			}
		}
	}
	return list
}

func (tree Tree) FindNode(hint string) *Node {
	pattern := regexp.QuoteMeta(hint)
	reg := regexp.MustCompile(pattern)

	if len(tree.Nodes) != 0 {
		for _, node := range tree.Nodes {
			for _, link := range node.Link {
				if reg.MatchString(link) {
					return node
				}
			}
			for _, alias := range node.Alias {
				if reg.MatchString(alias) {
					return node
				}
			}
		}
	}
	return nil
}

// hint means "link" or "alias"
func (tree Tree) FindAllNode(hint string) []*Node {
	list := []*Node{}
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
			for _, node := range subtree.FindAllNode(hint) {
				list = append(list, node)
			}
		}
	}
	return list
}
