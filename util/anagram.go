package util

import (
	"sort"
	"strings"
)

func Anagram(in []string) (dict map[string][]string) {
	if len(in) == 0 {
		return nil
	}

	dict = map[string][]string{}
	for _, v := range in {
		s := strings.Split(v, "")
		sort.Strings(s)
		kata := strings.Join(s, "")
		dict[kata] = append(dict[kata], v)
	}

	return
}
