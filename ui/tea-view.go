package ui

import (
	"github.com/charmbracelet/lipgloss"
	"strconv"
	"strings"
	"web-tree/utils"
)

var suggesCount, viewportCount = 0, 0

func (m Model) browserView() string {
	var s string
	s = m.browseInput.View()
	return s
}

func (m Model) searchView() string {
	var s string
	switch m.mode {
	case search:
		s = m.searchInput.View()
	case advancedSearch:
		for i := range m.adSearchInput {
			s = s + m.adSearchInput[i].View() + "\n"
		}
		s = s + adSearchSubmitStyle.Render("Search")
	case add:
		for i := range m.addInput {
			s = s + m.addInput[i].View() + "\n"
		}
		s = s + adSearchSubmitStyle.Render("Add")
	default:
		switch m.lastMode {
		case search:
			s = m.searchInput.View()
		case advancedSearch:
			for i := range m.adSearchInput {
				s = s + m.adSearchInput[i].View() + "\n"
			}
			s = s + adSearchSubmitStyle.Render("Search")
		case add:
			for i := range m.addInput {
				s = s + m.addInput[i].View() + "\n"
			}
			s = s + adSearchSubmitStyle.Render("Add")
		default:
			s = m.searchInput.View()
		}
	}
	return s
}

func (m *Model) suggestionListView() string {
	var b strings.Builder
	switch m.mode {
	case browser:
		input := m.browseInput.Value()
		renderTargets := []string{
			input,
		}
		styles := []lipgloss.Style{
			suggestionMatchedStyle,
		}
		for i, suggestion := range m.suggestionList {
			suggestion = searchAndRender(suggestion, renderTargets, styles)
			if i == m.sugSelected.index {
				suggestion = suggestionSelectedStyle.Render(suggestionSelectedSuroundLeft) +
					suggestion +
					suggestionSelectedStyle.Render(suggestionSelectedSuroundRight)
			}
			b.WriteString(suggestion + "\n")
		}
	case search:
		input := m.searchInput.Value()
		renderTargets := []string{
			treePrefix, nodePrefix, input,
			nameHint, linkHint, aliasHint,
			nodePathHint,
		}
		styles := []lipgloss.Style{
			suggestionTreeStyle, suggestionNodeStyle, suggestionMatchedStyle,
			suggestionQuoteStyle, suggestionQuoteStyle, suggestionQuoteStyle,
			suggestionQuoteStyle,
		}
		for i, suggestion := range m.suggestionList {
			suggestion = searchAndRender(suggestion, renderTargets, styles)
			if i == m.sugSelected.index {
				suggestion = suggestionSelectedStyle.Render(suggestionSelectedSuroundLeft) +
					suggestion +
					suggestionSelectedStyle.Render(suggestionSelectedSuroundRight)
			}
			b.WriteString(suggestion + "\n")
		}
	case advancedSearch:
		i := m.adInpSelected.index
		if i == len(m.adSearchInput) {
			break
		}
		for j, suggestion := range m.adsuggestionList[i] {
			if j == m.sugSelected.index {
				suggestion = suggestionSelectedStyle.Render(suggestionSelectedSuroundLeft) +
					suggestion +
					suggestionSelectedStyle.Render(suggestionSelectedSuroundRight)
			}
			b.WriteString(suggestion + "\n")
		}
	}
	return b.String()
}

func (m Model) allTreeView() string {
	return ""
}

func (m *Model) treeTabView() string {
	var b strings.Builder
	renderedTabs := []string{}

	start, end := m.paginator.GetSliceBounds(len(m.tabs))
	for i, item := range m.tabs[start:end] {
		if i+start == m.tabSelected.index {
			renderedTabs = append(renderedTabs, treeTabBoxSelectedStyle.Render(item))
		} else {
			renderedTabs = append(renderedTabs, treeTabBoxStyle.Render(item))
		}
	}

	b.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...))
	return b.String()
}

