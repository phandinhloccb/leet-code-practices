package main

import "fmt"

// func productExceptSelf(nums []int) []int {
// 	left := make([]int, len(nums))

// 	for i := 0; i < len(nums); i++ {
// 		if i==0 {
// 			left[i] = 1
// 		}	else {
// 			left[i] = nums[i-1] * left[i-1]
// 		}
// 	}

// 	right := make([]int, len(nums))

// 	for i := len(nums) - 1; i >= 0; i-- {
// 		if i == len(nums)-1 {
// 			right[i] = 1
// 		} else {
// 			right[i] = nums[i+1] * right[i+1]
// 		}
// 	}

// 	for i := 0; i < len(nums); i++ {
// 		nums[i] = left[i] * right[i]
// 	}

// 	return nums
// } //T:O(n) //S:O(n)

func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))

	left :=1
	for i, val := range nums {
		answer[i] = left
		left *= val
	}

	fmt.Println("answer", answer)

	right :=1
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] = right * answer[i]
		right *= nums[i]
	}

	fmt.Println("answer", answer)

	return answer
}

func main() {
	// Test the function
	nums := []int{1, 2, 3, 4}
	result := productExceptSelf(nums)
	fmt.Println("Output:", result)
}

/*
Example 1:
Input: nums = [1,2,3,4]
left = [1,1,2,6]
right = [24,12,4,1]

nums = [1,2,3,4]
answers = [1,1,2,6]
answer = [24,12,8,6]

Output: [24,12,8,6]

Example 2:
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]

solution:
Thay vì tính lại tích mỗi lần, ta có thể chia quá trình ra thành 2 phần:
- Tích từ bên trái (prefix product)
- Tích từ bên phải (suffix product)

→ answer[i] = left[i] * right[i], trong
*/