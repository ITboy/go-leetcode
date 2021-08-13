package lesson3

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	maxLen, sublen, length := 0, 1, len(s)
	charmap := map[rune]int{}
	chars := []rune(s)
	for start := 0; start+maxLen < length; start++ {
		i := start + sublen - 1
		for ; i < length; i++ {
			c := chars[i]
			if _, ok := charmap[c]; ok {
				sublen = i - start
				fmt.Println(i, sublen)
				break
			} else {
				charmap[c] = i
			}
		}
		if i == length {
			sublen = length - start
		}
		if sublen > maxLen {
			maxLen = sublen
		}
		delete(charmap, chars[start])
	}
	return maxLen
}
