package ui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
	"web-tree/utils"
)

// All tree name (expect root) and all nodes
func searchSuggestionInit() []string {
	treePathes := utils.RootTree.DeepGetAllSubtreeName()
	nodes, nodePathes := utils.RootTree.DeepGetAllNodeWithPath()
	treeSuggestionList := []string{}
	nodeSuggestionList := []string{}

	for _, treePath := range treePathes {
		s := treePrefix + treePath + nameHint
		treeSuggestionList = append(treeSuggestionList, s)
	}

	for i, node := range nodes {
		if node != nil {
			link := strings.Join(node.GetNodeLinks(), linkSep)
			alias := strings.Join(node.GetNodeAlias(), aliasSep)
			s := nodePrefix + nodePathes[i] + nodePathHint + link + linkHint + alias + aliasHint
			nodeSuggestionList = append(nodeSuggestionList, s)
		}
	}
	return utils.MergeList(treeSuggestionList, nodeSuggestionList).([]string)
}

func adSearchSuggestionInit() [][]string {
	suggestions := make([][]string, len(advancedSearchInputInit()))
	var suggestion []string
	nodes := utils.RootTree.FindAllNode("")
	for i := range advancedSearchInputInit() {
		switch i {
		case 0:
			suggestion = utils.RootTree.DeepGetAllSubtreeName()
		case 1:
			for _, node := range nodes {
				suggestion = append(suggestion, node.Link...)
			}
		case 2:
			for _, node := range nodes {
				suggestion = append(suggestion, node.Alias...)
			}
		case 3:
			for _, node := range nodes {
				suggestion = append(suggestion, node.Label...)
			}
		}
		if len(suggestion) != 0 {
			suggestion = utils.RemoveDup(suggestion)
		}
		suggestions[i] = suggestion
	}
	return suggestions
}

func searchInputInit() textinput.Model {
	ti := textinput.New()
	ti.KeyMap = textinput.KeyMap{}
	ti.Placeholder = "Input link or alias"
	ti.Focus()
	ti.SetSuggestions(searchSuggestionInit())
	ti.ShowSuggestions = true
	return ti
}

func advancedSearchInputInit() []textinput.Model {
	inputs := make([]textinput.Model, 4)
	var ti textinput.Model

	for i := range inputs {
		ti = textinput.New()
		ti.KeyMap = textinput.KeyMap{}
		ti.ShowSuggestions = true
		switch i {
		case 0:
			ti.Placeholder = "Tree path"
		case 1:
			ti.Placeholder = "Links, space to separate multiple links"
		case 2:
			ti.Placeholder = "Alias, space to separate multiple alias"
		case 3:
			ti.Placeholder = "Labels, space to separate multiple labels"
		}
		inputs[i] = ti
	}
	return inputs
}

func addInputInit() []textinput.Model {
	inputs := make([]textinput.Model, 6)
	var ti textinput.Model
	for i := range inputs {
		ti = textinput.New()
		ti.KeyMap = textinput.KeyMap{}
		ti.ShowSuggestions = true
		switch i {
		case 0:
			ti.Placeholder = "Tree"
		case 1:
			ti.Placeholder = "Links"
		case 2:
			ti.Placeholder = "Alias"
		case 3:
			ti.Placeholder = "Descriptions"
		case 4:
			ti.Placeholder = "Labels"
		case 5:
			ti.Placeholder = "Icon"
		}
		inputs[i] = ti
	}

	return inputs
}

func paginatorInit() paginator.Model {
	pt := paginator.New()
	pt.KeyMap = paginator.KeyMap{}
	pt.PerPage = treePerPage
	pt.Type = paginator.Dots
	pt.ActiveDot = activeDotStyle.Render("•")
	pt.InactiveDot = inactiveDotStyle.Render("•")
	pt.SetTotalPages(len(utils.RootTree.GetAllSubtreeName()))
	return pt
}

func delPopWinInit() popWin {
	vi := viewport.New(0, 0)
	vi.HighPerformanceRendering = false
	vi.KeyMap = viewport.KeyMap{}

	return popWin{
		viewport:      vi,
		bottons:       []string{"yes", "no"},
		hint:          "",
		selectedIndex: 0,
	}
}

func textareaInit() textarea.Model {
	ta := textarea.New()
	ta.KeyMap = textarea.KeyMap{
		CharacterForward: key.NewBinding(
			key.WithKeys("ctrl+l"),
			key.WithHelp("ctrl+l", "CharacterForward"),
		),
		CharacterBackward: key.NewBinding(
			key.WithKeys("ctrl+h"),
			key.WithHelp("ctrl+h", "CharacterBackward"),
		),
		DeleteCharacterBackward: key.NewBinding(
			key.WithKeys("backspace"),
			key.WithHelp("backspace", "DeleteWordBackward"),
		),
		InsertNewline: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "InsertNewline"),
		),
		LineEnd: key.NewBinding(
			key.WithKeys("ctrl+e"),
			key.WithHelp("ctrl+e", "LineEnd"),
		),
		LineStart: key.NewBinding(
			key.WithKeys("ctrl+q"),
			key.WithHelp("ctrl+q", "LineStart"),
		),
		LineNext: key.NewBinding(
			key.WithKeys("ctrl+j"),
			key.WithHelp("ctrl+j", "LineNext"),
		),
		LinePrevious: key.NewBinding(
			key.WithKeys("ctrl+k"),
			key.WithHelp("ctrl+k", "LinePrevious"),
		),
		Paste: key.NewBinding(
			key.WithKeys("ctrl+p"),
			key.WithHelp("ctrl+p", "Paste"),
		),
	}
	return ta
}

func InitialModel() Model {
	root := utils.RootTree
	h := help.New()
	h.ShowAll = false

	return Model{
		help:             h,
		searchInput:      searchInputInit(),
		adSearchInput:    advancedSearchInputInit(),
		addInput:         addInputInit(),
		paginator:        paginatorInit(),
		suggestionList:   searchSuggestionInit(),
		adsuggestionList: adSearchSuggestionInit(),
		textarea:         textareaInit(),

		tabSelected:     selected{index: 0},
		sugSelected:     selected{index: 0},
		adInpSelected:   selected{index: 0},
		addInpSelected:  selected{index: 0},
		subSelected:     point{x: 0, y: 0},
		preSelectedTree: []point{},

		subMsgs:    subMsg{ylen: []int{0}},
		tabs:       root.GetAllSubtreeName(),
		keymap:     keymap,
		toggle:     false,
		helpToggle: false,
		root:       root,
		curTree:    &root,
		mode:       search,
		lastMode:   search,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}
