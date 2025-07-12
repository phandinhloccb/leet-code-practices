package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
    // only 1 node
    if head == nil || head.Next == nil {
        return nil
    } 

    // total node
    length := 0
    for curr := head; curr != nil; curr = curr.Next {
        length++
    }

    middleIndex := length / 2
    curr := head
    for i := 0; i < middleIndex - 1; i++ {
        curr = curr.Next
    }
    
    curr.Next = curr.Next.Next

    return head

}

func main(){
	// [1,3,4,7,1,2,6]
	deleteMiddle(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{		
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 6,
							},
						},
					},
				},
			},
		},
	})
}