// not conclude tree name itself, first subtree, then node
func (m *Model) getTreeView(t *utils.Tree, y int) string {
	var b strings.Builder
	var x int = 0
	var nextTreeLock = 0
	var nodeSelected = 0
	var nextree *utils.Tree
	renderedLine := ""
	components := []string{}
	styles := []lipgloss.Style{}

	if len(t.GetAllSubtree()) > 0 {
		for _, sub := range t.GetAllSubtree() {
			if len(m.preSelectedTree) > y {
				if x == m.preSelectedTree[y].x && y == m.preSelectedTree[y].y {
					nextTreeLock = 1
					nextree = sub
				}
			}

			var style lipgloss.Style
			component := ""
			if x == m.subSelected.x && y == m.subSelected.y {
				m.curTree = t // Get fathertree address, using for deletion
				m.subSelected.content = sub
				nextTreeLock = 1
				nextree = sub
				treeName := sub.GetTreeName()

				style = treeBoxSelectedStyle
				switch m.mode {
				case edit:
					height := lipgloss.Height(treeName)
					width := lipgloss.Width(treeName + "          ")
					m.textarea.SetWidth(width)
					m.textarea.SetHeight(height)
					component = m.textarea.View()
				default:
					component = treeName
				}
			} else {
				if nextTreeLock == 0 {
					nextree = t.GetAllSubtree()[0]
				}

				style = treeBoxStyle
				component = sub.GetTreeBaseName()
			}

			styles = append(styles, style)
			components = append(components, component)
			x++
		}
	}
	if len(t.GetNodes()) > 0 {
		for _, node := range t.GetNodes() {
			// Is selected in search mode
			switch selectedNode := m.subMsgs.searchedContent.(type) {
			case *utils.Node:
				if selectedNode == node {
					m.subSelected.x = x
					m.subSelected.y = y
					m.subMsgs.searchedContent = nil
				}
			}

			var style lipgloss.Style
			component := ""
			if y == m.subSelected.y && x == m.subSelected.x {
				nodeSelected = 1
				m.curTree = t // Get fathertree address, using for deletion
				m.subSelected.content = node
				m.detail = true

				switch m.mode {
				case edit:
					height := lipgloss.Height(m.getNodeView(node))
					width := lipgloss.Width(strings.Join(node.Link, " ") + "          ")
					m.textarea.SetWidth(width)
					m.textarea.SetHeight(height)
					component = m.textarea.View()
				default:
					component = m.getNodeView(node)
				}
				style = nodeBoxSelected

				m.detail = false
			} else {
				style = nodeBoxStyle
				component = m.getNodeView(node)
			}

			styles = append(styles, style)
			components = append(components, component)
			x++
		}
	}

	m.subMsgs.ylen = append(m.subMsgs.ylen, x)
	renderedLine = renderTreeComponents(components, styles)
	lineWidth := lipgloss.Width(renderedLine)
	reducedWidth := 0
	if m.subSelected.y == y {
		reducedWidth = m.getCurLineAveReducedWidth(lineWidth)
	} else {
		reducedWidth = m.getAveReducedWidth(lineWidth)
	}
	newComponents := reduceWidth(components, reducedWidth)
	renderedLine = renderTreeComponents(newComponents, styles)

	b.WriteString(renderedLine)
	b.WriteString("\n")

	if len(t.GetAllSubtree()) > 0 && nodeSelected == 0 {
		y++
		b.WriteString(m.getTreeView(nextree, y))
	}

	// m.debug = "x: " + strconv.Itoa(x) + " y: " + strconv.Itoa(y)
	return b.String()
}

func (m Model) getRenderedTreeName(t *utils.Tree) string {
	s := t.GetTreeBaseName()
	// aveReducedWidth := m.getAveReducedWidth()
	// s = s[:len(s)-aveReducedWidth]
	return treeBoxStyle.Render(s)
}

func (m Model) getNodeView(n *utils.Node) string {
	s := ""
	// aveReducedWidth := m.getAveReducedWidth()
	switch {
	case m.detail:
		s = m.expandNodeView(n)
	case m.toggle:
		s = n.GetNodeAlias()[0]

		// hint := n.Alias[0]
		// s = aliasStyle.Render(hint)
		// s = aliasStyle.Render(hint[:len(hint)-aveReducedWidth])
		// s = s[:len(s)-aveReducedWidth]
	default:
		s = n.GetNodeLinks()[0]

		// hint := n.Link[0]
		// s = linkStyle.Render(hint)
		// s = linkStyle.Render(hint[:len(hint)-aveReducedWidth])
		// s = s[:len(s)-aveReducedWidth]
	}

	return s
}

func (m Model) expandNodeView(n *utils.Node) string {
	var b strings.Builder
	b.WriteString("links: " + strings.Join(n.Link, " "))
	b.WriteString("\n")
	b.WriteString("alias: " + strings.Join(n.Alias, " "))
	b.WriteString("\n")
	b.WriteString("description: " + strings.Join(n.Desc, " "))
	b.WriteString("\n")
	b.WriteString("label: " + strings.Join(n.Label, " "))
	return b.String()
}

func (m Model) paginatorView() string {
	var b strings.Builder

	// b.WriteString(strings.Repeat("\n", paginatorHeight))
	b.WriteString(m.paginator.View())
	return b.String()
}

