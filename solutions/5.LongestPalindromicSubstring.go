package solutions

/*
5. Longest Palindromic Substring
Medium
23.6K
1.4K
Companies
Given a string s, return the longest
palindromic

substring
 in s.



Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"
*/

func longestPalindrome(s string) string {
	if s == "" || len(s) < 1 {
		return ""
	}
	start := 0
	end := 0

	for i := 0; i < len(s); i++ {
		len1 := expandFromCenter(s, i, i)
		len2 := expandFromCenter(s, i, i+1)
		len := 0
		if len1 > len2 {
			len = len1
		} else {
			len = len2
		}
		if len > end-start {
			start = i - ((len - 1) / 2)
			end = i + (len / 2)
		}

	}

	return s[start : end+1]
}

func expandFromCenter(s string, left, right int) int {
	if s == "" || left > right {
		return 0
	}

	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

/*
 func longestPalindrome(s string) string {
    sub:=""
    ss:=""

    for i:=0; i<len(s); i++ {
        for j:=i; j<=len(s); j++{
        sub=s[i:j]
                   if invertString(sub) == sub{
                if len(ss) < len (sub){
                    ss=sub

                }

        }
        }


            }
    return ss
}

func invertString(s string) string{
    r:=""
    for i:=len(s)-1; i>=0; i--{
        r+=string(s[i])
    }
    return r
}

*/
