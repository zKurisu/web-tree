package utils

import (
	"strings"
)

// webtree add --tree=root/fir/sec,another
// Problem: "," in url
func Split2List(s string) []string {
	list := strings.Split(s, `,`)
	return list
}

func SplitTreeLevel(s string) []string {
	levels := strings.Split(s, "/")
	return levels
}
