package pkg0002

import (
	"testing"
)

func genNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	node := &ListNode{}
	temp := node

	for _, num := range nums {
		temp.Next = &ListNode{Val: num}
		temp = temp.Next
	}

	return node.Next
}

func next(n *ListNode) *ListNode {
	if n != nil {
		n = n.Next
		return n
	}
	return nil
}

func isSame(n1, n2 *ListNode) bool {
	if n1 == nil && n2 == nil {
		return true
	}

	for n1 != nil {
		if n1.Val != n2.Val {
			return false
		}
		n1 = next(n1)
		n2 = next(n2)
	}

	if n2 != nil {
		return false
	}

	return true
}

func TestAddTwoNumbersV1(t *testing.T) {
	l1, l2 := genNode([]int{2, 4, 3}), genNode([]int{5, 6, 4})
	if !isSame(addTwoNumbersV1(l1, l2), genNode([]int{7, 0, 8})) {
		t.Error("AddTwoNumbersV1 should return true")
	}
	if isSame(addTwoNumbersV1(l1, l2), genNode([]int{8, 0, 7})) {
		t.Error("AddTwoNumbersV1 should return false")
	}

	l1, l2 = genNode([]int{0, 1}), genNode([]int{0, 1, 2})
	if !isSame(addTwoNumbersV1(l1, l2), genNode([]int{0, 2, 2})) {
		t.Error("AddTwoNumbersV1 should return true")
	}

	l1, l2 = genNode([]int{}), genNode([]int{0, 1})
	if !isSame(addTwoNumbersV1(l1, l2), genNode([]int{0, 1})) {
		t.Error("AddTwoNumbersV1 should return true")
	}

	l1, l2 = genNode([]int{9, 9}), genNode([]int{1})
	if !isSame(addTwoNumbersV1(l1, l2), genNode([]int{0, 0, 1})) {
		t.Error("AddTwoNumbersV1 should return true")
	}
}

func TestAddTwoNumbersV2(t *testing.T) {
	l1, l2 := genNode([]int{2, 4, 3}), genNode([]int{5, 6, 4})
	if !isSame(addTwoNumbersV2(l1, l2), genNode([]int{7, 0, 8})) {
		t.Error("AddTwoNumbersV2 should return true")
	}
	if isSame(addTwoNumbersV2(l1, l2), genNode([]int{8, 0, 7})) {
		t.Error("AddTwoNumbersV2 should return false")
	}

	l1, l2 = genNode([]int{0, 1}), genNode([]int{0, 1, 2})
	if !isSame(addTwoNumbersV2(l1, l2), genNode([]int{0, 2, 2})) {
		t.Error("AddTwoNumbersV2 should return true")
	}

	l1, l2 = genNode([]int{}), genNode([]int{0, 1})
	if !isSame(addTwoNumbersV2(l1, l2), genNode([]int{0, 1})) {
		t.Error("AddTwoNumbersV2 should return true")
	}

	l1, l2 = genNode([]int{9, 9}), genNode([]int{1})
	if !isSame(addTwoNumbersV2(l1, l2), genNode([]int{0, 0, 1})) {
		t.Error("AddTwoNumbersV2 should return true")
	}

}
