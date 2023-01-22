package solutions

/**
283. Move Zeroes
Easy
12.4K
314
Companies
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.



Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
Example 2:

Input: nums = [0]
Output: [0]
*/

func moveZeroes(nums []int) {
	index := 0
	for i, v := range nums {
		if v != 0 {
			nums[index] = v
			if i != index {

				nums[i] = 0

			}
			index++
		}

	}
}
