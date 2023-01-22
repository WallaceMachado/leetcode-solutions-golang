package solutions

/*
1844. Replace All Digits with Characters
Easy

509

51

Add to List

Share
You are given a 0-indexed string s that has lowercase English letters in its even indices and digits in its odd indices.

There is a function shift(c, x), where c is a character and x is a digit, that returns the xth character after c.

For example, shift('a', 5) = 'f' and shift('x', 0) = 'x'.
For every odd index i, you want to replace the digit s[i] with shift(s[i-1], s[i]).

Return s after replacing all digits. It is guaranteed that shift(s[i-1], s[i]) will never exceed 'z'.



Example 1:

Input: s = "a1c1e1"
Output: "abcdef"
Explanation: The digits are replaced as follows:
- s[1] -> shift('a',1) = 'b'
- s[3] -> shift('c',1) = 'd'
- s[5] -> shift('e',1) = 'f'
Example 2:

Input: s = "a1b2c3d4e"
Output: "abbdcfdhe"
Explanation: The digits are replaced as follows:
- s[1] -> shift('a',1) = 'b'
- s[3] -> shift('b',2) = 'd'
- s[5] -> shift('c',3) = 'f'
- s[7] -> shift('d',4) = 'h'
*/

import (
	"strconv"
	"strings"
)

func replaceDigits(s string) string {
	output := ""
	for i := 1; i < len(s); i = i + 2 {
		output += string(s[i-1])
		pos, _ := strconv.Atoi(string(s[i]))
		output += string(int(s[i-1]) + pos)

	}
	if len(s)%2 != 0 {
		output += string(s[len(s)-1])
	}
	values := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for i, v := range s {
		strings.
		for i, _ := range s {
			strings.Index()
			if values[s[i]] < values[s[i+1]] {
				res += values[s[i+1]] - values[s[i]]
			} else {
				res += values[s[i]]
			}
			a

		}
	}
	return output

}
