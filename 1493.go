package main

import "fmt"

func longestSubarray(nums []int) int {
	left := 0
	zeroCount := 0
	maxLength := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}

		if zeroCount > 1 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		if right-left > maxLength {
			maxLength = right - left
		}
	}

	return maxLength

}

func main() {
	fmt.Println(longestSubarray([]int{1, 1, 0, 1}))
}

/*

Given a binary array nums, you should delete one element from it.

Return the size of the longest non-empty subarray containing only 1's in the resulting array. Return 0 if there is no such subarray.



Example 1:

Input: nums = [1,1,0,1]
Output: 3
Explanation: After deleting the number in position 2, [1,1,1] contains 3 numbers with value of 1's.
Example 2:

Input: nums = [0,1,1,1,0,1,1,0,1]
Output: 5
Explanation: After deleting the number in position 4, [0,1,1,1,1,1,0,1] longest subarray with value of 1's is [1,1,1,1,1].
Example 3:

Input: nums = [1,1,1]
Output: 2
Explanation: You must delete one element.

*/
