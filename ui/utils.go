package ui

import (
	"github.com/charmbracelet/lipgloss"
	"log"
	"os/exec"
	"regexp"
	"sort"
	"strings"
	"syscall"
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

func getLastWord(s string) string {
	// Remove space at head or tail
	s = strings.TrimSpace(s)

	// Split string by space
	words := strings.Fields(s)

	// Get last word
	if len(words) > 0 {
		return words[len(words)-1]
	}

	return ""
}

// func commandParse(input string) (string, []string) {
// 	elements := strings.Split(input, " ")
// 	return elements[0], elements[1:]
// }
//
// // Remove prefix "WT"
// func commandHandler(cmd string, args []string) bool {
// 	pattern := regexp.QuoteMeta("MT")
// 	re := regexp.MustCompile(pattern)
// 	cmd = re.ReplaceAllString(cmd, "")
//
// 	return false
// }

func removeEndSpace(s string) string {
	if s[len(s)-1] == ' ' {
		s = s[:len(s)-1]
	}
	return s
}

func getPageNumber(index int) int {
	return index/5 + 1
}

func getIndex(slice []string, s string) int {
	for i, elem := range slice {
		if s == elem {
			return i
		}
	}
	return 0
}

func openLink(browser string, link string) bool {
	cmd := exec.Command(browser, link)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	if err := cmd.Start(); err != nil {
		return false
	}
	return true
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

func (m *Model) blurConfirm() {
	m.confirm.ans.Blur()
}

func sequence(fns ...func()) {
	for _, fn := range fns {
		fn()
	}
}
