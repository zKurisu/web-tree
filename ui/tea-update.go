package ui

import (
	"github.com/charmbracelet/bubbles/key"
	// "github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strconv"
	"web-tree/utils"
)

func (m *Model) updateSuggestionList() {
	m.suggestionList = Fuzzy(m.searchInput.Value(), SuggestionInit())
	m.searchInput.SetSuggestions(m.suggestionList)
}

func (m *Model) updateUIComponents(msg tea.Msg) []tea.Cmd {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	m.searchInput, cmd = m.searchInput.Update(msg)
	cmds = append(cmds, cmd)

	m.paginator, cmd = m.paginator.Update(msg)
	cmds = append(cmds, cmd)

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)
	return cmds
}

func (m *Model) updateContent() {
	m.itemSelected.content = m.items[m.itemSelected.index]

	root := utils.RootTree
	selectedContent, _ := m.itemSelected.content.(string)
	t := root.FindSubTree(selectedContent)
	m.content = m.getTreeView(t)
	m.viewport.SetContent(m.content)
}

func (m *Model) AfterModeChange() {

}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.debug = "Width: " + strconv.Itoa(msg.Width) + "  " + "Height: " + strconv.Itoa(msg.Height)
		searchBoxHeight := lipgloss.Height(m.searchView())
		// suggestionListHeight := lipgloss.Height(m.suggestionListView())
		treeTabHeight := lipgloss.Height(m.treeTabView())
		helpHeight := lipgloss.Height(m.helpView())
		footerHeight := lipgloss.Height(m.footerView())

		verticalMarginHeight := searchBoxHeight +
			helpHeight + treeTabHeight + footerHeight

		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height/2)
			m.viewport.HighPerformanceRendering = false
			m.viewport.KeyMap = viewport.KeyMap{}
			m.viewport.SetContent(m.content)
			m.ready = true
		} else {
			m.viewport.Width = msg.Width / 2
			m.viewport.Height = msg.Height - verticalMarginHeight
			count++
		}

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.UP):
			switch m.mode {
			case search:
				if msg.String() == m.keymap.UP.Keys()[0] {
					m.sugSelected.index--
					if m.sugSelected.index < 0 {
						m.sugSelected.index = len(m.searchInput.AvailableSuggestions()) - 1
					}
				}
			}
		case key.Matches(msg, m.keymap.DOWN):
			switch m.mode {
			case search:
				if msg.String() == m.keymap.DOWN.Keys()[0] {
					m.sugSelected.index++
					if m.sugSelected.index >= len(m.searchInput.AvailableSuggestions()) {
						m.sugSelected.index = 0
					}
				}
			}
		case key.Matches(msg, m.keymap.LEFT):
			switch m.mode {
			case search:
				if msg.String() == m.keymap.LEFT.Keys()[0] {
					posi := m.searchInput.Position()
					m.searchInput.SetCursor(posi - 1)
				}
			case display:
				start, _ := m.paginator.GetSliceBounds(len(m.items))
				if msg.String() == m.keymap.LEFT.Keys()[0] {
					m.paginator.PrevPage()
				}
				if msg.String() == m.keymap.LEFT.Keys()[1] {
					m.itemSelected.index--
					if m.itemSelected.index == start-1 {
						m.paginator.PrevPage()
					}
					if m.itemSelected.index < 0 {
						m.itemSelected.index = 0
					}
				}
			}
		case key.Matches(msg, m.keymap.RIGHT):
			switch m.mode {
			case search:
				if msg.String() == m.keymap.RIGHT.Keys()[0] {
					posi := m.searchInput.Position()
					m.searchInput.SetCursor(posi + 1)
				}
			case display:
				_, end := m.paginator.GetSliceBounds(len(m.items))
				if msg.String() == m.keymap.RIGHT.Keys()[0] {
					m.paginator.NextPage()
				}
				if msg.String() == m.keymap.RIGHT.Keys()[1] {
					m.itemSelected.index++
					if m.itemSelected.index == end {
						m.paginator.NextPage()
					}
					if m.itemSelected.index > len(m.items)-1 {
						m.itemSelected.index = len(m.items) - 1
					}
				}
			}
		case key.Matches(msg, m.keymap.DELETE):
			switch m.mode {
			case search:
				if msg.String() == m.keymap.DELETE.Keys()[0] {
					v := m.searchInput.Value()
					if len(v) != 0 {
						m.searchInput.SetValue(v[:len(v)-1])
					}
				}
			}
		case key.Matches(msg, m.keymap.COMPLETE):
			switch m.mode {
			case search:
				if msg.String() == m.keymap.COMPLETE.Keys()[0] {
					if len(m.suggestionList) != 0 {
						if m.sugSelected.index == 0 && m.searchInput.Value() != m.searchInput.AvailableSuggestions()[0] {
							m.searchInput.SetValue(m.searchInput.AvailableSuggestions()[0])
						} else {
							m.sugSelected.index++
							if m.sugSelected.index > len(m.searchInput.AvailableSuggestions())-1 {
								m.sugSelected.index = 0
							}
							m.searchInput.SetValue(m.searchInput.AvailableSuggestions()[m.sugSelected.index])
						}
					}
				} else if msg.String() == m.keymap.COMPLETE.Keys()[1] {
					if len(m.suggestionList) != 0 {
						m.sugSelected.index--
						if m.sugSelected.index < 0 {
							m.sugSelected.index = len(m.searchInput.AvailableSuggestions()) - 1
						}
						m.searchInput.SetValue(m.searchInput.AvailableSuggestions()[m.sugSelected.index])
					}
				}
				m.searchInput.CursorEnd()
			}
		case key.Matches(msg, m.keymap.SELECT):
			switch m.mode {
			case search:
				if len(m.suggestionList) != 0 {
					m.searchInput.SetValue(m.searchInput.AvailableSuggestions()[m.sugSelected.index])
				}
				m.searchInput.CursorEnd()
				v := m.searchInput.Value()
				if v[:len(nodePrefix)] == nodePrefix {
					m.sugSelected.content = getNodeMsg(v)
				} else if v[:len(treePrefix)] == treePrefix {
					m.sugSelected.content = getTreeMsg(v)
				}
			}
		case key.Matches(msg, m.keymap.CLEAR):
			switch m.mode {
			case search:
				m.searchInput.SetValue("")
				m.searchInput.CursorStart()
			}
		case key.Matches(msg, m.keymap.OPEN):
		case key.Matches(msg, m.keymap.ADD):
		case key.Matches(msg, m.keymap.JUMP):
		case key.Matches(msg, m.keymap.TOGGLE):
			m.toggle = !m.toggle
		case key.Matches(msg, m.keymap.DETAIL):
		case key.Matches(msg, m.keymap.SINGLE):
		case key.Matches(msg, m.keymap.SWITCH):
			switch m.mode {
			case search, advancedSearch:
				if msg.String() == m.keymap.SWITCH.Keys()[0] {
					m.mode = display
					m.searchInput.Blur()
				}
			case display:
				if msg.String() == m.keymap.SWITCH.Keys()[1] {
					m.mode = search
					m.searchInput.Focus()
				}
				if msg.String() == m.keymap.SWITCH.Keys()[2] {
					m.mode = advancedSearch
				}
			}
		case key.Matches(msg, m.keymap.HELP):
		case key.Matches(msg, m.keymap.QUIT):
			return m, tea.Quit
		}
	}

	m.updateContent()
	m.updateSuggestionList()
	cmds = m.updateUIComponents(msg)

	return m, tea.Batch(cmds...)
}
