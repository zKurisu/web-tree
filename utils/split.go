package utils

import (
	"log"
	"regexp"
	"strings"
)

// webtree add --tree=root/fir/sec,another
// Problem: "," in url
func Split2List(s string) []string {
	if s[0] != '"' && s[0] != '\'' {
		log.Fatal("The args should be surround with quote")
	}

	patterns := []string{`",'`, `',"`, `','`}
	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		s = re.ReplaceAllString(s, `","`)
	}

	str := s[1 : len(s)-1]
	list := strings.Split(str, `","`)
	return list
}

func SplitTreeLevel(s string) []string {
	levels := strings.Split(s, "/")
	return levels
}
