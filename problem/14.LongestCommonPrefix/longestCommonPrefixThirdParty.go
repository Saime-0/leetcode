package main

import "sort"

func longestCommonPrefixThirdParty(strs []string) string {
	if len(strs) == 1 { // handle only 1 element
		return strs[0]
	}

	// sort them first, the most different one will be in first and last
	sort.Strings(strs)

	// compare first and last
	lastIndex := len(strs) - 1
	for i := range strs[0] {
		if strs[0][i] != strs[lastIndex][i] {
			return strs[0][:i]
		}
	}
	return strs[0]
}
