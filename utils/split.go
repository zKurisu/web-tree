package utils

import (
	"strings"
)

// webtree add --tree=root/fir/sec,another
// Problem: "," in url
func Split2List(s string) []string {
	list := strings.Split(s, `,`)
	return RemoveEmp(list)
}

func SplitTreeLevel(s string) []string {
	levels := strings.Split(s, "/")
	return RemoveEmp(levels)
}

func MergeList(master []string, slave []string) []string {
	for _, elem := range slave {
		master = append(master, elem)
	}
	return master
}

func RemoveEmp(list []string) []string {
	newList := []string{}
	for _, elem := range list {
		if len(elem) != 0 {
			newList = append(newList, elem)
		}
	}
	return newList
}

func List2String(list []string) string {
	var s string
	for _, elem := range list {
		s = s + elem + ","
	}
	return s
}
