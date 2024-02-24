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

func (m Model) treeTabView() string {
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
func (m Model) getTreeView(t *utils.Tree) string {
	var b strings.Builder
	rendered := []string{}

	if len(t.GetAllSubtree()) > 0 {
		for _, sub := range t.GetAllSubtree() {
			rendered = append(rendered, m.getRenderedTreeName(sub))
		}
	}
	if len(t.GetNodes()) > 0 {
		for _, node := range t.GetNodes() {
			rendered = append(rendered, m.getNodeView(node))
		}
	}
	b.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, rendered...))
	b.WriteString("\n")

	if len(t.GetAllSubtree()) > 0 {
		b.WriteString(m.getTreeView(t.GetAllSubtree()[0]))
	}

	return b.String()
}

func (m Model) getRenderedTreeName(t *utils.Tree) string {
	return treeBoxStyle.Render(t.GetTreeName())
}

func (m Model) getNodeView(n *utils.Node) string {
	s := ""
	if m.toggle == false {
		s = s + linkStyle.Render(n.Link[0])
	} else {
		s = s + aliasStyle.Render(n.Alias[0])
	}
	return nodeBoxStyle.Render(s)
}

func (m Model) getDetailNodeView(n *utils.Node) string {
	return ""
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
		"\n" + "selected index: " + strconv.Itoa(m.tabSelected.index) +
		"\n" + "viewport YPosition: " + strconv.Itoa(m.viewport.YPosition)
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
	searchBox.WriteString("\n")
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

	s := searchBox.String() + string(displayBytes)
	// nodemsg := nodeMsg{}
	// treemsg := treeMsg{}
	// if m.sugSelected.content != nil {
	// 	switch msg := m.sugSelected.content.(type) {
	// 	case nodeMsg:
	// 		nodemsg = msg
	// 	case treeMsg:
	// 		treemsg = msg
	// 	}
	// }

	// return lipgloss.PlaceHorizontal(150, 0.5, b.String())
	return s
}
