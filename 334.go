package main

import "fmt"

func increasingTriplet(nums []int) bool {
	first := int(^uint(0) >> 1)
	second := int(^uint(0) >> 1)

	for _, num := range nums {
		if num <= first {
			first = num
		} else if num < second {
			second = num
		} else if num > second {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(increasingTriplet([]int{1,2,3,4,5}))
	fmt.Println(increasingTriplet([]int{5,4,3,2,1}))
	fmt.Println(increasingTriplet([]int{2,1,5,0,4,6}))
}

/*
Example 1:

Input: nums = [1,2,3,4,5]
Output: true
Explanation: Any triplet where i < j < k is valid.
Example 2:

Input: nums = [5,4,3,2,1]
Output: false
Explanation: No triplet exists.
Example 3:

Input: nums = [2,1,5,0,4,6]
Output: true
Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.
*/