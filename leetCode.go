package leetCode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) String() {
	if head == nil {
		return "-> //"
	}
	return fmt.Sprintf("%d -> ", head.Val) + head.Next.String()
}

func Ints2List(nums ...int) ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := ListNode{Val: nums[0]}
	lastNode := &head
	for _, n := range nums[1:len(nums)] {
		lastNode.Next = &ListNode{Val: n}
		lastNode = lastNode.Next
	}
	return head
}
