package ui

import (
	"github.com/charmbracelet/bubbles/key"
	// "github.com/charmbracelet/bubbles/paginator"
	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"strconv"
	"strings"
	"web-tree/utils"
)

func (m *Model) updateSuggestionList() {
	switch m.mode {
	case browser:
		m.suggestionList = Fuzzy(m.browseInput.Value(), browseSuggestionInit())
		m.browseInput.SetSuggestions(m.suggestionList)
	case search:
		m.suggestionList = Fuzzy(m.searchInput.Value(), searchSuggestionInit())
		m.searchInput.SetSuggestions(m.suggestionList)
	case advancedSearch:
		for i := range m.adSearchInput {
			m.adsuggestionList[i] = Fuzzy(m.adSearchInput[i].Value(), adSearchSuggestionInit()[i])
			m.adSearchInput[i].SetSuggestions(m.adsuggestionList[i])
		}
	}
}

func (m *Model) updateUIComponents(msg tea.Msg) []tea.Cmd {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	m.browseInput, cmd = m.browseInput.Update(msg)
	cmds = append(cmds, cmd)

	m.searchInput, cmd = m.searchInput.Update(msg)
	cmds = append(cmds, cmd)

	for i := range m.adSearchInput {
		m.adSearchInput[i], cmd = m.adSearchInput[i].Update(msg)
		cmds = append(cmds, cmd)
	}

	for i := range m.addInput {
		m.addInput[i], cmd = m.addInput[i].Update(msg)
		cmds = append(cmds, cmd)
	}

	m.paginator, cmd = m.paginator.Update(msg)
	cmds = append(cmds, cmd)

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)

	m.confirm.ans, cmd = m.confirm.ans.Update(msg)
	cmds = append(cmds, cmd)
	return cmds
}

// call it after m.adInpSelected.index has changed
func (m *Model) updateAdSearch() {
	for i := range m.adSearchInput {
		if i == m.adInpSelected.index && m.mode == advancedSearch {
			m.adSearchInput[i].Focus()
			m.adSearchInput[i].TextStyle = activeStyle
			m.adSearchInput[i].PromptStyle = activeStyle
		} else {
			m.adSearchInput[i].Blur()
			m.adSearchInput[i].TextStyle = inactiveStyle
			m.adSearchInput[i].PromptStyle = inactiveStyle
		}
	}
}

func (m *Model) updateAddInput() {
	for i := range m.addInput {
		if i == m.addInpSelected.index && m.mode == add {
			m.addInput[i].Focus()
			m.addInput[i].TextStyle = activeStyle
			m.addInput[i].PromptStyle = activeStyle
		} else {
			m.addInput[i].Blur()
			m.addInput[i].TextStyle = inactiveStyle
			m.addInput[i].PromptStyle = inactiveStyle
		}
	}
}

func (m *Model) updateBrowseInput() {
	if m.mode == browser {
		m.browseInput.Focus()
		m.browseInput.TextStyle = activeStyle
		m.browseInput.PromptStyle = activeStyle
	} else {
		m.browseInput.Blur()
		m.browseInput.TextStyle = inactiveStyle
		m.browseInput.PromptStyle = inactiveStyle
	}
}

func (m *Model) updateConfirmInput() {
	if m.mode == confirm {
		m.confirm.ans.Focus()
		m.confirm.ans.TextStyle = activeStyle
		m.confirm.ans.PromptStyle = activeStyle
	} else {
		m.confirm.ans.Blur()
		m.confirm.ans.TextStyle = inactiveStyle
		m.confirm.ans.PromptStyle = inactiveStyle
	}
}

func (m *Model) updateContent() {
	m.tabSelected.content = m.tabs[m.tabSelected.index]

	root := utils.RootTree
	selectedContent, _ := m.tabSelected.content.(string)
	t := root.FindSubTree(selectedContent)
	m.subMsgs.ylen = []int{len(m.tabs)}
	m.content = m.getTreeView(t, 1)
	if m.subSelected.y == 0 {
		m.subSelected.content = t
	}
	m.viewport.SetContent(m.content)
}

func (m *Model) updateTextarea() {
	// Set info of current selected node
}

