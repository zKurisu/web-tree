package ui

import (
	"strings"
)

func (m Model) browserView() string {
	return ""
}

func (m Model) searchView() string {
	return ""
}

func (m Model) suggestionListView() string {
	var b strings.Builder
	for i, suggestion := range m.suggestionList {
		if i == m.sugSelectedIndex {
			suggestion = suggestionSelectedStyle.Render(suggestion)
		}
		b.WriteString(suggestion + "\n")
	}
	return b.String()
}

func (m Model) allTreeView() string {
	return ""
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

func (m Model) View() string {
	if !m.ready {
		return "Initializing..."
	}
	s := m.searchInput.View() + "\n" + m.suggestionListView()

	return s
}
