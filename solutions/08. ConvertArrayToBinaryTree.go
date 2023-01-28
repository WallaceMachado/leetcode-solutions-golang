package solutions

/**
108. Convert Sorted Array to Binary Search Tree
Easy
9K
453
Companies
Given an integer array nums where the elements are sorted in ascending order, convert it to a
height-balanced
 binary search tree.



Example 1:


Input: nums = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Explanation: [0,-10,5,null,-3,null,9] is also accepted:

Example 2:


Input: nums = [1,3]
Output: [3,1]
Explanation: [1,null,3] and [3,1] are both height-balanced BSTs.







 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/
func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil {
		return nil
	}
	return recursion(nums, 0, len(nums)-1)
}

func recursion(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := left + (right-left)/2
	root := &TreeNode{
		Val:   nums[mid],
		Left:  recursion(nums, left, mid-1),
		Right: recursion(nums, mid+1, right),
	}
	return root
}
