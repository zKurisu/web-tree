package ui

import (
	"github.com/charmbracelet/bubbles/key"
	// "github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m *Model) updateSuggestionList() {
	m.suggestionList = Fuzzy(m.searchInput.Value(), SuggestionInit())
	m.searchInput.SetSuggestions(m.suggestionList)
}

func (m *Model) AfterModeChange() {

}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	m.updateSuggestionList()

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		searchBoxHeight := lipgloss.Height(m.searchView())
		helpHeight := lipgloss.Height(m.helpView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := searchBoxHeight + helpHeight + footerHeight

		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = searchBoxHeight
			m.viewport.HighPerformanceRendering = false
			m.viewport.KeyMap = viewport.KeyMap{}
			m.viewport.SetContent(m.content)
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
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
				if msg.String() == m.keymap.LEFT.Keys()[0] {
					m.paginator.PrevPage()
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
				if msg.String() == m.keymap.RIGHT.Keys()[0] {
					m.paginator.NextPage()
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
		case key.Matches(msg, m.keymap.DETAIL):
		case key.Matches(msg, m.keymap.SINGLE):
		case key.Matches(msg, m.keymap.SWITCH):
			m.debug = msg.String()
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
	m.searchInput, cmd = m.searchInput.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
