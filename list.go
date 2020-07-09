package leetCode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) String() string {
	if head == nil {
		return "//"
	}
	return fmt.Sprintf("%d -> ", head.Val) + head.Next.String()
}

func Ints2List(nums ...int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := ListNode{Val: nums[0]}
	tail := &head
	for _, n := range nums[1:] {
		tail.Next = &ListNode{Val: n}
		tail = tail.Next
	}
	return &head
}