func (m Model) confirmView() string {
	var b strings.Builder

	if m.mode == confirm {
		b.WriteString(m.confirm.ans.View())
	}
	return b.String()
}

func (m Model) viewportView() string {
	return viewBoxStyle.Render(m.viewport.View() + "\n" + m.indexView())
	// return m.viewport.View()
}

func (m Model) indexView() string {
	var indexStr string
	var startX int
	sugCount := m.getSugCount()
	if sugCount != 0 {
		indexStr = strconv.Itoa(m.sugSelected.index+1) + "/" + strconv.Itoa(sugCount)
	} else {
		index := 0
		depth := 0
		switch content := m.subSelected.content.(type) {
		case *utils.Node:
			index = m.subSelected.y
			depth = index
		case *utils.Tree:
			index = m.subSelected.y
			depth = index + content.GetTreeDepth()
		}

		indexStr = strconv.Itoa(index) + "/" + strconv.Itoa(depth)
	}
	startX = m.viewport.Width - len(indexStr)

	return strings.Repeat(" ", startX) + indexStr
}

func (m Model) singleTreeView() string {
	return ""
}

func (m Model) helpView() string {
	return m.help.View(m.keymap)
}

func (m Model) headerView() string {
	// var searchBox strings.Builder
	// searchBox.WriteString(m.browserView()) // Browser box is needed
	// searchBox.WriteString("\n")
	// searchBox.WriteString(m.searchView())
	// return searchBox.String()
	return lipgloss.JoinVertical(lipgloss.Top, m.browserView(), m.searchView())
}

func (m Model) bodyView() string {
	// var bodyBox strings.Builder

	// bodyBox.WriteString(m.treeTabView())
	// bodyBox.WriteString("\n")
	// bodyBox.WriteString(m.viewportView())

	// delPopWinBytes := []byte(m.delPopWinView())
	// displayBytes := []byte(displayBox.String())
	// suggesCount = len(suggestionListBytes)
	// viewportCount = len(displayBytes)
	// if len(suggestionListBytes) > len(displayBytes) {
	// 	replaceLen = len(displayBytes)
	// } else {
	// 	replaceLen = len(suggestionListBytes)
	// }
	// displayBytes = append(suggestionListBytes[:replaceLen], displayBytes[replaceLen:]...)
	// return bodyBox.String()
	return lipgloss.JoinVertical(lipgloss.Top, m.treeTabView(), m.viewportView())
}

func (m Model) footerView() string {
	// var footerBox strings.Builder
	// footerBox.WriteString(m.paginatorView())
	// footerBox.WriteString("\n")
	// footerBox.WriteString(m.debugView())
	// footerBox.WriteString("\n")
	// footerBox.WriteString(m.confirmView())
	// footerBox.WriteString("\n")
	// footerBox.WriteString(m.helpView())
	// return footerBox.String()
	return lipgloss.JoinVertical(lipgloss.Top, m.paginatorView(), m.debugView(),
		m.confirmView(), m.helpView())
}

