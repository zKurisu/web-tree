package ui

import (
	"github.com/charmbracelet/lipgloss"
	"log"
	"math"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"web-tree/utils"
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

func intSum(list []int) int {
	total := 0
	for _, value := range list {
		total += value
	}
	return total
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

func removeSpace(s string) string {
	pattern := `\s+`
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(s, "")
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

func (m *Model) getVerticalMarginHeight() int {
	browseBoxHeight := lipgloss.Height(m.browserView())
	searchBoxHeight := lipgloss.Height(m.searchView())
	treeTabHeight := lipgloss.Height(m.treeTabView())
	helpHeight := lipgloss.Height(m.helpView())
	paginatorHeight := lipgloss.Height(m.paginatorView())
	debugHeight := lipgloss.Height(m.debugView())
	// debugHeight := 0
	confirmHeight := lipgloss.Height(m.confirmView())

	verticalMarginHeightComponents := []int{
		browseBoxHeight, searchBoxHeight, treeTabHeight, helpHeight,
		paginatorHeight, debugHeight, confirmHeight,
	}

	// Newline between components
	verticalMarginHeight :=
		intSum(verticalMarginHeightComponents) + len(verticalMarginHeightComponents) - 3 // 4: include header, body, footer

	m.debug = strconv.Itoa(verticalMarginHeight)

	if m.mode != confirm {
		verticalMarginHeight = verticalMarginHeight - confirmHeight
	}
	return verticalMarginHeight
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

func (m Model) getYPerPage() int {
	nodeHeight := lipgloss.Height(nodeBoxStyle.Render("x"))
	yPerPage := m.viewport.Height / nodeHeight
	return yPerPage
}

func (m Model) getCurDepth() int {
	depth := 0
	switch content := m.subSelected.content.(type) {
	case *utils.Node:
		depth = m.subSelected.y
	case *utils.Tree:
		depth = m.subSelected.y + content.GetTreeDepth()
	}
	return depth
}

func (m Model) isLineMove() bool {
	depth := m.getCurDepth()
	yPerPage := m.getYPerPage()
	curY := m.subSelected.y

	if curY > depth-yPerPage && curY < depth {
		return true
	}
	return false
}

func (m Model) isPageUp() bool {
	yPerPage := m.getYPerPage()
	curY := m.subSelected.y
	if curY%yPerPage == 0 && curY != 0 {
		return true
	}
	return false
}

func (m Model) isPageDown() bool {
	// curY % yPerPage == 1, curY != 1
	yPerPage := m.getYPerPage()
	curY := m.subSelected.y
	if curY%yPerPage == 1 && curY != 1 {
		return true
	}
	return false
}

func (m Model) getTotalXInLine() int {
	t := m.curTree
	return len(t.GetNodes()) + len(t.GetAllSubtreeName())
}

func (m Model) getCurLineWidth() int {
	return m.curLineWidth
}

func (m Model) getAveReducedWidth() int {
	width := 0.0
	curLineWidth := m.getCurLineWidth()
	totalXInLine := m.getTotalXInLine()
	if totalXInLine > 0 {
		if m.subSelected.y == 0 || totalXInLine == 1 {
			// no selected one
			width = float64(curLineWidth-m.viewport.Width) / float64(totalXInLine)
		} else {
			// except selected one
			width = float64(curLineWidth-m.viewport.Width) / float64(totalXInLine-1)
		}
	}

	if width > 0 {
		return int(math.Ceil(width))
	}
	return 0
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

func (m *Model) blurBrowse() {
	m.browseInput.Blur()
}

func sequence(fns ...func()) {
	for _, fn := range fns {
		fn()
	}
}
