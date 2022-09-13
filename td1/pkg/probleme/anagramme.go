package probleme

import (
	"sort"
	"strings"
)

func Footprint(s string) (footprint string) {
	// sort.Slice method accept only slices, so a transform from string to slice is required
	ssl := strings.Split(s, "")
	sort.Slice(ssl, func(i, j int) bool { return ssl[i]<ssl[j] })
	// re-transform slice to string after sort
	footprint = strings.Join(ssl, "")
	return
}

func Anagrams(words []string) (anagrams map[string][]string) {
	anagrams = make(map[string][]string)
	for _,word := range words {
		if list, ok := anagrams[Footprint(word)] ; ok {
			list = append(list, word)
			anagrams[Footprint(word)] = list
		} else {
			anagrams[Footprint(word)] = []string{word}
		}
	}
	return
}