func (m *Model) afterModeChange() {
	switch m.mode {
	case browser:
		sequence(m.blurSearch, m.blurAdsearch, m.blurAddInput, m.blurTextarea)

		m.browseInput.Focus()
		m.browseInput.ShowSuggestions = true
	case search:
		sequence(m.blurAdsearch, m.blurAddInput, m.blurTextarea, m.blurBrowse)

		m.searchInput.Focus()
		m.searchInput.ShowSuggestions = true
	case advancedSearch:
		sequence(m.blurSearch, m.blurAddInput, m.blurTextarea, m.blurBrowse)

		m.adSearchInput[0].Focus()
		for i := range m.adSearchInput {
			m.adSearchInput[i].ShowSuggestions = true
		}
	case add:
		sequence(m.blurSearch, m.blurAdsearch, m.blurTextarea, m.blurBrowse)

		m.addInput[0].Focus()
	case display:
		sequence(m.blurSearch, m.blurAdsearch, m.blurAddInput, m.blurTextarea, m.blurConfirm, m.blurBrowse)

		m.searchInput.ShowSuggestions = false
		m.searchInput.Blur()
		for i := range m.adSearchInput {
			m.adSearchInput[i].ShowSuggestions = false
		}
	case edit:
		var b strings.Builder
		switch content := m.subSelected.content.(type) {
		case *utils.Tree:
			b.WriteString(content.GetTreeName())
		case *utils.Node:
			b.WriteString("links: " + strings.Join(content.Link, " "))
			b.WriteString("\n")
			b.WriteString("alias: " + strings.Join(content.Alias, " "))
			b.WriteString("\n")
			b.WriteString("description: " + strings.Join(content.Desc, " "))
			b.WriteString("\n")
			b.WriteString("label: " + strings.Join(content.Label, " "))
		}

		m.textarea.Focus()
		m.textarea.SetValue(b.String())
	case confirm:
		targetType := ""
		targetHint := ""
		operation := ""
		ansHint := ""

		switch target := m.subSelected.content.(type) {
		case *utils.Node:
			targetType = "node"
			targetHint = target.GetNodeAlias()[0]
		case *utils.Tree:
			targetType = "tree"
			targetHint = target.GetTreeName()
		}
		if m.delete {
			operation = "detele"
			ansHint = " <yes|no> "
		} else if m.copy {
			operation = "copy"
			targetType = "Link"
			targetHint = "Enter order of link"
			ansHint = " <1|2|..> "
			// m.debug = "In copy?"
		} else if m.open {
			operation = "open"
			targetType = "Link"
			targetHint = "Enter order of link"
			ansHint = " <1|2|..> "
		}

		m.confirm.ans.SetValue("Confirm: Do you want to " + operation + " " +
			targetType + " [" + targetHint + "]?" + ansHint)
		m.confirm.hint = m.confirm.ans.Value()
		m.confirm.ans.CursorEnd()

		m.confirm.ans.Focus()
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// m.debug = "Width: " + strconv.Itoa(msg.Width) + "  " + "Height: " + strconv.Itoa(msg.Height)
		searchBoxHeight := lipgloss.Height(m.searchView())
		treeTabHeight := lipgloss.Height(m.treeTabView())
		helpHeight := lipgloss.Height(m.helpView())
		paginatorHeight := lipgloss.Height(m.paginatorView())
		debugHeight := lipgloss.Height(m.debugView())
		footerHeight := lipgloss.Height(m.footerView())

		verticalMarginHeight := searchBoxHeight +
			helpHeight + treeTabHeight + paginatorHeight +
			debugHeight + footerHeight
		m.winMsgs.Width = msg.Width
		m.winMsgs.Height = msg.Height - verticalMarginHeight

		if !m.ready {
			m.paginator = paginatorInit(msg.Width/8 - 2)
			m.viewport = viewport.New(msg.Width-2, msg.Height-verticalMarginHeight-2-2-2)
			m.viewport.HighPerformanceRendering = false
			m.viewport.KeyMap = viewport.KeyMap{}
			m.viewport.SetContent(m.content)
			m.browseInput.Width = msg.Width / 8
			// m.viewport.YPosition = searchBoxHeight + treeTabHeight + 1
			m.ready = true
		} else {
			m.paginator.PerPage = msg.Width/8 - 2
			m.viewport.Width = msg.Width - 2
			m.viewport.Height = msg.Height - verticalMarginHeight - 2 - 2 - 2
		}

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.UP):
			switch m.mode {
			case browser:
				if msg.String() == m.keymap.UP.Keys()[0] {
					m.sugSelected.index--
					if m.sugSelected.index < 0 {
						m.sugSelected.index = len(m.browseInput.AvailableSuggestions()) - 1
					}
				}
			case search:
				if msg.String() == m.keymap.UP.Keys()[0] {
					m.sugSelected.index--
					if m.sugSelected.index < 0 {
						m.sugSelected.index = len(m.searchInput.AvailableSuggestions()) - 1
					}
				}
			case advancedSearch:
				if msg.String() == m.keymap.UP.Keys()[0] {
					m.adInpSelected.index--
					if m.adInpSelected.index < 0 {
						m.adInpSelected.index = 0
					} else {
						m.sugSelected = selected{index: 0}
					}
				}
			case add:
				if msg.String() == m.keymap.UP.Keys()[0] {
					m.addInpSelected.index--
					if m.addInpSelected.index < 0 {
						m.addInpSelected.index = 0
					} else {
						m.sugSelected = selected{index: 0}
					}
				}
			case display:
				if msg.String() == m.keymap.UP.Keys()[1] {
					if len(m.preSelectedTree) > 0 {
						m.subSelected = m.preSelectedTree[len(m.preSelectedTree)-1]
						m.preSelectedTree = m.preSelectedTree[:len(m.preSelectedTree)-1]
						// if m.viewport.YOffset-m.subSelected.y > 0 {
						// 	m.viewport.YOffset -= 3
						// }
					}
				}
			}
		case key.Matches(msg, m.keymap.DOWN):
			switch m.mode {
			case browser:
				if msg.String() == m.keymap.DOWN.Keys()[0] {
					m.sugSelected.index++
					if m.sugSelected.index >= len(m.browseInput.AvailableSuggestions()) {
						m.sugSelected.index = 0
					}
				}
			case search:
				if msg.String() == m.keymap.DOWN.Keys()[0] {
					m.sugSelected.index++
					if m.sugSelected.index >= len(m.searchInput.AvailableSuggestions()) {
						m.sugSelected.index = 0
					}
				}
			case advancedSearch:
				if msg.String() == m.keymap.DOWN.Keys()[0] {
					m.adInpSelected.index++
					if m.adInpSelected.index >= len(m.adSearchInput) {
						// not -1 because a submit botton
						m.adInpSelected.index = len(m.adSearchInput)
					} else {
						m.sugSelected = selected{index: 0}
					}
				}
			case add:
				if msg.String() == m.keymap.DOWN.Keys()[0] {
					m.addInpSelected.index++
					if m.addInpSelected.index >= len(m.addInput) {
						// not -1 because a submit botton
						m.addInpSelected.index = len(m.addInput)
					} else {
						m.sugSelected = selected{index: 0}
					}
				}
			case display:
				if msg.String() == m.keymap.DOWN.Keys()[1] {
					if m.subSelected.y < len(m.subMsgs.ylen)-1 {
						m.preSelectedTree = append(m.preSelectedTree, m.subSelected)
						m.subSelected.x = 0
						m.subSelected.y++
						// if m.subSelected.y-m.viewport.YOffset > 1 {
						// 	m.viewport.YOffset += 3
						// }
					}
				}
			}
		case key.Matches(msg, m.keymap.LEFT):
			switch m.mode {
			case browser:
				if msg.String() == m.keymap.LEFT.Keys()[0] {
					posi := m.browseInput.Position()
					m.browseInput.SetCursor(posi - 1)
				}
			case search:
				if msg.String() == m.keymap.LEFT.Keys()[0] {
					posi := m.searchInput.Position()
					m.searchInput.SetCursor(posi - 1)
				}
			case advancedSearch:
				i := m.adInpSelected.index
				// Except when selected the botton
				if i == len(m.adSearchInput) {
					break
				}
				if msg.String() == m.keymap.LEFT.Keys()[0] {
					posi := m.adSearchInput[i].Position()
					m.adSearchInput[i].SetCursor(posi - 1)
				}
			case add:
				i := m.addInpSelected.index
				// Except when selected the botton
				if i == len(m.addInput) {
					break
				}
				if msg.String() == m.keymap.LEFT.Keys()[0] {
					posi := m.addInput[i].Position()
					m.addInput[i].SetCursor(posi - 1)
				}
			case display:
				start, _ := m.paginator.GetSliceBounds(len(m.tabs))
				if msg.String() == m.keymap.LEFT.Keys()[0] {
					if start > 0 {
						m.tabSelected.index = start - 1
						m.paginator.PrevPage()
					}
				}
				if msg.String() == m.keymap.LEFT.Keys()[1] {
					// tab selected part
					if m.subSelected.y == 0 {
						m.tabSelected.index--
						if m.tabSelected.index == start-1 {
							m.paginator.PrevPage()
						}
						if m.tabSelected.index < 0 {
							m.tabSelected.index = 0
						}
					} else {
						// point selected part
						m.subSelected.x--
						if m.subSelected.x < 0 {
							m.subSelected.x = 0
						}
					}

				}
			case confirm:
				if msg.String() == m.keymap.LEFT.Keys()[0] {
					posi := m.confirm.ans.Position()
					m.confirm.ans.SetCursor(posi - 1)
				}
			}
		case key.Matches(msg, m.keymap.RIGHT):
			switch m.mode {
			case browser:
				if msg.String() == m.keymap.RIGHT.Keys()[0] {
					posi := m.browseInput.Position()
					m.browseInput.SetCursor(posi + 1)
				}
			case search:
				if msg.String() == m.keymap.RIGHT.Keys()[0] {
					posi := m.searchInput.Position()
					m.searchInput.SetCursor(posi + 1)
				}
			case advancedSearch:
				i := m.adInpSelected.index
				// Except when selected the botton
				if i == len(m.adSearchInput) {
					break
				}
				if msg.String() == m.keymap.RIGHT.Keys()[0] {
					posi := m.adSearchInput[i].Position()
					m.adSearchInput[i].SetCursor(posi + 1)
				}
			case add:
				i := m.addInpSelected.index
				// Except when selected the botton
				if i == len(m.addInput) {
					break
				}
				if msg.String() == m.keymap.RIGHT.Keys()[0] {
					posi := m.addInput[i].Position()
					m.addInput[i].SetCursor(posi + 1)
				}
			case display:
				_, end := m.paginator.GetSliceBounds(len(m.tabs))
				if msg.String() == m.keymap.RIGHT.Keys()[0] {
					if end < len(m.tabs) {
						m.tabSelected.index = end
						m.paginator.NextPage()
					}
				}
				if msg.String() == m.keymap.RIGHT.Keys()[1] {
					if m.subSelected.y == 0 {
						// tab selected part
						m.tabSelected.index++
						if m.tabSelected.index == end {
							m.paginator.NextPage()
						}
						if m.tabSelected.index > len(m.tabs)-1 {
							m.tabSelected.index = len(m.tabs) - 1
						}
					} else {
						// point selected part
						m.subSelected.x++
						if m.subSelected.x > m.subMsgs.ylen[m.subSelected.y]-1 {
							m.subSelected.x = m.subMsgs.ylen[m.subSelected.y] - 1
						}
					}

				}
			case confirm:
				if msg.String() == m.keymap.RIGHT.Keys()[0] {
					posi := m.confirm.ans.Position()
					m.confirm.ans.SetCursor(posi + 1)
				}
			}
		case key.Matches(msg, m.keymap.DELETE):
			switch m.mode {
			case browser:
				if msg.String() == m.keymap.DELETE.Keys()[0] {
					v := m.browseInput.Value()
					p := m.browseInput.Position()
					if len(v) != 0 {
						if p > 0 {
							m.browseInput.SetValue(v[:p-1] + v[p:])
							m.browseInput.SetCursor(p - 1)
						}
					}
				}
			case search:
				if msg.String() == m.keymap.DELETE.Keys()[0] {
					v := m.searchInput.Value()
					p := m.searchInput.Position()
					if len(v) != 0 {
						if p > 0 {
							m.searchInput.SetValue(v[:p-1] + v[p:])
							m.searchInput.SetCursor(p - 1)
						}
					}
				}
			case advancedSearch:
				i := m.adInpSelected.index
				p := m.adSearchInput[i].Position()
				// Except when selected the botton
				if i == len(m.adSearchInput) {
					break
				}

				if msg.String() == m.keymap.DELETE.Keys()[0] {
					v := m.adSearchInput[i].Value()
					if len(v) != 0 {
						m.adSearchInput[i].SetValue(v[:len(v)-1])
						if p > 0 {
							m.adSearchInput[i].SetValue(v[:p-1] + v[p:])
							m.adSearchInput[i].SetCursor(p - 1)
						}
					}
				}
			case add:
				i := m.addInpSelected.index
				p := m.addInput[i].Position()
				// Except when selected the botton
				if i == len(m.addInput) {
					break
				}

				if msg.String() == m.keymap.DELETE.Keys()[0] {
					v := m.addInput[i].Value()
					if len(v) != 0 {
						m.addInput[i].SetValue(v[:len(v)-1])
						if p > 0 {
							m.addInput[i].SetValue(v[:p-1] + v[p:])
							m.addInput[i].SetCursor(p - 1)
						}
					}
				}
			case confirm:
				if msg.String() == m.keymap.DELETE.Keys()[0] {
					v := m.confirm.ans.Value()
					p := m.confirm.ans.Position()
					if len(v) != len(m.confirm.hint) {
						if p > len(m.confirm.hint) {
							m.confirm.ans.SetValue(v[:p-1] + v[p:])
							m.confirm.ans.SetCursor(p - 1)
						}
					}
				}
			}
		case key.Matches(msg, m.keymap.COPY):
			if m.mode == display {
				switch m.subSelected.content.(type) {
				case *utils.Node:
					m.lastMode = m.mode
					m.mode = confirm
					m.copy = true
					m.afterModeChange()
				}
			}
		case key.Matches(msg, m.keymap.COMPLETE):
			switch m.mode {
			case browser:
				if msg.String() == m.keymap.COMPLETE.Keys()[0] {
					if len(m.suggestionList) != 0 {
						if m.sugSelected.index == 0 && m.browseInput.Value() != m.browseInput.AvailableSuggestions()[0] {
							m.browseInput.SetValue(m.browseInput.AvailableSuggestions()[0])
						} else {
							m.sugSelected.index++
							if m.sugSelected.index > len(m.browseInput.AvailableSuggestions())-1 {
								m.sugSelected.index = 0
							}
							m.browseInput.SetValue(m.browseInput.AvailableSuggestions()[m.sugSelected.index])
						}
					}
				} else if msg.String() == m.keymap.COMPLETE.Keys()[1] {
					if len(m.suggestionList) != 0 {
						m.sugSelected.index--
						if m.sugSelected.index < 0 {
							m.sugSelected.index = len(m.browseInput.AvailableSuggestions()) - 1
						}
						m.browseInput.SetValue(m.browseInput.AvailableSuggestions()[m.sugSelected.index])
					}
				}
				m.browseInput.CursorEnd()
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
			case advancedSearch:
				i := m.adInpSelected.index
				// Except when selected the botton
				if i == len(m.adSearchInput) {
					break
				}
				if msg.String() == m.keymap.COMPLETE.Keys()[0] {
					if len(m.adsuggestionList[i]) != 0 {
						if m.sugSelected.index == 0 && m.adSearchInput[i].Value() != m.adSearchInput[i].AvailableSuggestions()[0] {
							m.adSearchInput[i].SetValue(m.adSearchInput[i].AvailableSuggestions()[0])
						} else {
							m.sugSelected.index++
							if m.sugSelected.index > len(m.adSearchInput[i].AvailableSuggestions())-1 {
								m.sugSelected.index = 0
							}
							m.adSearchInput[i].SetValue(m.adSearchInput[i].AvailableSuggestions()[m.sugSelected.index])
						}
					}
				} else if msg.String() == m.keymap.COMPLETE.Keys()[1] {
					if len(m.adsuggestionList[i]) != 0 {
						m.sugSelected.index--
						if m.sugSelected.index < 0 {
							m.sugSelected.index = len(m.adSearchInput[i].AvailableSuggestions()) - 1
						}
						m.adSearchInput[i].SetValue(m.adSearchInput[i].AvailableSuggestions()[m.sugSelected.index])
					}
				}
				m.adSearchInput[i].CursorEnd()
			case add:
				// current do nothing for this
			}
		case key.Matches(msg, m.keymap.SELECT):
			switch m.mode {
			case browser:
				// m.browseInput.SetValue(m.browseInput.AvailableSuggestions()[m.sugSelected.index])
				// m.browseInput.CursorEnd()
				m.browser = m.browseInput.Value()
				m.mode = m.lastMode
				m.lastMode = browser
				m.afterModeChange()
			case search:
				if len(m.suggestionList) != 0 {
					m.searchInput.SetValue(m.searchInput.AvailableSuggestions()[m.sugSelected.index])
					m.searchInput.CursorEnd()
					v := m.searchInput.Value()
					m.searchInput.SetValue("")

					tabTarget := ""
					if v[:len(nodePrefix)] == nodePrefix {
						m.sugSelected.content = getNodeMsg(v)
					} else if v[:len(treePrefix)] == treePrefix {
						m.sugSelected.content = getTreeMsg(v)
					}

					switch content := m.sugSelected.content.(type) {
					case nodeMsg:
						hint := append(content.link, content.alias...)
						m.subMsgs.searchedContent = utils.RootTree.DeepFindSubTree(content.path).FindNode(hint)
						tabTarget = strings.Split(content.path, "/")[0]
					case treeMsg:
						tabTarget = strings.Split(content.path, "/")[0]
					}

					for i, tab := range m.tabs {
						if tab == tabTarget {
							m.tabSelected.index = i
						}
					}
					// m.mode = display
				}
			case advancedSearch:
				i := m.adInpSelected.index
				// When selected the botton
				// set tabSelected.index and subSelected.content
				if i == len(m.adSearchInput) {
					var treeName string
					var hints []string
					for i := range m.adSearchInput {
						v := m.adSearchInput[i].Value()

						if v != "" {
							switch i {
							case 0:
								var selectedTab string
								treeName = v
								selectedTab = utils.SplitTreeLevel(v)[0]
								for i, tab := range m.tabs {
									if tab == selectedTab {
										m.tabSelected.index = i
									}
								}
							case 1, 2:
								hints = append(hints, strings.Split(v, " ")...)
							}
						}
					}
					hints = utils.RemoveEmp(hints)
					t := utils.RootTree.DeepFindSubTree(treeName)
					if t != nil {
						m.subMsgs.searchedContent = t.FindNode(hints)
					}
					m.lastMode = m.mode
					m.mode = display
					m.afterModeChange()
					break
				}
				if len(m.adsuggestionList[i]) != 0 {
					m.adSearchInput[i].SetValue(m.adSearchInput[i].AvailableSuggestions()[m.sugSelected.index])
					m.adSearchInput[i].CursorEnd()
					m.adInpSelected.index++
				}
			case add:
				i := m.addInpSelected.index
				// When select the botton
				if i == len(m.addInput) {
					root := utils.RootTree
					treeName := m.addInput[0].Value()

					links := strings.Split(m.addInput[1].Value(), " ")
					alias := strings.Split(m.addInput[2].Value(), " ")
					desc := strings.Split(m.addInput[3].Value(), ",")
					label := strings.Split(m.addInput[4].Value(), " ")
					icon := m.addInput[5].Value()

					root.DeepAddNewSubTree(treeName)
					t := root.DeepFindSubTree(treeName)
					n, _ := utils.NewNode(links, alias, desc, icon, label, "None")
					if len(links) != 0 && links[0] != "" {
						t.AppendNode(n)
					}
					utils.WriteAll()

					if len(utils.SplitTreeLevel(treeName)) == 1 {
						utils.RootTree = utils.Tree{
							Name:     "root",
							SubTrees: utils.GetAllRootSubTree(),
							Nodes:    []*utils.Node{},
						}
						m.tabs = utils.RootTree.GetAllSubtreeName()
						m.paginator.SetTotalPages(len(utils.RootTree.GetAllSubtreeName()))

						oriPageNumber := getPageNumber(m.tabSelected.index)
						m.tabSelected.index = getIndex(m.tabs, treeName)
						newPageNumber := getPageNumber(m.tabSelected.index)

						for i := 0; i < newPageNumber-oriPageNumber; i++ {
							m.paginator.NextPage()
						}

					}
				}
				m.addInpSelected.index = 0
				m.lastMode = m.mode
				m.mode = display
			case confirm:
				// When stroke "enter" in command mode...
				answer := getLastWord(m.confirm.ans.Value())
				if m.delete && answer == "yes" {
					// Here add delete operation
					root := utils.RootTree

					switch content := m.subSelected.content.(type) {
					case *utils.Tree:
						root.DeepDelSubTree(content.GetTreeName())
						if utils.IsRootSubTreeExist(content.GetTreeName()) {
							utils.RootTree = utils.Tree{
								Name:     "root",
								SubTrees: utils.GetAllRootSubTree(),
								Nodes:    []*utils.Node{},
							}

							start, _ := m.paginator.GetSliceBounds(len(m.tabs))
							if m.tabSelected.index < m.subMsgs.ylen[m.subSelected.y]-1 && m.tabSelected.index > 0 {
								m.tabSelected.content = m.tabs[m.tabSelected.index]
								m.tabs = utils.RootTree.GetAllSubtreeName()
							}
							if m.tabSelected.index == m.subMsgs.ylen[m.subSelected.y]-1 && m.tabSelected.index > 0 {
								m.tabSelected.index--
								m.tabSelected.content = m.tabs[m.tabSelected.index]
								m.tabs = utils.RootTree.GetAllSubtreeName()
								if m.tabSelected.index == start-1 {
									m.paginator.PrevPage()
									m.paginator.SetTotalPages(len(utils.RootTree.GetAllSubtreeName()))
								}
								if m.tabSelected.index < 0 {
									m.tabSelected.index = 0
								}
							}

							// no elements on right and no elements on left
							// if m.tabSelected.index == m.subMsgs.ylen[m.subSelected.y]-1 && m.tabSelected.index == 0 {
							// 	m.lastMode = m.mode
							// 	m.mode = search
							// }
						}
					case *utils.Node:
						hints := []string{}
						hints = append(hints, content.GetNodeLinks()...)
						hints = append(hints, content.GetNodeAlias()...)
						m.curTree.DelNode(hints)
						// log.Println(content)
					}
					utils.WriteAll()

					// no elements on right
					//if m.subSelected.x < m.subMsgs.ylen[m.subSelected.y]-1 {
					//	// Nothing happend
					//}

					m.delete = false
					// no elements on right but has elements on left
					if m.subSelected.x == m.subMsgs.ylen[m.subSelected.y]-1 && m.subSelected.x > 0 {
						m.subSelected.x--
					}

					// no elements on right and no elements on left
					if m.subSelected.x == m.subMsgs.ylen[m.subSelected.y]-1 && m.subSelected.x == 0 {
						m.subSelected = m.preSelectedTree[len(m.preSelectedTree)-1]
						m.preSelectedTree = m.preSelectedTree[:len(m.preSelectedTree)-1]
					}
				} else if m.copy {
					switch content := m.subSelected.content.(type) {
					case *utils.Node:
						index, _ := strconv.Atoi(answer)
						if index <= len(content.GetNodeLinks()) {
							clipboard.WriteAll(content.GetNodeLinks()[index-1])
							m.copy = false
						}
					}
				} else if m.open {
					switch content := m.subSelected.content.(type) {
					case *utils.Node:
						index, _ := strconv.Atoi(answer)
						m.debug = "Open link"
						if index <= len(content.GetNodeLinks()) {
							openLink(m.browser, content.GetNodeLinks()[index-1])
							m.open = false
						}
					}
				}

				m.lastMode = m.mode
				// m.debug = msg.String()
				m.mode = display

				if m.lastMode != m.mode {
					m.afterModeChange()
				}
			}
		case key.Matches(msg, m.keymap.CLEAR):
			switch m.mode {
			case search:
				m.searchInput.SetValue("")
				m.searchInput.CursorStart()
			case advancedSearch:
				i := m.adInpSelected.index
				m.adSearchInput[i].SetValue("")
				m.adSearchInput[i].CursorStart()
			}
		case key.Matches(msg, m.keymap.OPEN):
			if m.mode == display {
				m.lastMode = m.mode
				m.mode = confirm
				m.open = true
				m.afterModeChange()
			}
		case key.Matches(msg, m.keymap.JUMP):
		case key.Matches(msg, m.keymap.TOGGLE):
			m.toggle = !m.toggle
		case key.Matches(msg, m.keymap.DETAIL):
		case key.Matches(msg, m.keymap.SINGLE):
		case key.Matches(msg, m.keymap.SWITCH):
			m.lastMode = m.mode
			// m.debug = msg.String()
			switch msg.String() {
			case m.keymap.SWITCH.Keys()[0]:
				m.mode = display
			case m.keymap.SWITCH.Keys()[1]:
				m.mode = search
			case m.keymap.SWITCH.Keys()[2]:
				m.mode = advancedSearch
			case m.keymap.SWITCH.Keys()[3]:
				m.mode = add
			case m.keymap.SWITCH.Keys()[4]:
				m.mode = browser
			case m.keymap.SWITCH.Keys()[5]:
				if m.lastMode == display {
					m.mode = edit
				}
			case m.keymap.SWITCH.Keys()[6]:
				if m.lastMode == display {
					m.mode = confirm
					m.delete = true
				}
				// case m.keymap.SWITCH.Keys()[6]:
				// 	if m.lastMode == display {
				// 		m.mode = command
				// 	}
			}

			if m.lastMode != m.mode {
				m.afterModeChange()
			}
		case key.Matches(msg, m.keymap.SAVE):
			switch m.mode {
			case edit:
				m.lastMode = m.mode
				m.mode = display
				m.textarea.Blur()

				content := m.textarea.Value()
				switch target := m.subSelected.content.(type) {
				case *utils.Tree:
					levels := strings.Split(content, "/")
					nextFatherName := strings.Join(levels[0:len(levels)-1], "/")
					if nextFatherName == "" {
						nextFatherName = "root"
					}
					oriFatherName := target.GetFatherName()
					// m.debug = "next: " + nextFatherName + " ori: " + oriFatherName

					if nextFatherName != oriFatherName {
						nextFatherTree := utils.RootTree.DeepFindSubTree(nextFatherName)
						oriFatherTree := utils.RootTree.DeepFindSubTree(target.GetFatherName())

						if nextFatherTree == nil {
							utils.RootTree.DeepAddNewSubTree(nextFatherName)
							nextFatherTree = utils.RootTree.DeepFindSubTree(nextFatherName)
						}
						nextFatherTree.AppendSubTree(target)
						oriFatherTree.DelSubTree(target.GetTreeName())
					}
					target.Name = content
				case *utils.Node:
					lines := strings.Split(content, "\n")
					target.Link = strings.Split(lines[0], " ")[1:]
					target.Alias = strings.Split(lines[1], " ")[1:]
					target.Desc = strings.Split(lines[2], " ")[1:]
					target.Label = strings.Split(lines[3], " ")[1:]
				}

				utils.WriteAll()
			}

		case key.Matches(msg, m.keymap.HELP):
			m.helpToggle = !m.helpToggle
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keymap.QUIT):
			return m, tea.Quit
		}
	}

	m.updateAdSearch()
	m.updateBrowseInput()
	m.updateAddInput()
	m.updateContent()
	m.updateSuggestionList()
	m.updateTextarea()
	// m.updateConfirmInput()
	cmds = append(cmds, m.updateUIComponents(msg)...)

	return m, tea.Batch(cmds...)
}
