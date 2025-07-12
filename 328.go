package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}

	odd := head
	even := head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next 
		odd = odd.Next

		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead

	return head
}

func main() {
	oddEvenList(
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
							Next: nil,
						},
					},
				},
			},	
		},
	)
}

