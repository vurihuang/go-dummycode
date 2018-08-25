package pkg0002

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbersV1(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	temp, carry, v := res, 0, 0

	for l1 != nil || l2 != nil {
		v = carry
		if l1 != nil {
			v += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		}
		carry, v = v/10, v%10
		temp.Next = &ListNode{Val: v}
		temp = temp.Next
	}

	if carry == 1 {
		temp.Next = &ListNode{Val: 1}
	}

	return res.Next
}

func addTwoNumbersV2(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	temp, carry, v := res, 0, 0

	for l1 != nil || l2 != nil || carry != 0 {
		v = carry
		if l1 != nil {
			v += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		}
		carry, v = v/10, v%10
		temp.Next = &ListNode{Val: v}
		temp = temp.Next
	}

	return res.Next
}
