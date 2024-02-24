package ui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
	"web-tree/utils"
)

// All tree name (expect root) and all nodes
func SuggestionInit() []string {
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

func inputInit() textinput.Model {
	ti := textinput.New()
	ti.KeyMap = textinput.KeyMap{}
	ti.Placeholder = "Input link or alias"
	ti.Focus()
	ti.SetSuggestions(SuggestionInit())
	ti.ShowSuggestions = true
	return ti
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

func InitialModel() Model {
	root := utils.RootTree

	return Model{
		help:           help.New(),
		searchInput:    inputInit(),
		paginator:      paginatorInit(),
		suggestionList: SuggestionInit(),
		tabSelected:    selected{index: 0},
		sugSelected:    selected{index: 0},
		subSelected:    point{x: 0, y: 0},
		tabs:           root.GetAllSubtreeName(),
		keymap:         keymap,
		toggle:         false,
		root:           root,
		mode:           search,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}
