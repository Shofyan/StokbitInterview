package util

import (
	"sort"
	"strings"
)

func Anagram(in []string) (dict map[string][]string) {
	if len(in) == 0 {
		return
	}

	dict = map[string][]string{}

	// get all string in array
	for _, v := range in {
		s := strings.Split(v, "")
		sort.Strings(s)
		kata := strings.Join(s, " ")

		if len(dict) == 0 {
			dict[kata] = append(dict[kata], v)
		} else {
			dict[kata] = append(dict[kata], v)
		}

	}

	return
}
