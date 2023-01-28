package solutions

/*
217. Contains Duplicate
Easy
8.1K
1.1K
Companies
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.



Example 1:

Input: nums = [1,2,3,1]
Output: true
Example 2:

Input: nums = [1,2,3,4]
Output: false
Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true

func containsDuplicate(nums []int) bool {
   m:=map[int]int{}
   for i:=0; i<len(nums); i++{
       if m[nums[i]] > 0 {
           return true
       }else{
           m[nums[i]]++
       }
   }
   return false
}
*/

func containsDuplicate(nums []int) bool {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		_, exists := m[nums[i]]
		if exists {
			return true
		} else {
			m[nums[i]] = i
		}
	}
	return false
}
