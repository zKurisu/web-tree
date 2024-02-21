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
	treePathes := utils.RootTree.DeepGetAllSubtreePath("")
	nodes := utils.RootTree.FindAllNode("")
	treeSuggestionList := []string{}
	nodeSuggestionList := []string{}

	for _, treePath := range treePathes {
		treePrefix := suggestionTreeStyle.Render("*tree ")
		nameHint := suggestionQuoteStyle.Render(" (name) ")
		s := treePrefix + treePath + nameHint
		treeSuggestionList = append(treeSuggestionList, s)
	}

	for _, node := range nodes {
		nodePrefix := suggestionNodeStyle.Render("*node ")
		linkHint := suggestionQuoteStyle.Render(" (link) ")
		aliasHint := suggestionQuoteStyle.Render(" (alias) ")
		link := strings.Join(node.GetNodeLinks(), " ")
		alias := strings.Join(node.GetNodeAlias(), " ")
		s := nodePrefix + link + linkHint + alias + aliasHint
		nodeSuggestionList = append(nodeSuggestionList, s)
	}
	return utils.MergeList(treeSuggestionList, nodeSuggestionList)
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
	pt.SetTotalPages(len(utils.RootTree.GetSubtreesName()))
	return pt
}

func InitialModel() Model {
	root := utils.RootTree

	return Model{
		help:             help.New(),
		searchInput:      inputInit(),
		paginator:        paginatorInit(),
		suggestionList:   SuggestionInit(),
		sugSelectedIndex: 0,
		items:            root.GetSubtreesName(),
		keymap:           keymap,
		root:             root,
		mode:             search,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}
