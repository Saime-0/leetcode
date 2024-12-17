package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	shortestStr := strs[0]
	for _, str := range strs {
		if len(str) < len(shortestStr) {
			shortestStr = str
		}
	}
	if shortestStr == "" {
		return ""
	}
	var lastBeatPrefix string
	for i, r := range shortestStr {
		for _, str := range strs {
			if []rune(str)[i] != r {
				return lastBeatPrefix
			}
		}
		lastBeatPrefix += string(r)
	}

	return lastBeatPrefix
}
