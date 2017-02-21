package algorithms

import (
	"os"
	"fmt"
)

func LengthOfLongestSubstring(s string) int {
	start := 0
	max := 0
	dict := make(map[string]int)
	for i, j := range s {
		char := string(j)
		if _, ok := dict[char]; !ok {
			dict[char] = i
		} else {
			if start <= dict[char] {
				start = dict[char] + 1
			}
			dict[char] = i
		}
		if max <= i - start {
			max = i - start + 1
		}

	}
	return max
}
