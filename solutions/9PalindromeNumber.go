package solutions

/**
9. Palindrome Number
Easy
8.5K
2.4K
Companies
Given an integer x, return true if x is a
palindrome
, and false otherwise.



Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

func isPalindrome(x int) bool {
       r:=true
       if x < 0 || (x > 0 && x%10 == 0){
           return false
       }
       sX:=strconv.Itoa(x)
        c:=len(sX)-1
       for i:=0; i<len(sX); i++{
           if sX[i] != sX[c]{
               r=false
           }
           c--
       }
       return r
}

*/

func isPalindromes(x int) bool {
	if x < 0 || (x > 0 && x%10 == 0) {
		return false
	}

	var numeroRverso int
	tmp := x
	for tmp > 0 {
		numeroRverso = numeroRverso*10 + tmp%10
		tmp = tmp / 10
	}
	return numeroRverso == x
}
