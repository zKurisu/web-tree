package ui

import (
	"github.com/charmbracelet/lipgloss"
	"log"
	"math"
	"os/exec"
	"regexp"
	"sort"
	// "strconv"
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

func (m Model) getPageNumber(index int) int {
	tabPerPage := m.paginator.PerPage
	return index/tabPerPage + 1
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

	// m.debug = strconv.Itoa(verticalMarginHeight)

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

func (m Model) getAveReducedWidth(lineWidth int) int {
	width := 0.0
	totalXInLine := m.getTotalXInLine()
	if totalXInLine > 0 {
		width = float64(lineWidth-m.viewport.Width) / float64(totalXInLine)
	}

	if width > 0 {
		return int(math.Ceil(width))
	}
	return 0
}

func (m Model) getCurLineAveReducedWidth(lineWidth int) int {
	width := 0.0
	totalXInLine := m.getTotalXInLine()
	if totalXInLine > 1 {
		width = float64(lineWidth-m.viewport.Width) / float64(totalXInLine-1)
	}

	if width > 0 {
		return int(math.Ceil(width))
	}
	return 0
}

func (m Model) reduceWidth(components []string, width int) []string {
	reducedComponent := []string{}
	totalReducedWidth := width * len(components)

	for index, component := range components {
		if totalReducedWidth > 0 && index != m.subSelected.x {
			totalReducedWidth = totalReducedWidth - len(component) + 1
			component = component[:1]
		}
		// oriLen := len(component)
		// if oriLen > width {
		// 	component = component[:oriLen-width]
		// } else {
		// 	component = component[:1]
		// }
		reducedComponent = append(reducedComponent, component)
	}

	return reducedComponent
}

func (m Model) getSugCount() int {
	var totalCount int = 0
	if m.mode == search || m.mode == browser && len(m.suggestionList) != 0 {
		totalCount = len(m.suggestionList)
	} else if m.mode == advancedSearch && len(m.adsuggestionList[m.adInpSelected.index]) != 0 {
		totalCount = len(m.adsuggestionList[m.adInpSelected.index])
	}
	return totalCount
}

func (m Model) getSugPerPage() int {
	return m.viewport.Height
}

func (m *Model) jumpTabPage(treeName string) {
	root := utils.RootTree
	m.tabs = root.GetAllSubtreeName()
	m.paginator.SetTotalPages(len(root.GetAllSubtreeName()))

	oriPageNumber := m.getPageNumber(m.tabSelected.index)
	m.tabSelected.index = getIndex(m.tabs, treeName)
	newPageNumber := m.getPageNumber(m.tabSelected.index)
	// m.debug = strconv.Itoa(oriPageNumber) + " " + strconv.Itoa(newPageNumber)

	diff := newPageNumber - oriPageNumber
	for i := float64(0); i < math.Abs(float64(diff)); i++ {
		if diff < 0 {
			m.paginator.PrevPage()
		} else {
			m.paginator.NextPage()
		}
	}
}

func renderTreeComponents(components []string, styles []lipgloss.Style) string {
	if len(components) != len(styles) {
		log.Fatal("Length of components does not equal to renderHint!")
	}
	rendered := []string{}
	for i, style := range styles {
		component := components[i]

		rendered = append(rendered, style.Render(component))
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, rendered...)
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
