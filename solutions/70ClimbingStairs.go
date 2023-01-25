package solutions

/*
70. Climbing Stairs
Easy
16.8K
512
Companies
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?



Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
*/

func climbStairs(n int) int {
	c := 1
	a := 0
	for i := 0; i < n; i++ {
		b := c
		c = c + a
		a = b

	}
	return c
}

/*
func climbStairs(n int) int {
    array:=[]int{1,1}
    for i:=2; i<= n; i++ {
        array = append(array, (array[i-1]+array[i-2]))

    }
    return array[n]
}

*/
