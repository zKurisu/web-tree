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

func (m Model) allTreeView() string {
	return ""
}

func (m Model) singleTreeView() string {
	return ""
}

func (m Model) helpView() string {
	return ""
}

func (m Model) View() string {
	s := strings.Join(m.getItems(), " ")

	return s
}
