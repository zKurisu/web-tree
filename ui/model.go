package ui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	// tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"web-tree/utils"
)

type Model struct {
	help             help.Model
	browseInput      textinput.Model
	searchInput      textinput.Model
	adSearchInput    []textinput.Model
	addInput         []textinput.Model
	paginator        paginator.Model
	viewport         viewport.Model
	textarea         textarea.Model
	root             utils.Tree
	suggestionList   []string
	adsuggestionList [][]string
	tabs             []string

	tabSelected     selected
	sugSelected     selected
	adInpSelected   selected
	addInpSelected  selected
	subSelected     point
	preSelectedTree []point

	winMsgs winMsg
	subMsgs subMsg
	content string
	keymap  keyMap

	ready      bool
	toggle     bool
	detail     bool
	helpToggle bool
	delete     bool
	copy       bool
	open       bool
	paste      bool
	curTree    *utils.Tree

	mode     Mode
	lastMode Mode

	confirm      Confirm
	curLineWidth int

	browser string
	debug   string
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
	EDIT     key.Binding
	JUMP     key.Binding
	TOGGLE   key.Binding
	DETAIL   key.Binding
	SINGLE   key.Binding
	COPY     key.Binding
	COMPLETE key.Binding
	SELECT   key.Binding
	SWITCH   key.Binding
	SAVE     key.Binding
	QUIT     key.Binding
	HELP     key.Binding
}

type Confirm struct {
	ans  textinput.Model
	hint string
}

type selected struct {
	index   int
	content interface{}
}

type point struct {
	x, y    int
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

type subMsg struct {
	ylen            []int
	searchedContent interface{}
}

type winMsg struct {
	Width  int
	Height int
}

type Mode int

const (
	search Mode = iota
	advancedSearch
	display
	add
	edit
	confirm
	browser
)

var (
	treeWidth  = 5
	treeHeight = 5
	treeGap    = 5

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

	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder   = tabBorderWithBottom("┘", " ", "└")
	highlightColor    = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}

	searchBoxStyle = lipgloss.NewStyle()
	// suggestionBoxStyle      = lipgloss.NewStyle()
	suggestionBoxStyle = lipgloss.NewStyle().
				BorderStyle(lipgloss.RoundedBorder()).
				BorderTop(true).
				BorderBottom(true).
				BorderLeft(true).
				BorderRight(true)
	suggestionTreeStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#248"))
	suggestionNodeStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#2c8"))
	suggestionQuoteStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#4421f2"))
	suggestionMatchedStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#eeee92"))
	suggestionSelectedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#e78"))
	adSearchSubmitStyle     = lipgloss.NewStyle().Background(lipgloss.Color("#e78"))

	// viewBoxStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder())
	viewBoxStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderTop(true).
			BorderBottom(true).
			BorderLeft(true).
			BorderRight(true)

	// treeTabBoxStyle         = inactiveStyle.Copy().Border(lipgloss.RoundedBorder())
	// treeTabBoxSelectedStyle = activeStyle.Copy().Border(lipgloss.RoundedBorder())
	treeTabBoxStyle         = lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(highlightColor).Padding(0, 1)
	treeTabBoxSelectedStyle = treeTabBoxStyle.Copy().Border(activeTabBorder, true).Background(lipgloss.Color("#248"))

	treeBoxStyle         = inactiveStyle.Copy().Border(lipgloss.RoundedBorder())
	treeBoxSelectedStyle = activeStyle.Copy().Border(lipgloss.RoundedBorder()).Background(lipgloss.Color("#476"))

	BottonBoxStyle         = inactiveStyle.Copy().Border(lipgloss.RoundedBorder())
	BottonBoxSelectedStyle = activeStyle.Copy().Border(lipgloss.RoundedBorder()).Background(lipgloss.Color("#476"))

	nodeBoxStyle    = inactiveStyle.Copy().Border(lipgloss.RoundedBorder()).Foreground(lipgloss.Color("#a44"))
	nodeBoxSelected = activeStyle.Copy().Border(lipgloss.RoundedBorder()).Background(lipgloss.Color("#476"))
	linkStyle       = noStyle.Copy().Foreground(lipgloss.Color("#9338f9"))
	aliasStyle      = noStyle.Copy().Foreground(lipgloss.Color("#933909"))

	browserStyle = lipgloss.NewStyle()
	shakeStyle   = lipgloss.NewStyle()
	branchStyle  = lipgloss.NewStyle()

	helpStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder())

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
			key.WithKeys("ctrl+x"),
			key.WithHelp("ctrl+x", "Open the link"),
		),
		DELETE: key.NewBinding(
			key.WithKeys("backspace"),
			key.WithHelp("backspace", "Delete a character"),
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
			key.WithKeys("ctrl+["),
			key.WithHelp("ctrl+[", "Show all message of a node"),
		),
		SINGLE: key.NewBinding(
			key.WithKeys("ctrl+]"),
			key.WithHelp("ctrl+]", "Show single tree horizontally"),
		),
		SELECT: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "Select the one under cursor"),
		),
		SAVE: key.NewBinding(
			key.WithKeys("ctrl+s"),
			key.WithHelp("ctrl+s", "Save change to tree or node in TextArea"),
		),
		COMPLETE: key.NewBinding(
			key.WithKeys("tab", "shift+tab"),
			key.WithHelp("tab", "Autocomplete the input, index move forward"),
			key.WithHelp("shift+tab", "Autocomplete the input, index move backward"),
		),
		COPY: key.NewBinding(
			key.WithKeys("ctrl+y"),
			key.WithHelp("ctrl+y", "Yank in display mode"),
		),
		SWITCH: key.NewBinding(
			key.WithKeys("esc", "ctrl+n", "ctrl+u", "ctrl+a", "ctrl+o", "ctrl+e", "ctrl+d"),
			key.WithHelp("esc", "Display mode"),
			key.WithHelp("ctrl+n", "Normal search mode"),
			key.WithHelp("ctrl+u", "AdvancedSearch mode"),
			key.WithHelp("ctrl+a", "Add mode"),
			key.WithHelp("ctrl+o", "Browser mode"),
			key.WithHelp("ctrl+e", "Edit mode"),
			key.WithHelp("ctrl+d", "Set delete flag"),
			// key.WithHelp(":", "Command mode"),
		),
		QUIT: key.NewBinding(
			key.WithKeys("ctrl+c"),
			key.WithHelp("ctrl+c", "Quit"),
		),
		HELP: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "Toggle short or long help"),
		),
	}
)

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.HELP, k.QUIT, k.SWITCH,
	}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.UP, k.DOWN, k.LEFT, k.RIGHT},
		{k.HELP, k.QUIT, k.SWITCH},
	}
}

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}
