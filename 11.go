package main

import "fmt"

func maxArea(height []int) int {
	p1 := 0
	p2 := len(height) - 1
	maxArea := 0

	for p1 <= p2 {
		area := (p2 - p1) * min(height[p1], height[p2])
		maxArea = max(maxArea, area)

		if height[p1] > height[p2] {
			p2--
		} else {
			p1++
		}
	}

	return maxArea
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

/*

You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
Find two lines that together with the x-axis form a container, such that the container contains the most water.
Return the maximum amount of water a container can store.
Notice that you may not slant the container.

Example 1:
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
Example 2:

Input: height = [1,1]
Output: 1

*/
