package main

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := make(map[int]struct{})
	for _, num := range nums1 {
		set1[num] = struct{}{}
	}

	set2 := make(map[int]struct{})
	for _, num := range nums2 {
		set2[num] = struct{}{}
	}

	var diff1 []int
	for num := range set1 {
		if _, found := set2[num]; !found {
			diff1 = append(diff1, num)
		}
	}

	var diff2 []int
	for num := range set2 {
		if _, found := set1[num]; !found {
			diff2 = append(diff2, num)
		}
	}

	return [][]int{diff1, diff2}

}

func main() {
	fmt.Println(findDifference([]int{1, 2, 3}, []int{2, 4, 6}))

}

/*
2215. Find the Difference of Two Arrays
Easy
Topics
premium lock icon
Companies
Hint
Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:

answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
answer[1] is a list of all distinct integers in nums2 which are not present in nums1.
Note that the integers in the lists may be returned in any order.



Example 1:

Input: nums1 = [1,2,3], nums2 = [2,4,6]
Output: [[1,3],[4,6]]
Explanation:
For nums1, nums1[1] = 2 is present at index 0 of nums2, whereas nums1[0] = 1 and nums1[2] = 3 are not present in nums2. Therefore, answer[0] = [1,3].
For nums2, nums2[0] = 2 is present at index 1 of nums1, whereas nums2[1] = 4 and nums2[2] = 6 are not present in nums1. Therefore, answer[1] = [4,6].
Example 2:

Input: nums1 = [1,2,3,3], nums2 = [1,1,2,2]
Output: [[3],[]]
Explanation:
For nums1, nums1[2] and nums1[3] are not present in nums2. Since nums1[2] == nums1[3], their value is only included once and answer[0] = [3].
Every integer in nums2 is present in nums1. Therefore, answer[1] = [].
*/
