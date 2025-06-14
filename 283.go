package main

import "fmt"

func moveZeroes(nums []int) {
	lastNonZeroIndex := 0

	for _, num := range nums {
		if num != 0 {
			nums[lastNonZeroIndex] = num
			lastNonZeroIndex++
		}
	}

	for i := lastNonZeroIndex; i < len(nums); i++ {
		nums[i] = 0
	}

	fmt.Println(nums)

}

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
}

/*
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
Note that you must do this in-place without making a copy of the array.

Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
Example 2:

Input: nums = [0]
Output: [0]
*/
