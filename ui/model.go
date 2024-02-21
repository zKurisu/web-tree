package ui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	// tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"web-tree/utils"
)

type Model struct {
	help           help.Model
	searchInput    textinput.Model
	paginator      paginator.Model
	viewport       viewport.Model
	suggestionList []string
	items          []string
	content        string
	ready          bool
	keymap         keyMap
	root           utils.Tree
}

type keyMap struct {
	UP     key.Binding
	DOWN   key.Binding
	LEFT   key.Binding
	RIGHT  key.Binding
	OPEN   key.Binding
	ADD    key.Binding
	DELETE key.Binding
	JUMP   key.Binding
	TOGGLE key.Binding
	DETAIL key.Binding
	SINGLE key.Binding
	SWITCH key.Binding
	HELP   key.Binding
}

var (
	treePerPage = 5
	treeWidth   = 5
	treeHeight  = 5
	treeGap     = 5

	nodeWidth  = 5
	nodeHeight = 5

	searchBoxWidth  = 5
	searchBoxHeight = 5

	noStyle                 = lipgloss.NewStyle()
	searchBoxStyle          = lipgloss.NewStyle()
	suggestionBoxStyle      = lipgloss.NewStyle()
	suggestionTreeStyle     = lipgloss.NewStyle()
	suggestionNodeStyle     = lipgloss.NewStyle()
	suggestionQuoteStyle    = lipgloss.NewStyle()
	suggestionMatchedStyle  = lipgloss.NewStyle()
	suggestionSelectedStyle = lipgloss.NewStyle()

	browserStyle     = lipgloss.NewStyle()
	treeStyle        = lipgloss.NewStyle()
	nodeStyle        = lipgloss.NewStyle()
	shakeStyle       = lipgloss.NewStyle()
	branchStyle      = lipgloss.NewStyle()
	activeDotStyle   = lipgloss.NewStyle()
	inactiveDotStyle = lipgloss.NewStyle()

	keymap = keyMap{
		UP: key.NewBinding(
			key.WithKeys("ctrl+k"),
			key.WithHelp("ctrl+k", "Move up"),
		),
		DOWN: key.NewBinding(
			key.WithKeys("ctrl+j"),
			key.WithHelp("ctrl+j", "Move down"),
		),
		LEFT: key.NewBinding(
			key.WithKeys("ctrl+h"),
			key.WithHelp("ctrl+h", "Move left"),
		),
		RIGHT: key.NewBinding(
			key.WithKeys("ctrl+l"),
			key.WithHelp("ctrl+l", "Move right"),
		),
		OPEN: key.NewBinding(
			key.WithKeys("ctrl+o"),
			key.WithHelp("ctrl+o", "Open the link"),
		),
		ADD: key.NewBinding(
			key.WithKeys("ctrl+a"),
			key.WithHelp("ctrl+a", "Add a new tree/node"),
		),
		DELETE: key.NewBinding(
			key.WithKeys("ctrl+d"),
			key.WithHelp("ctrl+d", "Delete a tree/node"),
		),
		JUMP: key.NewBinding(
			key.WithKeys("ctrl+m"),
			key.WithHelp("ctrl+m", "Jump"),
		),
		TOGGLE: key.NewBinding(
			key.WithKeys("ctrl+t"),
			key.WithHelp("ctrl+t", "Toggle from link and alias"),
		),
		DETAIL: key.NewBinding(
			key.WithKeys("ctrl+s"),
			key.WithHelp("ctrl+s", "Show all message of a node"),
		),
		SINGLE: key.NewBinding(
			key.WithKeys("ctrl+x"),
			key.WithHelp("ctrl+x", "Show single tree horizontally"),
		),
		SWITCH: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "Switch between input and display region"),
		),
		HELP: key.NewBinding(
			key.WithKeys("ctrl+h"),
			key.WithHelp("ctrl+h", "Toggle short or long help"),
		),
	}
)

func (k keyMap) ShortHelp() []keyMap {
	return nil
}

func (k keyMap) FullHelp() [][]keyMap {
	return nil
}
