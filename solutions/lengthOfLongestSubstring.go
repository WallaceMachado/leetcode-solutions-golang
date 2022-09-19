package solutions

/*
3. Longest Substring Without Repeating Characters
Medium

28705

1223

Add to List

Share
Given a string s, find the length of the longest substring without repeating characters.



Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	var (
		longest int
		ss      string
	)
	for _, letter := range s {
		fmt.Println(string(letter), "  ", ss)
		index := strings.Index(ss, string(letter))
		if index != -1 {
			if len(ss) > longest {
				longest = len(ss)
			}
			fmt.Println(string(letter), "  ", ss)
			ss = ss[index+1:]
			fmt.Println(string(letter), "  ", ss)
		}
		ss += string(letter)
	}
	if len(ss) > longest {
		longest = len(ss)
	}

	return longest
}