func (m Model) debugView() string {
	// start, end := m.paginator.GetSliceBounds(len(m.tabs))
	lastMode := ""
	s := ""
	switch m.mode {
	case search:
		s = "search"
	case advancedSearch:
		s = "advancedSearch"
	case display:
		s = "display"
	case add:
		s = "add"
	case edit:
		s = "edit"
	case confirm:
		s = "confirm"
	case browser:
		s = "browser"
	}

	switch m.lastMode {
	case search:
		lastMode = "search"
	case advancedSearch:
		lastMode = "advancedSearch"
	case display:
		lastMode = "display"
	case add:
		lastMode = "add"
	case edit:
		lastMode = "edit"
	case confirm:
		lastMode = "confirm"
	case browser:
		lastMode = "browser"
	}

	// strYlens := []string{}
	// for _, ylen := range m.subMsgs.ylen {
	// 	strYlens = append(strYlens, strconv.Itoa(ylen))
	// }
	// // strYlen := strings.Join(strYlens, " ")

	// preStr := ""
	// for _, point := range m.preSelectedTree {
	// 	preStr += "{" + strconv.Itoa(point.x) + "," + strconv.Itoa(point.y) + "}"
	// }
	// searchBoxHeight := lipgloss.Height(m.searchView())
	// totalReduce := 0

	// if m.ready {
	// 	if m.isPageDown() {
	// 		m.debug = "pagedown"
	// 	} else {
	// 		m.debug = "wait"
	// 	}
	// 	totalReduce = m.getAveReducedWidth()
	// }

	// m.debug = strconv.Itoa(m.viewport.YOffset)
	// if m.copy {
	// 	m.debug = "copy.."
	// }

	nodePath := ""
	switch content := m.sugSelected.content.(type) {
	case nodeMsg:
		nodePath = content.path
		hint := append(content.link, content.alias...)
		m.subMsgs.searchedContent = utils.RootTree.DeepFindSubTree(nodePath).FindNode(hint)
		// tabTarget = strings.Split(content.path, "/")[0]
	case treeMsg:
		// tabTarget = strings.Split(content.path, "/")[0]
	}
	// posiList := [][]int{}
	// posiX, posiY := 0, 0
	// treeName := ""
	//nodePath := ""
	//switch content := m.subSelected.content.(type) {
	//case *utils.Tree:
	//	// posiX, posiY = m.root.DeepGetTreePosi(content.GetTreeName(), 0)
	//	// posiList = m.root.GetTreePosiList(content.GetTreeName(), [][]int{})
	//	// treeName = content.GetTreeName()
	//}
	// posiListView := ""
	// for _, elem := range posiList {
	// 	posiListView += "[" + strconv.Itoa(elem[0]) + " " + strconv.Itoa(elem[1]) + "] "
	// }

	// return s + " " + m.debug + "\n" + "start: " + strconv.Itoa(start) +
	// 	"\n" + "end: " + strconv.Itoa(end) +
	// 	"\n" + "selected index: " + strconv.Itoa(m.tabSelected.index) +
	// 	"\n" + "point x, y: " + strconv.Itoa(m.subSelected.x) + " " + strconv.Itoa(m.subSelected.y) +
	// 	"\n" + strconv.Itoa(len(m.subMsgs.ylen))
	// 	"\n" + strconv.Itoa(count)
	return "last: " + lastMode + " current: " + s + "\n" + nodePath + "\n" +
		m.debug
	// "PreSelectedTree.x:" + strconv.Itoa(m.preSelectedTree.x) + " PreSelectedTree.y:" + strconv.Itoa(m.preSelectedTree.y) + "\n" +
	// "PreSelectedTree: " + preStr + "\n" +
	// "posiX: " + strconv.Itoa(posiX) + "posiY: " + strconv.Itoa(posiY) + "\n" +
	// "posiListView: " + posiListView + "\n" +
	// "TreeName: " + treeName + "\n" +
	// "subSelected.x:" + strconv.Itoa(m.subSelected.x) + " subSelected.y:" + strconv.Itoa(m.subSelected.y) + "\n" +
	// "ylen: " + strconv.Itoa(len(m.subMsgs.ylen)) + "\n" +
	// "all ylen: " + strYlen + "\n" +
	// m.curTree.Name + "\n" +
	// strconv.Itoa(len(m.tabs)) + "\n" +
	// "Width: " + strconv.Itoa(m.winMsgs.Width) + " Height: " + strconv.Itoa(m.winMsgs.Height) + "\n" +
	// "searchBoxHeight: " + strconv.Itoa(searchBoxHeight) + "\n" + strconv.Itoa(m.getYPerPage()) + "\n" +
	// "totalX: " + strconv.Itoa(m.getTotalXInLine()) + " " + "totalW: " + strconv.Itoa(m.getCurLineWidth()) + "\n" +
	// "totalReduce: " + strconv.Itoa(totalReduce) + "\n" +
	// strconv.Itoa(m.getSugPerPage()) + " sugSe: " + strconv.Itoa(m.sugSelected.index+1)
	// "suggestion string: " + m.debug + "\n" +
	// "sugCount: " + strconv.Itoa(suggesCount) + " BodyCount: " + strconv.Itoa(viewportCount)
}

func (m Model) View() string {
	if !m.ready {
		return "Initializing..."
	}

	header := m.headerView()
	body := m.bodyView()
	footer := m.footerView()

	// if m.helpToggle == true {
	// 	count++
	// 	helpBytes := []byte(m.helpView())
	// 	if len(helpBytes) > len(displayBytes) {
	// 		replaceLen = len(displayBytes)
	// 	} else {
	// 		replaceLen = len(helpBytes)
	// 	}
	// 	displayBytes = append(helpBytes[:replaceLen], displayBytes[replaceLen:]...)
	// }

	// s := searchBox.String() + "\n" + string(displayBytes) + "\n" + m.confirmView() + "\n" + m.helpView()
	s := lipgloss.JoinVertical(0, header, body, footer)

	// return lipgloss.PlaceHorizontal(100, 0.5, s)
	return s
}
