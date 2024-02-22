package ui

import (
	"github.com/charmbracelet/lipgloss"
	"strings"
)

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
	var b strings.Builder
	start, end := m.paginator.GetSliceBounds(len(m.items))
	for _, item := range m.items[start:end] {
		b.WriteString(item + strings.Repeat(" ", treeGap))
	}
	b.WriteString("\n")
	b.WriteString(m.paginator.View())
	return b.String()
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
	s := ""
	switch m.mode {
	case search:
		s = "search"
	case advancedSearch:
		s = "advancedSearch"
	case display:
		s = "display"
	}
	return s + " " + m.debug
}

func (m Model) View() string {
	if !m.ready {
		return "Initializing..."
	}
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
	s := m.searchView() + "\n" + m.suggestionListView() + "\n" + m.allTreeView() + "\n" + m.debugView()

	return lipgloss.PlaceHorizontal(150, 0.5, s)
}
