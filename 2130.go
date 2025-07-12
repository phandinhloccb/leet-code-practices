package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func pairSum(head *ListNode) int {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    var prev *ListNode
    curr := slow
    for curr != nil {
        nextTemp := curr.Next
        curr.Next = prev
        prev = curr
        curr = nextTemp
    }


    maxSum := 0
    left := head
    right := prev 
    for right != nil {
        sum := left.Val + right.Val
        if sum > maxSum {
            maxSum = sum
        }
        left = left.Next
        right = right.Next
    }

    return maxSum
}

func main() {
	pairSum(
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 6,
							},
						},
					},
				},
			},
		},
	)
}
