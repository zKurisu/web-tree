package ui

import (
	"github.com/charmbracelet/lipgloss"
	"strconv"
	"strings"
	"web-tree/utils"
)

var count = 0

func (m Model) browserView() string {
	return ""
}

func (m Model) searchView() string {
	return m.searchInput.View()
}

func (m Model) suggestionListView() string {
	var b strings.Builder
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
		suggestion = suggestion
		suggestion = searchAndRender(suggestion, renderTargets, styles)
		if i == m.sugSelected.index {
			suggestion = suggestionSelectedStyle.Render(suggestionSelectedSuroundLeft) +
				suggestion +
				suggestionSelectedStyle.Render(suggestionSelectedSuroundRight)
		}
		b.WriteString(suggestion + "\n")
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
	rendered := []string{}

	if len(t.GetAllSubtree()) > 0 {
		for _, sub := range t.GetAllSubtree() {
			if y == m.subSelected.y && x == m.subSelected.x {
				rendered = append(rendered, treeBoxSelectedStyle.Render(sub.GetTreeName()))
			} else {
				rendered = append(rendered, m.getRenderedTreeName(sub))
			}
			x++
		}
	}
	if len(t.GetNodes()) > 0 {
		for _, node := range t.GetNodes() {
			// Is selected in search mode
			switch selectedNode := m.subSelected.content.(type) {
			case *utils.Node:
				if selectedNode == node {
					m.subSelected.x = x
					m.subSelected.y = y
					m.subSelected.content = nil
				}
			}

			if y == m.subSelected.y && x == m.subSelected.x {
				m.detail = true
				rendered = append(rendered, nodeBoxSelected.Render(m.getNodeView(node)))
				m.detail = false
			} else {
				rendered = append(rendered, nodeBoxStyle.Render(m.getNodeView(node)))
			}
			x++
		}
	}

	m.subMsgs.ylen = append(m.subMsgs.ylen, x)
	b.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, rendered...))
	b.WriteString("\n")

	if len(t.GetAllSubtree()) > 0 {
		y++
		b.WriteString(m.getTreeView(t.GetAllSubtree()[0], y))
	}

	return b.String()
}

func (m Model) getRenderedTreeName(t *utils.Tree) string {
	return treeBoxStyle.Render(t.GetTreeBaseName())
}

func (m Model) getNodeView(n *utils.Node) string {
	s := ""
	switch {
	case m.detail:
		s = s + m.expandNodeView(n)
	case m.toggle:
		s = s + aliasStyle.Render(n.Alias[0])
	default:
		s = s + linkStyle.Render(n.Link[0])
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

	b.WriteString(strings.Repeat("\n", paginatorHeight))
	b.WriteString(m.paginator.View())
	return b.String()
}

func (m Model) viewportView() string {
	// return viewBoxStyle.Render(m.viewport.View())
	return m.viewport.View()
}

func (m Model) singleTreeView() string {
	return ""
}

func (m Model) helpView() string {
	return m.help.View(m.keymap)
}

func (m Model) headerView() string {
	return ""
}

func (m Model) bodyView() string {
	return ""
}

func (m Model) footerView() string {
	return ""
}

func (m Model) debugView() string {
	start, end := m.paginator.GetSliceBounds(len(m.tabs))
	s := ""
	switch m.mode {
	case search:
		s = "search"
	case advancedSearch:
		s = "advancedSearch"
	case display:
		s = "display"
	}
	return s + " " + m.debug + "\n" + "start: " + strconv.Itoa(start) +
		"\n" + "end: " + strconv.Itoa(end) +
		// "\n" + "selected index: " + strconv.Itoa(m.tabSelected.index) +
		// "\n" + "point x, y: " + strconv.Itoa(m.subSelected.x) + " " + strconv.Itoa(m.subSelected.y) +
		// "\n" + strconv.Itoa(len(m.subMsgs.ylen))
		"\n" + strconv.Itoa(count)
}

func (m Model) View() string {
	if !m.ready {
		return "Initializing..."
	}

	var searchBox strings.Builder
	var displayBox strings.Builder
	var replaceLen int = 0
	searchBox.WriteString(m.searchView())
	searchBox.WriteString("\n")
	// searchBox.WriteString(m.suggestionListView())
	// searchBox.WriteString("\n")

	displayBox.WriteString(m.treeTabView())
	displayBox.WriteString("\n")
	displayBox.WriteString(m.viewportView())
	displayBox.WriteString("\n")
	displayBox.WriteString(m.paginatorView())
	displayBox.WriteString("\n")
	displayBox.WriteString(m.debugView())
	displayBox.WriteString("\n")

	suggestionListBytes := []byte(m.suggestionListView())
	displayBytes := []byte(displayBox.String())
	if len(suggestionListBytes) > len(displayBytes) {
		replaceLen = len(displayBytes)
	} else {
		replaceLen = len(suggestionListBytes)
	}
	displayBytes = append(suggestionListBytes[:replaceLen], displayBytes[replaceLen:]...)

	if m.helpToggle == true {
		count++
		helpBytes := []byte(m.helpView())
		if len(helpBytes) > len(displayBytes) {
			replaceLen = len(displayBytes)
		} else {
			replaceLen = len(helpBytes)
		}
		displayBytes = append(helpBytes[:replaceLen], displayBytes[replaceLen:]...)
	}

	s := searchBox.String() + string(displayBytes) + m.helpView()

	// return lipgloss.PlaceHorizontal(150, 0.5, b.String())
	return s
}
