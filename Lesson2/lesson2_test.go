package lesson2

import (
	"math"
	"testing"
)

type testcase struct {
	num1, num2 *ListNode
	expected   *ListNode
}

func num2list(num int) *ListNode {
	var p *ListNode = &ListNode{}
	var head = p
	for num != 0 {
		p.Next = &ListNode{Val: num % 10}
		p = p.Next
		num = num / 10
	}
	return head.Next
}

func list2num(l *ListNode) int {
	sum := 0
	jie := 0
	for l != nil {
		sum += int(math.Pow10(jie)) * l.Val
		l = l.Next
		jie++
	}
	return sum
}

func equalList(l1, l2 *ListNode) bool {
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil && l1.Val != l2.Val {
			return false
		}
		if (l1 != nil && l2 == nil) || (l2 != nil && l1 == nil) {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}

func TestAddTwoNumbers(t *testing.T) {
	cases := []testcase{
		{num1: num2list(342), num2: num2list(465), expected: num2list(807)},
		{num1: num2list(0), num2: num2list(0), expected: num2list(0)},
	}

	for _, c := range cases {
		if got := addTwoNumbers(c.num1, c.num2); !equalList(got, c.expected) {
			t.Errorf("addTwoNumbers(%v, %v) = %v, expected: %v", list2num(c.num1), list2num(c.num2), list2num(got), list2num(c.expected))
		}
	}
}
