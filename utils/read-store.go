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

var RootTree = Tree{
	Name:     "root",
	SubTrees: GetAllRootSubTree(),
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

func NewTree(name string) (*Tree, error) {
	if IsNameValid(name) {
		return &Tree{
			Name:     name,
			SubTrees: []*Tree{},
			Nodes:    []*Node{},
		}, nil
	}
	return nil, errors.New("Tree name does not valid, it can not be root of empty")
}

func NewNode(links []string, alias []string, desc []string, icon string, labels []string, style interface{}) (*Node, error) {
	for _, link := range links {
		err := IsUrl(link)
		if err != nil {
			return nil, err
		}
	}
	return &Node{
		Link:  links,
		Alias: alias,
		Desc:  desc,
		Icon:  icon,
		Label: labels,
		Style: style,
	}, nil
}

func getAllRootSubtreeName() []string {
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

func getRootSubTree(name string) *Tree {
	var t *Tree = nil
	allRootSubTreeName := getAllRootSubtreeName()

	if IsInList(allRootSubTreeName, name) {
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

func GetAllRootSubTree() []*Tree {
	list := []*Tree{}
	for _, name := range getAllRootSubtreeName() {
		list = append(list, getRootSubTree(name))
	}
	return list
}

func GetRootTree() *Tree {
	return &RootTree
}

func (tree *Tree) FindSubTree(name string) *Tree {
	for _, subtree := range tree.SubTrees {
		if subtree.Name == name {
			return subtree
		}
	}
	return nil
}

func (tree *Tree) DeepFindSubTree(name string) *Tree {
	if len(tree.GetAllSubtree()) == 0 {
		return nil
	}
	t := tree.FindSubTree(name)
	if t == nil {
		for _, sub := range tree.GetAllSubtree() {
			t = sub.DeepFindSubTree(name)
			if t != nil {
				return t
			}
		}
	}
	return t
}

//	func (tree *Tree) FindAllSubTree(name string) []*Tree {
//		list := []*Tree{}
//		if len(tree.SubTrees) == 0 {
//			return list
//		}
//		for _, subtree := range tree.SubTrees {
//			if subtree.Name == name {
//				list = append(list, subtree)
//			}
//			if len(subtree.SubTrees) != 0 {
//				for _, subsubtree := range subtree.FindAllSubTree(name) {
//					list = append(list, subsubtree)
//				}
//			}
//		}
//		return list
//	}
func (tree *Tree) FindNodes(hints []string) []*Node {
	nodeList := []*Node{}
	if len(tree.Nodes) != 0 {
		for _, node := range tree.Nodes {
			// log.Println("Node call")
			if node.MatchHint(hints) != 0 {
				nodeList = append(nodeList, node)
			}
		}
	}
	return nodeList
}

// Most matched node
func (tree *Tree) FindNode(hints []string) *Node {
	hints = RemoveDup(hints)
	nodeMap := make(map[*Node]int)

	for _, node := range tree.Nodes {
		nodeMap[node] = node.MatchHint(hints)
	}

	var maxNode *Node
	var maxMatch int = 0
	for node, matchInt := range nodeMap {
		if matchInt > maxMatch {
			maxNode = node
			maxMatch = matchInt
		}
	}
	return maxNode
}

// hint means "link" or "alias"
func (tree *Tree) FindAllNode(hint string) []*Node {
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

//	func (node *Node) MatchHint(hints []string) int {
//		matched := 0
//		if len(hints) == 0 {
//			return matched
//		}
//		pattern := regexp.QuoteMeta(hints[0])
//		reg := regexp.MustCompile(pattern)
//
//		for _, link := range node.Link {
//			if reg.MatchString(link) {
//				matched = 1 + node.MatchHint(hints[1:])
//			}
//		}
//		for _, alias := range node.Alias {
//			if reg.MatchString(alias) {
//				matched = 1 + node.MatchHint(hints[1:])
//			}
//		}
//		return matched
//	}
func (node *Node) MatchHint(hints []string) int {
	matched := 0
	if len(hints) == 0 {
		return matched
	}

	for _, link := range node.Link {
		if link == hints[0] {
			matched = 1 + node.MatchHint(hints[1:])
			// log.Println(matched)
		}
		// log.Println("Link is: " + link + " -- Hint is: " + hints[0])
	}
	for _, alias := range node.Alias {
		if alias == hints[0] {
			matched = 1 + node.MatchHint(hints[1:])
			// log.Println(matched)
		}
		// log.Println("Alias is: " + alias + " -- Hint is: " + hints[0])
	}

	if matched == 0 {
		node.MatchHint(hints[1:])
	}

	return matched
}

func RemoveDup(list []string) []string {
	newList := []string{}

	for _, elem := range list {
		if !IsInList(newList, elem) {
			newList = append(newList, elem)
		}
	}
	return newList
}
