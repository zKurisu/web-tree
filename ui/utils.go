package ui

import (
	"regexp"
	"sort"
	"strings"
)

func Fuzzy(src string, targets []string) []string {
	if src == "" {
		return nil
	}

	list := []string{}
	sort.Strings(targets)

	pattern := regexp.QuoteMeta(src)
	re := regexp.MustCompile(pattern)

	for _, elem := range targets {
		if re.MatchString(elem) {
			start := strings.Index(elem, re.String())
			end := start + len(re.String())
			matched := suggestionMatchedStyle.Render(elem[start:end])
			s := elem[:start] + matched + elem[end:]
			list = append(list, s)
		}
	}
	return list
}

// items there is all subtrees of root
func (m Model) getItems() []string {
	return m.root.GetSubtreesName()
}
