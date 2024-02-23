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
	adSearchInput  []textinput.Model
	paginator      paginator.Model
	viewport       viewport.Model
	root           utils.Tree
	suggestionList []string
	items          []string
	itemSelected   selected
	sugSelected    selected
	content        string
	keymap         keyMap
	ready          bool
	toggle         bool
	mode           Mode
	debug          string
}

type keyMap struct {
	UP       key.Binding
	DOWN     key.Binding
	LEFT     key.Binding
	RIGHT    key.Binding
	CLEAR    key.Binding
	OPEN     key.Binding
	ADD      key.Binding
	DELETE   key.Binding
	JUMP     key.Binding
	TOGGLE   key.Binding
	DETAIL   key.Binding
	SINGLE   key.Binding
	COMPLETE key.Binding
	SELECT   key.Binding
	SWITCH   key.Binding
	QUIT     key.Binding
	HELP     key.Binding
}

type selected struct {
	index   int
	content interface{}
}

type treeMsg struct {
	path string
}

type nodeMsg struct {
	path  string
	link  []string
	alias []string
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

	treePrefix                     = "*tree "
	nodePrefix                     = "*node "
	nameHint                       = " (name) "
	linkHint                       = " (link) "
	aliasHint                      = " (alias) "
	nodePathHint                   = " (path) "
	suggestionSelectedSuroundLeft  = "-}}"
	suggestionSelectedSuroundRight = "{{-"
	linkSep                        = " "
	aliasSep                       = " "

	nodeWidth  = 5
	nodeHeight = 5

	searchBoxWidth  = 5
	searchBoxHeight = 5

	paginatorHeight = 5

	noStyle       = lipgloss.NewStyle()
	activeStyle   = noStyle.Copy().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"})
	inactiveStyle = noStyle.Copy().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"})

	searchBoxStyle          = lipgloss.NewStyle()
	suggestionBoxStyle      = lipgloss.NewStyle()
	suggestionTreeStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#248"))
	suggestionNodeStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#2c8"))
	suggestionQuoteStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#4421f2"))
	suggestionMatchedStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#eeee92"))
	suggestionSelectedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#e78"))

	viewBoxStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder())

	treeTabBoxStyle         = inactiveStyle.Copy().Border(lipgloss.RoundedBorder())
	treeTabBoxSelectedStyle = activeStyle.Copy().Border(lipgloss.RoundedBorder())

	treeBoxStyle         = inactiveStyle.Copy().Border(lipgloss.RoundedBorder())
	treeBoxSelectedStyle = activeStyle.Copy().Border(lipgloss.RoundedBorder())

	nodeBoxStyle    = inactiveStyle.Copy().Border(lipgloss.RoundedBorder())
	nodeBoxSelected = activeStyle.Copy().Border(lipgloss.RoundedBorder())
	linkStyle       = noStyle.Copy().Foreground(lipgloss.Color("#9338f9"))
	aliasStyle      = noStyle.Copy().Foreground(lipgloss.Color("#933909"))

	browserStyle = lipgloss.NewStyle()
	shakeStyle   = lipgloss.NewStyle()
	branchStyle  = lipgloss.NewStyle()

	activeDotStyle   = activeStyle.Copy()
	inactiveDotStyle = inactiveStyle.Copy()

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
		CLEAR: key.NewBinding(
			key.WithKeys("ctrl+r"),
			key.WithHelp("ctrl+r", "Clear input text"),
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
		SELECT: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "Select the one under cursor"),
		),
		COMPLETE: key.NewBinding(
			key.WithKeys("tab", "shift+tab"),
			key.WithHelp("tab", "Autocomplete the input, index move forward"),
			key.WithHelp("shift+tab", "Autocomplete the input, index move backward"),
		),
		SWITCH: key.NewBinding(
			key.WithKeys("esc", "i", "u"),
			key.WithHelp("esc", "Display mode"),
			key.WithHelp("i", "Normal search mode"),
			key.WithHelp("u", "AdvancedSearch mode"),
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
