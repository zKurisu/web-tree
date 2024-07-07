package ui

import (
	"github.com/charmbracelet/lipgloss"
	"log"
	"regexp"
	"sort"
	"strings"
)

func Fuzzy(src string, targets []string) []string {
	if src == "" {
		return nil
	}

	list := []string{}
	sort.Strings(targets)

	pattern := regexp.QuoteMeta(src)
	re := regexp.MustCompile(pattern)

	for _, elem := range targets {
		if re.MatchString(elem) {
			list = append(list, elem)
		}
	}
	return list
}

// items there is all subtrees of root
func (m Model) getItems() []string {
	return m.root.GetAllSubtreeName()
}

func getTreeMsg(s string) treeMsg {
	treePathPattern := regexp.QuoteMeta(treePrefix) + `(.*)` + regexp.QuoteMeta(nameHint)

	treePathRe := regexp.MustCompile(treePathPattern)

	return treeMsg{
		path: treePathRe.FindStringSubmatch(s)[1],
	}
}

func getNodeMsg(s string) nodeMsg {
	nodePathPattern := regexp.QuoteMeta(nodePrefix) + `(.*)` + regexp.QuoteMeta(nodePathHint)
	nodelinkPattern := regexp.QuoteMeta(nodePathHint) + `(.*)` + regexp.QuoteMeta(linkHint)
	nodealiasPattern := regexp.QuoteMeta(linkHint) + `(.*)` + regexp.QuoteMeta(aliasHint)

	nodePathRe := regexp.MustCompile(nodePathPattern)
	nodeLinkRe := regexp.MustCompile(nodelinkPattern)
	nodealiasRe := regexp.MustCompile(nodealiasPattern)

	return nodeMsg{
		path:  nodePathRe.FindStringSubmatch(s)[1],
		link:  strings.Split(nodeLinkRe.FindStringSubmatch(s)[1], linkSep),
		alias: strings.Split(nodealiasRe.FindStringSubmatch(s)[1], aliasSep),
	}
}

func searchAndRender(src string, targets []string, styles []lipgloss.Style) string {
	if len(targets) != len(styles) {
		log.Fatal("In function [searchAndRender], the length of targets should be the same with styles")
	}
	for i := 0; i < len(targets); i++ {
		pattern := regexp.QuoteMeta(targets[i])
		re := regexp.MustCompile(pattern)
		src = re.ReplaceAllString(src, styles[i].Render(targets[i]))
	}

	return src
}

func removeEndSpace(s string) string {
	if s[len(s)-1] == ' ' {
		s = s[:len(s)-1]
	}
	return s
}

func (m *Model) blurSearch() {
	m.searchInput.Blur()
}

func (m *Model) blurAdsearch() {
	for i := range m.adSearchInput {
		m.adSearchInput[i].Blur()
	}
}

func (m *Model) blurAddInput() {
	for i := range m.addInput {
		m.addInput[i].Blur()
	}
}

func (m *Model) blurTextarea() {
	m.textarea.Blur()
}

func sequence(fns ...func()) {
	for _, fn := range fns {
		fn()
	}
}
