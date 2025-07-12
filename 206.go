
package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode = nil

    cur := head
    for cur != nil {
        next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
    }

	return prev
    
}

func main() {
	reverseList(
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
						},
					},
				},
			},
		},
	)
}