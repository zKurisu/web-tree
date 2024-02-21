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
	help             help.Model
	searchInput      textinput.Model
	adSearchInput    []textinput.Model
	paginator        paginator.Model
	viewport         viewport.Model
	root             utils.Tree
	suggestionList   []string
	items            []string
	sugSelectedIndex int
	content          string
	keymap           keyMap
	ready            bool
	mode             Mode
}

type keyMap struct {
	UP       key.Binding
	DOWN     key.Binding
	LEFT     key.Binding
	RIGHT    key.Binding
	OPEN     key.Binding
	ADD      key.Binding
	DELETE   key.Binding
	JUMP     key.Binding
	TOGGLE   key.Binding
	DETAIL   key.Binding
	SINGLE   key.Binding
	COMPLETE key.Binding
	SWITCH   key.Binding
	QUIT     key.Binding
	HELP     key.Binding
}

type Mode int

const (
	search Mode = iota
	advancedSearch
	display
)

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
	suggestionTreeStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#248"))
	suggestionNodeStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#2c8"))
	suggestionQuoteStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#4421f2"))
	suggestionMatchedStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#eeee92"))
	suggestionSelectedStyle = lipgloss.NewStyle().Background(lipgloss.Color("#e78"))

	viewBoxStyle = lipgloss.NewStyle()

	browserStyle     = lipgloss.NewStyle()
	treeStyle        = lipgloss.NewStyle()
	nodeStyle        = lipgloss.NewStyle()
	shakeStyle       = lipgloss.NewStyle()
	branchStyle      = lipgloss.NewStyle()
	activeDotStyle   = lipgloss.NewStyle()
	inactiveDotStyle = lipgloss.NewStyle()

	keymap = keyMap{
		UP: key.NewBinding(
			key.WithKeys("ctrl+k", "k"),
			key.WithHelp("ctrl+k", "Move up"),
			key.WithHelp("k", "Move up"),
		),
		DOWN: key.NewBinding(
			key.WithKeys("ctrl+j", "j"),
			key.WithHelp("ctrl+j", "Move down"),
			key.WithHelp("j", "Move down"),
		),
		LEFT: key.NewBinding(
			key.WithKeys("ctrl+h", "h"),
			key.WithHelp("ctrl+h", "Move left"),
			key.WithHelp("h", "Move left"),
		),
		RIGHT: key.NewBinding(
			key.WithKeys("ctrl+l", "l"),
			key.WithHelp("ctrl+l", "Move right"),
			key.WithHelp("l", "Move right"),
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
			key.WithKeys("backspace", "ctrl+d"),
			key.WithHelp("backspace", "Delete a character"),
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
		COMPLETE: key.NewBinding(
			key.WithKeys("tab", "shift+tab"),
			key.WithHelp("tab", "Autocomplete the input, index move forward"),
			key.WithHelp("shift+tab", "Autocomplete the input, index move backward"),
		),
		SWITCH: key.NewBinding(
			key.WithKeys("esc", "i", "ctrl+i"),
			key.WithHelp("esc", "Display mode"),
			key.WithHelp("i", "Normal search mode"),
			key.WithHelp("esc", "AdvancedSearch mode"),
		),
		QUIT: key.NewBinding(
			key.WithKeys("ctrl+c"),
			key.WithHelp("ctrl+c", "Quit"),
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
