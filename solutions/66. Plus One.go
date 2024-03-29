package solutions

/*
You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.



Example 1:

Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Incrementing by one gives 123 + 1 = 124.
Thus, the result should be [1,2,4].
Example 2:

Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
Incrementing by one gives 4321 + 1 = 4322.
Thus, the result should be [4,3,2,2].
Example 3:

Input: digits = [9]
Output: [1,0]
Explanation: The array represents the integer 9.
Incrementing by one gives 9 + 1 = 10.
Thus, the result should be [1,0].
*/

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	if digits[0] == 0 {
		digits = append([]int{1}, digits...)

	}

	return digits
}

/*
 func plusOne(digits []int) []int {
    if len(digits)<1{
        return digits
    }
    i:=len(digits)-1
    c:=digits[len(digits)-1]
    a:=1
    for a>0 && i >=0{
        if c == 9 && i ==0 {
            digits[0]=1
            digits= append(digits, 0)
        }else if  c  == 9{
            digits[i]=0
            c= digits[i-1]
            a=1
        }else {
            digits[i]=digits[i] +1
            a=0
        }
        i--
    }

    return digits
}
*/
